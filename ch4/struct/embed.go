package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

// 带有匿名成员的结构体
type Circle1 struct {
	Point
	Radius int
}

type Wheel1 struct {
	Circle1
	Spokes int
}

func main() {
	// Access nested struct
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Radius = 8

	// 访问带有匿名成员的结构体
	var w1 Wheel1
	w1.X = 8
	w1.Radius = 5

	var w2 = Wheel1{Circle1{Point{8, 8}, 5}, 20}

	w2 = Wheel1{
		Circle1: Circle1{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w2)

	w2.X = 42
	fmt.Printf("%#v\n", w)
}
