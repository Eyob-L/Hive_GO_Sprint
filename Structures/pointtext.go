package main
import "fmt"

// func main() {
// 	 fmt.Println(PointText(Point{X : 3.3, Y : 2.2, Text : "hey"}))
// }

type Point struct {
	X float32
	Y float32
	Text string
}

func PointText(p Point) Point {
	newx := p.X
	newy := p.Y

	newtext := fmt.Sprintf("Text at (%f, %f)", p.X, p.Y)
	return Point {X : newx, Y : newy, Text : newtext}

}
