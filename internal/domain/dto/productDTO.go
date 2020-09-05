package dto

import "pantrycontrol-backend/internal/domain/entities"

type ProductDTO struct {
	Name      string `json:"name"`
	Unit      string `json:"unit"`
	Brand     string `json:"brand"`
	Category  string `json:"category"`
}

func (p *ProductDTO) ToProduct() entities.Product {
	return entities.Product{
		Name:     p.Name,
		Unit:     p.Unit,
		Brand:    p.Brand,
		Category: p.Category,
	}
}