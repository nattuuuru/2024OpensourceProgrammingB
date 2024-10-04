package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13
	var f float64 = 12.9 //f := 12.9

	fmt.Printf("value i :%d, value f :%f\n", i, f)
	//fmt.Print("%d * %f = %f\n", i, f, i*f) //invalid operation: i * f (mismatched types int and float64)
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f) //conversion (casting)
	//fmt.Printf("%d * %f = %f\n", i, f, i*int(f)) //conversion (casting)
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
}
