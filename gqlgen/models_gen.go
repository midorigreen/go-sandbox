// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	User User   `json:"user"`
}
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}