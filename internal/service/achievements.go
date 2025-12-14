package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/lavatee/tracker_backend/internal/model"
	"github.com/lavatee/tracker_backend/internal/repository"
	"github.com/minio/minio-go/v7"
)

const (
	pendingAchievementStatus  = "pending"
	approvedAchievementStatus = "approved"
	rejectedAchievementStatus = "rejected"
)

type AchievementsService struct {
	repo   *repository.Repository
	s3     *minio.Client
	bucket string
}

func NewAchievementsService(repo *repository.Repository, s3 *minio.Client, bucket string) *AchievementsService {
	return &AchievementsService{
		repo:   repo,
		s3:     s3,
		bucket: bucket,
	}
}

func GetDocumentURL(key string) string {
	return fmt.Sprintf("https://5a1bc5f7-b5c2-4a61-969a-beacbd4d7999.selstorage.ru/%s", key)
}

func (s *AchievementsService) CreateAchievement(ctx context.Context, ach model.Achievement, fileName string, file multipart.File) (int, error) {
	documentKey := fmt.Sprintf("%d_%d.%s", ach.UserId, ach.NodeId, strings.Split(fileName, ".")[len(strings.Split(fileName, "."))-1])
	ach.DocumentUrl = GetDocumentURL(documentKey)
	data, err := io.ReadAll(file)
	if err != nil {
		return 0, err
	}
	id, err := s.repo.Achievements.CreateAchievement(ach)
	if err != nil {
		return 0, err
	}
	_, err = s.s3.PutObject(ctx, s.bucket, fileName, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{})
	if err != nil {
		return 0, err
	}
	return id, err
}

func (s *AchievementsService) DeleteAchievement(achId int, userId int) error {
	return s.repo.Achievements.DeleteAchievement(achId, userId)
}

func (s *AchievementsService) GetUserAchievements(userId int) ([]model.Achievement, error) {
	return s.repo.Achievements.GetUserAchievements(userId)
}

func (s *AchievementsService) GetAchievementsByStatus(status string, userId int) ([]model.Achievement, error) {
	return s.repo.Achievements.GetAchievementsByStatus(status)
}

func (s *AchievementsService) ApproveAchievement(achId int, userId int) error {
	if isAdmin := s.repo.Users.CheckIsAdmin(userId); !isAdmin {
		return fmt.Errorf("user is not an admin")
	}
	return s.repo.Achievements.SetAchievementStatus(achId, approvedAchievementStatus)
}

func (s *AchievementsService) RejectAchievement(achId int, userId int, comment string) error {
	if isAdmin := s.repo.Users.CheckIsAdmin(userId); !isAdmin {
		return fmt.Errorf("user is not an admin")
	}
	if err := s.repo.Achievements.SetRejectComment(achId, comment); err != nil {
		return err
	}
	return s.repo.Achievements.SetAchievementStatus(achId, rejectedAchievementStatus)
}

func (r *AchievementsService) GetAchievementById(achId int) (model.Achievement, error) {
	return r.repo.Achievements.GetAchievementById(achId)
}
