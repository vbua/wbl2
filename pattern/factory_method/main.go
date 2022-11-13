package main

import "fmt"

/*
Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания
объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.

Применимость:
- Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код.
- Когда вы хотите дать возможность пользователям расширять части вашего фреймворка или библиотеки.
- Когда вы хотите экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых.

Плюсы и минусы:
+ Избавляет класс от привязки к конкретным классам продуктов.
+ Выделяет код производства продуктов в одно место, упрощая поддержку кода.
+ Упрощает добавление новых продуктов в программу.
+ Реализует принцип открытости/закрытости.
- Может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать
свой подкласс создателя.
*/

type Pet interface {
	GetName() string
	GetAge() int
	GetSound() string
}

type pet struct {
	name  string
	age   int
	sound string
}

func (p *pet) GetName() string {
	return p.name
}

func (p *pet) GetSound() string {
	return p.sound
}

func (p *pet) GetAge() int {
	return p.age
}

type Dog struct {
	pet
}

type Cat struct {
	pet
}

func GetPet(petType string) Pet {
	if petType == "собака" {
		return &Dog{
			pet{
				name:  "Найда",
				age:   2,
				sound: "гав",
			},
		}
	}
	if petType == "кошка" {
		return &Cat{
			pet{
				name:  "Машка",
				age:   3,
				sound: "мяу",
			},
		}
	}
	return nil
}

func describePet(pet Pet) string {
	return fmt.Sprintf("%s %d лет. Ее/его звук %s", pet.GetName(), pet.GetAge(), pet.GetSound())
}

func main() {
	dog := GetPet("собака")
	petDescription := describePet(dog)

	fmt.Println(petDescription)
	fmt.Println("-------------")

	cat := GetPet("кошка")
	petDescription = describePet(cat)

	fmt.Println(petDescription)
}
