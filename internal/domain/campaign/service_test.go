package campaign

import (
	"errors"
	"ms-go-notification/internal/contract"
	"ms-go-notification/internal/internal_errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = contract.NewCampaignDTO{
		Name:    "Campaign X",
		Content: "Default Content",
		Emails:  []string{"email1@e.com", "email2@e.com"},
	}
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)

}

func (r *repositoryMock) Get() ([]Campaign, error) {
	//args := r.Called(campaign)
	return nil, nil

}

func Test_Create_Campaign(t *testing.T) {
	assertions := assert.New(t)
	repo := new(repositoryMock)
	service := Service{repo}
	repo.On("Save", mock.Anything).Return(nil)

	id, err := service.Create(newCampaign)

	repo.AssertExpectations(t)
	assertions.NotNil(id)
	assertions.Nil(err)
}

func Test_Create_Campaign_DomainError(t *testing.T) {
	assertions := assert.New(t)

	campaignDto := contract.NewCampaignDTO{
		Name:    "",
		Content: "Body",
		Emails:  []string{"email1@e.com"},
	}

	repo := new(repositoryMock)
	service := Service{repo}
	id, err := service.Create(campaignDto)

	assertions.False(errors.Is(internal_errors.ErrInternal, err))
	assertions.Empty(id)
	assertions.NotEmpty(err)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assertions := assert.New(t)

	repo := new(repositoryMock)
	service := Service{repo}
	repo.On("Save", mock.Anything).Return(internal_errors.ErrInternal)
	service.Repository = repo

	_, err := service.Create(newCampaign)
	assertions.True(errors.Is(err, internal_errors.ErrInternal))
}
