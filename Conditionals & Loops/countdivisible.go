package sprint
//import "fmt"

func CountDivisible(from, to, step, divisor int) int {
	if step <= 0 {
		return 0
	}else if divisor == 0 {
		return 0
	}else {
		 var count int = 0
		for i := from; i < to; i += step {
			if i % divisor == 0 {
				count = count + 1
			}
		}
		return count
	}
}
/*func main(){
	fmt.Println(CountDivisible(5, 17, 2, 3))
}*/
