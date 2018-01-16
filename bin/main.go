package main

import (
	. "github.com/APA42/fake_go_component_1"
	"fmt"
)


func main(){
	var fake = FakeComponentStruct{};
	fmt.Println(fake.DoSomething())
}
