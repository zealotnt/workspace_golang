package main

import (
	"fmt"
	"regexp"
)

func main() {
	DATABASE_URL := "postgres://ukwklxchohiwrt:m_9V5QFtURhM6JKjlkBAVyNvvm@ec2-54-83-202-64.compute-1.amazonaws.com:5432/dbbpbsb92frcn6"

	fmt.Println("1. Basic regex:")
	r, err := regexp.Compile(`postgres`)

	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return
	}

	// Will print 'Match'
	if r.MatchString(DATABASE_URL) == true {
		fmt.Printf("Match \r\n")
		fmt.Printf("%v" + "\r\n", r.FindAllString((DATABASE_URL), -1))
		fmt.Printf("%v" + "\r\n", r.FindAllStringIndex((DATABASE_URL), -1))

	} else {
		fmt.Printf("No match ")
	}

	fmt.Println("2.   Parenthesis finding:")
	fmt.Println("2.1. Basic:")
	//[[cat c] [sat s] [mat m]]
	re, err := regexp.Compile(`(.)at`) // want to know what is in front of 'at'
	res := re.FindAllStringSubmatch("The cat sat on the mat.", -1)
	fmt.Printf("%v" + "\r\n", res)
	fmt.Println(res[1][0])

	fmt.Println("2.2. More than one group:")
	// Prints [[ex e x] [ec e c] [e  e  ]]
	s := "Nobody expects the Spanish inquisition."
	re1, err := regexp.Compile(`(e)(.)`) // Prepare our regex
	result_slice := re1.FindAllStringSubmatch(s, -1)
	fmt.Printf("%v", result_slice)

	fmt.Println("2.3. Empty result group:")
	s = "Mr. Leonard Spock"
	re1, err = regexp.Compile(`(Mr)(s)?\. (\w+) (\w+)`)
	result:= re1.FindStringSubmatch(s)

	for k, v := range result {
		fmt.Printf("%d. %s\n", k, v)
	}
	// Prints
	// 0. Mr. Leonard Spock
	// 1. Mr
	// 2.
	// 3. Leonard
	// 4. Spock	
	result_slice = re1.FindAllStringSubmatch(s, -1)
	fmt.Printf("%v\n", result_slice)



	fmt.Println("Let do our exercise:")
	fmt.Printf("%s \r\n", DATABASE_URL)
	re1, err = regexp.Compile(`(\w*):\/\/`)
	result = re1.FindStringSubmatch(DATABASE_URL)
	for k, v := range result {
		fmt.Printf("%d. %s\n", k, v)
	}
	
	if result != nil {
		fmt.Printf("%s", result[1])
	} else {
		fmt.Println("Exercise: String not found !!!")
	}
}