package main

import (
	"fmt"
	"ms-proto/service"

	"google.golang.org/protobuf/proto"
)

func main() {
	user := &service.User{
		Username: "mszlu",
		Age:      18,
	}

	marshal, err := proto.Marshal(user) //序列化
	if err != nil {
		panic(err)
	}
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser) //反序列化
	if err != nil {
		panic(err)
	}

	fmt.Println(newUser.String())
}
