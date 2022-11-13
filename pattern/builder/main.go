package main

import (
	"fmt"
	"strings"
)

/*
Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

Применение:
- Когда вы хотите избавиться от «телескопического конструктора».
- Когда ваш код должен создавать разные представления какого-то объекта. Например, деревянные и железобетонные дома.
- Когда вам нужно собирать сложные составные объекты.

Плюсы и минусы:
+ Позволяет создавать продукты пошагово.
+ Позволяет использовать один и тот же код для создания различных продуктов.
+ Изолирует сложный код сборки продукта от его основной бизнес-логики.
- Усложняет код программы из-за введения дополнительных классов.
- Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения
результата.
*/

type burger struct {
	bread    string
	hasMeat  bool
	toppings []string
	sauces   []string
}

type iBurgerBuilder interface {
	setBread()
	setMeat()
	setToppings()
	setSauces()
	getBurger() burger
}

type veganBurger struct {
	burger
}

func (v *veganBurger) setBread() {
	v.burger.bread = "обычный белый хлеб"
}

func (v *veganBurger) setMeat() {
	v.burger.hasMeat = false
}

func (v *veganBurger) setToppings() {
	v.burger.toppings = []string{"помидоры", "огурцы", "лук"}
}

func (v *veganBurger) setSauces() {
	v.burger.sauces = []string{"сырный"}
}

func (v *veganBurger) getBurger() burger {
	return v.burger
}

type meatBurger struct {
	burger
}

func (v *meatBurger) setBread() {
	v.burger.bread = "обычный белый хлеб"
}

func (v *meatBurger) setMeat() {
	v.burger.hasMeat = true
}

func (v *meatBurger) setToppings() {
	v.burger.toppings = []string{"помидоры", "огурцы", "лук"}
}

func (v *meatBurger) setSauces() {
	v.burger.sauces = []string{"сырный"}
}

func (v *meatBurger) getBurger() burger {
	return v.burger
}

type director struct {
	builder iBurgerBuilder
}

func (d *director) setBuilder(builder iBurgerBuilder) {
	d.builder = builder
}

func (d *director) buildBurger() burger {
	d.builder.setBread()
	d.builder.setMeat()
	d.builder.setToppings()
	d.builder.setSauces()

	return d.builder.getBurger()
}

func describeBurger(burger burger) {
	fmt.Printf("хлеб: %s, мясо: %t, начинка: %s, соусы: %s\n",
		burger.bread, burger.hasMeat, strings.Join(burger.toppings, ", "), strings.Join(burger.sauces, ", "))
}

func main() {
	director := &director{}
	director.setBuilder(&veganBurger{})
	veggieDelightSub := director.buildBurger()
	describeBurger(veggieDelightSub)

	fmt.Println("------------")

	director.setBuilder(&meatBurger{})
	chickenTeriyakiSub := director.buildBurger()

	describeBurger(chickenTeriyakiSub)
}
