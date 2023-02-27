package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID                      string    `json:"id,omitempty" gorm:"primarykey"`
	Email                   string    `json:"email" validate:"required"`
	Password                string    `json:"password" validate:"required"`
	FirstName               string    `json:"first_name" validate:"required"`
	LastName                string    `json:"last_name" validate:"required"`
	Phone                   *Phone    `json:"phone,omitempty"`
	Address                 *Address  `json:"address,omitempty"`
	AccountVerified         bool      `json:"account_verified"`
	AgreeTermsAndConditions bool      `json:"agree_terms_and_conditions"`
	ProfileCompleted        bool      `json:"profile_completed"`
	Role                    string    `json:"role,omitempty"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	DeletedAt               time.Time `json:"deleted_at" gorm:"index"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
