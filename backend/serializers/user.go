package serializers

type UserCreateSerializer struct {
	Username string `json:"username" validate:"required,min=3,max=10"`
	Email    string `json:"email" validate:"required,custom_email"`
	Password string `json:"password" validate:"required,min=8,max=40"`
}
