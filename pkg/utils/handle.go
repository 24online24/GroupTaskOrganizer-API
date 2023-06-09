package utils

import "fmt"

func HandleError(err error) { // afișează o eroare dacă există
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
