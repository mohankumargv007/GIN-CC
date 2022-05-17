package Models

import "time"

type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (todo *Todo) TableName() string {
	return "todo"
}

type User struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) TableName() string {
	return "user"
}

type Catergory struct {
	ID                uint      `json:"id"`
	CatName           string    `json:"cat_name" validate:"required"`
	CatDesc           string    `json:"cat_desc" validate:"required"`
	IsActive          bool      `json:"is_active"`
	MarkedForDeletion bool      `json:"marked_for_deletion"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (catergory *Catergory) TableName() string {
	return "catergory"
}
