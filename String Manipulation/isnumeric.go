package main
//import "fmt"

func IsNumeric(s string) bool {

	for i := 0; i < len(s); i++ {

		if s[i] < '0' || s[i] >'9' {
			return false
		}
	}
	return true
	
}

/*func main() {
	fmt.Println(IsNumeric("012345"))
	fmt.Println(IsNumeric("01+23-45"))

}*/