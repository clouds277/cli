package v7

import (
	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
	v6shared "code.cloudfoundry.org/cli/command/v6/shared"
	"code.cloudfoundry.org/cli/command/v7/shared"
)

//go:generate counterfeiter . RestartActor

type RestartActor interface {
	GetApplicationByNameAndSpace(appName string, spaceGUID string) (v7action.Application, v7action.Warnings, error)
	GetApplicationSummaryByNameAndSpace(appName string, spaceGUID string, withObfuscatedValues bool, routeActor v7action.RouteActor) (v7action.ApplicationSummary, v7action.Warnings, error)
	PollStart(appGUID string) (v7action.Warnings, error)
	StartApplication(appGUID string) (v7action.Warnings, error)
	StopApplication(appGUID string) (v7action.Warnings, error)
}

type RestartCommand struct {
	RequiredArgs        flag.AppName `positional-args:"yes"`
	usage               interface{}  `usage:"CF_NAME restart APP_NAME"`
	relatedCommands     interface{}  `related_commands:"restage, restart-app-instance"`
	envCFStagingTimeout interface{}  `environmentName:"CF_STAGING_TIMEOUT" environmentDescription:"Max wait time for buildpack staging, in minutes" environmentDefault:"15"`
	envCFStartupTimeout interface{}  `environmentName:"CF_STARTUP_TIMEOUT" environmentDescription:"Max wait time for app instance startup, in minutes" environmentDefault:"5"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       RestartActor
	RouteActor  v7action.RouteActor
}

func (cmd *RestartCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	cmd.SharedActor = sharedaction.NewActor(config)

	ccClient, _, err := shared.NewClients(config, ui, true, "")
	if err != nil {
		return err
	}

	ccClientV2, uaaClientV2, err := v6shared.NewClients(config, ui, true)
	if err != nil {
		return err
	}

	cmd.Actor = v7action.NewActor(ccClient, config, nil, nil)
	v2Actor := v2action.NewActor(ccClientV2, uaaClientV2, config)
	cmd.RouteActor = v2Actor

	return nil
}

func (cmd RestartCommand) Execute(args []string) error {
	err := cmd.SharedActor.CheckTarget(true, true)
	if err != nil {
		return err
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}

	app, warnings, err := cmd.Actor.GetApplicationByNameAndSpace(cmd.RequiredArgs.AppName, cmd.Config.TargetedSpace().GUID)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}
	cmd.UI.DisplayTextWithFlavor("Restarting app {{.AppName}} in org {{.OrgName}} / space {{.SpaceName}} as {{.Username}}...", map[string]interface{}{
		"AppName":   cmd.RequiredArgs.AppName,
		"OrgName":   cmd.Config.TargetedOrganization().Name,
		"SpaceName": cmd.Config.TargetedSpace().Name,
		"Username":  user.Name,
	})

	if app.Started() {
		cmd.UI.DisplayText("\nStopping app...")

		warnings, err = cmd.Actor.StopApplication(app.GUID)
		cmd.UI.DisplayWarnings(warnings)
		if err != nil {
			return err
		}
	}

	warnings, err = cmd.Actor.StartApplication(app.GUID)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}

	cmd.UI.DisplayText("\nWaiting for app to start...")
	warnings, err = cmd.Actor.PollStart(app.GUID)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}

	appSummaryDisplayer := shared.NewAppSummaryDisplayer(cmd.UI)
	summary, warnings, err := cmd.Actor.GetApplicationSummaryByNameAndSpace(
		cmd.RequiredArgs.AppName,
		cmd.Config.TargetedSpace().GUID,
		false,
		cmd.RouteActor,
	)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}

	appSummaryDisplayer.AppDisplay(summary, false)

	return nil
}
