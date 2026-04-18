package main
//import "fmt"

func SortIntegerTable(table []int) []int {

	//result := make([]int, 0, len(table))

	l := len(table)

	for i := 0; i < l-1; i++ {
		for j := 0; j < l - i - 1; j++ {
			if table[j] > table[j + 1] {
				table[j], table[j + 1] = table[j + 1], table[j]
			}
		} 
	}
	return table
} 

/*func main () {
	fmt.Println(SortIntegerTable([]int{2, 0, 5, 4, 1, 3}))
}*/
