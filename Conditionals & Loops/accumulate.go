package sprint
//import "fmt"

func Accumulate(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum = sum + i
	}
	return sum
}
/*func main (){
	fmt.Print(Accumulate(4))
}*/
