package main
import "fmt"

func Combinations() string {

	output := ""
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			for k := j + 1; k < 10; k++ {
				if i == 7 && j == 8 && k == 9{
					output += fmt.Sprintf("%d%d%d", i, j, k)
					return output
				}
				output += fmt.Sprintf("%d%d%d, ", i, j, k)
			}
		}
	}
	return output
}

/*func main(){
	fmt.Println(Combinations())
}*/
