package dto

type Forgot struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
