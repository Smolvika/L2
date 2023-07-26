```go
package main

import (
	"fmt"
)

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```
Ответ:
```
Вывод: 
2  
1
Если функция возвращается через явный оператор return, то defer выполняется после того, как значение в return уже установлено
```