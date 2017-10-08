package model

import "container/list"

type User struct {
	Name string
	Age int
}

func ListUser() list.List{

	l := list.List{};
	var  user1, user2, user3  User
	user1.Age=16
	user1.Name="ths1"

	user2.Age = 17
	user2.Name = "ths2"

	user3.Age = 18
	user3.Name = "ths3"

	l.PushBack(user1)
	l.PushBack(user2)
	l.PushBack(user3)
	return l
}
