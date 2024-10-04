package mocks

import (
	"apiblog/src/domain/entities"
	"apiblog/src/infrastructure/commons/models/page"

	"github.com/stretchr/testify/mock"
)

type MockPublicationRepository struct {
	mock.Mock
	InsertNewPublicationsCalled bool
	InsertNewPublicationsResult error

	ListPublicationsCalled bool
	ListPublicationsResult []entities.Publication
	ListPublicationsError  error

	GetPublicationByTitleAndAutorCalled bool
	GetPublicationByTitleAndAutorResult entities.Publication
	GetPublicationByTitleAndAutorError  error

	UpdatePublicationCalled bool
	UpdatePublicationResult error

	DeletePublicationByTitleAndAutorCalled bool
	DeletePublicationByTitleAndAutorResult error
}

func (m *MockPublicationRepository) Insert(publication entities.Publication) error {
	m.InsertNewPublicationsCalled = true
	return m.InsertNewPublicationsResult
}

func (m *MockPublicationRepository) List(autor, order string, pagination page.Pagination) ([]entities.Publication, error) {
	m.ListPublicationsCalled = true
	m.ListPublicationsResult = append(m.ListPublicationsResult, entities.Publication{Title: "Teste1 "})
	m.ListPublicationsResult = append(m.ListPublicationsResult, entities.Publication{Title: "Teste2"})
	return m.ListPublicationsResult, nil
}

func (m *MockPublicationRepository) PublicationByTitleAndAutor(title, autor string) (entities.Publication, error) {
	m.GetPublicationByTitleAndAutorCalled = true
	return m.GetPublicationByTitleAndAutorResult, nil
}

func (m *MockPublicationRepository) Update(publication entities.Publication) error {
	m.UpdatePublicationCalled = true
	return nil
}

func (m *MockPublicationRepository) DeleteByTitleAndAutor(title, autor string) error {
	m.DeletePublicationByTitleAndAutorCalled = true
	return nil
}
