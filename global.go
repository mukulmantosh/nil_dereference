package main

func foo() *int {
	return nil
}

func exampleOfGlobalAnalysis() {
	x := foo()
	_ = *x
}
