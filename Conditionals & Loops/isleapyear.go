package sprint
//import "fmt"

func IsLeapYear(year int) bool {
	if year % 100 == 0 && year % 400 != 0 {
		return false
	}else if year % 4 == 0 {
		return true
	}else {
		return false
	}
}
/*func main(){
	fmt.Println(IsLeapYear(2100))
}*/
