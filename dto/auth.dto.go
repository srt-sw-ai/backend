package dto

type SignupDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
}

type AuthDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditProfileDto struct {
	NickName string `json:"nickName"`
	Email    string `json:"email"`
}

type ProfileDto struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	NickName string `json:"nickName"`
	ImageUri string `json:"imageUri,omitempty"`
}