package validators

type RegisterCustomerBody struct {
	Name     string `json:"name" validate:"required,min=3,max=128"`
	Email    string `json:"email" validate:"required,email,max=128"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}
