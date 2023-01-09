package main

import (
	"fmt"

	"github.com/aleksey-lihodievski/course/arraysAndSlices"
	"github.com/aleksey-lihodievski/course/customErrors"
	"github.com/aleksey-lihodievski/course/features"
	"github.com/aleksey-lihodievski/course/functions"
	"github.com/aleksey-lihodievski/course/generics"
	"github.com/aleksey-lihodievski/course/interfaces"
	"github.com/aleksey-lihodievski/course/types"
)

func main() {
	features.Run()
	types.Run()
	arraysAndSlices.Run()
	functions.Run()
	customErrors.Run()
	generics.Run()
	interfaces.Run()

	fmt.Println("After all program")
}