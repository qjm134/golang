package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func printJson(w http.ResponseWriter, r *http.Request) {
	var p person

	r.ParseForm()
	//	fmt.Println(r.URL.Path)
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", v)
		//		switch k
		//		case "name":
		//			p.name = v
		//		case "age":
		//			p.age = v
		//		default:
		//			fmt.Println("unkown title")
	}
	//	fmt.Println(r.Form["name"])
	p.Name = r.Form["name"][0]
	p.Age = r.Form["age"][0]

	j, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err")
	}

	fmt.Fprintf(w, string(j))
}

func main() {
	http.HandleFunc("/", printJson)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Listen and serve: ", err)
	}
}
