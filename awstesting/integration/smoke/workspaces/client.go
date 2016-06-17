//Package workspaces provides gucumber integration tests support.
package workspaces

import (
	"github.com/djannot/aws-sdk-go/awstesting/integration/smoke"
	"github.com/djannot/aws-sdk-go/service/workspaces"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@workspaces", func() {
		World["client"] = workspaces.New(smoke.Session)
	})
}
