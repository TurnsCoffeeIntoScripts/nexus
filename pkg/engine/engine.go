package engine

import (
	"fmt"
	"github.com/google/uuid"
)

// Engine ...
// TODO comment
type Engine struct {
	ID uuid.UUID
}

// CreateEngine ...
// TODO comment
func CreateEngine() *Engine {
	e := &Engine{}
	e.ID = uuid.New()

	return e
}

// Start ...
// The main method that loads, initialize and start the main threads of the engine
func (e *Engine) Start() {
	fmt.Printf("New engine UUID: %v\n", e.ID)
}
