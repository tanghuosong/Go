package main

import (
	"./function"
	"fmt"
	"./model"
	"container/list"
)

func main() {
	var user  model.User
	user.Name = "THS"
	user.Age =18
	fmt.Println(user)
	fmt.Println(function.Add(2,2))

	l := list.New()
	var  user1, user2, user3  model.User
	user1.Age=16
	user1.Name="ths1"

	user2.Age = 17
	user2.Name = "ths2"

	user3.Age = 18
	user3.Name = "ths3"

	l.PushBack(user1)
	l.PushBack(user2)
	l.PushBack(user3)

	fmt.Println(l.Len())

	//for i:= model.ListUser().Front() ; i != nil; i.Next() {
	//	fmt.Println(i.Value)
	//}
}
