// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreaminfluxdb

import (
	"context"
	"errors"
	"github.com/aws/smithy-go/middleware"
	"slices"
	"strings"
	"testing"
)

var errTestReturnEarly = errors.New("errTestReturnEarly")

func captureMiddlewareStack(stack *middleware.Stack) func(*middleware.Stack) error {
	return func(inner *middleware.Stack) error {
		*stack = *inner
		return errTestReturnEarly
	}
}
func TestOpCreateDbClusterSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.CreateDbCluster(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpCreateDbInstanceSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.CreateDbInstance(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpCreateDbParameterGroupSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.CreateDbParameterGroup(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpDeleteDbClusterSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.DeleteDbCluster(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpDeleteDbInstanceSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.DeleteDbInstance(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpGetDbClusterSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.GetDbCluster(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpGetDbInstanceSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.GetDbInstance(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpGetDbParameterGroupSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.GetDbParameterGroup(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpListDbClustersSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.ListDbClusters(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpListDbInstancesSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.ListDbInstances(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpListDbInstancesForClusterSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.ListDbInstancesForCluster(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpListDbParameterGroupsSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.ListDbParameterGroups(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpListTagsForResourceSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.ListTagsForResource(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpTagResourceSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.TagResource(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpUntagResourceSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.UntagResource(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpUpdateDbClusterSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.UpdateDbCluster(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
func TestOpUpdateDbInstanceSRAOperationOrder(t *testing.T) {
	expect := []string{
		"OperationSerializer",
		"Retry",
		"ResolveAuthScheme",
		"GetIdentity",
		"ResolveEndpointV2",
		"Signing",
		"OperationDeserializer",
	}

	var captured middleware.Stack
	svc := New(Options{
		APIOptions: []func(*middleware.Stack) error{
			captureMiddlewareStack(&captured),
		},
	})
	_, err := svc.UpdateDbInstance(context.Background(), nil)
	if err != nil && !errors.Is(err, errTestReturnEarly) {
		t.Fatalf("unexpected error: %v", err)
	}

	var actual, all []string
	for _, step := range strings.Split(captured.String(), "\n") {
		trimmed := strings.TrimSpace(step)
		all = append(all, trimmed)
		if slices.Contains(expect, trimmed) {
			actual = append(actual, trimmed)
		}
	}

	if !slices.Equal(expect, actual) {
		t.Errorf("order mismatch:\nexpect: %v\nactual: %v\nall: %v", expect, actual, all)
	}
}
