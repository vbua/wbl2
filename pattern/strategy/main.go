package main

import (
	"fmt"
	"log"
)

/*
Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает
каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

Применимость:
- Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
- Когда у вас есть множество похожих классов, отличающихся только некоторым поведением.
- Когда вы не хотите обнажать детали реализации алгоритмов для других классов.
- Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора.
Каждая ветка такого оператора представляет собой вариацию алгоритма.

Плюсы и минусы:
+ Горячая замена алгоритмов на лету.
+ Изолирует код и данные алгоритмов от остальных классов.
+ Уход от наследования к делегированию.
+ Реализует принцип открытости/закрытости.
- Усложняет программу за счёт дополнительных классов.
- Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

type strategy interface {
	push(key, value string, ttl int64) error
	pop(key string) error
}

type File struct{}

func (File) push(key, value string, ttl int64) error {
	fmt.Println("Pushing [", key, ":", value, "] to file...")

	return nil
}

func (File) pop(key string) error {
	fmt.Println("Popping [", key, "] from file...")

	return nil
}

type Redis struct{}

func (Redis) push(key, value string, ttl int64) error {
	fmt.Println("Pushing [", key, ":", value, "] to redis...")

	return nil
}

func (Redis) pop(key string) error {
	fmt.Println("Popping [", key, "] from redis...")

	return nil
}

type Cache struct {
	Strategy strategy
}

func (c *Cache) Push(key, value string, ttl int64) error {
	return c.Strategy.push(key, value, ttl)
}

func (c *Cache) Pop(key string) error {
	return c.Strategy.pop(key)
}

func main() {
	var c = &Cache{}

	c.Strategy = File{}
	if err := c.Push("key-f", "value-f", 3600); err != nil {
		log.Fatalln(err)
	}

	c.Strategy = Redis{}
	if err := c.Pop("key-r"); err != nil {
		log.Fatalln(err)
	}
}
