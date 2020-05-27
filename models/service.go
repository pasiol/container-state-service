package models

import (
	"errors"
	"time"
)

type Service struct {
	Name      string
	Started   bool
	Ended     bool
	CreatedAt time.Time
	EndedAt   time.Time
}

var (
	services []*Service
	nextID   int
)

func AddService(s Service) (Service, error) {
	if s.Name == "" {
		return Service{}, errors.New("the service name cannot be empty")
	}
	if !s.Started && s.Ended {
		return Service{}, errors.New("the service cannot ended if it not have started")
	}
	s.CreatedAt = time.Now()
	if s.Ended {
		s.EndedAt = time.Now()
	} else {
		s.EndedAt = time.Time{}
	}
	services = append(services, &s)
	return s, nil
}
