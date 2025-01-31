package validators

type RegisterCustomerBody struct {
	Name     string `json:"name" validate:"required,min=3,max=128"`
	Email    string `json:"email" validate:"required,email,max=128"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}

type RegisterOrganizerBody struct {
	RegisterCustomerBody
	CompanyName string `json:"company_name" validate:"required,min=1,max=128"`
	ContactInfo string `json:"contact_info" validate:"required,min=1,max=16"`
}
