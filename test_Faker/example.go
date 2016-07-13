package main

import (
	"fmt"
	"github.com/wawandco/fako"
)

type User struct {
	Name     string `fako:"full_name"`
	Username string `fako:"username"`
	Email    string `fako:"email_address"` //Notice the fako:"email_address" tag
	Phone    string `fako:"phone"`
	Password string `fako:"simple_password"`
	Address  string `fako:"street_address"`
}

func main() {
	var user User
	fako.Fill(&user)

	fmt.Println("fako.Fill(&user) result: \r\n\t", user)
	// This prints something like AnthonyMeyer@Twimbo.biz
	// or another valid email

	var userWithOnlyEmail User
	fako.FillOnly(&userWithOnlyEmail, "Email")
	fmt.Println("fako.FillOnly(&userWithOnlyEmail, \"Email\") result: \r\n\t", userWithOnlyEmail)
	//This will fill all only the email

	var userWithoutEmail User
	fako.FillExcept(&userWithoutEmail, "Email")
	fmt.Println("fako.FillExcept(&userWithoutEmail, \"Email\") result: \r\n\t", userWithoutEmail)
	//This will fill all the fields except the email

}
