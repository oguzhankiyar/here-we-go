package main

import "fmt"

type Component interface {
	Accept(visitor Visitor)
}

type Monitor struct {
	Id string
}

func (m Monitor) Accept(visitor Visitor) {
	visitor.VisitMonitor(m)
}

type Keyboard struct {
	Id string
}

func (k Keyboard) Accept(visitor Visitor) {
	visitor.VisitKeyboard(k)
}

type Mouse struct {
	Id string
}

func (m Mouse) Accept(visitor Visitor) {
	visitor.VisitMouse(m)
}

type Visitor interface {
	VisitMouse(mouse Mouse)
	VisitKeyboard(keyboard Keyboard)
	VisitMonitor(monitor Monitor)
}

type ConnectComponentVisitor struct {

}

func (v ConnectComponentVisitor) VisitMouse(mouse Mouse) {
	fmt.Println("Connected mouse:", mouse.Id)
}

func (v ConnectComponentVisitor) VisitKeyboard(keyboard Keyboard) {
	fmt.Println("Connected keyboard:", keyboard.Id)
}

func (v ConnectComponentVisitor) VisitMonitor(monitor Monitor) {
	fmt.Println("Connected monitor:", monitor.Id)
}

type Computer struct {
	Components 	[]Component
}

func NewComputer(components ...Component) Computer {
	return Computer{components}
}

func (c Computer) Accept(visitor Visitor) {
	for _, v := range c.Components {
		v.Accept(visitor)
	}
}

func main() {
	computer := NewComputer(Monitor{"dell"}, Keyboard{"microsoft"}, Mouse{"logitech"})

	connectVisitor := ConnectComponentVisitor{}

	computer.Accept(connectVisitor)
}