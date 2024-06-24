package main

import (
	"fmt"

	goinject "github.com/d1360-64rc14/go-inject"
)

func main() {
	goinject.Register[Printer](&LoremPrinter{
		lorem: "Lorem ipsum dolor sit amet, ",
	})

	print()
	printAgain()

	printUses()
}

func print() {
	var p Printer

	goinject.Inject(&p)

	p.Print("Diego")
}

func printAgain() {
	var p Printer

	goinject.Inject(&p)

	p.Print("Garcia")
}

func printUses() {
	var p Printer

	goinject.Inject(&p)

	if lp, ok := p.(*LoremPrinter); ok {
		fmt.Println(lp.uses)
	}
}

type Printer interface {
	Print(string)
}

type LoremPrinter struct {
	lorem string
	uses  int
}

func (s *LoremPrinter) Print(str string) {
	fmt.Println(s.lorem + str + "!")
	s.uses++
}

type HelloPrinter struct{}

func (s *HelloPrinter) Initialize() {
}

func (s *HelloPrinter) Print(str string) {
	fmt.Println("Hello World, " + str + "!")
}
