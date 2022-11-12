package models

type VerifyMail struct {
	Email       string `json:"email"`
	UserName    string `json:"userName"`
	VerifyToken string `json:"verifyToken"`
}
