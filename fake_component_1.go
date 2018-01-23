package fake_go_component_1

import "fmt"

const(
	COMPONENT = "FakeComponent_1"
	VERSION = "0.0.1"
)

type FakeComponent1 interface {
	DoSomething() string
}

type FakeComponentStruct struct {
}

func (f FakeComponentStruct) DoSomething() string {
	return fmt.Sprintf("Component: %s  Version: %s", COMPONENT, VERSION)
}
