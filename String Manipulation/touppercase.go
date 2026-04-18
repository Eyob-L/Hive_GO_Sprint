package main
//import "fmt"

func ToUpperCase(s string) string {

	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {

		if s[i] >= 'a' && s[i] <= 'z' {
			result[i] = s[i] - 32
		}else {
			result[i] = s[i]
		}
	}
	return string(result)
}

/*func main() {
	fmt.Println(ToUpperCase("Hello! How's your day going?"))
}*/