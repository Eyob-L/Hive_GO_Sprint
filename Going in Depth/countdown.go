package main
//import "fmt"

func Countdown(n int) string {
	str := ""
	for r := n; r > 0; r = r - 2 {
		
		str += string(rune(int32(r + 48))) + ", "
	}
	return str + "0!"
}
/*func main(){
	fmt.Println(Countdown(7))
}*/
