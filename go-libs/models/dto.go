package models

type Dto interface {
	IsValid() error
}
