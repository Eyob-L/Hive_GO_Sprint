package main
//import "fmt"

func SplitWhitespaces(s string) []string {

	    var output []string
		var tmp []rune

		for j := range s {

			if s[j] == ' ' || s[j] == '\t' || s[j] == '\v' || s[j] == '\n' {
				if len(tmp) > 0 {
					output = append(output, string(tmp))
					tmp = nil
				}
			}else {
			tmp = append(tmp, rune(s[j]))
			}
		}
		if len(tmp) > 0 {
			output = append(output, string(tmp))
		}
		return output
	
}

/*func main () {
	fmt.Println(SplitWhitespaces("\nHello! How have you been? "))
}*/