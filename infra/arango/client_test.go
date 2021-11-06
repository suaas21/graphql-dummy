package arango

import (
	"context"
	"fmt"
	"github.com/suaas21/graphql-dummy/config"
	"testing"
)

func TestNewClient(t *testing.T) {
	db, err := NewArangoDB(context.Background(), &config.Arango{
		Host:       "localhost",
		Port:       "8529",
		DBUser:     "root",
		DBPassword: "admin",
		DBName:     "_system",
	})
	if err != nil {
		t.Fatal(err)
	}

	client := NewClient(db)

	query := fmt.Sprintf(`
	for x in Book
		filter x.ok == true
		LIMIT 0, 2
		return x
	`)

	res, count, err := client.Query(context.Background(), query, nil)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res, count)
}
