package repository

import (
	"blog/ent"
	"blog/ent/enttest"
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeEntClient(t *testing.T) (*ent.Client, error) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	return client, client.Schema.Create(context.Background())
}
