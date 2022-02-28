package utils

import (
	"fmt"
)

func HandleError(e error) {
	if e != nil {
		PrintError(fmt.Sprint(e))
	}
}

func PrintError(err string){
	fmt.Println("")
	fmt.Println("!!-------------------------------------------!!")
	fmt.Println("	" + err)
	fmt.Println("	Try again")
	fmt.Println("!!-------------------------------------------!!")
	fmt.Println("")
}
