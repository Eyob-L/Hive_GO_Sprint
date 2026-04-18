package main
//import "fmt"

func StrLength(s string) []int {

	r := []rune(s)
	return []int{len(r), len(s)}
}

/*func main() {
	fmt.Println(StrLength("Hello World!"))
}*/