package model

import (
	"github.com/go-errors/errors"
	"time"
)

type ErrorIncident struct {
	IncidentTime time.Time

	DeploymentID string
	Level        string
	Message      string
}

func (e *ErrorIncident) New() error {
	return errors.New(e)
}
