Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
Выведет:
nil
false

В данном случае динамический тип равен os.PathError, а значение равно nil. Интерфейс равен nil, только если
и динамический тип, и значение равно nil.

Под капотом интерфейс будут храниться в eface структуре, если это пустой интерфейс, или в iface, если у него есть методы.
В _type содержится информация о типе значения, которое хранит интерфейс, а data - указатель на реальные данные. 
В случае iface, itab содержит информацию как об интерфейсе, так и о его целевом типе, а fun[0] указывает, 
реализует ли тип интерфейс или нет. fun[0]!=0 означает, что тип реализует интерфейс.

type eface struct {
	_type *_type
	data  unsafe.Pointer
}

type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type itab struct {
	inter *interfacetype
	_type *_type
    ...
    ...
	fun  
}

```
