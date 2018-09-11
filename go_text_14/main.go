package main

import (
	"fmt"
)

type Imp struct {
	Value int
}

//  change tostring() method
func (i Imp) String() string {
	return fmt.Sprintf(" stringrt()= %d  ", i.Value)
}

func main() {
	//fmt.Println()
	i := Imp{Value: 3}
	fmt.Println(i)

	fmt.Print(fmt.Sprintf(" stringrt()= %d", i.Value))
}
