package main

import "database/sql"

func loadUser(db *sql.DB, name string) error {
	// TODO: use parameterized SQL and handle the returned error correctly.
	_, err := db.Exec("SELECT id FROM users WHERE name = $1", name)
	return err
}

func main() {}
