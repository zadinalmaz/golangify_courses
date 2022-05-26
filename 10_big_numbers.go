package main
 
import (
    "fmt"
)
 
func main() {
    const distance = 236000000000000000
    const lightSpeed = 299792
    const secondsPerDay = 86400
    const daysPerYear = 365
 
    const years = distance / lightSpeed / secondsPerDay / daysPerYear
 
    fmt.Println("Расстояние в световых годах до Карликовой галактики в Большом Псе составляет:", years) // Выводит: Расстояние в световых годах до Карликовой галактики в Большом Псе составляет: 24962
}