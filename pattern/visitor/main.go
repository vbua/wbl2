package main

import "fmt"

/*
Посетитель — это поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции,
не изменяя классы объектов, над которыми эти операции могут выполняться.

Применимость:
- Когда вам нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов, например, деревом.
- Когда над объектами сложной структуры объектов надо выполнять некоторые не связанные между собой операции, но вы не
хотите «засорять» классы такими операциями.
- Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.


Плюсы и минусы:
+ Упрощает добавление операций, работающих со сложными структурами объектов.
+ Объединяет родственные операции в одном классе.
+ Посетитель может накапливать состояние при обходе структуры элементов.
- Паттерн не оправдан, если иерархия элементов часто меняется.
- Может привести к нарушению инкапсуляции элементов.
*/

type CarPart interface {
	Accept(CarPartVisitor)
}

type Wheel struct {
	Name string
}

func (w *Wheel) Accept(visitor CarPartVisitor) {
	visitor.visitWheel(w)
}

type Engine struct{}

func (e *Engine) Accept(visitor CarPartVisitor) {
	visitor.visitEngine(e)
}

type Car struct {
	parts []CarPart
}

func NewCar() *Car {
	this := new(Car)
	this.parts = []CarPart{
		&Wheel{"переднее слева"},
		&Wheel{"переднее справа"},
		&Wheel{"заднее слева"},
		&Wheel{"заднее справа"},
		&Engine{}}
	return this
}

func (c *Car) Accept(visitor CarPartVisitor) {
	for _, part := range c.parts {
		part.Accept(visitor)
	}
}

type CarPartVisitor interface {
	visitWheel(wheel *Wheel)
	visitEngine(engine *Engine)
}

type GetMessageVisitor struct {
	Messages []string
}

func (g *GetMessageVisitor) visitWheel(wheel *Wheel) {
	g.Messages = append(g.Messages, fmt.Sprintf("Посещаю %v колесо\n", wheel.Name))
}

func (g *GetMessageVisitor) visitEngine(engine *Engine) {
	g.Messages = append(g.Messages, fmt.Sprintf("Посещаю движок\n"))
}

func main() {
	car := NewCar()
	visitor := new(GetMessageVisitor)
	car.Accept(visitor)
	fmt.Println(visitor.Messages)
}
