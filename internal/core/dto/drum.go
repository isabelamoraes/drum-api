package dto

import "drum-api/internal/infra/gorm/model"

type UpsertDrumRequest struct {
	Name          string `json:"name" binding:"required"`
	Mark          string `json:"mark" binding:"required"`
	Configuration string `json:"configuration" binding:"required"`
	IsEletronic   *bool  `json:"isEletronic" binding:"required"`
}

type DrumDTO struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Mark          string `json:"mark"`
	Configuration string `json:"configuration"`
	IsEletronic   bool   `json:"isEletronic"`
}

type DrumReviewDTO struct {
	Rating  int    `json:"id"`
	Comment string `json:"comment,omitempty"`
}

type DrumWithReviewDTO struct {
	DrumDTO
	Reviews []DrumReviewDTO `json:"reviews,omitempty"`
}

func ToDrumDTO(d *model.Drum) DrumDTO {
	drum := DrumDTO{
		ID:            int(d.ID),
		Name:          d.Name,
		Mark:          d.Mark,
		Configuration: d.Configuration,
		IsEletronic:   d.IsEletronic,
	}

	return drum
}

func ToDrumReviewDTO(r *model.Review) DrumReviewDTO {
	review := DrumReviewDTO{
		Rating:  r.Rating,
		Comment: r.Comment,
	}

	return review
}

func ToDrumWithReviewDTO(d *model.Drum) DrumWithReviewDTO {
	reviews := make([]DrumReviewDTO, len(d.Reviews))

	for i := range d.Reviews {
		reviews[i] = ToDrumReviewDTO(&d.Reviews[i])
	}

	drum := DrumDTO{
		ID:            int(d.ID),
		Name:          d.Name,
		Mark:          d.Mark,
		Configuration: d.Configuration,
		IsEletronic:   d.IsEletronic,
	}

	return DrumWithReviewDTO{
		drum,
		reviews,
	}
}
