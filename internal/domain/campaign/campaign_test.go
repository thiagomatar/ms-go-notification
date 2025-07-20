package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign("1", name, content, contacts)

	if campaign.ID != "2" {
		t.Errorf("Expected ID to be '1', got '%s'", campaign.ID)
	} else if campaign.Name != name {
		t.Errorf("Expected Name to be '%s', got '%s'", name, campaign.Name)
	} else if campaign.Content != content {
		t.Errorf("Expected Content to be '%s', got '%s'", content, campaign.Content)
	} else if len(campaign.Contacts) != len(contacts) {
		t.Errorf("Expected correct contacts")
	}
}
