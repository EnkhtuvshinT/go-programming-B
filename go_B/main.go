package main

import (
	"fmt"
	"time"

)

func main() {
	HotSpurs := "hm ? j madi"
	replacePlayer := strings.NewReplacer("?", "son")
	player := replacePlayer.Replace(HotSpurs)
	fmt.Println(player)
}