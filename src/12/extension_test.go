package main

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	fmt.Print("dog")
}
func (d *Dog) SpeakTo(host string) {
	d.p.SpeakTo(host)
}
func TestDog(t *testing.T) {
	dog := new (Dog)
	dog.SpeakTo("xsss")
}
