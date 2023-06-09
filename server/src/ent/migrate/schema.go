// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "name", Type: field.TypeString, Size: 50},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "title", Type: field.TypeString},
		{Name: "body", Type: field.TypeString},
		{Name: "category_id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_categories_posts",
				Columns:    []*schema.Column{PostsColumns[6]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		PostsTable,
		UsersTable,
	}
)

func init() {
	PostsTable.ForeignKeys[0].RefTable = CategoriesTable
	PostsTable.ForeignKeys[1].RefTable = UsersTable
}
