package domain

import (
	"api/first-go/common"
	"api/first-go/util"
	"time"

	"errors"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       	string 		`db:"id"`
	Username 	string		`db:"username"`
	Email    	string		`db:"email"`
	Password 	string		`db:"password"`
	CreatedAt 	time.Time 	`db:"created_at"`
	UpdatedAt 	time.Time 	`db:"updated_at"`
}

func NewUser(r RegisterUserSchema) (*User, error) {
	if r.Email == "" || r.Password == "" || r.Username == "" {
		return nil, errors.New("input cannot be null")
	}

	return &User{
		Id: uuid.NewString(), 
		Email: r.Email, 
		Username: r.Username, 
		Password: r.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		}, nil
}

func (u *User) Validate() (err error) {
	if err = u.ValidateEmail(); err != nil {
		return
	}

	if err = u.ValidatePassword(); err != nil {
		return
	}

	if err = u.ValidateUsername(); err != nil {
		return
	}

	return
}

func (u *User) ValidateEmail() (err error) {
	if u.Email == "" {
		return common.ErrEmailRequired
	}

	emails := strings.Split(u.Email, "@")
	if len(emails) != 2 {
		return common.ErrEmailInvalid
	}
	return
}

func (u *User) ValidatePassword() (err error) {
	if u.Password == "" {
		return common.ErrPasswordRequired
	}

	if len(u.Password) < 6 {
		return common.ErrPasswordInvalidLength
	}
	return
}

func (u *User) ValidateUsername() (err error) {
	if u.Username == "" {
		return common.ErrUsernameRequired
	}

	if len(u.Username) < 6 {
		return common.ErrUsernameInvalidLength
	}
	return
}

func (u *User) EncryptPassword(salt int) (err error) {

	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), salt)
	if err != nil {
		return
	}
	u.Password = string(encryptedPass)
	return nil
}

func (u User) IsExists() bool {
	
	return u.Id != ""
}

func (u User) ComparePassword(plain string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
}

func (u User) GenerateToken(exp float64) (tokenString string, err error) {
	return util.GenerateToken(u.Id, string(u.Email), exp)
}