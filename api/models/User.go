package models

import (
	"cloudgobackend/api/security"
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:20;not null;unique" json:"name"`
	Email     string    `gorm:"size:50;not null;unique" json:"email"`
	Password  string    `gorm:"size:60;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	Posts     []Post    `gorm:"foreignkey:AuthorID" json:"posts"`
}

func (User) TableName() string {
	return "public.USERS"
}

func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Password required")
		}
		if u.Email == "" {
			return errors.New("Required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Password required")
		}
		if u.Email == "" {
			return errors.New("Required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Password required")
		}
		if u.Email == "" {
			return errors.New("Required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		return nil
	}
}
