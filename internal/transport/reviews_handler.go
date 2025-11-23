package transport

import (
	"pharmacy-team/internal/models"
	"pharmacy-team/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewsHandler struct {
	reviews service.ReviewService
}

func NewReviewsHandler(reviews service.ReviewService) *ReviewsHandler {
	return &ReviewsHandler{reviews: reviews}
}

func (p *ReviewsHandler) Routes(r *gin.Engine) {
	r.POST("/pharmacy/:id/reviews", p.Create)
	r.GET("/pharmacy/:id/reviews", p.GetListReviews)
	r.PATCH("/reviews/:id", p.Update)
	r.DELETE("/reviews/:id", p.Delete)
}

func (p *ReviewsHandler) Create(c *gin.Context) {

	medicineIDParam := c.Param("id")
	medicineIDS, err := strconv.ParseUint(medicineIDParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid medicine id"})
		return
	}
	medicineID := uint(medicineIDS)

	
	var req models.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}


	req.MedicineID = medicineID
	review, err := p.reviews.CreateReview(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, review)
}

func (r *ReviewsHandler) GetListReviews (c *gin.Context) {
	idstr := c.Param("id")

	pharmacyID, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "неправильный id"})
		return
	}

	reviews, err := r.reviews.ListPharmacyReview (uint(pharmacyID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, reviews)
}

func (r *ReviewsHandler) Update(c *gin.Context) {
	idstr := c.Param("id")
	reviewID, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "неправильный id"})
		return
	}
	var req models.UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	review, err := r.reviews.UpdateReview(uint(reviewID), req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, review)
}
func (r *ReviewsHandler) Delete(c *gin.Context) {
	idstr := c.Param("id")
	reviewID, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "неправильный id"})
		return
	}
	if err := r.reviews.DeleteReview(uint(reviewID)); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}	
	c.JSON(200, gin.H{"message": "отзыв успешно удален"})
}


