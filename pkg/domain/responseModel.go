package domain

type UserResponse struct {
	UserId       uint   `json:"userid"`
	UserName     string `json:"username"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Verification bool   `json:"verification" `
	Vip          bool   `json:"vip" `
	PhoneNumber  string `json:"phonenumber"`
	Profile      string `json:"profile"`
	Token        string `json:"token"`
}

type AdminResponse struct {
	AdminId      uint   `json:"adminid"`
	AdminName    string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Verification bool   `json:"verification"`
	PhoneNumber  string `json:"phonenumber"`
	Token        string `json:"token"`
}
