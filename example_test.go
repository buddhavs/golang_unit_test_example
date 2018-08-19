package main

import "testing"

// NewMockService ...
func NewMockService() *Service {
	mq := &mq{}
	storage := &mockStorage{}
	status := &mockStatus{}

	service := Service{
		Mq:      mq,
		Storage: storage,
		Status:  status}

	mq.s = &service
	storage.s = &service
	status.s = &service

	return &service
}

type (
	mockStorage struct {
		service
	}

	mockStatus struct {
		service
	}
)

func (p *mockStorage) backup() error {
	return nil
}

func (p *mockStorage) snapshot() error {
	return nil
}

func (p *mockStatus) update() {
}

func (p *mockStatus) updatedb() {
}

func (p *mockStatus) check() {
}

func TestStatusExpectFailed(t *testing.T) {
	service := NewMockService()
	err := service.process(false)
	if err != nil {
		t.Error(err)
	}
}

func TestStatus(t *testing.T) {
	service := NewMockService()
	err := service.process(true)
	if err != nil {
		t.Error(err)
	}
}
