package main 
import "fmt"

func wrapper() func() int {
	x:=0
	return func() int {
		x++
		return x
	}
}

func main(){
incr := wrapper()
fmt.Println(incr())
fmt.Println(incr())

incr2 := func() int {
	y:=0
	y++
	return y
}
fmt.Println(incr2())
fmt.Println(incr2())
}