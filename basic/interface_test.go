package basic

import (
	"testing"
	"fmt"
)


type DeptModeFull interface {
	Name() string
	SetName(name string)
	Relocate(building string, floor uint8)
}

type DeptModeA interface {
	Name() string
	SetName(name string)
}

type DeptModeB interface {
	Relocate(building string, floor uint8)
}

type Dept struct {
	name     string
	building string
	floor    uint8
	Key      string
}

func (self Dept) Name() string {
	return self.name
}

func (self Dept) SetName(name string) {
	self.name = name
}

func (self *Dept) Relocate(building string, floor uint8) {
	self.building = building
	self.floor = floor
}


//1、结构Dept的方法集中仅包含方法接收者为Dept的方法，即：Name()和SetName()）,
//	所以，结构Dept的实例仅为DeptModeA的实现。
//2、结构的指针*Dept的方法集包含了方法接受者为Dept和*Dept的方法，即：Name()、SetName()和Relocate(),
//	所以，接口Dept的实例的指针为全部三个接口——DeptModeFull、DeptModeA和DeptModeB的实现。
func TestInterfaceImplement(t *testing.T){
	dept1 := Dept{
		name:     "MySohu",
		building: "Internet",
		floor:    7}

	switch v := interface{}(dept1).(type) {
	case DeptModeFull:
		fmt.Printf("The dept1 is a DeptModeFull.\n")
	case DeptModeB:
		fmt.Printf("The dept1 is a DeptModeB.\n")
	case DeptModeA:
		fmt.Printf("The dept1 is a DeptModeA.\n")
	default:
		fmt.Printf("The type of dept1 is %v\n", v)
	}

	deptPtr1 := &dept1
	if _, ok := interface{}(deptPtr1).(DeptModeFull); ok {
		fmt.Printf("The deptPtr1 is a DeptModeFull.\n")
	}
	if _, ok := interface{}(deptPtr1).(DeptModeA); ok {
		fmt.Printf("The deptPtr1 is a DeptModeA.\n")
	}
	if _, ok := interface{}(deptPtr1).(DeptModeB); ok {
		fmt.Printf("The deptPtr1 is a DeptModeB.\n")
	}
}