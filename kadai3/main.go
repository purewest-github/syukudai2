package main

import "fmt"

type Turtle struct {
	name string
	x    float64
	y    float64
	a    float64
}

// 課題3-1
func (t *Turtle) move(a float64) {
	t.a -= a
}

// 課題3-2
func (n *Turtle) moven(name string) {
	n.name = name
}
func (k *Turtle) movek(name string) {
	k.name += name

}
func main() {
	// structの宣言の仕方
	var k1 Turtle = Turtle{"ふうた先生",0,0,0}
	var t1 Turtle = Turtle{"師匠", 1000, 5, 180.5}
	var t2 Turtle = Turtle{"弟子", 10, 250, 270.3}
	fmt.Println(t1)
	fmt.Println(t2)
	// 課題3-1
	t1.move(35)
	fmt.Println(t1)
	// 課題3-2
	t2.moven("西")
	fmt.Println(t2)
	k1.movek(":こんにちは！")
	fmt.Println(k1)

}
