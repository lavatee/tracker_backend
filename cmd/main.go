package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	trackerbackend "github.com/lavatee/tracker_backend"
	"github.com/lavatee/tracker_backend/internal/endpoint"
	"github.com/lavatee/tracker_backend/internal/repository"
	"github.com/lavatee/tracker_backend/internal/service"
	_ "github.com/lib/pq"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	logrus.Infof("Database configuration: host=%s, port=%s, user=%s, dbname=%s, sslmode=%s",
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.user"),
		viper.GetString("db.dbname"),
		viper.GetString("db.sslmode"))

	db, err := repository.NewPostgresDB(repository.PostgresConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error initializing db: %s", err.Error())
	}
	defer db.Close()
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		logrus.Fatalf("Failed to create migrate driver: %s", err.Error())
	}

	migrationsPath := "file://schema"
	migrations, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		logrus.Fatalf("Failed to create migrate instance: %s", err.Error())
	}
	if err = migrations.Up(); err != nil && err != migrate.ErrNoChange {
		logrus.Fatalf("Migrations error: %s", err.Error())
	}
	neoDb, err := repository.ConnectNeoDB(viper.GetString("neo4j.user"), viper.GetString("neo4j.password"))
	if err != nil {
		logrus.Fatalf("neo connect error: %s", err.Error())
	}
	if err := RunNeoMigrations(context.Background(), neoDb, "file://node_schema"); err != nil {
		logrus.Fatalf("neo migration error: %s", err.Error())
	}
	repo := repository.NewRepository(db, neoDb)
	services := service.NewService(repo)
	endp := endpoint.NewEndpoint(services)
	server := &trackerbackend.Server{}
	go func() {
		if err := server.Run(viper.GetString("port"), endp.InitRoutes()); err != nil {
			logrus.Fatalf("error running http server: %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("Shutting down server...")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error shutting down server: %s", err.Error())
	}

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func RunNeoMigrations(ctx context.Context, driver neo4j.DriverWithContext, path string) error {
	files, err := filepath.Glob(filepath.Join(path, "*.cypher"))
	if err != nil {
		return err
	}

	sort.Strings(files)

	appliedMigrations, err := getAppliedMigrations(ctx, driver)
	if err != nil {
		return err
	}

	for _, file := range files {
		name := filepath.Base(file)

		if appliedMigrations[name] {
			fmt.Println("Already applied:", name)
			continue
		}

		fmt.Println("Applying:", name)

		content, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		if err := applyMigration(ctx, driver, string(content)); err != nil {
			return fmt.Errorf("migration %s failed: %w", name, err)
		}

		if err := saveMigrationRecord(ctx, driver, name); err != nil {
			return fmt.Errorf("failed to save migration record for %s: %w", name, err)
		}

		fmt.Println("Applied:", name)
	}

	return nil
}

func getAppliedMigrations(ctx context.Context, driver neo4j.DriverWithContext) (map[string]bool, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `MATCH (m:Migration) RETURN m.id AS id`

	res, err := session.Run(ctx, query, nil)
	if err != nil {
		return nil, err
	}

	migrations := make(map[string]bool)

	for res.Next(ctx) {
		id := res.Record().Values[0].(string)
		migrations[id] = true
	}

	return migrations, nil
}

func applyMigration(ctx context.Context, driver neo4j.DriverWithContext, query string) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	commands := strings.Split(query, ";")

	for _, cmd := range commands {
		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}

		_, err := session.Run(ctx, cmd, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func saveMigrationRecord(ctx context.Context, driver neo4j.DriverWithContext, name string) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	query := `
		CREATE (:Migration {
			id: $id,
			applied_at: $ts
		})
	`

	_, err := session.Run(ctx, query, map[string]interface{}{
		"id": name,
		"ts": time.Now().Format(time.RFC3339),
	})

	return err
}
