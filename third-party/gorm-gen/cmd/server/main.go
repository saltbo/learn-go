package main

import "gorm-gen-example/query"

func main() {
	q := query.Use()

	q.User.Find()

	q.Company.First()
}
