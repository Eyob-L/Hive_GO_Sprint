package main
import "fmt"

func GenerateRange(min, max int) []int {
	
	if min > max {
		return nil
	}

	len := max - min
	output := make([]int, len)

	for i := 0; i < len; i++ {
		output[i] = min + i
	}
	return output
}


func main (){
	fmt.Println(GenerateRange(6, 4))
}
