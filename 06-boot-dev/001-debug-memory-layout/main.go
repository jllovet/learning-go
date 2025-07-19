package main

import (
	"fmt"
	"reflect"
)

type contact struct {
	sendingLimit int32
	age          int32
	userID       string
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

func main() {

	typ := reflect.TypeOf(1)
	fmt.Printf("Struct is %d bytes\n", typ.Size())

	typ = reflect.TypeOf("")
	fmt.Printf("Struct is %d bytes\n", typ.Size())

	typ = reflect.TypeOf(contact{})
	fmt.Printf("Struct is %d bytes\n", typ.Size())

	typ = reflect.TypeOf(perms{})
	fmt.Printf("Struct is %d bytes\n", typ.Size())
}
