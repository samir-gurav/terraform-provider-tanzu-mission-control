---
Title: "Git Repository Resource"
Description: |-
    Creating the Git Repository resource.
---

# Git Repository

The `tanzu-mission-control_git_repository` resource allows you to add, update, and delete git repository to a particular scope through Tanzu Mission Control.

Git repositories are used to store kustomizations that will be synced to your cluster.
To add a repository, you must be associated with the cluster.admin or clustergroup.admin role.

[git-repository]: https://docs.vmware.com/en/VMware-Tanzu-Mission-Control/services/tanzumc-using/GUID-26C2D2F3-0E5C-4E56-B875-B7FB003267E4.html

## Git Repository Scope

In the Tanzu Mission Control resource hierarchy, there are two levels at which you can specify git repository resources:
- **object groups** - `cluster_group` block under `scope` sub-resource
- **Kubernetes objects** - `cluster` block under `scope` sub-resource

**Note:**
The scope parameter is mandatory in the schema and the user needs to add one of the defined scopes to the script for the provider to function.
Only one scope per resource is allowed.

## Cluster group scoped Git Repository

### Example Usage

```terraform
# Create Tanzu Mission Control git repository with attached set as default value.
resource "tanzu-mission-control_git_repository" "cluster_group_git_repository" {
  name = "tf-git-repository-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster_group {
      name = "default" # Required
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    url                = "testGitRepositoryURL" # Required
    secret_ref         = "testSourceSecret"
    interval           = "10m"    # Default: 5m
    git_implementation = "GO_GIT" # Default: GO_GIT
    ref {
      branch = "testBranchName"
      tag    = "testTag"
      semver = "testSemver"
      commit = "testCommit"
    }
  }
}
```

## Cluster scoped Git Repository

### Example Usage

```terraform
# Create Tanzu Mission Control git repository with attached set as default value.
resource "tanzu-mission-control_git_repository" "cluster_git_repository" {
  name = "tf-git-repository-name" # Required

  namespace_name = "tf-namespace" #Required

  scope {
    cluster {
      name                    = "testcluster" # Required
      provisioner_name        = "attached"    # Default: attached
      management_cluster_name = "attached"    # Default: attached
    }
  }

  meta {
    description = "Create namespace through terraform"
    labels      = { "key" : "value" }
  }

  spec {
    url                = "testGitRepositoryURL" # Required
    secret_ref         = "testSourceSecret"
    interval           = "10m"    # Default: 5m
    git_implementation = "GO_GIT" # Default: GO_GIT
    ref {
      branch = "testBranchName"
      tag    = "testTag"
      semver = "testSemver"
      commit = "testCommit"
    }
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the Repository.
- `namespace_name` (String) Name of Namespace.
- `scope` (Block List, Min: 1, Max: 1) Scope for the git repository, having one of the valid scopes: cluster, cluster_group. (see [below for nested schema](#nestedblock--scope))
- `spec` (Block List, Min: 1, Max: 1) Spec for the Repository. (see [below for nested schema](#nestedblock--spec))

### Optional

- `meta` (Block List, Max: 1) Metadata for the resource (see [below for nested schema](#nestedblock--meta))

### Read-Only

- `id` (String) The ID of this resource.
- `status` (Map of String) Status for the Repository.

<a id="nestedblock--scope"></a>
### Nested Schema for `scope`

Optional:

- `cluster` (Block List, Max: 1) The schema for cluster full name (see [below for nested schema](#nestedblock--scope--cluster))
- `cluster_group` (Block List, Max: 1) The schema for cluster group full name (see [below for nested schema](#nestedblock--scope--cluster_group))

<a id="nestedblock--scope--cluster"></a>
### Nested Schema for `scope.cluster`

Required:

- `name` (String) Name of this cluster

Optional:

- `management_cluster_name` (String) Name of the management cluster
- `provisioner_name` (String) Provisioner of the cluster


<a id="nestedblock--scope--cluster_group"></a>
### Nested Schema for `scope.cluster_group`

Required:

- `name` (String) Name of the cluster group



<a id="nestedblock--spec"></a>
### Nested Schema for `spec`

Required:

- `url` (String) URL of the git repository. Repository URL should begin with http, https, or ssh

Optional:

- `git_implementation` (String) GitImplementation specifies which client library implementation to use. go-git is the default git implementation.
- `interval` (String) Interval at which to check gitrepository for updates. This is the interval at which Tanzu Mission Control will attempt to reconcile changes in the repository to the cluster. A sync interval of 0 would result in no future syncs. If no value is entered, a default interval of 5 minutes will be applied as `5m`.
- `ref` (Block List, Max: 1) Reference specifies git reference to resolve. (see [below for nested schema](#nestedblock--spec--ref))
- `secret_ref` (String) Reference to the secret. Repository credential.

<a id="nestedblock--spec--ref"></a>
### Nested Schema for `spec.ref`

Optional:

- `branch` (String) Branch from git to checkout. When branch is given, then that branch from the git repository will be checked out. If the given branch doesn’t exist in the git repository, then adding the git repository will fail. If no branch is given, the `master` branch will be used.
- `commit` (String) Commit SHA to checkout. Takes precedence over all other reference fields. When git_implementation is `GO_GIT`, this can be combined with branch to shallow clone branch in which the commit is expected to exist.
- `semver` (String) SemVer expression to checkout from git tags. Takes precedence over tag. When semver is given, then the latest tag matching that semver will be checked out from the git repository. If no tag in the git repository matches semver, then adding the git repository will fail. If semver is given, tag and branch will be ignored if they are populated.
- `tag` (String) Tag from git to checkout. Takes precedence over branch. When a tag is given, that tag from the git repository will be checked out. If the given tag doesn’t exist in the git repository, then adding the git repository will fail. If both tag and branch are given, tag overrides branch and the branch value will be ignored.



<a id="nestedblock--meta"></a>
### Nested Schema for `meta`

Optional:

- `annotations` (Map of String) Annotations for the resource
- `description` (String) Description of the resource
- `labels` (Map of String) Labels for the resource

Read-Only:

- `resource_version` (String) Resource version of the resource
- `uid` (String) UID of the resource