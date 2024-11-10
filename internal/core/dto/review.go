package dto

type CreateReviewDTO struct {
	Rating  int
	Comment string
}

type CreateReviewRequest struct {
	Rating  int    `json:"rating" binding:"required"`
	Comment string `json:"comment"`
}
