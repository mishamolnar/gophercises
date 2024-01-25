package main

import (
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5"
	"gophercises/phone_number_normalizer"
	"gophercises/phone_number_normalizer/db"
)

func oldMain() {
	//conn := db.PgConnection("postgresql://postgres:Passw0rd!@localhost:5432/archive?currentSchema=archive1")
	conn := db.PgConnection("user=postgres password=Passw0rd! host=localhost port=5432 dbname=archive")
	input := flag.String("n", "(123) 456-7893", "Number to add into db, any format")
	flag.Parse()
	normalizeAndSave(*input, conn)
}

func normalizeAndSave(input string, connection db.PgConnection) {
	number := phone_number_normalizer.Normalize(input)
	_, err := connection.GetRecord(number)
	if err != nil && err != pgx.ErrNoRows {
		fmt.Println(err)
		return
	} else if err == nil {
		fmt.Println("Phone number was present")
		return
	}
	fmt.Println("Saving phone number")
	err = connection.CreateRecord(db.PhoneNumber{
		Input:  input,
		Number: number,
	})
	if err != nil {
		fmt.Println(err)
	}

}
