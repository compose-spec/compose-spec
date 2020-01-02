# Compose Specification

The Compose specification establish a standard for definition of multi-container, platform-agnostic applications. The specification can be found [here](spec.md).

![logo](logo.jpg)   

## Table of Contents

Additional documentation about how this group operates:
* [Governance](GOVERNANCE.md)
* [Contribution Guidelines](GUIDELINES.md)
* [Implementations](#Implementations)
* [Releases](https://github.com/docker/compose-spec/releases)

## Use cases

To provide more context on the Compose specification the following section gives example use cases for each part of the spec.

### Developement tools

Developers can rely on a Compose file to specify a container-based application that will run as local set of containers on a single engine. The Compose implementation in this specific scenario could offer some specific features (local volume  binding, live-reload) to better address specific development needs, but the global application definition and Compose abstraction is the same used for other use cases. Platform features expected by the specification (like configs and secrets) can be mocked by local resources.

### Kubernetes deployment

Kubernetes container orchestration relies on a set of abstract concepts and related APIs to manage services, deployments and relate lifecycle. This offers flexibility to address many use-cases, but also make the 90% simplest usages pretty complex with hundred-lines yaml files under end-user charge. As demonstrated by projects [Kompose)(https://github.com/kubernetes/kompose) and [Compose on kubernete](https://github.com/docker/compose-on-kubernetes) the simpler Compose model can be translated into Kubernetes API payloads to cover most usages, and let user rely on the extact same compose file for local development.

### Cloud providers

Cloud providers do offer proprietary container hosting solutions, based on internal orchestrators and proprietary APIs. For larger adoption and an easier first-time experience they can rely on the Compose spec to support application deployment without the need to write custom configuration files. Provider-specific features can later be enabled using either custom extensions or a dedicated configuration file along side the compose file.


## Contributing

Development happens on GitHub for the spec. Issues are used for bugs and actionable items and longer discussions can happen on the [mailing list](https://groups.google.com/forum/#!forum/compose-spec).

The specification and code is licensed under the Apache 2.0 license found in the [LICENSE](LICENSE) file.

## Implementations

* docker-compose
* Docker CLI (`stack` command)
* [Compose on kubernetes](https://github.com/docker/compose-on-kubernetes)
* [Kompose](https://github.com/kubernetes/kompose)



| Metadata |                  |
| -------- | ---------------: |
| Version  | 3.9              |
| Status   | Work in progress |
| Created  | 2020-01-02       |

## Table of Contents
  * [Status of this document](#status-of-this-document)
  * [Version](#version)
  * [Application model](#application-model)
  * [Services](#services)
  * [Networks](#networks)
  * [Volumes](#volumes)
  * [Configs](#configs)
  * [Secrets](#secrets)

