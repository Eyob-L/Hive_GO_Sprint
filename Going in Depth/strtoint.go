package main
//import "fmt"

func StrToInt(s string) int {

	result := 0
	sign := 1
	i := 0

	if len(s) > 0 {
		if s[0] == '-' {
			sign = -1
			i = 1
		}else if s[0] == '+' {
			i = 1
		}
	}
	
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			result = result*10 + int(s[i] - '0')
		}else {
			return 0
		}
		
	}
	return result * sign
}

/*func main () {
	fmt.Println(StrToInt("10203"))
	fmt.Println(StrToInt("0000000010203"))
	fmt.Println(StrToInt("010 203"))
	fmt.Println(StrToInt("Hello World!"))
	fmt.Println(StrToInt("+10203"))
	fmt.Println(StrToInt("-10203"))
	fmt.Println(StrToInt("++10203"))
	fmt.Println(StrToInt("--10203"))
}*/