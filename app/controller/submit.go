package controller

import (
	"fmt"
	"net/http"
	"net/smtp"
	"path/filepath"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// Обработчик для отправки заявки
func SubmitForm(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == http.MethodPost {
		// Парсинг формы
		r.ParseForm()
		name := r.FormValue("name")
		email := r.FormValue("email")
		phone := r.FormValue("phone")

		// Отправка email
		err := sendEmail(name, email, phone)
		if err != nil {
			http.Error(rw, "Ошибка отправки заявки", http.StatusInternalServerError)
			return
		}

		// Успешный ответ
		path := filepath.Join("public", "html", "succes.html")
		tmpl, err := template.ParseFiles(path)
		if err != nil {
			http.Error(rw, err.Error(), 400)
			return
		}

		err = tmpl.Execute(rw, nil)
		if err != nil {
			http.Error(rw, err.Error(), 400)
			return
		}

	} else {
		http.Error(rw, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

// Функция для отправки email
func sendEmail(name, email, phone string) error {
	from := "ваш-системный-email@mail.ru" // Ваша системная почта
	to := "alexxxa2005@mail.ru"           // email для получения сообщений

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: Новая заявка\r\n\r\nИмя: %s\nEmail: %s\nТелефон: %s", from, to, name, email, phone)

	// Настройка SMTP - сервера
	smtpHost := "smtp.gmail.com" // SMTP сервер
	smtpPort := "465"
	smtpUser := "alexpospelov87@gmail.com"
	smtpPassword := "XZ1987cdM"

	// Настройка аутентификации
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		fmt.Println("Error sending email:", err)
		return err
	}
	return err
}
