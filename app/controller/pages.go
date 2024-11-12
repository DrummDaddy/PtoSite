package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func StartPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	text := "Приевет инженерм ПТО)!"
	//возвращаем простой текст
	fmt.Fprint(rw, text)
}
