package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"email1@e.com", "email2@e.com"}
)

func Test_NewCampaign(t *testing.T) {
	assertions := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assertions.NotEmpty(campaign.ID)
	assertions.NotEmpty(campaign.CreatedOn)
	assertions.Equal(campaign.Name, name)
	assertions.Equal(campaign.Content, content)
	assertions.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assertions := assert.New(t)

	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assertions.Equal(campaign.Name, name)
	assertions.Equal(campaign.Content, content)
	assertions.Equal(len(campaign.Contacts), len(contacts))
	assertions.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assertions.Equal(err.Error(), "name is required")
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assertions.Equal(err.Error(), "content is required")
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assertions.Equal(err.Error(), "contacts is required")
}
