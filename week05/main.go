package main

import (
	"fmt" // c language #include <stdio.h>
	"math"
	"reflect"
	"strings"
)

func main() {
	//var i int
	//i = 55

	//var i int = 55
	var f float64 = 1.92

	//f := 1.92
	i := "55"
	fmt.Println(strings.Title("kim inha"))
	fmt.Printf("%f %f\n", f, math.Ceil(f))
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
	fmt.Println("i is", f)
	fmt.Println("i is", i)
	fmt.Print("i is ", i, "\n")
	fmt.Println("i is", i)
	//fmt.Printf("i is %d\n", i) // 숫자로 출력하려면 i := "55"를 55로 수정
	fmt.Printf("i is %s\n", i) //숫자가 아닌 문자로 "55" 출력
}
