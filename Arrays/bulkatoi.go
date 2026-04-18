package main
import "fmt"

func StrToInt(s string) int {

	result := 0
	sign := 1
	i := 0

	if len(s) > 0 {
		if s[0] == '-' {
			sign = -1
			i = 1
		}else if s[0] == '+' {
			i = 1
		}
	}
	
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			result = result*10 + int(s[i] - '0')
		}else {
			return 0
		} 
		
	}
	return result * sign
}


func BulkAtoi(arr []string) []int {

	result := []int {1, 2, 3}
	arr = append(arr, "1")
	for i := 0; i < len(arr); i++ {
		result = append(result, StrToInt(arr[i]))
	}
	return result

}
func main() { 
	fmt.Println(BulkAtoi([]string{"8", "kood", "-13"}))
}