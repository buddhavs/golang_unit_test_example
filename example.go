package main

import (
	"errors"
	"fmt"
)

type (
	// Service ...
	Service struct {
		Mq
		Storage
		Status
	}

	// Mq ...
	Mq interface {
		process(bool) error
	}

	// Storage ...
	Storage interface {
		snapshot() error
		backup() error
	}

	// Status ...
	Status interface {
		update()
		updatedb()
		check()
	}

	service struct {
		s *Service
	}

	mq struct {
		service
	}

	storage struct {
		service
	}

	status struct {
		service
	}
)

// NewService ...
func NewService() *Service {
	mq := &mq{}
	storage := &storage{}
	status := &status{}

	service := Service{
		Mq:      mq,
		Storage: storage,
		Status:  status}

	mq.s = &service
	storage.s = &service
	status.s = &service

	return &service
}

func (p *mq) process(b bool) error {
	fmt.Println("Processing...")

	err := p.s.backup()
	if err != nil {
		fmt.Println("Backup error...")
	}

	err = p.s.snapshot()
	if err != nil {
		fmt.Println("Snapshot error...")
	}

	switch b {
	case true:
		return nil
	default:
		return errors.New("Process error")
	}
}

func (p *storage) backup() error {
	fmt.Println("Backing up...")
	p.s.update()
	p.s.updatedb()
	p.s.check()
	return nil
}

func (p *storage) snapshot() error {
	fmt.Println("Snap shotting XD...")
	return nil
}

func (p *status) update() {
	fmt.Println("Updaing status...")
}

func (p *status) updatedb() {
	fmt.Println("Updaing Db...")
}

func (p *status) check() {
	fmt.Println("Checking status...")
}

func main() {
	service := NewService()
	service.process(true)
}
