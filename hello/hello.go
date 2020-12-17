package main

import "fmt"

type Bread struct {
	val string
}

type StrawberryJam struct {
	open bool
}

type SpoonOfStrawberryJam struct {
}

type Sandwitch struct {
	val string
}

func GetBreads(n int) []*Bread {
	breads := make([]*Bread, n)
	breads[0] = &Bread{val: "breadLower"}
	breads[1] = &Bread{val: "breadUpper"}
	return breads
}

func OpenStrawberryJam(jam *StrawberryJam) {
	jam.open = true
}

func GetOneSpoon(_ *StrawberryJam) *SpoonOfStrawberryJam {
	return &SpoonOfStrawberryJam{}
}

func PutJamOnBread(b *Bread, spoon *SpoonOfStrawberryJam) {
	b.val += " with jam"
}

func MakeSandwitch(breads []*Bread) *Sandwitch {
	return &Sandwitch{val: (breads[0].val + breads[1].val)}
}

func main() {
	// 1. 빵 두개를 꺼낸다.
	breads := GetBreads(2)
	fmt.Println("This is breads after 1 : ", breads[0].val, breads[1].val)

	// 2. 딸기잼 뚜껑을 연다.
	jam := &StrawberryJam{}
	OpenStrawberryJam(jam)
	fmt.Println("This is breads after 2 : ", breads[0].val, breads[1].val)

	// 3. 딸기잼을 한스푼 뜬다.
	spoon := GetOneSpoon(jam)
	fmt.Println("This is breads after 3 : ", breads[0].val, breads[1].val)

	// 4. 딸기잼을 빵에 바른다.
	PutJamOnBread(breads[0], spoon)
	fmt.Println("This is breads after 4 : ", breads[0].val, breads[1].val)

	// 5. 빵을 덮는다.
	sandwitch := MakeSandwitch(breads)

	// 6. 완성
	fmt.Println(sandwitch.val)
}
