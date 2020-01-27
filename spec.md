# The Compose Specification

version: 3.9

## Table of Contents

- [Status of this document](#status-of-this-document)
- [The Compose application model](#the-compose-application-model)
- [Compose file](#compose-file)
- [Version top-level element](#version-top-level-element)
- [Services top-level element](#services-top-level-element)
- [Networks top-level element](#networks-top-level-element)
- [Volumes top-level element](#volumes-top-level-element)
- [Configs top-level element](#configs-top-level-element)
- [Secrets top-level element](#secrets-top-level-element)
- [Fragments](#fragments)
- [Extension](#extension)
- [Interpolation](#interpolation)


## Status of this document

This document specifies the Compose file format used to define multi-containers applications. Distribution of this document is unlimited.

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [RFC 2119](https://tools.ietf.org/html/rfc2119).

## The Compose application model

The Compose specification allows one to define a platform-agnostic container based application. Such an application is designed as a set of containers which have to both run together with adequate shared resources and communication channels.

Computing components of an application are defined as [Services](#Services-top-level-element). A Service is an abstract concept implemented on platforms by running the same container image (and configuration) one or more times.

Services communicate with each other through [Networks](#Networks-top-level-element). In this specification, a Network is a platform capability abstraction to establish an IP route between containers within services connected together. Low-level, platform-specific networking options are grouped into the Network definition and MAY be partially implemented on some platforms.

Services store and share persistent data into [Volumes](#Volumes-top-level-element). The specification describes such a persistent data as a high-level filesystem mount with global options. Actual platform-specific implementation details are grouped into the Volumes definition and MAY be partially implemented on some platforms.

Some services require configuration data that is dependent on the runtime or platform. For this, the specification defines a dedicated concept: [Configs](Configs-top-level-element). From a Service container point of view, Configs are comparable to Volumes, in that they are files mounted into the container. But the actual definition involves distinct platform resources and services, which are abstracted by this type.

A [Secret](#Secrets-top-level-element) is a specific flavour of configuration data for sensitive data that SHOULD NOT be exposed without security considerations. Secrets are made available to services as files mounted into their containers, but the platform-specific resources to provide sensitive data are specific enough to deserve a distinct concept and definition within the Compose specification.

Distinction within Volumes, Configs and Secret allows implementations to offer a comparable abstraction at service level, but cover the specific configuration of adequate platform resources for well identified data usages.

A **Project** is an individual deployment of an application specification on a platform. A project's name is used to group
resources together and isolate them from other applications or other installation of the same Compose specified application with distinct parameters. A Compose implementation creating resources on a platform MUST prefix resource names by project and
set the label `com.docker.compose.project`.

### Illustrative example

The following example illustrates Compose specification concepts with a concrete example application. The example is non-normative.

Consider an application split into a frontend web application and a backend service.

The frontend is configured at runtime with an HTTP configuration file managed by infrastructure, providing an external domain name, and an HTTPS server certificate injected by the platform's secured secret store.

The backend stores data in a persistent volume.

Both services communicate with each other on an isolated back-tier network, while frontend is also connected to a front-tier network and exposes port 443 for external usage.

```
(External user) --> 443 [frontend network]
                            |
                  +--------------------+
                  |  frontend service  |...ro...<HTTP configuration>
                  |      "webapp"      |...ro...<server certificate> #secured
                  +--------------------+
                            |
                        [backend network]
                            |
                  +--------------------+
                  |  backend service   |  r+w   ___________________
                  |     "database"     |=======( persistent volume )
                  +--------------------+        \_________________/
```


The example application is composed of the following parts:
- 2 services, backed by Docker images: `webapp` and `database`
- 1 secret (HTTPS certificate), injected into the frontend
- 1 configuration (HTTP), injected into the frontend
- 1 persistent volume, attached to the backend
- 2 networks

```yml
version: "3"
services:
  frontend:
    image: awesome/webapp
    ports:
      - "443:8043"
    networks:
      - front-tier
      - back-tier
    configs:
      - httpd-config
    secrets:
      - server-certificate

  backend:
    image: awesome/database
    volumes:
      - db-data:/etc/data
    networks:
      - back-tier

volumes:
  db-data:
    driver: flocker
    driver_opts:
      size: "10GiB"

config:
  httpd-config:
    external: true

secrets:
  server-certificate:
    external: true

networks:
  # The presence of these objects is sufficient to define them
  front-tier: {}
  back-tier: {}
```

This example illustrates the distinction between volumes, configs and secrets. While all of them are all exposed
to service containers as mounted files or directories, only a volume can be configured for read+write access.
Secrets and configs are read-only. The volume configuration allows you to select a volume driver and pass driver options
to tweak volume management according to the actual infrastructure. Configs and Secrets rely on platform services,
and are declared `external` as they are not managed as part of the application lifecycle: the Compose implementation
will use a platform-specific lookup mechanism to retrieve runtime values.

## Compose file

The Compose file is a [YAML](http://yaml.org/) file defining
[version](#version) (REQUIRED),
[services](#service-top-level-element) (REQUIRED),
[networks](#network-top-level-element),
[volumes](#volume-top-level-element),
[configs](#configs-top-level-element) and
[secrets](#secrets-top-level-element).
The default path for a Compose file is `compose.yaml` (preferred) or `compose.yml` in working directory.
Compose implementations SHOULD also support `docker-compose.yaml` and `docker-compose.yml` for backward compatibility. 
If both files exist, Compose implementations MUST prefer canonical `compose.yaml` one. 

Multiple Compose files can be combined together to define the application model. The combination of YAML files
MUST be implemented by appending/overriding YAML elements based on Compose file order set by the user. Simple
attributes and maps get overridden by the highest order Compose file, lists get merged by appending. Relative
paths MUST be resolved based on the **first** Compose file's parent folder, whenever complimentary files being
merged are hosted in other folders.

As some Compose file elements can both be expressed as single strings or complex objects, merges MUST apply to
the expanded form.

## Version top-level element

A top-level version property is required by the specification. Version MUST be 3.x or later, legacy docker-compose 1.x and 2.x are not included as part of this specification. Implementations MAY accept such legacy formats for compatibility purposes.

The specification format follows [Semantic Versioning](https://semver.org), which means that the file format is backward compatible within a major version set. As the specification evolves, minor versions MAY introduce new elements and MAY deprecate others for removal in the next major version.

Implementations MAY ignore attributes used in a configuration file that are not supported by the declared version, whenever they are valid for a more recent version. If they do, a warning message MUST inform the user.

## Services top-level element

A Service is an abstract definition of a computing resource within an application which can be scaled/replaced
independently from other components. Services are backed by a set of containers, run by the platform
according to replication requirements and placement constraints. Being backed by containers, Services are defined
by a Docker image and set of runtime arguments. All containers within a service are identically created with these
arguments.

A Compose file MUST declare a `services` root element as a map whose keys are string representations of service names,
and whose values are service definitions. A service  definition contains the configuration that is applied to each
container started for that service.

Each service MAY also include a Build section, which defines how to create the Docker image for the service.
Compose implementations MAY support building docker images using this service definition. If not implemented
the Build section SHOULD be ignored and the Compose file MUST still be considered valid.

Build support is an OPTIONAL aspect of the Compose specification, and is described in detail [here](build.md)

Each Service defines runtime constraints and requirements to run its containers. The `deploy` section groups 
these constraints and allows the platform to adjust the deployment strategy to best match containers' needs with
available resources.

Deploy support is an OPTIONAL aspect of the Compose specification, and is described in detail [here](deploy.md). If
not implemented the Deploy section SHOULD be ignored and the Compose file MUST still be considered valid.

### deploy

`deploy` specifies the configuration for the deployment and lifecycle of services, as defined [here](deploy.md).

### build

`build` specifies the build configuration for creating container image from source, as defined [here](build.md).

### cap_add

`cap_add` specifies additional container [capabilities](http://man7.org/linux/man-pages/man7/capabilities.7.html)
as strings.

```
cap_add:
  - ALL
```

### cap_drop

`cap_drop` specifies container [capabilities](http://man7.org/linux/man-pages/man7/capabilities.7.html) to drop
as strings.

```
cap_drop:
  - NET_ADMIN
  - SYS_ADMIN
```

### cgroup_parent

`cgroup_parent` specifies an OPTIONAL parent [cgroup](http://man7.org/linux/man-pages/man7/cgroups.7.html) for the container.

```
cgroup_parent: m-executor-abcd
```

### command

`command` overrides the the default command declared by the container image (i.e. by Dockerfile's `CMD`).

```
command: bundle exec thin -p 3000
```

The command can also be a list, in a manner similar to [Dockerfile](https://docs.docker.com/engine/reference/builder/#cmd):

```
command: [ "bundle", "exec", "thin", "-p", "3000" ]
```

### configs

`configs` grant access to configs on a per-service basis using the per-service `configs`
configuration. Two different syntax variants are supported.

Compose implementations MUST report an error if config doesn't exist on platform or isn't defined in the
[`configs`](#configs-top-level-element) section of this Compose file.

There are two syntaxes defined for configs. To remain compliant to this specification, an implementation
MUST support both syntaxes. Implementations MUST allow use of both short and long syntaxes within the same document.

#### Short syntax

The short syntax variant only specifies the config name. This grants the
container access to the config and mounts it at `/<config_name>`
within the container. The source name and destination mount point are both set
to the config name.

The following example uses the short syntax to grant the `redis` service
access to the `my_config` and `my_other_config` configs. The value of
`my_config` is set to the contents of the file `./my_config.txt`, and
`my_other_config` is defined as an external resource, which means that it has
already been defined in the platform. If the external config does not exist,
the deployment MUST fail.

```yml
version: "3"
services:
  redis:
    image: redis:latest
    configs:
      - my_config
configs:
  my_config:
    file: ./my_config.txt
  my_other_config:
    external: true
```

#### Long syntax

The long syntax provides more granularity in how the config is created within the service's task containers.

- `source`: The name of the config as it exists in the platform.
- `target`: The path and name of the file to be mounted in the service's
  task containers. Defaults to `/<source>` if not specified.
- `uid` and `gid`: The numeric UID or GID that owns the mounted config file
  within the service's task containers. Default value when not specified is USER running container.
- `mode`: The [permissions](http://permissions-calculator.org/) for the file that is mounted within the service's
  task containers, in octal notation. Default value is world-readable (`0444`).
  Writable bit MUST be ignored. The executable bit can be set.

The following example sets the name of `my_config` to `redis_config` within the
container, sets the mode to `0440` (group-readable) and sets the user and group
to `103`. The `redis` service does not have access to the `my_other_config`
config.

```yml
version: "3"
services:
  redis:
    image: redis:latest
    configs:
      - source: my_config
        target: /redis_config
        uid: "103"
        gid: "103"
        mode: 0440
configs:
  my_config:
    external: true
```

You can grant a service access to multiple configs, and you can mix long and short syntax.

### container_name

`container_name` is a string that specifies a custom container name, rather than a generated default name.

```yml
container_name: my-web-container
```

Compose implementation MUST NOT scale a service beyond one container if the Compose file specifies a 
`container_name`. Attempting to do so MUST result in an error.

### credential_spec

`credential_spec` configures the credential spec for a managed service account.

Compose implementations that support services using Windows containers MUST support `file:` and
`registry:` protocols for credential_spec. Compose implementations MAY also support additional
protocols for custom use-cases.

The `credential_spec` must be in the format `file://<filename>` or `registry://<value-name>`.

```yml
credential_spec:
  file: my-credential-spec.json
```

When using `registry:`, the credential spec is read from the Windows registry on
the daemon's host. A registry value with the given name must be located in:

    HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Virtualization\Containers\CredentialSpecs

The following example loads the credential spec from a value named `my-credential-spec`
in the registry:

```yml
credential_spec:
  registry: my-credential-spec
```

#### Example gMSA configuration

When configuring a gMSA credential spec for a service, you only need
to specify a credential spec with `config`, as shown in the following example:

```yml
version: "3"
services:
  myservice:
    image: myimage:latest
    credential_spec:
      config: my_credential_spec

configs:
  my_credentials_spec:
    file: ./my-credential-spec.json|
```

### depends_on

`depends_on` expresses a startup and shutdown dependencies between services, Service dependencies cause the following
behaviors:

- Compose implementations MUST create services in dependency order. In the following
  example, `db` and `redis` are created before `web`.

- Compose implementations MUST remove services in dependency order. In the following
  example, `web` is removed before `db` and `redis`.

Simple example:

```yml
version: "3"
services:
  web:
    build: .
    depends_on:
      - db
      - redis
  redis:
    image: redis
  db:
    image: postgres
```

Compose implementations MUST guarantee dependency services have been started before
starting a dependent service.
Compose implementations MAY wait for dependency services to be "ready" before
starting a dependent service.

### devices

`devices` defines a list of device mappings for created containers.

```yml
devices:
  - "/dev/ttyUSB0:/dev/ttyUSB0"
```

### dns

`dns` defines custom DNS servers to set on the container network interface configuration. Can be a single value or a list.

```yml
dns: 8.8.8.8
```

```yml
dns:
  - 8.8.8.8
  - 9.9.9.9
```

### dns_search

`dns` defines custom DNS search domains to set on container network interface configuration. Can be a single value or a list.

```yml
dns_search: example.com
```

```yml
dns_search:
  - dc1.example.com
  - dc2.example.com
```

### domainname

`domainname` TODO

### entrypoint

`entrypoint` overrides the default entrypoint for the Docker image (i.e. `ENTRYPOINT` set by Dockerfile).
Compose implementations MUST clear out any default command on the Docker image - both `ENTRYPOINT` and `CMD` instruction
in the Dockerfile - when `entrypoint` is configured by a Compose file. If [`command](#command) is also set,
it is used as parameter to `entrypoint` as a replacement for Docker image's `CMD`

```yml
entrypoint: /code/entrypoint.sh
```

The entrypoint can also be a list, in a manner similar to
[Dockerfile](https://docs.docker.com/engine/reference/builder/#cmd):

```yml
entrypoint:
  - php
  - -d
  - zend_extension=/usr/local/lib/php/extensions/no-debug-non-zts-20100525/xdebug.so
  - -d
  - memory_limit=-1
  - vendor/bin/phpunit
```

### env_file

`env_file` adds environment variables to the container based on file content.

```yml
env_file: .env
```

`env_file` can also be a list. The files in the list MUST be processed from the top down. For the same variable
specified in two env files, the value from the last file in the list MUST stand.

```yml
env_file:
  - ./a.env
  - ./b.env
```

Relative path MUST be resolved from the Compose file's parent folder. As absolute paths prevent the Compose
file from being portable, Compose implementations SHOULD warn users when such a path is used to set `env_file`.

Environment variables declared in the [environment](#environment) section
MUST override these values &ndash; this holds true even if those values are
empty or undefined.

#### Env_file format

Each line in an env file MUST be in `VAR[=[VAL]]` format. Lines beginning with `#` MUST be ignored.
Blank lines MUST also be ignored.

The value of `VAL` is used as a raw string and not modified at all. If the value is surrounded by quotes
(as is often the case for shell variables), the quotes MUST be **included** in the value passed to containers
created by the Compose implementation.

`VAL` MAY be omitted, in such cases the variable value is empty string.
`=VAL` MAY be omitted, in such cases the variable is **unset**.

```bash
# Set Rails/Rack environment
RACK_ENV=development
VAR="quoted"
```

### environment

`environment` defines environment variables set in the container. `environment` can use either an array or a
map. Any boolean values; true, false, yes, no, MUST be enclosed in quotes to ensure
they are not converted to True or False by the YAML parser.

Environment variables MAY be declared by a single key (no value to equals sign). In such a case Compose
implementations SHOULD rely on some user interaction to resolve the value. If they do not, the variable
is unset and will be removed from the service container environment.

Map syntax:

```yml
environment:
  RACK_ENV: development
  SHOW: "true"
  USER_INPUT:
```

Array syntax:

```yml
environment:
  - RACK_ENV=development
  - SHOW=true
  - USER_INPUT
```

When both `env_file` and `environment` are set for a service, values set by `environment` have precedence.

### expose

`expose` defines the ports that Compose implementations MUST expose from container. These ports MUST be
accessible to linked services and SHOULD NOT be published to the host machine. Only the internal container
ports can be specified.

```yml
expose:
  - "3000"
  - "8000"
```

### external_links

`external_links` link service containers to services managed outside this Compose application.
`external_links` define the name of an existing service to retrieve using the platform lookup mechanism.
An alias of the form `SERVICE:ALIAS` can be specified.

```yml
external_links:
  - redis
  - database:mysql
  - database:postgresql
```

### extra_hosts

`extra_hosts` adds hostname mappings to the container network interface configuration (`/etc/hosts` for Linux).
Values MUST set hostname and IP address for additional hosts in the form of `HOSTNAME:IP`.

```yml
extra_hosts:
  - "somehost:162.242.195.82"
  - "otherhost:50.31.209.229"
```

Compose implementations MUST create matching entry with the IP address and hostname in the container's network
configuration, which means for Linux `/etc/hosts` will get extra lines:

```
162.242.195.82  somehost
50.31.209.229   otherhost
```

### healthcheck

`healthcheck` declares a check that's run to determine whether or not containers for this
service are "healthy". This overrides
[HEALTHCHECK Dockerfile instruction](https://docs.docker.com/engine/reference/builder/#healthcheck)
set by the service's Docker image.

```yml
healthcheck:
  test: ["CMD", "curl", "-f", "http://localhost"]
  interval: 1m30s
  timeout: 10s
  retries: 3
  start_period: 40s
```

`interval`, `timeout` and `start_period` are [specified as durations](#specifying-durations).

`test` defines the command the Compose implementation will run to check container health. It can be
either a string or a list. If it's a list, the first item must be either `NONE`, `CMD` or `CMD-SHELL`.
If it's a string, it's equivalent to specifying `CMD-SHELL` followed by that string.

```yml
# Hit the local web app
test: ["CMD", "curl", "-f", "http://localhost"]
```

Using `CMD-SHELL` will run the command configured as a string using the container's default shell
(`/bin/sh` for Linux). Both forms below are equivalent:

```yml
test: ["CMD-SHELL", "curl -f http://localhost || exit 1"]
```

```yml
test: curl -f https://localhost || exit 1
```

`NONE` disable the healthcheck, and is mostly useful to disable Healthcheck set by image. Alternatively
the healthcheck set by the image can be disabled by setting `disable: true`:

```yml
healthcheck:
  disable: true
```

### image

`image` specifies the image to start the container from. Image MUST follow the Open Container Specification 
[addressable image format](https://github.com/opencontainers/org/blob/master/docs/docs/introduction/digests.md),
as `[<registry>/][<project>/]<image>[:<tag>|@<digest>]`.


```yml
    image: redis
    image: redis:5
    image: redis@sha356:0ed5d5928d4737458944eb604cc8509e245c3e19d02ad83935398bc4b991aac7
    image: library/redis
    image: docker.io/library/redis
    image: my_private.registry:5000/redis
```    

If the image does not exist on the platform, Compose implementations MUST attempt to pull it. Compose
implementations with build support MAY offer alternative options for the end user to control precedence of
pull over building the image from source, however pulling the image MUST be the default behaviour.

`image` MAY be omitted from a Compose file as long as a `build` section is declared. Compose implementations
without build support MUST fail when `image` is missing from the Compose file.

### init

`init` run an init process (PID 1) inside the container that forwards signals and reaps processes.
Set this option to `true` to enable this feature for the service.

```yml
version: "3"
services:
  web:
    image: alpine:latest
    init: true
```

The init binary that is used is platform specific.

### isolation

`isolation` specifies a container’s isolation technology. Supported values are platform-specific.

### labels

`labels` add metadata to containers. You can use either an array or a map.

It's recommended that you use reverse-DNS notation to prevent your labels from conflicting with
those used by other software.

```yml
labels:
  com.example.description: "Accounting webapp"
  com.example.department: "Finance"
  com.example.label-with-empty-value: ""
```

```yml
labels:
  - "com.example.description=Accounting webapp"
  - "com.example.department=Finance"
  - "com.example.label-with-empty-value"
```

Compose implementations MUST create containers with canonical labels:

- `com.docker.compose.project` set on all resources created by Compose implementation to the user project name
- `com.docker.compose.service` set on service containers with service name as defined in the Compose file

The `com.docker.compose` label prefix is reserved. Specifying labels with this previx in the Compose file MUST
result in a runtime error.

### links

`links` defines a network link to containers in another service. Either specify both the service name and
a link alias (`SERVICE:ALIAS`), or just the service name.

```yml
web:
  links:
    - db
    - db:database
    - redis
```

Containers for the linked service MUST be reachable at a hostname identical to the alias, or the service name
if no alias was specified.

Links are not required to enable services to communicate - when no specific network configuration is set,
any service MUST be able to reach any other service at that service’s name on the `default` network. If services
do declare networks they are attached to, `links` SHOULD NOT override the network configuration and services not
attached to a shared network SHOULD NOT be able to communicate. Compose implementations MAY NOT warn the user
about this configuration mismatch.

Links also express implicit dependency between services in the same way as
[depends_on](#depends_on), so they determine the order of service startup.

### logging

`logging` defines the logging configuration for the service.

```yml
logging:
  driver: syslog
  options:
    syslog-address: "tcp://192.168.0.42:123"
```

The `driver` name specifies a logging driver for the service's containers. The default and available values
are platform specific. Driver specific options can be set with `options` as key-value pairs.

### network_mode

`network_mode` set service containers network mode. Available values are platform specific, but Compose
specification define specific values which MUST be implemented as described if supported:

- `none` which disable all container networking
- `host` which gives the container raw access to host's network interface
- `service:{name}` which gives the containers access to the specified service only

```yml
    network_mode: "host"
    network_mode: "none"
    network_mode: "service:[service name]"
```

### networks

`networks` defines the networks that service containers are attached to, referencing entries under the
[top-level `networks` key](#networks-top-level-element).

```yml
services:
  some-service:
    networks:
      - some-network
      - other-network
```

#### aliases

`aliases` declares alternative hostnames for this service on the network. Other containers on the same
network can use either the service name or this alias to connect to one of the service's containers.

Since `aliases` are network-scoped, the same service can have different aliases on different networks.

> **Note**: A network-wide alias can be shared by multiple containers, and even by multiple services.
> If it is, then exactly which container the name resolves to is not guaranteed.

The general format is shown here:

```yml
services:
  some-service:
    networks:
      some-network:
        aliases:
          - alias1
          - alias3
      other-network:
        aliases:
          - alias2
```

In the example below, service `frontend` will be able to reach the `backend` service at
the hostname `backend` or `database` on the `back-tier` network, and service `monitoring`
will be able to reach same `backend` service at `db` or `mysql` on the `admin` network.

```yml
version: "3"

services:
  frontend:
    image: awesome/webapp
    networks:
      - front-tier
      - back-tier

  monitoring:
    image: awesome/monitoring
    networks:
      - admin

  backend:
    image: awesome/backend
    networks:
      back-tier:
        aliases:
          - database
      admin:
        aliases:
          - mysql

networks:
  front-tier:
  back-tier:
  admin:
```

#### ipv4_address, ipv6_address

Specify a static IP address for containers for this service when joining the network.

The corresponding network configuration in the [top-level networks section](#networks) MUST have an
`ipam` block with subnet configurations covering each static address.

```yml
version: "3"

services:
  frontend:
    image: awesome/webapp
    networks:
      front-tier:
        ipv4_address: 172.16.238.10
        ipv6_address: 2001:3984:3989::10

networks:
  front-tier:
    ipam:
      driver: default
      config:
        - subnet: "172.16.238.0/24"
        - subnet: "2001:3984:3989::/64"
```

### pid

`pid` sets the PID mode for container created by the Compose implementation.
Supported values are platform specific.

### ports

Exposes container ports.
Port mapping MUST NOT be used with `network_mode: host` and doing so MUST result in a runtime error.

#### Short syntax

The short syntax is a comma-separated string to set host IP, host port and container port
in the form:

`[HOST:]CONTAINER[/PROTOCOL]` where:

- `HOST` is `[IP:](port | range)`
- `CONTAINER` is `port | range`
- `PROTOCOL` to restrict port to specified protocol. `tcp` and `udp` values are defined by the specification,
  Compose implementations MAY offer support for platform-specific protocol names.

Host IP, if not set, MUST bind to all network interfaces. Port can be either a single
value or a range. Host and container MUST use equivalent ranges.

Either specify both ports (`HOST:CONTAINER`), or just the container port. In the latter case, the
Compose implementation SHOULD automatically allocate and unassigned host port.

`HOST:CONTAINER` SHOULD always be specified as a (quoted) string, to avoid conflicts
with [yaml base-60 float](https://yaml.org/type/float.html).

Samples:

```yml
ports:
  - "3000"
  - "3000-3005"
  - "8000:8000"
  - "9090-9091:8080-8081"
  - "49100:22"
  - "127.0.0.1:8001:8001"
  - "127.0.0.1:5000-5010:5000-5010"
  - "6060:6060/udp"
```

> **Note**: Host IP mapping MAY not be supported on the platform, in such case Compose implementations SHOULD reject
> the Compose file and MUST inform the user they will ignore the specified host IP.

#### Long syntax

The long form syntax allows the configuration of additional fields that can't be
expressed in the short form.

- `target`: the container port
- `published`: the publicly exposed port
- `protocol`: the port protocol (`tcp` or `udp`), unspecified means any protocol
- `mode`: `host` for publishing a host port on each node, or `ingress` for a port to be load balanced.

```yml
ports:
  - target: 80
    published: 8080
    protocol: tcp
    mode: host
```

### restart

`restart` defines the policy that the platform will apply on container termination.

- `no`: The default restart policy. Does not restart a container under any circumstances.
- `always`: The policy always restarts the container until its removal.
- `on-failure`: The policy restarts a container if the exit code indicates an error.
- `unless-stopped`: The policy restarts a container irrespective of the exit code but will stop
  restarting when the service is stopped or removed.

```yml
    restart: "no"
    restart: always
    restart: on-failure
    restart: unless-stopped
```

### secrets

`secrets` grants access to sensitive data defined by [secrets](secrets) on a per-service basis. Two
different syntax variants are supported: the short syntax and the long syntax.

Compose implementations MUST report an error if the secret doesn't exist on the platform or isn't defined in the
[`secrets`](#secrets-top-level-element) section of this Compose file.

#### Short syntax

The short syntax variant only specifies the secret name. This grants the
container access to the secret and mounts it as read-only to `/run/secrets/<secret_name>`
within the container. The source name and destination mountpoint are both set
to the secret name.

The following example uses the short syntax to grant the `frontend` service
access to the `server-certificate` secret. The value of `server-certificate` is set
to the contents of the file `./server.cert`.

```yml
version: "3"
services:
  frontend:
    image: awesome/webapp
    secrets:
      - server-certificate
secrets:
  server-certificate:
    file: ./server.cert
```

#### Long syntax

The long syntax provides more granularity in how the secret is created within
the service's containers.

- `source`: The name of the secret as it exists on the platform.
- `target`: The name of the file to be mounted in `/run/secrets/` in the
  service's task containers. Defaults to `source` if not specified.
- `uid` and `gid`: The numeric UID or GID that owns the file within
  `/run/secrets/` in the service's task containers. Default value is USER running container.
- `mode`: The [permissions](http://permissions-calculator.org/) for the file to be mounted in `/run/secrets/`
  in the service's task containers, in octal notation.
  Default value is world-readable permissions (mode `0444`).
  The writable bit MUST be ignored if set. The executable bit MAY be set.

The following example sets the name of the `server-certificate` secret file to `server.crt`
within the container, sets the mode to `0440` (group-readable) and sets the user and group
to `103`. The value of `server-certificate` secret is provided by the platform through a lookup and
the secret lifecycle not directly managed by the Compose implementation.

```yml
version: "3"
services:
  frontend:
    image: awesome/webapp
    secrets:
      - source: server-certificate
        target: server.cert
        uid: "103"
        gid: "103"
        mode: 0440
secrets:
  server-certificate:
    external: true
```

Services MAY be granted access to multiple secrets. Long and short syntax for secrets MAY be used in the
same Compose file. Defining a secret in the top-level `secrets` MUTS NOT imply granting any service access to it.
Such grant must be explicit within service specification as [secrets](#secrets) service element.

### security_opt

`security_opt` overrides the default labeling scheme for each container.

```yml
security_opt:
  - label:user:USER
  - label:role:ROLE
```

### stop_grace_period

`stop_grace_period` specifies how long the Compose implementation MUST wait when attempting to stop a container if it doesn't
handle SIGTERM (or whichever stop signal has been specified with
[`stop_signal`](#stopsignal)), before sending SIGKILL. Specified
as a [duration](#specifying-durations).

```yml
    stop_grace_period: 1s
    stop_grace_period: 1m30s
```

Default value is 10 seconds for the container to exit before sending SIGKILL.

### stop_signal

`stop_signal` defines the signal that the Compose implementation MUST use to stop the service containers.
If unset containers are stopped by the Compose Implementation by sending `SIGTERM`.

```yml
stop_signal: SIGUSR1
```

### sysctls

`sysctls` defines kernel parameters to set in the container. `sysctls` can use either an array or a map.

```yml
sysctls:
  net.core.somaxconn: 1024
  net.ipv4.tcp_syncookies: 0
```

```yml
sysctls:
  - net.core.somaxconn=1024
  - net.ipv4.tcp_syncookies=0
```

You can only use sysctls that are namespaced in the kernel. Docker does not
support changing sysctls inside a container that also modify the host system.
For an overview of supported sysctls, refer to [configure namespaced kernel
parameters (sysctls) at runtime](https://docs.docker.com/engine/reference/commandline/run/#configure-namespaced-kernel-parameters-sysctls-at-runtime).

### tmpfs

`tmpfs` mounts a temporary file system inside the container. Can be a single value or a list.

```yml
tmpfs: /run
```

```yml
tmpfs:
  - /run
  - /tmp
```

### ulimits

`ulimits` overrides the default ulimits for a container. Either specifies as a single limit as an integer or
soft/hard limits as a mapping.

```yml
ulimits:
  nproc: 65535
  nofile:
    soft: 20000
    hard: 40000
```

### userns_mode

`userns_mode` sets the user namespace for the service. Supported values are platform specific and MAY depend
on platform configuration

```yml
userns_mode: "host"
```

### volumes

`volumes` defines mount host paths or named volumes that MUST be accessible by service containers.

If the mount is a host path and only used by a single service, it MAY be declared as part of the service
definition instead of the top-level `volumes` key.

To reuse a volume across multiple services, a named
volume MUST be declared in the [top-level `volumes` key](#volumes-top-level-element).

This example shows a named volume (`db-data`) being used by the `backend` service,
and a bind mount defined for a single service

```yml
version: "3"
services:
  backend:
    image: awesome/backend
    volumes:
      - type: volume
        source: db-data
        target: /data
        volume:
          nocopy: true
      - type: bind
        source: /var/run/postgres/postgres.sock
        target: /var/run/postgres/postgres.sock

volumes:
  db-data:
```

#### Short syntax

The short syntax uses a single string with comma-separated values to specify a volume mount
(`VOLUME:CONTAINER_PATH`), or an access mode (`VOLUME:CONTAINER:ACCESS_MODE`).

`VOLUME` MAY be either a host path on the platform hosting containers (bind mount) or a volume name.
`ACCESS_MODE` MAY be set as read-only by using `ro` or read and write by using `rw` (default).

> **Note**: Relative host paths MUST only be supported by Compose implementations that deploy to a
> local container runtime. This is because the relative path is resolved from the Compose file’s parent
> directory which is only applicable in the local case. Compose Implementations deploying to a non-local
> platform MUST reject Compose files which use relative host paths with an error. To avoid ambiguities
> with named volumes, relative paths SHOULD always begin with `.` or `..`.

#### Long syntax

The long form syntax allows the configuration of additional fields that can't be
expressed in the short form.

- `type`: the mount type `volume`, `bind`, `tmpfs` or `npipe`
- `source`: the source of the mount, a path on the host for a bind mount, or the
  name of a volume defined in the
  [top-level `volumes` key](#volumes-top-level-element). Not applicable for a tmpfs mount.
- `target`: the path in the container where the volume is mounted
- `read_only`: flag to set the volume as read-only
- `bind`: configure additional bind options
  - `propagation`: the propagation mode used for the bind
- `volume`: configure additional volume options
  - `nocopy`: flag to disable copying of data from a container when a volume is created
- `tmpfs`: configure additional tmpfs options
  - `size`: the size for the tmpfs mount in bytes
- `consistency`: the consistency requirements of the mount. Available values are platform specific

### domainname

`domainname` declares a custom domain name to use for the service container. MUST be a valid RFC 1123 hostname.

### hostname

`hostname` declares a custom host name to use for the service container. MUST be a valid RFC 1123 hostname.

### ipc

`ipc` configures the IPC isolation mode set by service container.

### mac_address

`mac_address` sets a MAC address for service container.

### privileged

`privileged` configures the service container to run with elevated privileges. Support and actual impacts are platform-specific.

### read_only

`read_only` configures service container to be created with a read-only filesystem.

### shm_size

`shm_size` configures the size of the shared memory (`/dev/shm` partition on Linux) allowed by the service container.
Specified as a [byte value](#specifying-byte-values).

### stdin_open

`stdin_open` configures service containers to run with an allocated stdin.

### tty

`tty` configure service container to run with a TTY.

### user

`user` overrides the user used to run the container process. Default is that set by image (i.e. Dockerfile `USER`),
if not set, `root`.

### working_dir

`working_dir` overrides the container's working directory from that specified by image (i.e. Dockerfile `WORKDIR`).

## Networks top-level element

Networks are the layer that allow services to communicate with each other. The networking model exposed to a service
is limited to a simple IP connection with target services and external resources, while the Network definition allows
fine-tuning the actual implementation provided by the platform.

Networks can be created by specifying the network name under a top-level `networks` section.
Services can connect to networks by specifying the network name under the service [`networks`](#networks) subsection

In the following example, at runtime, networks `front-tier` and `back-tier` will be created and the `frontend` service
connected to the `front-tier` network and the `back-tier` network.

```yml
version: "3"
services:
  frontend:
    image: awesome/webapp
    networks:
      - front-tier
      - back-tier

networks:
  front-tier:
  back-tier:
```

### driver

`driver` specifies which driver should be used for this network. Compose implementations MUST return an error if the
driver is not available on the platform.

```yml
driver: overlay
```

Default and available values are platform specific. Compose specification MUST support the following specific drivers:
`none` and `host`

- `host` use the host's networking stack
- `none` disable networking

#### host or none

The syntax for using built-in networks such as `host` and `none` is different, as such networks implicitly exists outside
the scope of the Compose implementation. To use them one MUST define an external network with the name `host` or `none` and
an alias that the Compose implementation can use (`hostnet` or `nonet` in the following examples), then grant the service
access to that network using its alias.

```yml
version: "3"
services:
  web:
    networks:
      hostnet: {}

networks:
  hostnet:
    external: true
    name: host
```

```yml
services:
  web:
    ...
    networks:
      nonet: {}

networks:
  nonet:
    external: true
    name: none
```

### driver_opts

`driver_opts` specifies a list of options as key-value pairs to pass to the driver for this network. These options are
driver-dependent - consult the driver's documentation for more information. Optional.

```yml
driver_opts:
  foo: "bar"
  baz: 1
```

### attachable

If `attachable` is set to `true`, then standalone containers SHOULD be able attach to this network, in addition to services.
If a standalone container attaches to the network, it can communicate with services and other standalone containers
that are also attached to the network.

```yml
networks:
  mynet1:
    driver: overlay
    attachable: true
```

### enable_ipv6

`enable_ipv6` enable IPv6 networking on this network.

### ipam

`ipam` specifies custom a IPAM configuration. This is an object with several properties, each of which is optional:

- `driver`: Custom IPAM driver, instead of the default.
- `config`: A list with zero or more configuration elements, each containing:
  - `subnet`: Subnet in CIDR format that represents a network segment

A full example:

```yml
ipam:
  driver: default
  config:
    - subnet: 172.28.0.0/16
```

### internal

By default, Compose implementations MUST provides external connectivity to networks. `internal` when set to `true` allow to
create an externally isolated network.

### labels

Add metadata to containers using Labels. Can use either an array or a dictionary.

Users SHOULD use reverse-DNS notation to prevent labels from conflicting with those used by other software.

```yml
labels:
  com.example.description: "Financial transaction network"
  com.example.department: "Finance"
  com.example.label-with-empty-value: ""
```

```yml
labels:
  - "com.example.description=Financial transaction network"
  - "com.example.department=Finance"
  - "com.example.label-with-empty-value"
```

Compose implementations MUST set `com.docker.compose.project` and `com.docker.compose.network` labels.

### external

If set to `true`, `external` specifies that this network’s lifecycle is maintained outside of that of the application.
Compose Implementations SHOULD NOT attempt to create these networks, and raises an error if one doesn't exist.

In the example below, `proxy` is the gateway to the outside world. Instead of attempting to create a network, Compose
implementations SHOULD interrogate the platform for an existing network simply called `outside` and connect the
`proxy` service's containers to it.

```yml
version: "3"

services:
  proxy:
    image: awesome/proxy
    networks:
      - outside
      - default
  app:
    image: awesome/app
    networks:
      - default

networks:
  outside:
    external: true
```

### name

`name` sets a custom name for this network. The name field can be used to reference networks which contain special characters.
The name is used as is and will **not** be scoped with the project name.

```yml
version: "3"
networks:
  network1:
    name: my-app-net
```

It can also be used in conjunction with the `external` property to define the platform network that the Compose implementation
should retrieve, typically by using a parameter so the Compose file doesn't need to hard-code runtime specific values:

```yml
version: "3"
networks:
  network1:
    external: true
    name: "${NETWORK_ID}"
```

## Volumes top-level element

Volumes are persistent data stores implemented by the platform. The Compose specification offers a neutral abstraction
for services to mount volumes, and configuration parameters to allocate them on infrastructure.

The `volumes` section allows the configuration of named volumes that can be reused across multiple services. Here's
an example of a two-service setup where a database's data directory is shared with another service as a volume so
that it can be periodically backed up:

```yml
version: "3"

services:
  backend:
    image: awesome/database
    volumes:
      - db-data:/etc/data

  backup:
    image: backup-service
    volumes:
      - db-data:/var/lib/backup/data

volumes:
  db-data:
```

An entry under the top-level `volumes` key can be empty, in which case it uses the platform's default configuration for
creating a volume. Optionally, you can configure it with the following keys:

### driver

Specify which volume driver should be used for this volume. Default and available values are platform specific. If the driver is not available, the Compose implementation MUST return an error and stop application deployment.

```yml
driver: foobar
```

### driver_opts

`driver_opts` specifies a list of options as key-value pairs to pass to the driver for this volume. Those options are driver-dependent.

```yml
volumes:
  example:
    driver_opts:
      type: "nfs"
      o: "addr=10.40.0.199,nolock,soft,rw"
      device: ":/docker/example"
```

### external

If set to `true`, `external` specifies that this volume already exist on the platform and its lifecycle is managed outside
of that of the application. Compose implementations MUST NOT attempt to create these volumes, and MUST return an error they
do not exist.

In the example below, instead of attempting to create a volume called
`{project_name}_data`, Compose looks for an existing volume simply
called `data` and mount it into the `db` service's containers.

```yml
version: "3"

services:
  backend:
    image: awesome/database
    volumes:
      - db-data:/etc/data

volumes:
  db-data:
    external: true
```

### labels

`labels` are used to add metadata to volumes. You can use either an array or a dictionary.

It's recommended that you use reverse-DNS notation to prevent your labels from
conflicting with those used by other software.

```yml
labels:
  com.example.description: "Database volume"
  com.example.department: "IT/Ops"
  com.example.label-with-empty-value: ""
```

```yml
labels:
  - "com.example.description=Database volume"
  - "com.example.department=IT/Ops"
  - "com.example.label-with-empty-value"
```

Compose implementation MUST set `com.docker.compose.project` and `com.docker.compose.volmume` labels.

### name

`name` set a custom name for this volume. The name field can be used to reference volumes that contain special
characters. The name is used as is and will **not** be scoped with the stack name.

```yml
version: "3"
volumes:
  data:
    name: "my-app-data"
```

It can also be used in conjunction with the `external` property. Doing so the name of the volume used to lookup for
actual volume on platform is set separately from the name used to refer to it within the Compose file:

```yml
volumes:
  db-data:
    external:
      name: actual-name-of-volume
```

This make it possible to make this lookup name a parameter of a Compose file, so that the model ID for volume is
hard-coded but the actual volume ID on platform is set at runtime during deployment:

```yml
volumes:
  db-data:
    external:
      name: ${DATABASE_VOLUME}
```

## Configs top-level element

Configs allow services to adapt their behaviour without the need to rebuild a Docker image. Configs are comparable to Volumes from a service point of view as they are mounted into service's containers filesystem. The actual implementation detail to get configuration provided by the platform can be set from the Configuration definition.

When granted accessto a config, the config content is mounted as a file in the container. The location of the mount point within the container defaults to `/<config-name>` in Linux containers and `C:\<config-name>` in Windows containers.

By default, the config MUST be owned by the user running the container command but can be overriden by service configuration.
By default, the config MUST have world-readable permissions (mode 0444), unless service is configured to override this.

Services can only access configs when explicitly granted by a [`configs`](#configs) subsection.

The top-level `configs` declaration defines or references
configuration data that can be granted to the services in this
application. The source of the config is either `file` or `external`.

- `file`: The config is created with the contents of the file at the specified path.
- `external`: If set to true, specifies that this config has already been created. Compose implementation does not
  attempt to create it, and if it does not exist, an error occurs.
- `name`: The name of config object on Platform to lookup. This field can be used to
  reference configs that contain special characters. The name is used as is
  and will **not** be scoped with the project name.

In this example, `http_config` is created (as `<project_name>_http_config)`when the application is deployed,
and `my_second_config` MUST already exists on Platform and value will be obtained by lookup.

In this example, `server-http_config` is created as `<project_name>_http_config` when the application is deployed,
by registering content of the `httpd.conf` as configuration data.

```yml
configs:
  http_config:
    file: ./httpd.conf
```

Alternatively, `http_config` can be declared as external, doing so Compose implementation will lookup `server-certificate` to expose configuration data to relevant services.

```yml
configs:
  http_config:
    external: true
```

External configs lookup can also use a distinct key by specifying a `name`. The following
example modifies the previous one to lookup for config using a parameter `HTTP_CONFIG_KEY`. Doing
so the actual lookup key will be set at deployment time by [interpolation](#interpolation) of
variables, but exposed to containers as hard-coded ID `http_config`.

```yml
configs:
  http_config:
    external: true
    name: "${HTTP_CONFIG_KEY}"
```

Compose file need to explicitly grant access to the configs to relevant services in the application.

## Secrets top-level element

Secrets are a flavour of Configs focussing on sensitive data, with specific constraint for this usage. As the platform implementation may significally differ from Configs, dedicated Secrets section allows to configure the related resources.

The top-level `secrets` declaration defines or references sensitive data that can be granted to the services in this
application. The source of the secret is either `file` or `external`.

- `file`: The secret is created with the contents of the file at the specified path.
- `external`: If set to true, specifies that this secret has already been created. Compose implementation does
  not attempt to create it, and if it does not exist, an error occurs.
- `name`: The name of the secret object in Docker. This field can be used to
  reference secrets that contain special characters. The name is used as is
  and will **not** be scoped with the project name.

In this example, `server-certificate` is created as `<project_name>_server-certificate` when the application is deployed,
by registering content of the `server.cert` as a platform secret.

```yml
secrets:
  server-certificate:
    file: ./server.cert
```

Alternatively, `server-certificate` can be declared as external, doing so Compose implementation will lookup `server-certificate` to expose secret to relevant services.

```yml
secrets:
  server-certificate:
    external: true
```

External secrets lookup can also use a distinct key by specifying a `name`. The following
example modifies the previous one to look up for secret using a parameter `CERTIFICATE_KEY`. Doing
so the actual lookup key will be set at deployment time by [interpolation](#interpolation) of
variables, but exposed to containers as hard-coded ID `server-certificate`.

```yml
secrets:
  server-certificate:
    external: true
    name: "${CERTIFICATE_KEY}"
```

Compose file need to explicitly grant access to the secrets to relevant services in the application.

## Fragments

It is possible to re-use configuration fragments using [YAML anchors](http://www.yaml.org/spec/1.2/spec.html#id2765878).

```yml
volumes:
  db-data: &default-volume
    driver: default
  metrics: *default-volume
```

In previous sample, an _anchor_ is created as `default-volume` based on `db-data` volume specification. It is later reused by _alias_ `*default-volume` to define `metrics` volume. Same logic can apply to any element in a Compose file. Anchor resolution MUST take place
before [variables interpolation](#interpolation), so variables can't be used to set anchors or aliases.

It is also possible to partially override values set by anchor reference using the
[YAML merge type](http://yaml.org/type/merge.html). In following example, `metrics` volume specification uses alias
to avoid repetition but override `name` attribute:

```yml
version: "3"

services:
  backend:
    image: awesome/database
    volumes:
      - db-data
      - metrics
volumes:
  db-data: &default-volume
    driver: default
    name: "data"
  metrics:
    <<: *default-volume
    name: "metrics"
```

## Extension

Special extensions fields can be of any format as long as they are located at the root of your Compose file, or first level element, and their name starts with the `x-` character sequence.

```yml
version: "3"
x-custom:
  foo:
    - bar
    - zot
```

The contents of such fields are unspecified by Compose specification, and can be used to enable custom features. Compose implementation to encounter an unknown extension field MUST NOT fail, but COULD warn about unknown field.

For platform extensions, it is highly recommended to prefix extension by platform/vendor name, the same way browsers add
support for [custom CSS features](https://www.w3.org/TR/2011/REC-CSS2-20110607/syndata.html#vendor-keywords)

```yml
service:
  backend:
    deploy:
      placement:
        x-aws-role: "arn:aws:iam::XXXXXXXXXXXX:role/foo"
        x-aws-region: "eu-west-3"
        x-azure-region: "france-central"
```

### Informative Historical Notes

This section is informative. At the time of writing, the following prefixes are known to exist:

| prefix     | vendor/organization |
| ---------- | ------------------- |
| docker     | Docker              |
| kubernetes | Kubernetes          |

### Using extensions as fragments

With the support for extension fields, Compose file can be written as follows to improve readability of reused fragments:

```yml
version: "3"
x-logging: &default-logging
  options:
    max-size: "12m"
    max-file: "5"
  driver: json-file

services:
  frontend:
    image: awesome/webapp
    logging: *default-logging
  backend:
    image: awesome/database
    logging: *default-logging
```

### specifying- byte values

Value express a byte value as a string in `{amount}{byte unit}` format:
The supported units are `b` (bytes), `k` or `kb` (kilo bytes), `m` or `mb` (mega bytes) and `g` or `gb` (giga bytes).

```
    2b
    1024kb
    2048k
    300m
    1gb
```

### specifying durations

Value express a duration as a string in thte in the form of `{value}{unit}`.
The supported units are `us` (microseconds), `ms` (milliseconds), `s` (seconds), `m` (minutes) and `h` (hours).
Value can can combine mutiple values and using without separator.

```
  10ms
  40s
  1m30s
  1h5m30s20ms
```

## Interpolation

Values in a Compose file can be set by variables, and interpolated at runtime. Compose files use a Bash-like
syntax `${VARIABLE}`

Both `$VARIABLE` and `${VARIABLE}` syntax are supported. Default values can be defined inline using typical shell syntax:
latest

- `${VARIABLE:-default}` evaluates to `default` if `VARIABLE` is unset or
  empty in the environment.
- `${VARIABLE-default}` evaluates to `default` only if `VARIABLE` is unset
  in the environment.

Similarly, the following syntax allows you to specify mandatory variables:

- `${VARIABLE:?err}` exits with an error message containing `err` if
  `VARIABLE` is unset or empty in the environment.
- `${VARIABLE?err}` exits with an error message containing `err` if
  `VARIABLE` is unset in the environment.

Other extended shell-style features, such as `${VARIABLE/foo/bar}`, are not
supported by the Compose specification.

You can use a `$$` (double-dollar sign) when your configuration needs a literal
dollar sign. This also prevents Compose from interpolating a value, so a `$$`
allows you to refer to environment variables that you don't want processed by
Compose.

```yml
web:
  build: .
  command: "$$VAR_NOT_INTERPOLATED_BY_COMPOSE"
```

If the Compose implementation can't resolve a substituted variable and no default value is defined, it MUST warn
the user and substitute the variable with an empty string.

As any values in a Compose file can be interpolated with variable substitution, including compact string notation
for complex elements, interpolation MUST be applied _before_ merge on a per-file-basis.
