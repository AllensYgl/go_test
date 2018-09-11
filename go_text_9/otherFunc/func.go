package otherFunc

import (
	"fmt"
)

func Prints(ex *Example) {
	fmt.Println(ex.Value)
}

func (ex *Example) P() {
	fmt.Print(" ", ex.Value)
}
