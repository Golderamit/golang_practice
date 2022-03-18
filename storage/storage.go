package storage

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	ID        int32     `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	IsActive  bool      `db:"is_active"`
	IsAdmin   bool      `db:"is_admin"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type AdminHomeDB struct {
	ID               int32     `db:"id"`
	Title            string    `db:"title"`
	Venue            string    `db:"venue"`
	Address          string    `db:"address"`
	Country          string    `db:"country"`
	Email            string    `db:"email"`
	PhoneNumber      string    `db:"phone_number"`
	ShortDescription string    `db:"short_description"`
	Description      string    `db:"description"`
	Image            string    `db:"image"`
	Status           bool      `db:"status"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
	FromDate         time.Time `db:"from_date"`
	ToDate           time.Time `db:"to_date"`
}

func (ug User) ValidateUser() error {
	return validation.ValidateStruct(&ug,
		validation.Field(&ug.Email,
			validation.Required.Error("email is required"),
			is.Email,
		),
		validation.Field(&ug.Password,
			validation.Required.Error("Password is required"),
			validation.Length(3, 10).Error("Password Lenght must be 3 to 10"),
		),
	)
}
