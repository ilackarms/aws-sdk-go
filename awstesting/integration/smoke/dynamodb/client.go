//Package dynamodb provides gucumber integration tests support.
package dynamodb

import (
	"github.com/djannot/aws-sdk-go/awstesting/integration/smoke"
	"github.com/djannot/aws-sdk-go/service/dynamodb"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@dynamodb", func() {
		World["client"] = dynamodb.New(smoke.Session)
	})
}
