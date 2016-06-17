//Package configservice provides gucumber integration tests support.
package configservice

import (
	"github.com/djannot/aws-sdk-go/awstesting/integration/smoke"
	"github.com/djannot/aws-sdk-go/service/configservice"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@configservice", func() {
		World["client"] = configservice.New(smoke.Session)
	})
}
