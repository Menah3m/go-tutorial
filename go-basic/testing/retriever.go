package testing

/*
   @Auth: menah3m
   @Desc:
*/

type Retriever struct {
}

func (Retriever) Get(url string) string {
	return "test message."
}
