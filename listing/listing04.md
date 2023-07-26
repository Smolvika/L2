```go
package main
 
func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
 
    for n := range ch {
        println(n)
    }
}

```
Ответ:
```
Вывод: 
0
1
2
3
4
5
6
7
8
9
deadlock
Так как канал не закрыт, range будет ждать новыйх данных, но никто больше не напишет,так как цикл завершен, поэтому дедлок
```