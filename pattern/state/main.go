package main

import "fmt"

/*
Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости
от своего состояния. Извне создаётся впечатление, что изменился класс объекта.

Применимость:
- Когда у вас есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния,
причём типов состояний много, и их код часто меняется.
- Когда код класса содержит множество больших, похожих друг на друга, условных операторов, которые выбирают поведения
в зависимости от текущих значений полей класса.
- Когда вы сознательно используете табличную машину состояний, построенную на условных операторах,
но вынуждены мириться с дублированием кода для похожих состояний и переходов.

Плюсы и минусы:
+ Избавляет от множества больших условных операторов машины состояний.
+ Концентрирует в одном месте код, связанный с определённым состоянием.
+ Упрощает код контекста.
- Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

type state interface {
	open(c *Connection)
	close(c *Connection)
}

type CloseState struct{}

func (cs CloseState) open(c *Connection) {
	fmt.Println("open the connection")
	c.setState(OpenState{})
}

func (cs CloseState) close(c *Connection) {
	fmt.Println("connection is already closed")
}

type OpenState struct{}

func (os OpenState) open(c *Connection) {
	fmt.Println("connection is already open")
}

func (os OpenState) close(c *Connection) {
	fmt.Println("close the connection")
	c.setState(CloseState{})
}

type Connection struct {
	_state state
}

func (c *Connection) Open() {
	c._state.open(c)
}

func (c *Connection) Close() {
	c._state.close(c)
}

func (c *Connection) setState(state state) {
	c._state = state
}

func main() {
	con := Connection{CloseState{}}
	con.Open()
	con.Open()
	con.Close()
	con.Close()
}
