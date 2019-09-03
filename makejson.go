package main

import "fmt"
import "encoding/json"

func main() {

	var Amap = make(map[string]string)
	var name string
	var addr string
	fmt.Scanln(&name)
	fmt.Scanln(&addr)
	fmt.Println(name, addr)
	Amap["name"] = name
	Amap["address"] = addr
	d, _ := json.Marshal(Amap)
	fmt.Println(string(d))
}
