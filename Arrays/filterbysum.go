package main
//import "fmt"

func FilterBySum(arr [][]int, limit int) [][]int {

	result := make([][]int, 0, len(arr)) 
	for _, sub := range arr {
		sum := 0

		for _, num := range sub {
			
			sum += num
		}
		if sum >=limit {
		result = append(result, sub)
		}
	}
	return result

}


/*func main () {
	fmt.Println(FilterBySum([][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, 9))
}*/