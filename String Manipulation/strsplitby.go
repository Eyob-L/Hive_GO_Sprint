package main
//import "fmt"

func StrSplitBy(s, sep string) []string {

	    var output []string
		var tmp string

		for j := 0; j < len(s); j++ {

			if  j <= len(s) - len(sep) && s[j:j+len(sep)] == sep {
				output = append(output, tmp)
				tmp = ""
				j += (len(sep) - 1)
			}else {
			 tmp += string(s[j])
			}
		}
		if len(sep) > 0 {
			output = append(output, tmp)
		}
		if len(s) == 0 {
			return nil
		}
		return output
}

// func main () {
// 	fmt.Println(StrSplitBy("HowYOUhaveYOUyouYOUbeen?", "YOU"))
// }