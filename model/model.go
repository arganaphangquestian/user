package model

// InputUser or Register model
type InputUser struct {
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Referral *string `json:"referral"`
}

// User model
type User struct {
	ID       uint64  `json:"id"`
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Referral *string `json:"referral"`
	Role     string  `json:"role"`
}
