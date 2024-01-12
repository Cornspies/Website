package user

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type User struct {
	UserID         string
	IsLoggedIn     bool
	HasProfileIcon bool
}

type Account struct {
	UserID         string
	Username       string
	Email          string
	Password       string
	HasProfileIcon bool
}

var Accounts []Account

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits
)

func GetUserData(r *http.Request) User {
	cookie, err := r.Cookie("SESSION")
	if err == nil {
		value := cookie.Value
		if strings.HasPrefix(value, "UserID:") {
			value = strings.TrimPrefix(value, "UserID:")
			for _, account := range Accounts {
				if account.UserID == value {
					return User{
						UserID:         value,
						IsLoggedIn:     true,
						HasProfileIcon: account.HasProfileIcon,
					}
				}
			}
		}
	}
	return User{
		UserID:         "",
		IsLoggedIn:     false,
		HasProfileIcon: false,
	}
}

func login(w http.ResponseWriter, username string, password string) error {
	if username == "" || password == "" {
		return errors.New("required field empty")
	}
	for _, account := range Accounts {
		if account.Username == username {
			if account.Password == password {
				log.Println("User " + username + " logged in")
				cookie := http.Cookie{
					Name:  "SESSION",
					Value: "UserID:" + account.UserID,
				}
				http.SetCookie(w, &cookie)
				return nil
			}
		}
	}
	return errors.New("username or password incorrect")
}

func register(username string, email string, password string, repeatPassword string) error {
	if username == "" || email == "" || password == "" || repeatPassword == "" {
		return errors.New("required field empty")
	}
	if password != repeatPassword {
		return errors.New("passwords do not match")
	}
	for _, account := range Accounts {
		if account.Username == username {
			return errors.New("username taken")
		}
		if account.Email == email {
			return errors.New("email already in use")
		}
	}
	var account = Account{
		UserID:         createUniqueUserID(),
		Username:       username,
		Email:          email,
		Password:       password,
		HasProfileIcon: false,
	}
	log.Println("New Account created: " + account.Username + " " + account.Email + " " + account.UserID)
	Accounts = append(Accounts, account)
	return nil
}

// https://stackoverflow.com/a/31832326
func createUniqueUserID() string {
	var src = rand.NewSource(time.Now().UnixNano())
	for {
		sb := strings.Builder{}
		sb.Grow(63)
		// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
		for i, cache, remain := 62, src.Int63(), letterIdxMax; i >= 0; {
			if remain == 0 {
				cache, remain = src.Int63(), letterIdxMax
			}
			if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
				sb.WriteByte(letterBytes[idx])
				i--
			}
			cache >>= letterIdxBits
			remain--
		}
		var userID = sb.String()
		for _, account := range Accounts {
			if account.UserID == userID {
				continue
			}
		}
		return userID
	}
}
