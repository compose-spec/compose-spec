# The Compose Specification 
version: 3.9

## Status of this document

This document specifies the Compose file format used to define multi-containers applications. Distribution of this document is unlimited.

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED",  "MAY", and "OPTIONAL" in this document are to be interpreted as described in [RFC 2119](https://tools.ietf.org/html/rfc2119).

## Version

A top-level version property is required by the specification. Version MUST be 3.x or later, legacy docker-compose 1.x and 2.x are not included as part of this specification. Implementations MAY accept such legacy formats for compatibility purposes.

The specification format follows [Semantic Versioning](https://semver.org), which means that the file format is backward compatible within a major version set. As the specification evolves, minor versions MAY introduce new elements and MAY deprecate others for removal in the next major version. Implementations MUST support features until they are removed in the next major release.

Implementations MAY ignore attributes used in a configuration file that are not supported by the declared version, whenever then are valid for a more recent version. If they do, a warning message MUST inform user. 


## Application model

The Compose specification allows one to define a platform-agnostic container based application. Such an application is designed as a set of containers which have to both run together with adequate shared resources and communication channels.

Computing components of an application are defined as [Services](#Services), which is an abstract concept implemented on platforms by running containers with using the same container image and configuration but replicated on one or more times.

Services communicate with each other through [Networks](#Networks). Those, within the specification, are just an abstraction of platform capability to establish an IP route between containers within services connected together. Low-level, platform-specific networking options are grouped into the Network definition and MAY be partially implemented on some platforms.

Services store and share persistent data into [Volumes](#Volumes). The specification describes such a persistent data as a high-level filesystem mount with global options, actual platform-specific implementation details are grouped into the Volumes definition and MAY be partially implemented on some platform.

Some services require configuration data that is dependent on the runtime or platform. For this, the specification defines a dedicated concept: [Configs](Configs). From a Service container point of view Configs are very comparable to Volumes, in that they are files mounted into the container, but the actual definition involves distinct platform resources and services, which are abstracted by this type.

A [Secrets](#Secrets) is a specific flavour of configuration data for sensible data that SHOULD not be exposed without security considerations. They are exposed to services as files mounted into their containers but the platform-specific resources to provide sensible data are specific enough to deserve a distinct concept and definition within the Compose specification.

Distinction within Volumes, Configs and Secret allows to offer a comparable abstraction at service level, but cover the specific configuration of adequate platform resources for well identified data usages.



## Services

A Service is an abstract definition of a computing resource within an application which can be scaled/replaced independently from other components. Services are actually backed by a set of containers, run by the platform according to replication requirements and placement constraints. Being backed by containers, Services are defined by a Docker image and set of runtime arguments. All containers within a service are identically created from those arguments.

Service also includes a Build section, defining how to create the Docker image for a service. Support by Compose implementations to build docker images according to this service definition is OPTIONAL. If not implemented the Build section SHOULD be ignored and the Compose file MUST still be considered valid. 

Build support is an OPTIONAL aspect of the Compose specification, and is described in detail [here](build.md)

Service define runtime constraints and requirement to run service's containers. The Deploy section do group those and allow plaform to adjust deployment strategy to best match containers needs with available resources.

Deploy support is an OPTIONAL aspect of the Compose specification, and is described in detail [here](deploy.md). If not implemented the Deploy section SHOULD be ignored and the Compose file MUST still be considered valid. 


*TODO* describe configuration attributes 


## Networks

Networks are communication channels between services managed by the platform. The networking model exposed to a service is limited to a simple IP connection with target services and external resources, while the Network definition allows to fine-tune the actual implementation provided by the platform.

*TODO* describe configuration attributes 

## Volumes

Volumes are persistent data stored implemented by the platform. The Compose specification offers a neutral abstraction for services to mount volumes, and configuration parameters to allocate them on infrastructure.


*TODO* describe configuration attributes 


## Configs

Configs allow services to adapt their behaviour without the need to rebuild a Docker image. Configs are comparable to Volumes from a service point of view as they are mounted into service's containers filesystem. The actual implementation detail to get configuration provided by the platform can be set from the Configuration definition. 

*TODO* describe configuration attributes 

## Secrets

Secrets are a flavour of Configs focussing on sensitive data, with specific constraint for this usage. As the platform implementation may significally differ from Configs, dedicated Secrets section allows to configure the related resources.

*TODO* describe configuration attributes 
