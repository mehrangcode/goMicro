package user_model

type User struct {
	Id         string  `json:"id"`
	Name       string  `json:"name,omitempty"`
	Email      string  `json:"email,omitempty"`
	Password   string  `json:"-"`
	Created_at *string `json:"created_at,omitempty"`
	Updated_at *string `json:"updated_at,omitempty"`
}
