package sprint
//import "fmt"

func BetweenLimits(from, to rune) string {
	start := from
	end := to
	if start > end {
		start = to
		end = from
	}
	var output string
	for  r := start + 1; r < end; r++{
		output = output + string(r)
	}
	return output
}
/*func main (){
	fmt.Println(BetweenLimits('j', 'f'))
}*/
