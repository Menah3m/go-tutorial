package main

import (
	"encoding/json"
	"fmt"
)

/*
@Auth: menah3m
@Desc:
*/
type Order struct {
	ID         string      `json:"id"`
	Items      []OrderItem `json:"items"`
	Quantity   int         `json:"quantity"`
	TotalPrice float64     `json:"total_price"`
}

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price"`
}

func marshal() {
	o := Order{
		ID:         "1234",
		Quantity:   3,
		TotalPrice: 30,
		Items: []OrderItem{{
			ID:    "1",
			Name:  "book",
			Price: 15,
		},
			{
				ID:    "2",
				Name:  "pen",
				Price: 5,
			},
		},
	}

	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)
}

func unmarshal() {
	s := `{"id":"1234","items":[{"id":"1","name":"book","price":15},{"id":"2","name":"pen","price":5}],"quantity":3,"total_price":30}
`
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}

func parseNLP() {
	res := `{
  "result": [
    {
      "synonym": "",
      "weight": "0.100000",
      "tag": "普通词",
      "word": "请"
    },
    {
      "synonym": "",
      "weight": "0.100000",
      "tag": "普通词",
      "word": "输入"
    },
    {
      "synonym": "",
      "weight": "1.000000",
      "tag": "品类",
      "word": "文本"
    }
  ],
  "success": true
}`
	// m := make(map[string]interface{})
	m := struct {
		Result []struct {
			Synonym string `json:"synonym"`
			Tag     string `json:"tag"`
			Word    string `json:"word"`
		} `json:"result"`
		Success bool `json:"success"`
	}{}
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}
	// 如果使用map来装数据，则取想要的值时需要判断类型，比较复杂，
	// 推荐使用定义 匿名结构体方式来装数据
	fmt.Printf("%+v", m.Result)
}

func main() {
	// marshal()
	// unmarshal()
	// parseNLP()
	ss := string(1)
	s := fmt.Sprintf("%d", 100)
	fmt.Println(s)
}
