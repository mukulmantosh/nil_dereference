package main

type B struct {
	id int
}

type A struct {
	b *B
}

func main() {
	var x = &A{}
	_ = x.b.id
}
