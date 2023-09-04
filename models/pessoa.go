package models

import (
	"github.com/google/uuid"
)

type Pessoa struct {
	Apelido     string
	Nome        string
	Nascimento  string
	Stack       []string
	Id          uuid.UUID
}
