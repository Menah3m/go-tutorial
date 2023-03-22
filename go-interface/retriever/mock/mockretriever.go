package mock

/*
   @Auth: menah3m
   @Desc:
*/

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	return r.Contents
}
