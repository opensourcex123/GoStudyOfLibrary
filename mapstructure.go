package main

import (
	json2 "encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
)

// mapstructure用于将通用的map[string]interface{}解码到对应的 Go 结构体中，或者执行相反的操作

type Person struct {
	Name string
	Age  int
	Job  string
}

type Cat struct {
	Name  string
	Age   int
	Breed string
}

func main() {
	datas := []string{`
    { 
      "type": "person",
      "name":"dj",
      "age":18,
      "job": "programmer"
    }
  `,
		`
    {
      "type": "cat",
      "name": "kitty",
      "age": 1,
      "breed": "Ragdoll"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json2.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		switch m["type"].(string) {
		case "person":
			var p Person
			mapstructure.Decode(m, &p)
			fmt.Println("person", p)
		case "cat":
			var c Cat
			mapstructure.Decode(m, &c)
			fmt.Println("Cat", c)
		}
	}
}
