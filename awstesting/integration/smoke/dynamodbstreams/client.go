//Package dynamodbstreams provides gucumber integration tests support.
package dynamodbstreams

import (
	"github.com/djannot/aws-sdk-go/awstesting/integration/smoke"
	"github.com/djannot/aws-sdk-go/service/dynamodbstreams"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@dynamodbstreams", func() {
		World["client"] = dynamodbstreams.New(smoke.Session)
	})
}
