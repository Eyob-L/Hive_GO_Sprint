package sprint
//import "fmt"

func ReverseAlphabet(step int) string {
	if step <= 0 {
		step = 1
	}
	
	alpha := "abcdefghijklmnopqrstuvwxyz"
	output := ""
	for i := 25; i >= 0; i = i - step {
		output = output + string(alpha[i])
	}
	return output

}
/*func main (){
	fmt.Println(ReverseAlphabet(5))
}*/
