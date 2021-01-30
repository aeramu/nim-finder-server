package search

// StudentRepo adapter
type StudentRepo interface {
	Search(keyword string, limit int, after string) StudentConnection
}
