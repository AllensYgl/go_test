package main

import "go_text/go_text_9/otherFunc"

//extend the function of the type Example
type examples struct {
	exam  *otherFunc.Example
	value int
	//....
}

// start extend
func (ex *examples) printss() {
	if ex == nil || ex.exam == nil {
		return
	}
	right := examples{ex.exam.Right, 0}
	left := examples{ex.exam.Left, 0}
	right.printss()
	left.printss()
	ex.exam.P()
}

func main() {
	text := otherFunc.Example{Value: 2}
	otherFunc.Prints(&text)

	roots := otherFunc.Example{Value: 3}
	roots.Left = &otherFunc.Example{Value: 5}
	roots.Right = &otherFunc.Example{Value: 4}
	roots.Left.Right = &otherFunc.Example{Value: 6}
	roots.Right.Left = &otherFunc.Example{Value: 7}

	exa := examples{&roots, 0}

	exa.printss()
}
