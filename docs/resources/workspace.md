---
Title: "Workspace Resource"
Description: |-
    Creating the workspace resource.
---

# Workspace

Group namespaces together into a workspace for more efficient management using this Terraform module.

The workspace is an organizational tool that helps you monitor and manage your Kubernetes namespaces within and across clusters.
Workspaces allow you to organize your managed namespaces into logical groups across clusters, perhaps to align with development projects.

To create a workspace, you must have `organization.edit` permissions in Tanzu Mission Control.
For more information, please refer [create a Workspace.][workspace]

[workspace]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-473F0C1F-DA60-4B04-9783-E6057A405604.html

## Example Usage

```terraform
# Create Tanzu Mission Control workspace
resource "tanzu-mission-control_workspace" "workspace" {
  name = "tf-workspace-test"

  meta {
    description = "Create workspace through terraform"
    labels = {
      "key1" : "value1",
      "key2" : "value2"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource
