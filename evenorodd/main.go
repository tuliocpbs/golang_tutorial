package main

type intslice [11]int

func main() {
	numbers := newIntSlice()

	numbers.printEvenOrOdd()
}

func newIntSlice() intslice {

	var numbers intslice

	for i := 0; i < 11; i++ {
		numbers[i] = i
	}

	return numbers
}

func (is intslice) printEvenOrOdd() {
	for _, n := range is {
		if n%2 == 0 {
			println(n, "is even")
		} else {
			println(n, "is odd")
		}
	}
}
