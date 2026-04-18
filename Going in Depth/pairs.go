package main
import "fmt"

func Pairs() string {

	output := ""
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			if i == 98 && j == 99 {
			output += fmt.Sprintf("%02d %02d", i, j)
			return output
			}
			
			output += fmt.Sprintf("%02d %02d, ", i, j)
		}
	}
	return output
}

/*func main(){
	fmt.Println(Pairs())
}*/
