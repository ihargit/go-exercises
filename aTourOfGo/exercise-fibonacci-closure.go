import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	call := 0
	prev1 := 0
	prev2 := 1
	return func() int {
		if call == 0 {
			call++
			return 0
		}
		if call == 1 {
			call++
			return 1
		}
		curr := prev1 + prev2
		prev1 = prev2
		prev2 = curr
		return curr
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}