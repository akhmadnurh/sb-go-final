package validators

type UpdateProfileBody struct {
	Name string `json:"name" validate:"required,min=3,max=128"`
}

type UpdateOrganizerProfileBody struct {
	UpdateProfileBody
	CompanyName string `json:"company_name" validate:"required,min=1,max=128"`
	ContactInfo string `json:"contact_info" validate:"required,min=1,max=16"`
}

type UpdatePasswordBody struct {
	Password           string `json:"password" validate:"required,min=8,max=128"`
	NewPassword        string `json:"new_password" validate:"required,min=8,max=128"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required,min=8,max=128"`
}
