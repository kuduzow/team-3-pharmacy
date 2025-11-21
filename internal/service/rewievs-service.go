package service

import (
	"errors"
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/repository"
)

// /
type ReviewService interface {
	CreateReview(req models.CreateReviewRequest) (*models.Review, error)
	UpdateReview(id uint,req models.UpdateReviewRequest) (*models.Review, error)
	DeleteReview(id uint) error
	ListPharmacyReview(pharmacyID uint) ([]models.Review, error)
}

type reviewService struct {
	reviews  repository.ReviewRepository
	pharmacy repository.PharmacyRepository
}

func NewReviewService(
	reviews repository.ReviewRepository,
	pharmacy repository.PharmacyRepository,
) ReviewService {
	return &reviewService{
		reviews:  reviews,
		pharmacy: pharmacy,
	}
}

func (r *reviewService) CreateReview(req models.CreateReviewRequest) (*models.Review, error) {
	if err := r.validateReview(req); err != nil {
		return nil, err
	}

	review := &models.Review{
		UserID:     req.UserID,
		MedicineID: req.MedicineID,
		Rating:     req.Rating,
		Text:       req.Text,
	}
	if err := r.reviews.Create(review); err != nil {
		return nil, err
	}

	return review, nil
}

func (r *reviewService) validateReview(req models.CreateReviewRequest) error {
	if req.UserID == 0 {
		return errors.New("поле user_id должно быть больше 0")
	}
	if req.MedicineID == 0 {
		return errors.New("поле medicine_id должно быть больше 0")
	}
	if req.Rating < 1 && req.Rating > 5 {
		return errors.New("Нет такой оценки")
	}
	if req.Text == "" {
		return errors.New("text не может быть пустым")
	}
	return nil
}

func (r *reviewService) UpdateReview(id uint, req models.UpdateReviewRequest) (*models.Review, error) {
	review, err := r.reviews.GetByID(id)
	if err != nil {
		return nil, errors.New("отзыв не найден")
	}

	// Обновляем только те поля, которые пришли
	if req.Rating != nil {
		review.Rating = *req.Rating
	}

	if req.Text != nil {
		review.Text = *req.Text
	}

	if err := r.reviews.Update(review); err != nil {
		return nil, err
	}

	return review, nil
}

func (r *reviewService) DeleteReview(id uint) error {
	if _, err := r.reviews.GetByID(id); err != nil {
		return err
	}

	return r.reviews.Delete(id)
}

func (r *reviewService) ListPharmacyReview(pharmacyID uint) ([]models.Review, error) {
	if _, err := r.reviews.GetByID(pharmacyID); err != nil {
		return nil, err
	}

	return r.reviews.GetReviewsByPharmacyID(pharmacyID)
}
