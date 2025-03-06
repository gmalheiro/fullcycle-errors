package main

import "fmt"

func main() {
	val, err := safeDiv(7, 0)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Printf("val %d\n", val)
	}
}

func safeDiv(a, b int) (val int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("division by zero")
		}
	}()
	return div(a, b), nil
}

func div(a, b int) int {
	return a / b
}
