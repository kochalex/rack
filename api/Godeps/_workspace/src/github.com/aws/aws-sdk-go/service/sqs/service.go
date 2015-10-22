// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package sqs

import (
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws"
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/request"
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/service"
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/service/serviceinfo"
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/aws/aws-sdk-go/internal/protocol/query"
	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/aws/aws-sdk-go/internal/signer/v4"
)

// Welcome to the Amazon Simple Queue Service API Reference. This section describes
// who should read this guide, how the guide is organized, and other resources
// related to the Amazon Simple Queue Service (Amazon SQS).
//
// Amazon SQS offers reliable and scalable hosted queues for storing messages
// as they travel between computers. By using Amazon SQS, you can move data
// between distributed components of your applications that perform different
// tasks without losing messages or requiring each component to be always available.
//
// Helpful Links:  Current WSDL (2012-11-05) (http://queue.amazonaws.com/doc/2012-11-05/QueueService.wsdl)
// Making API Requests (http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/MakingRequestsArticle.html)
// Amazon SQS product page (http://aws.amazon.com/sqs/) Using Amazon SQS Message
// Attributes (http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSMessageAttributes.html)
// Using Amazon SQS Dead Letter Queues (http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html)
// Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/rande.html#sqs_region)
//
//
// We also provide SDKs that enable you to access Amazon SQS from your preferred
// programming language. The SDKs contain functionality that automatically takes
// care of tasks such as:
//
//   Cryptographically signing your service requests Retrying requests Handling
// error responses
//
// For a list of available SDKs, go to Tools for Amazon Web Services (http://aws.amazon.com/tools/).
type SQS struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new SQS client.
func New(config *aws.Config) *SQS {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "sqs",
			APIVersion:  "2012-11-05",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &SQS{service}
}

// newRequest creates a new request for a SQS operation and runs any
// custom request initialization.
func (c *SQS) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}