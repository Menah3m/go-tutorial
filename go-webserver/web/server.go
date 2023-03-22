package main

import (
	"github.com/menah3m/go-tutorial/go-webserver/filelisting"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

/*
   @Auth: menah3m
   @Desc:
*/

type userError interface {
	error
	Message() string
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// panic error
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}

		}()
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error occured:%s", err.Error())

			// user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
				http.Error(writer, http.StatusText(code), code)
			case os.IsPermission(err):
				code = http.StatusForbidden
				http.Error(writer, http.StatusText(code), code)
			default:
				code = http.StatusInternalServerError
				http.Error(writer, http.StatusText(code), code)

			}
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandlerFilelist))

	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		panic(err)
	}
}
