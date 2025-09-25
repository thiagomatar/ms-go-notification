package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Default Content"
	contacts = []string{"email1@e.com", "email2@e.com"}
	fake     = faker.New()
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

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assertions.Equal("name must be greater than 5", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assertions.Equal("name must be less than 24", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assertions.Equal("content must be greater than 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(2000), contacts)

	assertions.Equal("content must be less than 1024", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assertions.Equal("contacts must be greater than 1", err.Error())
}

func Test_NewCampaign_MustValidateValidEmail(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assertions.Equal("email is invalid", err.Error())
}
