package accountuntils

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendPasswordResetEmail(email string, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "your-email@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset Request")
	resetLink := fmt.Sprintf("https://localhost:8080/reset-password?token=%s", token)
	m.SetBody("text/html", fmt.Sprintf("Click <a href=\"%s\">here</a> to reset your password.", resetLink))

	d := gomail.NewDialer("smtp.gmail.com", 587, "tranvietanh1803@gmail.com", "teln hzlw dtgc bvyv")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // Chỉ sử dụng cho thử nghiệm

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
