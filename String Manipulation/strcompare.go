package main
// import "fmt"

func StrCompare(a, b string) int {
	min := len(a)
	if len(a) > len(b) {
		min = len(b)
	}

	for i := 0; i < min; i++ {
		if a[i] > b[i] {
			return 1
		}else if a[i] < b[i] {
			return -1
		}
	}

	if len(a) > len(b) {
		return 1
	}else if len(a) < len(b) {
		return -1
	}
	return 0
}

// func main () {
// 	fmt.Println(StrCompare("Hi!", "Hi!"))
// 	fmt.Println(StrCompare("Day", "ay"))
// 	fmt.Println(StrCompare("weekday", "week"))
// }