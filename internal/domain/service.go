package domain

import v1 "k8s.io/api/apps/v1"

// Service holds information about a running service
type Service struct {
	Name             string `json:"name"`
	ApplicationGroup string `json:"applicationGroup"`
	RunningPodsCount int32  `json:"runningPodsCount"`
} // @name Service

// NewServiceFromDpl creates a service from a K8S deployment
func NewServiceFromDpl(dpl v1.Deployment) Service {
	return Service{
		Name:             dpl.GetName(),
		ApplicationGroup: dpl.GetLabels()["applicationGroup"],
		RunningPodsCount: *dpl.Spec.Replicas, // Defaults to 1, no need to check for nil pointer
	}
}
