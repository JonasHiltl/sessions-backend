// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "blurhash", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"ADMIN", "USER"}, Default: "USER"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "friend_count", Type: field.TypeInt, Default: 0},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_username",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[1]},
			},
			{
				Name:    "user_id",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[0]},
			},
		},
	}
	// UserFriendsColumns holds the columns for the "user_friends" table.
	UserFriendsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeString},
		{Name: "friend_id", Type: field.TypeString},
	}
	// UserFriendsTable holds the schema information for the "user_friends" table.
	UserFriendsTable = &schema.Table{
		Name:       "user_friends",
		Columns:    UserFriendsColumns,
		PrimaryKey: []*schema.Column{UserFriendsColumns[0], UserFriendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_friends_user_id",
				Columns:    []*schema.Column{UserFriendsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_friends_friend_id",
				Columns:    []*schema.Column{UserFriendsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
		UserFriendsTable,
	}
)

func init() {
	UserFriendsTable.ForeignKeys[0].RefTable = UsersTable
	UserFriendsTable.ForeignKeys[1].RefTable = UsersTable
}
