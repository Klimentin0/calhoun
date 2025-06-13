package controllers

import (
	"calhoun/views"
	"html/template"
	"net/http"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func Docs(tpl views.Template) http.HandlerFunc {
	topics := []struct {
		Topic        string
		Instructions template.HTML
	}{
		{
			Topic:        "Как получить пароль?",
			Instructions: "Перейти по ссылке",
		},
		{
			Topic:        "Как подкоючиться с нового устройства?",
			Instructions: "Использовать ...",
		},
		{
			Topic:        "Как получить пароль?",
			Instructions: "Перейти по ссылке",
		},
		{
			Topic:        "Как получить пароль?",
			Instructions: "Перейти по ссылке",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, topics)
	}
}
