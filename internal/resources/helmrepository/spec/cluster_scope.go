/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	helmrepoclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrepository"
)

func FlattenSpecForClusterScope(spec *helmrepoclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmRepositorySpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	flattenSpecData[URLKey] = spec.URL

	return []interface{}{flattenSpecData}
}
