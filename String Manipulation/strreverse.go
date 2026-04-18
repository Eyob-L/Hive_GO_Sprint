package main
//import "fmt"

func StrReverse(s string) string {

	result := ""

	for i := len(s) - 1; i >= 0; i-- {
		result += string(rune(int32(s[i])))
	}
	return result

}

/*func main() {
	fmt.Println(StrReverse("Hello Coder!"))
}*/