package service

import (
	"context"
	"fmt"

	"github.com/lavatee/tracker_backend/internal/model"
	"github.com/lavatee/tracker_backend/internal/repository"
)

const (
	prizeNode = "Награды"
)

type NodesService struct {
	repo *repository.Repository
}

func NewNodesService(repo *repository.Repository) *NodesService {
	return &NodesService{
		repo: repo,
	}
}

func (s *NodesService) GetNextNodes(ctx context.Context, id int64) ([]model.Node, error) {
	nodes, err := s.repo.Nodes.GetNextNodes(ctx, id)
	if err != nil {
		return nil, err
	}
	if id == 0 {
		achievementsNodes := make([]model.Node, 0)
		for _, node := range nodes {
			if node.Name != prizeNode {
				achievementsNodes = append(achievementsNodes, node)
			}
		}
		return achievementsNodes, nil
	}
	return nodes, nil
}

func (s *NodesService) GetPreviousNodes(ctx context.Context, id int64) ([]model.Node, error) {
	return s.repo.Nodes.GetPreviousNodes(ctx, id)
}

func (s *NodesService) UpdateNode(ctx context.Context, id int64, name string, points int, userId int) error {
	isAdmin := s.repo.Users.CheckIsAdmin(userId)
	if !isAdmin {
		return fmt.Errorf("user is not admin")
	}
	return s.repo.Nodes.UpdateNode(ctx, id, name, points)
}

func (s *NodesService) AddNode(ctx context.Context, parentID int64, name string, points int, userId int) (int64, error) {
	isAdmin := s.repo.Users.CheckIsAdmin(userId)
	if !isAdmin {
		return 0, fmt.Errorf("user is not admin")
	}
	id, err := s.repo.Nodes.AddNode(ctx, parentID, name, points)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *NodesService) GetNodeByID(ctx context.Context, id int64) (model.Node, error) {
	return s.repo.Nodes.GetNodeByID(ctx, id)
}
