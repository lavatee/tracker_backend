package service

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"strings"

	"fmt"

	"github.com/hashicorp/go-uuid"
	"github.com/lavatee/tracker_backend/internal/model"
	"github.com/lavatee/tracker_backend/internal/repository"
	"github.com/minio/minio-go/v7"
)

type ProductsService struct {
	repo   *repository.Repository
	s3     *minio.Client
	bucket string
}

func NewProductsService(repo *repository.Repository, s3 *minio.Client, bucket string) *ProductsService {
	return &ProductsService{
		repo:   repo,
		s3:     s3,
		bucket: bucket,
	}
}

func (s *ProductsService) CreateProduct(ctx context.Context, userId int, product model.Product, fileName string, file multipart.File) (int, error) {
	if isAdmin := s.repo.Users.CheckIsAdmin(userId); !isAdmin {
		return 0, fmt.Errorf("user is not an admin")
	}
	photoId, err := uuid.GenerateUUID()
	if err != nil {
		return 0, err
	}
	product.PhotoUrl = GetDocumentURL(fmt.Sprintf("%s.%s", photoId, strings.Split(fileName, ".")[len(strings.Split(fileName, "."))-1]))
	id, err := s.repo.Products.CreateProduct(product)
	if err != nil {
		return 0, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return 0, err
	}
	_, err = s.s3.PutObject(ctx, s.bucket, fileName, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{})
	if err != nil {
		return 0, err
	}
	return id, err
}

func (s *ProductsService) GetProducts() ([]model.Product, error) {
	return s.repo.Products.GetProducts()
}

func (s *ProductsService) GetProductById(productId int) (model.Product, error) {
	return s.repo.Products.GetProductById(productId)
}

func (s *ProductsService) DeleteProduct(productId int, userId int) error {
	if isAdmin := s.repo.Users.CheckIsAdmin(userId); !isAdmin {
		return fmt.Errorf("user is not an admin")
	}
	return s.repo.Products.DeleteProduct(productId)
}

func (s *ProductsService) UpdateProduct(product model.Product, userId int) error {
	if isAdmin := s.repo.Users.CheckIsAdmin(userId); !isAdmin {
		return fmt.Errorf("user is not an admin")
	}
	return s.repo.Products.UpdateProduct(product)
}
