package domain

type User struct {
	IIN       uint64 `json:"iin"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}
