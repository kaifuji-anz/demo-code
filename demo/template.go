package main

import "fmt"

type iInterfaceWithT[T any] interface {
	*T
	iInterface
}

func genericsReturnPtr[T iInterface](input T) T {
	var a T
	if input == nil {

	}
	return a
	// This is invalid
	// return nil
}

func genericsReturnPtr2[T any](_ T) *T {
	return nil
}

func genericsReturnPtr3[T iInterface](_ T) *T {
	return nil
}

func genericsReturnPtr4[T any, PT iInterfaceWithT[T]](_ PT) PT {
	return nil
}

func testFunc() {

	struct1 := s1{value: 10}
	struct1Ptr := &struct1

	return1 := genericsReturnPtr(struct1Ptr)
	return2 := genericsReturnPtr2(struct1)

	fmt.Printf("%T, %T", return1, return2)

	// This is invalid
	//genericsReturnPtr3[s1](struct1)
	return3 := genericsReturnPtr3[*s1](struct1Ptr)
	// This is invalid
	//return3.value = 10

	return4 := genericsReturnPtr4(struct1Ptr)
	//return4 := genericsReturnPtr4(struct1)
	return4.value = 10

	_ = return3
}
