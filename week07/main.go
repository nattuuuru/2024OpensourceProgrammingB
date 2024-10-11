package main

import (
	"fmt"
	"strings"
)

func main() {
	var army string = "해군!"
	armyFixed := strings.NewReplacer("!", "국")
	fmt.Println(army)
	fmt.Println(armyFixed.Replace(army))
}
