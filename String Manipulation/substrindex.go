package main
// import "fmt"

func SubstrIndex(s string, toFind string) int {
	
	if len(toFind) == 0 { //|| len(toFind) > len(s) {
		return 0
	}

	for i := 0; i <= len(s) - len(toFind); i++ {
		
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1

}

// func main () {
// 	fmt.Println(SubstrIndex("How are you?", "o"))
// 	fmt.Println(SubstrIndex("How are you doing?", "ou"))
// 	fmt.Println(SubstrIndex("You can do it!", " od"))
// }