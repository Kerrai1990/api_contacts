package models

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	u "github.com/kerrai1990/api_contacts/utils"
	"golang.org/x/crypto/bcrypt"
)

//Token -
type Token struct {
	UserID uint
	jwt.StandardClaims
}

//Account -
type Account struct {
	ID        uint       `gorm:"primary_key"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Token     string     `json:"token";sql:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at";sql:"index"`
}

//Validate user
func (acc *Account) Validate() (map[string]interface{}, bool) {

	fmt.Printf("Validating new user: %v : %v", acc.Email, acc.Password)

	if !strings.Contains(acc.Email, "@") {
		return u.Message(false, "Email Address is not valid"), false
	}

	if len(acc.Password) < 6 {
		return u.Message(false, "Password is not valid"), false
	}

	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", acc.Email).First(temp).Error

	if err != gorm.ErrRecordNotFound {
		fmt.Println(err)
		return u.Message(false, "Email already in use"), false
	}

	return u.Message(false, "Success"), true
}

//Create -
func (acc *Account) Create() map[string]interface{} {

	if resp, ok := acc.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	acc.Password = string(hashedPassword)

	GetDB().Create(acc)

	if acc.ID <= 0 {
		return u.Message(false, "Connection Error")
	}

	//Generate new JWT for user
	tk := &Token{UserID: acc.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))
	acc.Token = tokenString

	//Delete password
	acc.Password = ""

	response := u.Message(true, "Account Created")
	response["account"] = acc

	return response
}

//Login -
func Login(email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email Not Found")
		}
		return u.Message(false, "Connection Error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid Username/Password")
	}

	// Logged in Successfully
	account.Password = ""

	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))
	account.Token = tokenString

	response := u.Message(true, "Authenticated User")
	response["account"] = account

	return response

}

// GetUser -
func GetUser(u uint) *Account {

	account := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(account)
	if account.Email == "" {
		return nil
	}

	account.Password = ""

	return account
}
