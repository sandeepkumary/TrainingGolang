package main

import (
	"fmt"
	//import the redis client library
	"github.com/mediocregopher/radix.v2/redis"
)

//create a struct to capture all tags from the has datastructure
type user struct {
	first_name string
	last_name  string
	dob        string
	gender     string
	status     string
	email      string
}

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

	fmt.Println("Value for the key2 :", resp_get_str.String())

	// some more advanced stuff

	// 1 ***************** HASH DATASTRUCTURE with indivdual tags *****************************

	//setting a hash datastructure
	resp_set_hash := conn.Cmd("HMSET", "users:tom", "first_name", "Tom", "last_name", "Moody", "dob", "1981/09/08", "gender", "M", "status", "ACTIVE", "email", "abc.xyz@examples.com")
	if resp_set_hash.Err != nil {
		fmt.Println("Error : ", resp_set_hash.Err)
	}

	//getting hash datastructure individual tags
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

	//getting hash datastructer, all tags in one go
	// 2 ***************** HASH DATASTRUCTURE with all tags *****************************

	reply_hash_all, err := conn.Cmd("HGETALL", "users:tom").Map()
	if err != nil {
		fmt.Println("Error : ", err)
	}

	us1 := new(user)
	us1, err = populate_user_details(reply_hash_all)

	if err != nil {
		fmt.Println("Error : ", err)
	}

	print_user_details(us1)
}

// Function to set the values of hash tags from map to struct variable
func populate_user_details(usr map[string]string) (*user, error) {
	var err error
	us := new(user)
	us.first_name = usr["first_name"]
	us.last_name = usr["last_name"]
	us.dob = usr["dob"]
	us.gender = usr["gender"]
	us.status = usr["status"]
	us.email = usr["email"]

	return us, err
}

// Function to print the user details using printing of struct variable
func print_user_details(usr *user) {
	fmt.Println("..................................................................................")
	fmt.Printf("first name \t:\t%s \n", usr.first_name)
	fmt.Printf("last name \t:\t%s \n", usr.last_name)
	fmt.Printf("dob \t\t:\t%s \n", usr.dob)
	fmt.Printf("gender \t\t:\t%s \n", usr.gender)
	fmt.Printf("status \t\t:\t%s \n", usr.status)
	fmt.Printf("email \t\t:\t%s \n", usr.email)
	fmt.Println("..................................................................................")
}
