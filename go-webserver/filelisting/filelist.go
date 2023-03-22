package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

/*
   @Auth: menah3m
   @Desc:
*/

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandlerFilelist(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {

		return userError("path must start with " + prefix)
	}
	path := request.URL.Path[len("/list/"):] // 获取path
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
