package controller

import (
	"fmt"
	"net/http"
	"net/smtp"
	"text/template"
)

// Обработчик для отправки заявки
func SubmitForm(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		email := r.FormValue("email")
		phone := r.FormValue("phone")

		// Здесь можно добавить отправку email на указанный адрес
		err := sendEmail(name, email, phone)
		if err != nil {
			http.Error(rw, "Ошибка отправки заявки", http.StatusInternalServerError)
			return
		}

		// Отправка успешного сообщения
		tmpl, _ := template.New("success").Parse("<h1>Заявка отправлена!</h1>")
		tmpl.Execute(rw, nil)
	} else {
		http.Error(rw, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

// Функция для отправки email
func sendEmail(name, email, phone string) error {
	from := "your-email@example.com"  // укажите ваш email
	password := "your-email-password" //укажите ваш пароль
	to := "recipient@example.com"     // почта для приема сообщений

	msg := fmt.Sprintf("Имя: %s\nEmail: %s\nТелефон: %s", name, email, phone)
	subject := "Новая заявка"
	body := "Subject: " + subject + "\n\n" + msg

	//Настройка SMTP - сервера
	smtpHost := "smtp.example.com" // SMTP - сервер
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(body))
	return err
}
