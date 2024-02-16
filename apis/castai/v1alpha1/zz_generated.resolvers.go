/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AutoScaler.
func (mg *AutoScaler) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &EksClusterIdList{},
			Managed: &EksClusterId{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterIDRef,
		Selector:     mg.Spec.InitProvider.ClusterIDSelector,
		To: reference.To{
			List:    &EksClusterIdList{},
			Managed: &EksClusterId{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterID")
	}
	mg.Spec.InitProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EksUserArn.
func (mg *EksUserArn) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &EksClusterList{},
			Managed: &EksCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterIDRef,
		Selector:     mg.Spec.InitProvider.ClusterIDSelector,
		To: reference.To{
			List:    &EksClusterList{},
			Managed: &EksCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterID")
	}
	mg.Spec.InitProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EvictorAdvancedConfig.
func (mg *EvictorAdvancedConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &EksClusterIdList{},
			Managed: &EksClusterId{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterIDRef,
		Selector:     mg.Spec.InitProvider.ClusterIDSelector,
		To: reference.To{
			List:    &EksClusterIdList{},
			Managed: &EksClusterId{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterID")
	}
	mg.Spec.InitProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NodeConfiguration.
func (mg *NodeConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &EksClusterIdList{},
			Managed: &EksClusterId{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterIDRef,
		Selector:     mg.Spec.InitProvider.ClusterIDSelector,
		To: reference.To{
			List:    &EksClusterIdList{},
			Managed: &EksClusterId{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterID")
	}
	mg.Spec.InitProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NodeConfigurationDefault.
func (mg *NodeConfigurationDefault) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &EksClusterIdList{},
			Managed: &EksClusterId{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConfigurationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ConfigurationIDRef,
		Selector:     mg.Spec.ForProvider.ConfigurationIDSelector,
		To: reference.To{
			List:    &NodeConfigurationList{},
			Managed: &NodeConfiguration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConfigurationID")
	}
	mg.Spec.ForProvider.ConfigurationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConfigurationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterIDRef,
		Selector:     mg.Spec.InitProvider.ClusterIDSelector,
		To: reference.To{
			List:    &EksClusterIdList{},
			Managed: &EksClusterId{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterID")
	}
	mg.Spec.InitProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConfigurationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ConfigurationIDRef,
		Selector:     mg.Spec.InitProvider.ConfigurationIDSelector,
		To: reference.To{
			List:    &NodeConfigurationList{},
			Managed: &NodeConfiguration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ConfigurationID")
	}
	mg.Spec.InitProvider.ConfigurationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ConfigurationIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NodeTemplate.
func (mg *NodeTemplate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &EksClusterIdList{},
			Managed: &EksClusterId{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConfigurationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ConfigurationIDRef,
		Selector:     mg.Spec.ForProvider.ConfigurationIDSelector,
		To: reference.To{
			List:    &NodeConfigurationList{},
			Managed: &NodeConfiguration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConfigurationID")
	}
	mg.Spec.ForProvider.ConfigurationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConfigurationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterIDRef,
		Selector:     mg.Spec.InitProvider.ClusterIDSelector,
		To: reference.To{
			List:    &EksClusterIdList{},
			Managed: &EksClusterId{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterID")
	}
	mg.Spec.InitProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConfigurationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ConfigurationIDRef,
		Selector:     mg.Spec.InitProvider.ConfigurationIDSelector,
		To: reference.To{
			List:    &NodeConfigurationList{},
			Managed: &NodeConfiguration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ConfigurationID")
	}
	mg.Spec.InitProvider.ConfigurationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ConfigurationIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RebalancingJob.
func (mg *RebalancingJob) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &EksClusterIdList{},
			Managed: &EksClusterId{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RebalancingScheduleID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RebalancingScheduleIDRef,
		Selector:     mg.Spec.ForProvider.RebalancingScheduleIDSelector,
		To: reference.To{
			List:    &RebalancingScheduleList{},
			Managed: &RebalancingSchedule{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RebalancingScheduleID")
	}
	mg.Spec.ForProvider.RebalancingScheduleID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RebalancingScheduleIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterIDRef,
		Selector:     mg.Spec.InitProvider.ClusterIDSelector,
		To: reference.To{
			List:    &EksClusterIdList{},
			Managed: &EksClusterId{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterID")
	}
	mg.Spec.InitProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RebalancingScheduleID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RebalancingScheduleIDRef,
		Selector:     mg.Spec.InitProvider.RebalancingScheduleIDSelector,
		To: reference.To{
			List:    &RebalancingScheduleList{},
			Managed: &RebalancingSchedule{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RebalancingScheduleID")
	}
	mg.Spec.InitProvider.RebalancingScheduleID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RebalancingScheduleIDRef = rsp.ResolvedReference

	return nil
}
