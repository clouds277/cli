package commands_test

import (
	"cf"
	"cf/api"
	. "cf/commands"
	"cf/configuration"
	term "cf/terminal"
	"github.com/codegangsta/cli"
	"github.com/stretchr/testify/assert"
	"testhelpers"
	"testing"
)

func TestSuccessfullyLoggingIn(t *testing.T) {
	config := logout(t)

	ui := new(testhelpers.FakeUI)
	ui.Inputs = []string{"foo@example.com", "bar"}
	auth := &testhelpers.FakeAuthenticator{
		AccessToken: "my_access_token",
	}
	callLogin(
		nil,
		ui,
		config,
		&testhelpers.FakeOrgRepository{},
		&testhelpers.FakeSpaceRepository{},
		auth,
	)

	config, err := configuration.Load()
	assert.NoError(t, err)

	assert.Contains(t, ui.Outputs[0], config.Target)
	assert.Contains(t, ui.Outputs[2], "OK")
	assert.Contains(t, ui.Prompts[0], "Email")
	assert.Contains(t, ui.Prompts[1], "Password")

	assert.Equal(t, config.AccessToken, "my_access_token")
	assert.Equal(t, auth.Email, "foo@example.com")
	assert.Equal(t, auth.Password, "bar")
}

func TestLoggingInWithMultipleOrgsAndSpaces(t *testing.T) {
	config := logout(t)

	ui := new(testhelpers.FakeUI)
	ui.Inputs = []string{"foo@example.com", "bar", "2", "1"}

	orgs := []cf.Organization{
		cf.Organization{"FirstOrg", "org-1-guid"},
		cf.Organization{"SecondOrg", "org-2-guid"},
	}
	spaces := []cf.Space{
		cf.Space{"FirstSpace", "space-1-guid"},
		cf.Space{"SecondSpace", "space-2-guid"},
	}

	callLogin(
		nil,
		ui,
		config,
		&testhelpers.FakeOrgRepository{Organizations: orgs},
		&testhelpers.FakeSpaceRepository{Spaces: spaces},
		&testhelpers.FakeAuthenticator{},
	)

	assert.Contains(t, ui.Outputs[0], config.Target)

	assert.Contains(t, ui.Prompts[0], "Email")
	assert.Contains(t, ui.Prompts[1], "Password")
	assert.Contains(t, ui.Outputs[2], "OK")

	assert.Contains(t, ui.Outputs[3], "FirstOrg")
	assert.Contains(t, ui.Outputs[4], "SecondOrg")

	assert.Contains(t, ui.Prompts[2], "Organization")
	assert.Contains(t, ui.Outputs[5], "SecondOrg")
	assert.Contains(t, ui.Outputs[7], "FirstSpace")
	assert.Contains(t, ui.Outputs[8], "SecondSpace")

	assert.Contains(t, ui.Prompts[3], "Space")
	assert.Contains(t, ui.Outputs[9], "FirstSpace")

	config, err := configuration.Load()
	assert.NoError(t, err)
	assert.Equal(t, orgs[1], config.Organization)
	assert.Equal(t, spaces[0], config.Space)
}

func TestWhenUserPicksInvalidOrgNumberAndSpaceNumber(t *testing.T) {
	config := logout(t)

	orgs := []cf.Organization{
		cf.Organization{"Org1", "org-1-guid"},
		cf.Organization{"Org2", "org-2-guid"},
	}

	spaces := []cf.Space{
		cf.Space{"FirstSpace", "space-1-guid"},
		cf.Space{"SecondSpace", "space-2-guid"},
	}

	ui := new(testhelpers.FakeUI)
	ui.Inputs = []string{"foo@example.com", "bar", "3", "2", "3", "1"}

	callLogin(
		nil,
		ui,
		config,
		&testhelpers.FakeOrgRepository{Organizations: orgs},
		&testhelpers.FakeSpaceRepository{Spaces: spaces},
		&testhelpers.FakeAuthenticator{},
	)

	assert.Contains(t, ui.Prompts[2], "Organization")
	assert.Contains(t, ui.Outputs[5], "FAILED")

	assert.Contains(t, ui.Prompts[3], "Organization")
	assert.Contains(t, ui.Outputs[9], "Targeting org")

	assert.Contains(t, ui.Prompts[4], "Space")
	assert.Contains(t, ui.Outputs[13], "FAILED")

	assert.Contains(t, ui.Prompts[5], "Space")
	assert.Contains(t, ui.Outputs[17], "Targeting space")

	config, err := configuration.Load()
	assert.NoError(t, err)
	assert.Equal(t, orgs[1], config.Organization)
	assert.Equal(t, spaces[0], config.Space)
}

