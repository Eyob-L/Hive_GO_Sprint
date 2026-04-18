package sprint
//import "fmt"

func AlphabetMastery(n int) string {

	str := "abcdefghijklmnopqrstuvwxyz"
	str2 := ""
	for i := 0; i < n; i++ {
		str2 = str2 + string(str[i])
	}
	return str2
}
/*func main() {
	fmt.Println(AlphabetMastery(6))
}*/
