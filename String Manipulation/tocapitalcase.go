package main
//import "fmt"

func ToCapitalCase(s string) string {

	r := make([]byte, len(s))

	if  len(s) > 0 && s[0] >= 'a' && s[0] <= 'z' {
		r[0] = s[0] - 32
	}else if len(s) > 0 {
		r[0] = s[0]
	}

	for i := 1; i < len(s); i++ {

		r[i] = s[i]
		if s[i] >= 'A' && s[i] <= 'Z' {
			r[i] = s[i] + 32
		}
	}

	for i := 1; i < len(r); i++ {

		//result[i] = s[i]

		if !(r[i-1] >= 'A' && r[i-1] <= 'Z') && !(r[i-1] >= 'a' && r[i-1] <= 'z') && !(r[i-1] >= '0' && r[i-1] <= '9') && r[i] >= 'a' && r[i] <= 'z' {
			r[i] -= 32
		}


	}
	return string(r)
}

/*func main() {
	fmt.Println(ToCapitalCase("Hello! Great to sEe you! How-are-you-doing-2day?"))
}*/