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
		fmt.Println("Error : ", err)
	}

	defer conn.Close()

	//Write the redis command using Cmd() function
	//First arguement is always a valid redis command , foolowed by optionally keys and values

	// Setting the string data structure
	resp_set_str := conn.Cmd("SET", "key2", "this is set value from golang")

	if resp_set_str.Err != nil {
		fmt.Println("Error : ", resp_set_str.Err)
	}

	//Getting the value back from redis server
	resp_get_str := conn.Cmd("GET", "key2")
	if resp_get_str.Err != nil {
		fmt.Println("Error :", resp_get_str.Err)
	}

	fmt.Println("Value for the key1 :", resp_get_str.String())

	// some more advanced stuff

	//setting a hash datastructure
	resp_set_hash := conn.Cmd("HMSET", "users:tom", "first_name", "Tom", "last_name", "Moody", "dob", "1981/09/08", "gender", "M", "status", "ACTIVE", "email", "abc.xyz@examples.com")
	if resp_set_hash.Err != nil {
		fmt.Println("Error : ", resp_set_hash.Err)
	}

	first_name, err := conn.Cmd("HGET", "users:tom", "first_name").Str()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	last_name, err := conn.Cmd("HGET", "users:tom", "last_name").Str()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	dob, err := conn.Cmd("HGET", "users:tom", "dob").Str()
	if err != nil {
		fmt.Println("Error :", err)
	}

	gender, err := conn.Cmd("HGET", "users:tom", "gender").Str()
	if err != nil {
		fmt.Println("Error : ", err)
	}

	status, err := conn.Cmd("HGET", "users:tom", "status").Str()
	if err != nil {
		fmt.Println("Error :", err)
	}

	email, err := conn.Cmd("HGET", "users:tom", "email").Str()
	if err != nil {
		fmt.Println("Error : ", err)
	}
	fmt.Println("..................................................................................")
	fmt.Printf("first name \t:\t%s \nlast name \t:\t%s\ndob \t\t:\t%s\ngender \t\t:\t%s\nstatus \t\t:\t%s\nemail \t\t:\t%s\n", first_name, last_name, dob, gender, status, email)
	fmt.Println("..................................................................................")
}
