package main
import "fmt"

func ShiftBy(r rune, step int) rune {
	step = step % 26
	i := int(r - 'a')
	is := i + step
	return 'a' + rune(is%26)

}
func main (){
	fmt.Printf("%c\n", ShiftBy('x', 48))
}
