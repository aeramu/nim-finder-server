package service

import (
	"github.com/aeramu/nim-finder-server/usecase/adapter"
)

type interactor struct {
	r adapter.Repository
}
