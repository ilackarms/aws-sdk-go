//Package elastictranscoder provides gucumber integration tests support.
package elastictranscoder

import (
	"github.com/djannot/aws-sdk-go/awstesting/integration/smoke"
	"github.com/djannot/aws-sdk-go/service/elastictranscoder"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@elastictranscoder", func() {
		World["client"] = elastictranscoder.New(smoke.Session)
	})
}
