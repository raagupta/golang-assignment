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

func main() {
	var shape = os.Args[1]

	if shape == "circle" {
		var r_ip = os.Args[2]
		r, _ := strconv.ParseFloat(r_ip, 64)
		var c = Circle{r}

		fmt.Println("Area:", c.Area())
		fmt.Println("Perimeter:", c.Perimeter())

	} else if shape == "rectangle" {
		var l_ip = os.Args[2]
		var b_ip = os.Args[3]

		l, _ := strconv.ParseFloat(l_ip, 64)
		b, _ := strconv.ParseFloat(b_ip, 64)

		var rec = Rectangle{l, b}

		fmt.Println("Area:", rec.Area())
		fmt.Println("Perimeter:", rec.Perimeter())

	}
}
