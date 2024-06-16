package main

import "fmt"

type Creature struct {
	Species string
}

func (creature *Creature) Reset() {
	creature.Species = ""
}

func Pointer_demo() {
	creature := Creature{
		Species: "shark",
	}

	var src *Creature

	fmt.Printf("Species %+v\n", creature)
	changeCreature(src)
	fmt.Printf("Species %+v\n", creature)
}

func changeCreature(creature *Creature) {
	if creature == nil {
		fmt.Println("The creature is null please check")
		return
	}
	creature.Species = "jellyfish"
	fmt.Printf("Species inside change : %+v\n", *creature)
}

func Method_pointers() {
	creature := Creature{Species: "shark"}
	fmt.Printf("Species :%+v\n", creature)
	creature.Reset()
	fmt.Printf("Species : %+v\n", creature)
}

func Misc() {
	var i int = 10
	var j *int = &i

	var k **int = &j // pointer to a pointer

	fmt.Println(**k)
}
