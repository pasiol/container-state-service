package models

import (
	"fmt"
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
	for _, existing := range services {
		if s.Name == existing.Name {
			return Service{}, fmt.Errorf("service with name '%s' already exists", s.Name)
		}
	}
	if s.Name == "" {
		return Service{}, fmt.Errorf("the service name cannot be empty")
	}
	if !s.Started && s.Ended {
		return Service{}, fmt.Errorf("the service cannot ended if it not have started")
	}
	s.Started = true
	s.CreatedAt = time.Now()
	if s.Ended {
		s.EndedAt = time.Now()
	} else {
		s.EndedAt = time.Time{}
	}
	services = append(services, &s)
	return s, nil
}

func GetServiceByName(name string) (Service, error) {
	for _, s := range services {
		if name == s.Name {
			return *s, nil
		}
	}
	return Service{}, fmt.Errorf("service with name '%s' not founded", name)
}

func SetEnded(name string) (Service, error) {
	for _, s := range services {
		if name == s.Name {
			s.Ended = true
			s.EndedAt = time.Now()
			return *s, nil
		}
	}
	return Service{}, fmt.Errorf("service with name '%s' not founded", name)
}
