# Services top-level element

A service is an abstract definition of a computing resource within an application which can be scaled or replaced
independently from other components. Services are backed by a set of containers, run by the platform
according to replication requirements and placement constraints. As services are backed by containers, they are defined
by a Docker image and set of runtime arguments. All containers within a service are identically created with these
arguments.

A Compose file must declare a `services` top-level element as a map whose keys are string representations of service names,
and whose values are service definitions. A service  definition contains the configuration that is applied to each
service container.

Each service may also include a `build` section, which defines how to create the Docker image for the service.
Compose supports building docker images using this service definition. If not used, the `build` section is ignored and the Compose file is still considered valid. Build support is an optional aspect of the Compose Specification, and is
described in detail in the [Compose Build Specification](build.md) documentation.

Each service defines runtime constraints and requirements to run its containers. The `deploy` section groups
these constraints and allows the platform to adjust the deployment strategy to best match containers' needs with
available resources. Deploy support is an optional aspect of the Compose Specification, and is
described in detail in the [Compose Deploy Specification](deploy.md) documentation.
If not implemented the `deploy` section is ignored and the Compose file is still considered valid.

A service definition supports all attributes defined in the [Container Specification](container_spec.md),
as well as the following service-specific attributes.

## deploy

`deploy` specifies the configuration for the deployment and lifecycle of services, as defined [in the Compose Deploy Specification](deploy.md).

## develop

[![Compose v2.22.0](https://img.shields.io/badge/compose-v2.22.0-blue?style=flat-square)](https://github.com/docker/compose/releases/v2.22.0)


`develop` specifies the development configuration for maintaining a container in sync with source, as defined in the [Development Section](develop.md).

## profiles

`profiles` defines a list of named profiles for the service to be enabled under. If unassigned, the service is always started but if assigned, it is only started if the profile is activated.

If present, `profiles` follow the regex format of `[a-zA-Z0-9][a-zA-Z0-9_.-]+`.

```yaml
services:
  frontend:
    image: frontend
    profiles: ["frontend"]

  phpmyadmin:
    image: phpmyadmin
    depends_on:
      - db
    profiles:
      - debug
```

## scale

`scale` specifies the default number of containers to deploy for this service.
When both are set, `scale` must be consistent with the `replicas` attribute in the [Deploy Specification](deploy.md#replicas).
