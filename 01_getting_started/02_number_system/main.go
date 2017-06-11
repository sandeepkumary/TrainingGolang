package main 

import "fmt"

func main(){
	fmt.Printf("%d \n",42)
	fmt.Printf("%b \n",42)
	fmt.Printf("%x \n",42)
	fmt.Printf("%q \n",42)

	fmt.Printf("decimal \t binary \t hex \t character \n")
	for i:=0;i<100; i++{
		fmt.Printf("%d \t %b \t %x \t %c \n",i,i,i,i)
	}

}