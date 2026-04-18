package main
import "fmt"


type Point struct {
	X float32
	Y float32
	Text string
}

func MakePoint(x, y float32, text string) Point {

	return Point {X : x, Y : y, Text :text}
}

func main () {
	fmt.Println(MakePoint(2.1, 2.0, "grade"))
}