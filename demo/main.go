package main

import "fmt"

type iInterface interface {
	interfaceFunc()
}

type s1 struct {
	value int
}

func (s *s1) interfaceFunc() {
	fmt.Println("Hello, I am s1 and my value is: ", s.value)
}

func testFunc1(i iInterface) {
	fmt.Println("This is testFunc1")
	i.interfaceFunc()
}

func testFunc2[T iInterface](i T) {
	fmt.Println("This is testFunc2")
	i.interfaceFunc()
}

func interfaceForSlice(is []iInterface) {
	fmt.Println("This is interfaceForSlice")
	for _, i := range is {
		i.interfaceFunc()
	}
}

func genericsForSlice[T iInterface](is []T) []T {
	fmt.Println("This is genericsForSlice")
	for _, i := range is {
		i.interfaceFunc()
	}
	return is
}

func interfaceWithReturn(i iInterface) iInterface {
	return i
}

func genericsWithReturn[T iInterface](i T) T {
	return i
}

func main() {
	struct1, struct2 := &s1{1}, &s1{2}

	testFunc1(struct1)

	testFunc2(struct2)

	interfaceForSlice([]iInterface{&s1{3}, &s1{4}})

	// This is invalid
	//interfaceForSlice([]*s1{&s1{3}, &s1{4}})

	genericsForSlice([]*s1{&s1{3}, &s1{4}})

	returnFromInterface, returnFromgenerics := interfaceWithReturn(struct1), genericsWithReturn(struct1)

	fmt.Printf("%T, %T", returnFromInterface, returnFromgenerics)

	returnFromgenerics.value = 10

	// This is invalid
	returnFromInterface.(*s1).value = 10
}
