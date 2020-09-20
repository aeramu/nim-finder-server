package graphql

type UserConnection interface {
	Edges() []User
	PageInfo() PageInfo
}
