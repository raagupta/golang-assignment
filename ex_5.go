package main

import (
	"fmt"
	"os"
	"strconv"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

//float64 Area(Circle c)
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

//float64 Perimeter(Circle c)
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

//float64 Area(Rectangle r)
func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

//float64 Perimeter(Rectangle r)
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func Measure(s Shape) {
	area := s.Area()
	perimeter := s.Perimeter()

	fmt.Println("Area:", area)
	fmt.Println("Perimeter:", perimeter)
}

func main() {
	var shape = os.Args[1]

	if shape == "circle" {
		var r_ip = os.Args[2]
		r, _ := strconv.ParseFloat(r_ip, 64)

		c := Circle{r}
		Measure(c)

	} else if shape == "rectangle" {
		var l_ip = os.Args[2]
		var b_ip = os.Args[3]

		l, _ := strconv.ParseFloat(l_ip, 64)
		b, _ := strconv.ParseFloat(b_ip, 64)

		rec := Rectangle{l, b}
		Measure(rec)
	}
}
