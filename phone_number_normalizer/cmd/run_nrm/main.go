package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/ssh/knownhosts"
	"gophercises/phone_number_normalizer/db"
	"log"
)

func main() {
	input := flag.String("n", "(123) 456-7893", "Number to add into db, any format")
	flag.Parse()
	ctx := context.Background()
	conn, err := 	pgxpool.New(ctx, "user=postgres password=Passw0rd! host=localhost port=5432 dbname=archive search_path=archive1")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queries := db.New(conn)
	saveIfNotPresent(ctx, *input, queries, conn)

}

func saveIfNotPresent(ctx context.Context, input string, queries *db.Queries, conn *pgxpool.Pool) db.Number {
	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit(ctx)
	number := knownhosts.Normalize(input)
	num, err := queries.WithTx(tx).GetNumber(ctx, number)
	if err == nil {
		fmt.Println("Number is alredy present in the db ")
		return num
	} else if err != pgx.ErrNoRows {
		fmt.Println(err)
		return db.Number{}
	}
	fmt.Println("Saving number to the db")
	createdNumber, err := queries.CreateNumber(ctx, db.CreateNumberParams(db.Number{
		Input:  pgtype.Text{String: input, Valid: true},
		Number: number,
	}))
	if err != nil {
		fmt.Println(err)
	}
	return createdNumber
}
