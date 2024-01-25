package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type PgConnection string

func (pc *PgConnection) CreateRecord(pn PhoneNumber) error {
	conn, err := pgx.Connect(context.Background(), string(*pc))
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	_, err = conn.Exec(context.Background(), "INSERT INTO archive1.NUMBERS(input, number) values ($1, $2)", pn.Input, pn.Number)
	if err != nil {
		return err
	}
	return nil
}

func (pc *PgConnection) GetRecord(number string) (PhoneNumber, error) {
	conn, err := pgx.Connect(context.Background(), string(*pc))
	if err != nil {
		return PhoneNumber{}, err
	}
	defer conn.Close(context.Background())
	pn := PhoneNumber{}
	err = conn.QueryRow(context.Background(), "SELECT input, number FROM archive1.NUMBERS WHERE NUMBER = $1", number).Scan(&pn.Input, &pn.Number)
	if err != nil {
		return PhoneNumber{}, err
	}
	return pn, nil
}

type PhoneNumber struct {
	Input, Number string
}
