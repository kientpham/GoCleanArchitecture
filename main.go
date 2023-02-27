package main

import (
	"fmt"
	"gocleanarchitecture/controller"
)

func main() {
	fmt.Println("Hello")
	user := controller.NewUser("John", 20)
	fmt.Println(user.GetName())

}
