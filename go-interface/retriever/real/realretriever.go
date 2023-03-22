package real

import (
	"io/ioutil"
	"net/http"
)

/*
   @Auth: menah3m
   @Desc:
*/

type Retriever struct {
	Url string
}

func (r Retriever) Get(url string) string {
	if r.Url != "" {
		url = r.Url
	}
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}
