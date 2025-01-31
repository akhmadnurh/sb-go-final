package validators

type LoginBody struct {
	Email    string `json:"email" validate:"required,email,max=128"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}

type LoginOrganizerBody struct {
	LoginBody
}
