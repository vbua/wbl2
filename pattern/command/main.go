package main

import "fmt"

/*
Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты, позволяя передавать их
как аргументы при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.

Применимость:
- Когда вы хотите параметризовать объекты выполняемым действием.
- Когда вы хотите ставить операции в очередь, выполнять их по расписанию или передавать по сети.
- Когда вам нужна операция отмены.

Плюсы и минусы:
+ Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
+ Позволяет реализовать простую отмену и повтор операций.
+ Позволяет реализовать отложенный запуск операций.
+ Позволяет собирать сложные команды из простых.
+ Реализует принцип открытости/закрытости.
- Усложняет код программы из-за введения множества дополнительных классов.
*/

type Command interface {
	Execute() string
}

type PingCommand struct{}

func (p *PingCommand) Execute() string {
	return "ping"
}

type StatusCommand struct{}

func (p *StatusCommand) Execute() string {
	return "status"
}

func execByName(name string) string {
	// Register commands
	commands := map[string]Command{
		"ping":   &PingCommand{},
		"status": &StatusCommand{},
	}

	if command := commands[name]; command == nil {
		return "Нет такой команды"
	} else {
		return command.Execute()
	}
}

func main() {
	fmt.Println(execByName("status"))
	fmt.Println(execByName("ping"))

	fmt.Println(execByName("test"))
}
