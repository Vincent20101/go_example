package main

import (
	"fmt"
	"log"

	"github.com/vmihailenco/msgpack/v4"
)

type Person struct {
	Name string
}

type Friend1 struct {
	Person
}

type Friend2 struct {
	Person `mapstructure:",squash"`
}

func main() {
	datas := []string{`
    { 
      "type": "friend1",
      "person": {
        "name":"dj"
      }
    }
  `,
		`
    {
      "type": "friend2",
      "name": "dj2"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}
		fmt.Println(data)
		err := msgpack.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		//switch m["type"].(string) {
		//case "friend1":
		//	var f1 Friend1
		//	mapstructure.Decode(m, &f1)
		//	fmt.Printf("%+v\n",f1)
		//	fmt.Println("friend1", f1)
		//
		//case "friend2":
		//	var f2 Friend2
		//	mapstructure.Decode(m, &f2)
		//	fmt.Printf("%+v\n",f2)
		//	fmt.Println("friend2", f2)
		//}
	}
}
