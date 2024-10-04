package services_test

import (
	"apiblog/src/domain/entities"
	"apiblog/src/infrastructure/commons/models/page"
	"apiblog/test/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertNewPublications(t *testing.T) {
	mockRepo := &mocks.MockPublicationRepository{}
	service := mocks.NewPublicationService(mockRepo)

	publication := entities.Publication{Title: "Test Publication"}
	err := service.InsertNewPublications(publication)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !mockRepo.InsertNewPublicationsCalled {
		t.Error("InsertNewPublications was not called")
	}
}
func TestListPublications(t *testing.T) {
	mockRepo := &mocks.MockPublicationRepository{}
	service := mocks.NewPublicationService(mockRepo)

	slice, err := service.ListPublications("autor", "order", page.Pagination{})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !mockRepo.ListPublicationsCalled {
		t.Error("InsertNewPublications was not called")
	}

	assert.NotNil(t, slice)
}

func TestGetPublicationByTitleAndAutor(t *testing.T) {
	mockRepo := &mocks.MockPublicationRepository{}
	service := mocks.NewPublicationService(mockRepo)

	structs, err := service.GetPublicationByTitleAndAutor("autor", "order")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !mockRepo.GetPublicationByTitleAndAutorCalled {
		t.Error("InsertNewPublications was not called")
	}

	assert.NotNil(t, structs)
}

func TestUpdatePublication(t *testing.T) {
	mockRepo := &mocks.MockPublicationRepository{}
	service := mocks.NewPublicationService(mockRepo)

	err := service.UpdatePublication(entities.Publication{})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !mockRepo.UpdatePublicationCalled {
		t.Error("InsertNewPublications was not called")
	}
}

func TestDeletePublicationByTitleAndAutor(t *testing.T) {
	mockRepo := &mocks.MockPublicationRepository{}
	service := mocks.NewPublicationService(mockRepo)

	err := service.DeletePublicationByTitleAndAutor("autor", "order")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !mockRepo.DeletePublicationByTitleAndAutorCalled {
		t.Error("InsertNewPublications was not called")
	}
}
