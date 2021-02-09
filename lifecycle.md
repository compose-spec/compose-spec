# The Compose Specification - Lifecycle
version: 3.9

## Introduction

This document describes how Compose implementations MUST interact with platforms to apply the Compose 
application model and deploy applications.

### Project

Compose implementations are responsible for managing application deployment and management, including updates and
removal. The Compose Specification uses the term "_project_" for all resources allocated when deploying a 
Compose file. The actual mechanism used to identify such resources is not specified by the specification and can 
be platform specific.

## Up

The "up" lifecycle defines initial creation or update of a project from a Compose file. It handles changes made 
to the Compose file from the initial creation of the project.

### Resources

Compose implementations MUST first check that referenced Volumes, Networks, Configs and Secrets exist on the Platform. 

- Resources declared as `external` MUST exist on the Platform prior to deployment and MUST be addressable using the 
 `name` parameter. Compose implementations MUST NOT manage the lifecycle of external resources. 
- Compose implementations MUST identify and reuse resource previously created on the platform for the project. If 
  resources do not exist, implementations MUST create them.

A Compose implementation MAY NOT check existing resources do match their definition by Compose file. For example, as 
a Compose file is updated to change Volume specification by editing `driver_opts`, a Compose Implementation MAY NOT
warn user that the identified existing volume has been created with distinct options.

### Convergence

Compose implementations MUST manage all containers on the Platform which form part of the project. As some platforms
offer high-level abstraction over containers, a Compose implementation MAY interact with such abstractions and assume 
actual containers running on platform are consistent.

- containers that form part of the project MUST be associated with a service defined by project's Compose file. 
  Compose implementations MUST implement mechanisms to check that containers match the configuration defined by 
  service in the Compose file. Containers matching the service specification are considered to be "_running_". 
  Containers with a distinct configuration are considered to be "_diverged_" from the specification
- existing containers found for project but matching none of the configured service are considered to be "_orphaned_"


Compose implementations SHOULD create containers for each service by following their dependency order, as defined by
`depends_on`, `links` or an attached network reference `service:name`. Compose implementations SHOULD wait for all
dependant containers to be created on Platform, but it MAY not wait for container to return a valid `healthcheck`
response.

Compose implementations MUST remove diverged containers and replace them with new containers matching service 
configuration. The replacement of containers SHOULD follow the service update policy as defined by the 
[`deploy`](deploy.md) section.

Compose implementations SHOULD remove Orphaned containers.

Failure to create a service MUST NOT remove other successfully deployed services and resources. If it occurs, the
project MUST be marked as in an unhealthy state. The user MUST be warned about project creation failure.


## Down

The "down" lifecycle defines project removal. A Compose file is NOT REQUIRED for this lifecycle to be ran, but user
must provide a projet reference which can be used to find all project resources on the platform.

### Service Removal

As for "up" convergence, Compose implementations MUST search for existing containers on the platform created for the
project, and start a graceful shutdown sequence (`stop_signal`). Compose implementations MUST wait for a delay defined 
by  `stop_grace_period` for containers to handle the termination request. If they don't, Compose implementations MUST
force remove the containers.

Compose implementations SHOULD remove containers for service by following their dependency order. Most dependent containers SHOULD be stopped first, and their dependencies SHOULD be removed once they are fully removed.

### Resources Removal

Networks, Configs and Secrets created by Compose implementations (i.e not set as `external`) MUST be removed when the
project is removed. Volumes MUST NOT be removed unless the user explicitly requested for them to be.