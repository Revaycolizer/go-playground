package main
import (
    "fmt"
    "maps"
)

func main(){
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v3 := m["k3"]
    fmt.Println("v3:", v3)	

    fmt.Println("len:", len(m))

	delete(m, "k2")
    fmt.Println("map:", m)	

    clear(m)
    fmt.Println("map:", m)
}