package main
//import "fmt"

func BalanceOut(arr []bool) []bool {
	tc := 0
	fc := 0
	//output := []bool {}
	for i := 0; i < len(arr); i++ {
		if arr[i] == true {
			tc++
		}else if arr[i] == false {
			fc++
		}
		
		
	}
	d := tc - fc

	if d > 0 {
		for i := 0; i < d; i++ {
			arr = append(arr, false)
		}
		
	}else if d < 0 {
		for i := 0; i < -d; i++ {
			arr = append(arr, true)
		}
	}
	return arr

}

/*func main() {
	fmt.Println(BalanceOut([]bool{true, false, false, false}))
}*/
