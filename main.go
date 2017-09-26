package main

import (
	"fmt"
	"os"
	sc "strconv"

	"github.com/icrowley/fake"
)

func main() {

	currentDir, _ := os.Getwd()
	fmt.Println(currentDir)
	fileWriterPTR, ferr := os.Create(currentDir + "/testdata")
	if ferr != nil {
		fmt.Println("file creation error", ferr)
	}
	defer fileWriterPTR.Close()

	err := fake.SetLang("en")
	if err != nil {
		panic(err)
	}
	//password := fake.SimplePassword()

	for i := 0; i < 10; i++ {

		ssn := fake.DigitsN(9)
		phone := fake.DigitsN(9)
		city := fake.City()
		state := fake.State()
		address := fake.StreetAddress()
		s, _ := sc.Atoi(ssn)
		firstname := fake.MaleFirstName()
		if s%2 == 0 {
			firstname = fake.FemaleFirstName()
		}

		lastname := fake.LastName()
		fmt.Fprintf(fileWriterPTR, "%s,%s,%s,%s,%s,%s,%s\n", ssn, phone, city, state, address, firstname, lastname)
	}

}
