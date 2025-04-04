// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

//	A model representing the state of a resource within an account according to
//
// telemetry config.
type TelemetryConfiguration struct {

	//  The account ID which contains the resource managed in telemetry configuration.
	// An example of a valid account ID is 012345678901 .
	AccountIdentifier *string

	//  The timestamp of the last change to the telemetry configuration for the
	// resource. For example, 1728679196318 .
	LastUpdateTimeStamp *int64

	//  The identifier of the resource, for example i-0b22a22eec53b9321 .
	ResourceIdentifier *string

	//  Tags associated with the resource, for example { Name: "ExampleInstance",
	// Environment: "Development" } .
	ResourceTags map[string]string

	//  The type of resource, for example AWS::EC2::Instance .
	ResourceType ResourceType

	//  The configuration state for the resource, for example { Logs: NotApplicable;
	// Metrics: Enabled; Traces: NotApplicable; } .
	TelemetryConfigurationState map[string]TelemetryState

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
