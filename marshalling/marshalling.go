package marshalling

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Date   string `json:"date"`
}

func Run() {
	b := Book{
		Title:  "The Hitchhiker's Guide to the Galaxy",
		Author: "Douglas Adams",
		Date:   "1979",
	}

	byteArray, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling", err)
	}

	fmt.Printf("Unmarshalled: \n%+v\nMarshalled: \n%s\n\n", b, byteArray)

	// should be - `{"title":"The Hitchhiker's Guide to the Galaxy","author":"Douglas Adams","date":"1979"}`
	str := `{"title":"The Hitchhiker's Guide to the Galaxy","author":"Douglas Adams","date":"1979"}`

	newBook := Book{}
	// var newBook interface{}
	err = json.Unmarshal([]byte(str), &newBook)

	if err != nil {
		fmt.Println(fmt.Errorf("error unmarshalling %+v", err))
	}

	fmt.Println("Unmarshalled: ", newBook)

	// does not work with structs
	// switch newBook.(type) {
	// case Book:
	// 	var date interface{}
	// 	fmt.Println("Is book")

	// 	date = newBook.(Book).Date

	// 	switch date.(type) {
	// 	case string:
	// 		fmt.Println("Missing field Date is a string")
	// 		fmt.Printf("Unmarshalled book:\n%+v\n\n", newBook)
	// 	}
	// }

}
