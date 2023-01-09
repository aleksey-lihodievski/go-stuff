package customErrors

import (
	"errors"
	"fmt"
	"log"
)

type SomeError struct {
	Msg          string
	someNewField string
}

func (s SomeError) Error() string {
	return s.Msg + " Custom Error"
}

func thatTriggersError() error {
	return &SomeError{Msg: "Hello error message", someNewField: "Error handling is cool"}
	// return errors.New("Error")
}

func thatHandlesError() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered")
		}
	}()

	err := thatTriggersError()
	artificialErr := new(SomeError)

	if err != nil {
		switch {
		case errors.As(err, &artificialErr):
			fmt.Println("Error is artificial", artificialErr.someNewField)
			log.Panic(err)
		default:
			fmt.Println("Error is not artificial")
			log.Panic(err)
		}
	}
	fmt.Println("No error")
}

func errorFuncLevelThree() error {
	err := errorFuncLevelTwo()

	returnErr := fmt.Errorf("error in function level three: %w", err)

	fmt.Println(returnErr)
	fmt.Println(errors.Unwrap(returnErr))
	fmt.Println(errors.Unwrap(errors.Unwrap(returnErr)))

	return returnErr
}

func errorFuncLevelTwo() error {
	err := errorFuncLevelOne()
	return fmt.Errorf("error in function level two: %w", err)
}

func errorFuncLevelOne() error {
	err := errors.New("error in function level one")
	return err
}

func Run() {
	// thatHandlesError()
	// errorFuncLevelThree()
}
