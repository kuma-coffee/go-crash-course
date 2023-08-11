package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	person := &Person{
		Firstname: "John",
		Lastname:  "Doe",
	}

	serializedPerson, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}

	err = ioutil.WriteFile("person.data", serializedPerson, 0644)
	if err != nil {
		log.Fatal("error writing the file:", err)
	}

	person2 := &Person{}
	err = proto.Unmarshal(serializedPerson, person2)
	if err != nil {
		log.Fatal("unmarshaling error:", err)
	}
	fmt.Println(person2.Firstname)
	fmt.Println(person2.Lastname)
}
