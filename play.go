package main

import (
	"fmt"
	"time"
)

type User struct {
  name string
  arg int
}

func (u *User) userAdpter() {
  fmt.Println(u.name)
}

func sendMessage(msg string) {
  fmt.Println("send message: ", msg)
}

func sub() {
  fmt.Println("sub is running")
  time.Sleep(2 * time.Second)
  fmt.Println("sub is finished")
}

// type Person struct {
// 	Name string
// 	Arg  int
// }

type Animal struct {
  Name string
  Arg int
}

// var person = Person{
//         Name: "sora",
//         Arg: 23,
//     }

var animal = Animal{
  Name: "sora",
  Arg: 23,
}

func main() {
  // c := make(chan string)
  // go func() {
  //   c <- "hello"
  // }()
  //
  // close(c)

  fmt.Println(animal)
}

