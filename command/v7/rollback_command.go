package v7

import (
	"context"
	"fmt"

	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/cf/errors"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/v7/shared"
	"code.cloudfoundry.org/cli/resources"
)

type RollbackCommand struct {
	Actor RollbackActor
	BaseCommand

	Force           bool                 `short:"f" description:"Force rollback without confirmation"`
	RequiredArgs    flag.AppName         `positional-args:"yes"`
	Version         flag.PositiveInteger `long:"revision" required:"true" description:"Roll back to the given app revision"`
	relatedCommands interface{}          `related_commands:"revisions"`
	usage           interface{}          `usage:"CF_NAME rollback APP_NAME [--revision REVISION_NUMBER] [-f]"`

	LogCacheClient sharedaction.LogCacheClient
	Stager         shared.AppStager
	//stopStreamingFunc func()
}

//go:generate counterfeiter . RollbackActor

type RollbackActor interface {
	CreateDeploymentByApplicationAndRevision(appGUID string, revisionGUID string) (string, v7action.Warnings, error)
	GetApplicationByNameAndSpace(name string, spaceGUID string) (resources.Application, v7action.Warnings, error)
	GetDetailedAppSummary(appName string, spaceGUID string, withObfuscatedValues bool) (v7action.DetailedApplicationSummary, v7action.Warnings, error)
	GetRevisionByApplicationAndVersion(appGUID string, revisionVersion int) (resources.Revision, v7action.Warnings, error)
	GetRevisionsByApplicationNameAndSpace(appName string, spaceGUID string) ([]resources.Revision, v7action.Warnings, error)
	GetStreamingLogsForApplicationByNameAndSpace(appName string, spaceGUID string, client sharedaction.LogCacheClient) (<-chan sharedaction.LogMessage, <-chan error, context.CancelFunc, v7action.Warnings, error)
}

func (cmd RollbackCommand) Execute(args []string) error {
	// cmd.stopStreamingFunc = nil
	cmd.UI.DisplayWarning(command.ExperimentalWarning)
	cmd.UI.DisplayNewline()

	targetRevision := int(cmd.Version.Value)
	err := cmd.SharedActor.CheckTarget(true, true)
	if err != nil {
		return err
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}

	app, warnings, _ := cmd.Actor.GetApplicationByNameAndSpace(cmd.RequiredArgs.AppName, cmd.Config.TargetedSpace().GUID)
	cmd.UI.DisplayWarnings(warnings)

	revisions, warnings, _ := cmd.Actor.GetRevisionsByApplicationNameAndSpace(app.Name, cmd.Config.TargetedSpace().GUID)

	cmd.UI.DisplayWarnings(warnings)
	if len(revisions) == 0 {
		return errors.New(fmt.Sprintf("No revisions for app %s", cmd.RequiredArgs.AppName))
	}
	newRevision := revisions[len(revisions)-1].Version + 1
	revision, warnings, _ := cmd.Actor.GetRevisionByApplicationAndVersion(app.GUID, targetRevision)
	cmd.UI.DisplayWarnings(warnings)

	// TODO Localization?

	if !cmd.Force {
		cmd.UI.DisplayTextWithFlavor("Rolling '{{.AppName}}' back to revision '{{.TargetRevision}}' will create a new revision. The new revision '{{.NewRevision}}' will use the settings from revision '{{.TargetRevision}}'.", map[string]interface{}{
			"AppName":        cmd.RequiredArgs.AppName,
			"TargetRevision": targetRevision,
			"NewRevision":    newRevision,
		})

		prompt := "Are you sure you want to continue?"
		response, promptErr := cmd.UI.DisplayBoolPrompt(false, prompt, nil)

		if promptErr != nil {
			return promptErr
		}

		if !response {
			cmd.UI.DisplayText("App '{{.AppName}}' has not been rolled back to revision '{{.TargetRevision}}'.", map[string]interface{}{
				"AppName":        cmd.RequiredArgs.AppName,
				"TargetRevision": targetRevision,
			})
			return nil
		}
	}
	cmd.UI.DisplayTextWithFlavor("Rolling back to revision {{.TargetRevision}} for app {{.AppName}} in org {{.OrgName}} / space {{.SpaceName}} as {{.Username}}...", map[string]interface{}{
		"AppName":        cmd.RequiredArgs.AppName,
		"TargetRevision": targetRevision,
		"OrgName":        cmd.Config.TargetedOrganization().Name,
		"SpaceName":      cmd.Config.TargetedSpace().Name,
		"Username":       user.Name,
	})

	_, warnings, _ = cmd.Actor.CreateDeploymentByApplicationAndRevision(app.GUID, revision.GUID)
	cmd.UI.DisplayWarnings(warnings)

	cmd.UI.DisplayOK()

	

// 	cmd.UI.DisplayText("Staging app and tracing logs...")
// 	logStream, errStream, cancelFunc, warnings, err := cmd.Actor.GetStreamingLogsForApplicationByNameAndSpace(cmd.RequiredArgs.AppName, cmd.Config.TargetedSpace().GUID, cmd.LogCacheClient)
// 	cmd.UI.DisplayWarnings(warnings)
// 	if err != nil {
// 		return err
// 	}
// 	if cmd.stopStreamingFunc != nil {
// 		cmd.stopStreamingFunc()
// 	}
// 	cmd.stopStreamingFunc = cancelFunc
// 	go cmd.getLogs(logStream, errStream)
	return nil
}

// func (cmd RollbackCommand) getLogs(logStream <-chan sharedaction.LogMessage, errStream <-chan error) {
// 	for {
// 		select {
// 		case logMessage, open := <-logStream:
// 			if !open {
// 				return
// 			}
// 			if logMessage.Staging() {
// 				cmd.UI.DisplayLogMessage(logMessage, false)
// 			}
// 		case err, open := <-errStream:
// 			if !open {
// 				return
// 			}
// 			_, ok := err.(actionerror.LogCacheTimeoutError)
// 			if ok {
// 				cmd.UI.DisplayWarning("timeout connecting to log server, no log will be shown")
// 			}
// 			cmd.UI.DisplayWarning("Failed to retrieve logs from Log Cache: " + err.Error())
// 		}
// 	}
// }
