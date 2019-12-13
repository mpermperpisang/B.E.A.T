package helper

import (
	"fmt"
)

/*ErrorMessage : global error*/
func ErrorMessage(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
