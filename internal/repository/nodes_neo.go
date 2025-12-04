package repository

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type NodesNeo struct {
	driver neo4j.DriverWithContext
}

func (r *NodesNeo) GetNextNodes(ctx context.Context, driver neo4j.DriverWithContext, id int64) ([]Node, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	if id == 0 {
		query := `
			MATCH (root:Node {name:"ROOT"})-[:NEXT]->(child)
			RETURN ID(child) AS id, child.name AS name, child.points AS points
		`

		result, err := session.Run(ctx, query, nil)
		if err != nil {
			return nil, err
		}

		var nodes []Node

		for result.Next(ctx) {
			rec := result.Record()
			nodes = append(nodes, Node{
				ID:     rec.Values[0].(int64),
				Name:   rec.Values[1].(string),
				Points: int(rec.Values[2].(int64)),
			})
		}

		return nodes, result.Err()
	}

	query := `
		MATCH (n:Node)-[:NEXT]->(child)
		WHERE ID(n) = $id
		RETURN ID(child) AS id, child.name AS name, child.points AS points
	`

	result, err := session.Run(ctx, query, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	var nodes []Node

	for result.Next(ctx) {
		rec := result.Record()
		nodes = append(nodes, Node{
			ID:     rec.Values[0].(int64),
			Name:   rec.Values[1].(string),
			Points: int(rec.Values[2].(int64)),
		})
	}

	return nodes, result.Err()
}

func (r *NodesNeo) GetPreviousNodes(ctx context.Context, driver neo4j.DriverWithContext, id int64) ([]Node, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
		MATCH (parent)-[:NEXT]->(n:Node)
		WHERE ID(n) = $id
		RETURN ID(parent) AS id, parent.name AS name, parent.points AS points
	`

	result, err := session.Run(ctx, query, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	var parents []Node

	for result.Next(ctx) {
		rec := result.Record()
		parents = append(parents, Node{
			ID:     rec.Values[0].(int64),
			Name:   rec.Values[1].(string),
			Points: int(rec.Values[2].(int64)),
		})
	}

	return parents, result.Err()
}

func (r *NodesNeo) UpdateNode(ctx context.Context, driver neo4j.DriverWithContext, id int64, name string, points int) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	query := `
		MATCH (n:Node)
		WHERE ID(n) = $id
		SET n.name = $name,
		    n.points = $points
	`

	_, err := session.Run(ctx, query, map[string]interface{}{
		"id":     id,
		"name":   name,
		"points": points,
	})

	return err
}

func (r *NodesNeo) AddNode(ctx context.Context, driver neo4j.DriverWithContext, parentID int64, name string, points int) (int64, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	query := `
		MATCH (p:Node)
		WHERE ID(p) = $parentID
		CREATE (n:Node {name: $name, points: $points})
		CREATE (p)-[:NEXT]->(n)
		RETURN ID(n)
	`

	result, err := session.Run(ctx, query, map[string]interface{}{
		"parentID": parentID,
		"name":     name,
		"points":   points,
	})

	if err != nil {
		return 0, err
	}

	if result.Next(ctx) {
		return result.Record().Values[0].(int64), nil
	}

	return 0, fmt.Errorf("no result returned")
}

func (r *NodesNeo) GetNodeByID(ctx context.Context, driver neo4j.DriverWithContext, id int64) (*Node, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
		MATCH (n:Node)
		OPTIONAL MATCH (parent)-[:NEXT]->(n)
		WHERE ID(n) = $id
		RETURN ID(n) AS id, n.name AS name, n.points AS points, ID(parent) AS parentID
	`

	result, err := session.Run(ctx, query, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, err
	}

	if !result.Next(ctx) {
		return nil, fmt.Errorf("node not found")
	}

	rec := result.Record()

	node := Node{
		ID:     rec.Values[0].(int64),
		Name:   rec.Values[1].(string),
		Points: int(rec.Values[2].(int64)),
	}

	if rec.Values[3] != nil {
		node.ParentID = rec.Values[3].(int64)
	}

	return &node, nil
}
