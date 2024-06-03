package main

import (
	"fmt"

	"github.com/opa-oz/go-todo/todo"
)

type TestServer struct {
	name    string
	version string
}

type TestEntity struct {
	name string
}

func (t *TestServer) callFunc(entity *TestEntity, extractName func(*TestEntity) string) string {
	return extractName(entity)
}

func (t *TestServer) unimplementedFunc() (*TestEntity, error) {
	gotValue := todo.String("Make https request here", "return this for now")

	fmt.Println(fmt.Sprintf("\t%s", gotValue))

	handler := todo.StringF("Return current date as string", "2024-01-01")

	fmt.Println(fmt.Sprintf("\t%s", handler("smth")))

	return todo.Ptr[TestEntity]("Implement this function later"), todo.Error("Make up some error or something")
}
func (t *TestServer) implementedFunc() *TestEntity {
	entity, err := t.unimplementedFunc()

	todo.Opts.ChangePrefix("~~ fancy todo ~~")

	if err != nil {
		fmt.Println(fmt.Sprintf("\t%s", err))
	}

	doubleCheck := todo.PtrF[TestServer]("Get pointer lately")

	doubleCheck()

	return entity
}

func main() {
	server := TestServer{}

	server.implementedFunc()
}
