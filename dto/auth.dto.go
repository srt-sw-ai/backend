package dto

type SignupDto struct {
	NickName string `json:"nickName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Phone    string `json:"phone"`
	EmergencyPhone string `json:"emergencyPhone"`
	Address  string `json:"address"`
	Allergys string `json:"allergys"`
	UnderlyingDiseases string `json:"underlyingDiseases"`
	Medicines string `json:"medicines"`
	BloodType string `json:"bloodType"`
	Weight string `json:"weight"`
	Height string `json:"height"`
}

type AuthDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditProfileDto struct {
	NickName string `json:"nickName"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Phone    string `json:"phone"`
	EmergencyPhone string `json:"emergencyPhone"`
	Address  string `json:"address"`
	Allergys string `json:"allergys"`
	UnderlyingDiseases string `json:"underlyingDiseases"`
	Medicines string `json:"medicines"`
	BloodType string `json:"bloodType"`
	Weight string `json:"weight"`
	Height string `json:"height"`
}

type ProfileDto struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	NickName string `json:"nickName"`
	ImageUri string `json:"imageUri,omitempty"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Phone    string `json:"phone"`
	EmergencyPhone string `json:"emergencyPhone"`
	Address  string `json:"address"`
	Allergys string `json:"allergys"`
	UnderlyingDiseases string `json:"underlyingDiseases"`
	Medicines string `json:"medicines"`
	BloodType string `json:"bloodType"`
	Weight string `json:"weight"`
	Height string `json:"height"`
}