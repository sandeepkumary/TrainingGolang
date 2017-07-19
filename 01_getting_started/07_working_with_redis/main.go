package main

import (
	"fmt"
	//import the redis client library
	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	//connect to the already running redis server ( here localhost)
	// default port 6379 is used here
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Error :", err)
	}

	defer conn.Close()

	//Write the redis command using Cmd() function
	//First arguement is always a valid redis command , foolowed by optionally keys and values

	// Setting the string data structure
	resp := conn.Cmd("SET", "key2", "this is set value from golang")

	if resp.Err != nil {
		fmt.Println("Error : ", resp.Err)
	}

	//Getting the value back from redis server
	resp = conn.Cmd("GET", "key2")
	if resp.Err != nil {
		fmt.Println("Error :", resp.Err)
	}

	fmt.Println("Value for the key1 :", resp.String())
}
