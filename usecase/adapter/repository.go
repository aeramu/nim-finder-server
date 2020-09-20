package adapter

import "github.com/aeramu/nim-finder-server/entity"

//Repository adapter
type Repository interface {
	Search(keyword string) []entity.User
}
