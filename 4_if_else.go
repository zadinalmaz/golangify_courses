package main
 
import (
	"fmt"
)
 
func main() {
	fmt.Println("На дворе 2100 год. Он високосный?")
 
	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
 
	if leap {
    	    fmt.Println("Этот год високосный!")
	} else {
    	    fmt.Println("К сожалению, нет. Этот год не високосный.")
	}
}