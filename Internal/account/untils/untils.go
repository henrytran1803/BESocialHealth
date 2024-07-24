package accountuntils

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

func SendPasswordResetEmail(email string, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "your-email@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset Request")
	resetLink := fmt.Sprintf("https://localhost:8080/reset-password?token=%s", token)
	body := fmt.Sprintf(
		`Click <a href="%s">here</a> to reset your password.<br><br>
         If you prefer, you can copy and paste this token into the password reset page: <b>%s</b>`,
		resetLink, token,
	)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "tranvietanh1803@gmail.com", "teln hzlw dtgc bvyv")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

var jwtSecret = []byte("your_global_secret_key")

func GenerateJWT(userID int, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"secret":  secret,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomString(length int) string {
	return StringWithCharset(length, charset)
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Secret string `json:"secret"`
	jwt.StandardClaims
}

func VerifyJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
