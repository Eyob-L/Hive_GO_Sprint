package main
//import "fmt"

func StrConcatWith(strs []string, sep string) string {

	output := ""

	for i, str := range(strs) {
		if i > 0 {
			output += sep
		}
		output += str
	}
	return output
	
}

// func main () {
// 	toConcat := []string{"Three", " Two", " One", " Go!"}
// 	fmt.Println(StrConcatWith(toConcat, "."))
// }