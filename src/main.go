package main

import (
	"fmt"
	"transform"
)

func main() {
	inputValue := "Aspiration.com"
	//Below code is to Check given input value is AlphaNumeric or not
	/*isAlphaNumeric := regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(inputValue)
	if !isAlphaNumeric {
		fmt.Printf("%s is not an Alphanumeric string", inputValue)
		return
	}*/
	s := transform.NewSkipString(3, inputValue)
	transform.MapString(&s)
	fmt.Println("Output using with interfaces:", s)
	fmt.Println("Output using function:", transform.CapitalizeEveryThirdAlphanumericChar(inputValue))

}
