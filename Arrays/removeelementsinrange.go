package main
//import "fmt"

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	
	start := from
	end := to

	if from > to {
		start = to
		end = from
	}
	new := []float64{}
	for i := 0; i < len(arr); i++ {
		if i < start || i >= end {
			new = append(new, arr[i])
		}
		
	}
	return new
	
}

/*func main () {
	fmt.Println(RemoveElementsInRange([]float64{10., .8, -.4, 20., 7.7, 3.}, 4, 1))
}*/