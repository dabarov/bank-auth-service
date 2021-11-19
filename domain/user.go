package domain

type User struct {
	ID        int64  `json:"id"`
	IIN       int64  `json:"iin"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}
