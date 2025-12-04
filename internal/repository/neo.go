package repository

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ConnectNeoDB() (neo4j.DriverWithContext, error) {
	uri := "neo4j://localhost:7687"
	username := "neo4j"
	password := "password"

	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, fmt.Errorf("neo4j connection error: %w", err)
	}

	return driver, nil
}
