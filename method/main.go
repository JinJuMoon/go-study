package main

import "fmt"

type Jam interface {
	GetOneSpoon() string
}

type Bread struct {
	val string
}

type StrawberryJam struct {
	val string
}

type OrangeJam struct {
	val string
}

func (b *Bread) PutJam(jam Jam) {
	b.val += " + " + jam.GetOneSpoon()
}

func (b *Bread) String() string {
	return b.val
}

func (sjam *StrawberryJam) GetOneSpoon() string {
	return "strawberryJam"
}

func (ojam *OrangeJam) GetOneSpoon() string {
	return "orangeJam"
}

func main() {
	bread := &Bread{val: "bread"}
	jam := &StrawberryJam{}

	bread.PutJam(jam)

	fmt.Println(bread.String())
}
