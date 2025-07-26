package endpoints

import "ms-go-notification/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
