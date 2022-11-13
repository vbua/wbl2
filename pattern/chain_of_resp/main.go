package main

import "fmt"

/*
Цепочка обязанностей — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно
по цепочке обработчиков. Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать
запрос дальше по цепи.

Применимость:
- Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно,
какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
- Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
- Когда набор объектов, способных обработать запрос, должен задаваться динамически.

Плюсы и минусы:
+ Уменьшает зависимость между клиентом и обработчиками.
+ Реализует принцип единственной обязанности.
+ Реализует принцип открытости/закрытости.
- Запрос может остаться никем не обработанным.
*/

type step interface {
	run(*customer)
	setNextStep(step)
}

type customer struct {
	name           string
	isHighPriority bool
}

type voiceAssistant struct {
	next step
}

func (v *voiceAssistant) run(cust *customer) {
	fmt.Println("[Автоответчик робот] Общаюсь с покупателем: ", cust.name)
	v.next.run(cust)
}

func (v *voiceAssistant) setNextStep(next step) {
	v.next = next
}

type associate struct {
	next step
}

func (a *associate) run(cust *customer) {
	if cust.isHighPriority {
		fmt.Println("Отправляю покупателя к менеджеру")
		a.next.run(cust)
		return
	}
	fmt.Println("[Ассистент] Общаюсь с покупателем: ", cust.name)
	a.next.run(cust)
}

func (a *associate) setNextStep(next step) {
	a.next = next
}

type manager struct {
	next step
}

func (a *manager) run(cust *customer) {
	fmt.Println("[Менеджер] Общаюсь с клиентом: ", cust.name)
}

func (a *manager) setNextStep(next step) {
	a.next = next
}

func main() {
	m := &manager{}

	assoc := &associate{}
	assoc.setNextStep(m)

	va := &voiceAssistant{}
	va.setNextStep(assoc)

	normalCust := &customer{
		name: "Bob",
	}

	va.run(normalCust)

	fmt.Println("===================")

	highPriorityCust := &customer{
		name:           "John",
		isHighPriority: true,
	}

	va.run(highPriorityCust)
}
