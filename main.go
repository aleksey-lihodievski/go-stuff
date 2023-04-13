package main

import (
	"fmt"

	"github.com/aleksey-lihodievski/course/arraysAndSlices"
	"github.com/aleksey-lihodievski/course/customContext"
	"github.com/aleksey-lihodievski/course/customErrors"
	"github.com/aleksey-lihodievski/course/features"
	"github.com/aleksey-lihodievski/course/functions"
	"github.com/aleksey-lihodievski/course/generics"
	"github.com/aleksey-lihodievski/course/goroutines"
	"github.com/aleksey-lihodievski/course/interfaces"
	"github.com/aleksey-lihodievski/course/marshalling"
	"github.com/aleksey-lihodievski/course/mutexes"
	"github.com/aleksey-lihodievski/course/orm"
	"github.com/aleksey-lihodievski/course/problems"
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
	customContext.Run()
	goroutines.Run()
	mutexes.Run()
	marshalling.Run()
	orm.Run()

	problems.SolveProblems()

	fmt.Println("After all program")
}
