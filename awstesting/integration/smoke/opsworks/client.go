//Package opsworks provides gucumber integration tests support.
package opsworks

import (
	"github.com/djannot/aws-sdk-go/awstesting/integration/smoke"
	"github.com/djannot/aws-sdk-go/service/opsworks"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@opsworks", func() {
		World["client"] = opsworks.New(smoke.Session)
	})
}
