package uuid

import "github.com/google/uuid"

type UUIDGenerator interface {
	Generate() (string, error)
}

type Generator struct {
}

//Generates a Random
func (v *Generator) Generate() (string, error) {
	uuid, err := uuid.NewRandom()
	return uuid.String(), err
}
