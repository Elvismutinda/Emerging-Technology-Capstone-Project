package categories

type CreateCategoryRequest struct {
	Name         string `json:"name" validate:"required"`
	CategoryType string `json:"category_type" validate:"required,oneof=INCOME EXPENSE"`
}

type UpdateCategoryRequest struct {
	Name         string `json:"name" validate:"required_without=CategoryType"`
	CategoryType string `json:"category_type" validate:"required_without=Name,omitempty,oneof=INCOME EXPENSE"`
}
