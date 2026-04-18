package main
//import "fmt"

func GetLastRune(s string) rune {
	r := []rune(s)
	return r[len(r) - 1]
}

/*func main() {

	fmt.Println(GetLastRune("kood"))
}*/