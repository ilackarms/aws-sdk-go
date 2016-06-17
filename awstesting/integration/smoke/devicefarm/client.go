//Package devicefarm provides gucumber integration tests support.
package devicefarm

import (
	"github.com/djannot/aws-sdk-go/aws"
	"github.com/djannot/aws-sdk-go/awstesting/integration/smoke"
	"github.com/djannot/aws-sdk-go/service/devicefarm"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@devicefarm", func() {
		// FIXME remove custom region
		World["client"] = devicefarm.New(smoke.Session,
			aws.NewConfig().WithRegion("us-west-2"))
	})
}