func TestLoggingInWitOneOrgAndOneSpace(t *testing.T) {
	config := logout(t)

	ui := new(testhelpers.FakeUI)
	ui.Inputs = []string{"foo@example.com", "bar"}

	orgs := []cf.Organization{
		cf.Organization{"FirstOrg", "org-1-guid"},
	}
	spaces := []cf.Space{
		cf.Space{"FirstSpace", "space-1-guid"},
	}

	callLogin(
		nil,
		ui,
		config,
		&testhelpers.FakeOrgRepository{Organizations: orgs},
		&testhelpers.FakeSpaceRepository{Spaces: spaces},
		&testhelpers.FakeAuthenticator{},
	)

	assert.Contains(t, ui.Outputs[0], config.Target)

	assert.Contains(t, ui.Prompts[0], "Email")
	assert.Contains(t, ui.Prompts[1], "Password")
	assert.Contains(t, ui.Outputs[2], "OK")
	assert.Contains(t, ui.Outputs[3], "FirstOrg")
	assert.Contains(t, ui.Outputs[4], "OK")

	assert.Contains(t, ui.Outputs[5], "API endpoint:")
	assert.Contains(t, ui.Outputs[7], "FirstOrg")
	assert.Contains(t, ui.Outputs[8], "FirstSpace")

	config, err := configuration.Load()
	assert.NoError(t, err)
	assert.Equal(t, orgs[0], config.Organization)
	assert.Equal(t, spaces[0], config.Space)
}

func TestLoggingInWithoutOrg(t *testing.T) {
	config := logout(t)

	ui := new(testhelpers.FakeUI)
	ui.Inputs = []string{"foo@example.com", "bar"}
	orgs := []cf.Organization{}
	spaces := []cf.Space{}

	callLogin(
		nil,
		ui,
		config,
		&testhelpers.FakeOrgRepository{Organizations: orgs},
		&testhelpers.FakeSpaceRepository{Spaces: spaces},
		&testhelpers.FakeAuthenticator{},
	)

	assert.Contains(t, ui.Outputs[0], config.Target)

	assert.Contains(t, ui.Prompts[0], "Email")
	assert.Contains(t, ui.Prompts[1], "Password")
	assert.Contains(t, ui.Outputs[2], "OK")
	assert.Contains(t, ui.Outputs[3], "No orgs found.")

	config, err := configuration.Load()
	assert.NoError(t, err)
	assert.Equal(t, cf.Organization{}, config.Organization)
	assert.Equal(t, cf.Space{}, config.Space)
}

func TestLoggingInWithOneOrgAndNoSpace(t *testing.T) {
	config := logout(t)

	ui := new(testhelpers.FakeUI)
	ui.Inputs = []string{"foo@example.com", "bar"}
	orgs := []cf.Organization{
		cf.Organization{"FirstOrg", "org-1-guid"},
	}
	spaces := []cf.Space{}

	callLogin(
		nil,
		ui,
		config,
		&testhelpers.FakeOrgRepository{Organizations: orgs},
		&testhelpers.FakeSpaceRepository{Spaces: spaces},
		&testhelpers.FakeAuthenticator{},
	)

	assert.Contains(t, ui.Outputs[0], config.Target)

	assert.Contains(t, ui.Prompts[0], "Email")
	assert.Contains(t, ui.Prompts[1], "Password")
	assert.Contains(t, ui.Outputs[2], "OK")

	assert.Contains(t, ui.Outputs[5], "API endpoint:")
	assert.Contains(t, ui.Outputs[7], "FirstOrg")
	assert.Contains(t, ui.Outputs[8], "No spaces found")

	config, err := configuration.Load()
	assert.NoError(t, err)
	assert.Equal(t, orgs[0], config.Organization)
	assert.Equal(t, cf.Space{}, config.Space)
}

func TestUnsuccessfullyLoggingIn(t *testing.T) {
	config := logout(t)

	ui := new(testhelpers.FakeUI)
	ui.Inputs = []string{
		"foo@example.com",
		"bar",
		"bar",
		"bar",
		"bar",
	}

	callLogin(
		nil,
		ui,
		config,
		&testhelpers.FakeOrgRepository{},
		&testhelpers.FakeSpaceRepository{},
		&testhelpers.FakeAuthenticator{AuthError: true},
	)

	assert.Contains(t, ui.Outputs[0], config.Target)
	assert.Equal(t, ui.Outputs[1], "Authenticating...")
	assert.Equal(t, ui.Outputs[2], "FAILED")
	assert.Equal(t, ui.Outputs[5], "Authenticating...")
	assert.Equal(t, ui.Outputs[6], "FAILED")
	assert.Equal(t, ui.Outputs[9], "Authenticating...")
	assert.Equal(t, ui.Outputs[10], "FAILED")
}

func callLogin(c *cli.Context, ui term.UI, config *configuration.Configuration, orgRepo api.OrganizationRepository, spaceRepo api.SpaceRepository, auth api.Authenticator) {
	l := NewLogin(ui, config, orgRepo, spaceRepo, auth)
	l.Run(c)
}

func logout(t *testing.T) (config *configuration.Configuration) {
	config, err := configuration.Load()
	assert.NoError(t, err)
	err = config.ClearSession()
	assert.NoError(t, err)
	return
}
