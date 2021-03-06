//Package codepipeline provides gucumber integration tests support.
package codepipeline

import (
	"github.com/djannot/aws-sdk-go/awstesting/integration/smoke"
	"github.com/djannot/aws-sdk-go/service/codepipeline"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@codepipeline", func() {
		World["client"] = codepipeline.New(smoke.Session)
	})
}
