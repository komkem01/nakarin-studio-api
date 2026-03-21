package entitiesdto

import "nakarin-studio/app/utils/base"

type GenderCreateRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=50"`
	IsActive bool   `json:"is_active"`
}

type GenderUpdateRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=50"`
	IsActive bool   `json:"is_active"`
}

type GenderResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GenderListResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	base.ResponsePaginate
}

type GenderInfoRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}
