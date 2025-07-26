package database

import (
	"ms-go-notification/internal/domain/campaign"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)
	return nil
}

func (c *CampaignRepository) Get() []campaign.Campaign {
	return c.campaigns
}
