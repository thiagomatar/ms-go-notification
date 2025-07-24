package campaign

import (
	"ms-go-notification/internal/contract"
	"ms-go-notification/internal/internal_errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaignDTO) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internal_errors.ErrInternal
	}

	return campaign.ID, nil
}
