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

A **Project** is an individual deployment of an application specification on a platform. Project name is used to group
resources together and isolate them from other applications or other installation of the same Compose specified application with distinct parameters. Compose implementation creating resources on platform MUST prefix resource names by project and
set label `com.docker.compose.project`.

### Illustration sample

The following sample illustrates Compose specification concepts with a concrete sample application. The sample is non-normative.

Consider an application split into a frontend web application and a backend service.

Frontend is configured at runtime with http configuration file managed by infrastucture, providing external domain name, and https server certificate injected from platform's secured secret store. 

Backend stores data in a persistent volume. 

Both services communicate together on an isolated back-tier network, while webapp is also connected to a front-tier network and expose port 443 to external usage. 

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


so we have :
- 2 services, backed by docker images `webapp` and `database`
- 1 secret (https certificate), injected in fronted
- 1 configuration (HTTP), injected in frontend
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
  front-tier:
  back-tier:
```

This sample illustrate the distinction between volumes, configs and secrets. While all of them are all exposed to service containers as mounted files or directories, only a volume can be configured for read+write access - secrets and config are read-only. Volume configuration allows you to select a volume driver and pass driver options to tweak volume management according to the actual infrastructure. Configs an Secrets rely on platform services, and are declared `external` as not being managed as part of the application: Compose implementation will use platform-specific lookup mechanism to retrieve runtime values.


## Compose file 

The Compose file is a [YAML](http://yaml.org/) file defining
[version](#version) (REQUIRED),
[services](#service) (REQUIRED),
[networks](#network),
[volumes](#volume),
[configs](#configs) and
[secrets](#secrets).
The default path for a Compose file is `./docker-compose.yml

Multiple Compose file can be combined together to define application model. Combination of yaml files MUST be implemented by appending/overriding yaml elements based on compose file order set by user. Simple attributes an Maps get overriden by latest compose file, lists get merged by appending.

As some Compose file elements can both be expressed as single strings or complex object, merge MUST apply to the expended form.

## Interpolation

Values in a Compose file can be set by variables, and interpolated at runtime. Compose file uses a Bash-like syntax `${VARIABLE}`

Both `$VARIABLE` and `${VARIABLE}` syntax are supported. Default values can be define inline using typical shell syntax:

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

If Compose implementation can't resolve a substitution variable and no default value is defined, it MUST warn user and substitute variable by an empty string.

As any values in a Compose file can be interpolated with variable substitution, including compact string notation for complex elements, interpolation MUST be applied _before_ merge on a per-file-basis.

## Services

A Service is an abstract definition of a computing resource within an application which can be scaled/replaced independently from other components. Services are actually backed by a set of containers, run by the platform according to replication requirements and placement constraints. Being backed by containers, Services are defined by a Docker image and set of runtime arguments. All containers within a service are identically created from those arguments.

Service also includes a Build section, defining how to create the Docker image for a service. Support by Compose implementations to build docker images according to this service definition is OPTIONAL. If not implemented the Build section SHOULD be ignored and the Compose file MUST still be considered valid. 

Build support is an OPTIONAL aspect of the Compose specification, and is described in detail [here](build.md)

Service define runtime constraints and requirement to run service's containers. The Deploy section do group those and allow plaform to adjust deployment strategy to best match containers needs with available resources.

Deploy support is an OPTIONAL aspect of the Compose specification, and is described in detail [here](deploy.md). If not implemented the Deploy section SHOULD be ignored and the Compose file MUST still be considered valid. 

A Compose file MUST declare a `services` root element as a map or service names to service definitions. A service definition contains configuration that is applied to each container started for that service.

### deploy

`deploy` specify configuration related to the deployment and running of services, as defined [here](deploy.md). 

### build

`build`	specify the build process to create container image from source, as defined [here](build.md)

### cap_add

`cap_add`	allows to add container [capabilities](http://man7.org/linux/man-pages/man7/capabilities.7.html).

```
cap_add:
  - ALL
```

### cap_drop

`cap_drop` allows to drop container [capabilities](http://man7.org/linux/man-pages/man7/capabilities.7.html) as strings. 
```
cap_drop:
  - NET_ADMIN
  - SYS_ADMIN
```

### cgroup_parent

`cgroup_parent` Specify an OPTIONAL parent [cgroup](http://man7.org/linux/man-pages/man7/cgroups.7.html) for the container.

```
cgroup_parent: m-executor-abcd
```

### command
`command` override the the default command declared by container image (i.e. by Dockerfile's `CMD`).
```
command: bundle exec thin -p 3000
```

The command can also be a list, in a manner similar to [Dockerfile](https://docs.docker.com/engine/reference/builder/#cmd):
```
command: bundle exec thin -p 3000
```

### configs

`configs` grant access to configs on a per-service basis using the per-service `configs`
configuration. Two different syntax variants are supported.

> **Note**: The config must already exist or be
> [defined in the top-level `configs` configuration](#configs)
> of this Compose file.

#### Short syntax

The short syntax variant only specifies the config name. This grants the
container access to the config and mounts it at `/<config_name>`
within the container. The source name and destination mountpoint are both set
to the config name.

The following example uses the short syntax to grant the `redis` service
access to the `my_config` and `my_other_config` configs. The value of
`my_config` is set to the contents of the file `./my_config.txt`, and
`my_other_config` is defined as an external resource, which means that it has
already been defined in Docker, either by running the `docker config create`
command or by another stack deployment. If the external config does not exist,
the stack deployment fails with a `config not found` error.

```yml
version: "3"
services:
  redis:
    image: redis:latest
    configs:
      - my_config
configs:
  my_config:
    external: true
```

#### Long syntax

The long syntax provides more granularity in how the config is created within the service's task containers.

- `source`: The name of the config as it exists in Docker.
- `target`: The path and name of the file to be mounted in the service's
  task containers. Defaults to `/<source>` if not specified.
- `uid` and `gid`: The numeric UID or GID that owns the mounted config file
  within in the service's task containers. Both default to `0` on Linux if not
  specified. Not supported on Windows.
- `mode`: The permissions for the file that is mounted within the service's
  task containers, in octal notation. For instance, `0444`
  represents world-readable. The default is `0444`. Configs cannot be writable
  because they are mounted in a temporary filesystem, so if you set the writable
  bit, it is ignored. The executable bit can be set. If you aren't familiar with
  UNIX file permission modes, you may find this
  [permissions calculator](http://permissions-calculator.org/){: target="_blank" class="_" }
  useful.

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
        uid: '103'
        gid: '103'
        mode: 0440
configs:
  my_config:
    external: true
```

You can grant a service access to multiple configs and you can mix long and short syntax. 
Defining a config does not imply granting a service access to it.

### container_name

`container_name` specify a custom container name, rather than a generated default name.

```yml
container_name: my-web-container
```

Because Docker container names must be unique, Compose implementatoin cannot scale a service beyond
1 container if Compose file specify a container_name. Attempting to do so MUST results in
an error.

### credential_spec

`credential_spec` configure the credential spec for managed service account. 

Compose implementation to support services using Windows containers MUST support `file:` and `registry:` protocols on credential_spec. 
Compose implementation MAY also support additional protocols for custom use-cases


The `credential_spec` must be in the format `file://<filename>` or `registry://<value-name>`.

When using `file:`, the referenced file must be present in the `CredentialSpecs`
subdirectory in the Docker data directory, which defaults to `C:\ProgramData\Docker\`
on Windows. The following example loads the credential spec from a file named
`C:\ProgramData\Docker\CredentialSpecs\my-credential-spec.json`:

```yml
credential_spec:
  file: my-credential-spec.json
```

When using `registry:`, the credential spec is read from the Windows registry on
the daemon's host. A registry value with the given name must be located in:

    HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Virtualization\Containers\CredentialSpecs

The following example load the credential spec from a value named `my-credential-spec`
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

`depends_on` express dependency between services, Service dependencies cause the following
behaviors:

- Compose implementation MUST create services in dependency order. In the following
  example, `db` and `redis` are created before `web`. 

- Compose implementation MUST remove services in dependency order. In the following
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

Compose implementation MAY NOT wait for dependency services to be "ready" before
starting a dependent service, only guarantee they have been started. 

### devices

`devices` defines a list of device mappings for created containers. 

```yml
devices:
  - "/dev/ttyUSB0:/dev/ttyUSB0"
```

### dns

`dns` define custom DNS servers to set on container network interface configuration. Can be a single value or a list.

```yml
dns: 8.8.8.8
```

```yml
dns:
  - 8.8.8.8
  - 9.9.9.9
```

### dns_search

`dns` define custom DNS search domainsto set on container network interface configuration. Can be a single value or a list.

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

`entrypoint` override the default entrypoint or docker image (i.e. `ENTRYPOINT` set by Dockerfile) 
**and** clears out any default command on the image - meaning that if there's a `CMD` instruction 
in the Dockerfile, it is ignored..

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

`env_file` add environment variables to the container based on file content. 


```yml
env_file: .env
```

`env_file` can also be a list. The files in the list are processed from the top down. 
For the same variable specified in file `a.env` and assigned a different value in file 
`b.env`, if `b.env` is listed below (after), then the value from `b.env` stands. 

```yml
env_file:
  - ./a.env
  - ./b.env
```

Relative path MUST be resolved from the Compose file parent folder. As absolute paths prevent the Compose 
file to be portable, Compose implementation SHOULD warn user when such a path is used to set `env_file`.

Environment variables declared in the [environment](#environment) section
_override_ these values &ndash; this holds true even if those values are
empty or undefined.

#### Env_file format

Each line in an env file to be in `VAR=VAL` format. Lines beginning with `#` are treated as comments and 
are ignored. Blank lines are also ignored.

The value of `VAL` is used as raw string and not modified at all. If the value is surrounded by quotes 
(as is often the case of shell variables), the quotes MUST be **included** in the value passed to containers
created by Compose implementation.

```bash
# Set Rails/Rack environment
RACK_ENV=development
VAR="quoted"
```


### environment

`environment` define environment variables set into container. `environment` can use either an array or a 
map. Any boolean values; true, false, yes no, MUST be enclosed in quotes to ensure
they are not converted to True or False by the YML parser.

Environment variables can be declare by a single key (no value ot equal sign). In such case Compose 
implementation SHOULD rely on some user interaction to resolve value.

Map syntax:
```yml
environment:
  RACK_ENV: development
  SHOW: 'true'
  USER_INPUT:
```

Array syntax:
```yml
environment:
  - RACK_ENV=development
  - SHOW=true
  - USER_INPUT
```

### expose

`expose` define ports the Compose implementation MUST expose from container. Such port SHOULD not be 
published to the host machine - they'll only be accessible to linked services. Only the internal port 
can be specified.

```yml
expose:
 - "3000"
 - "8000"
```

### external_links

`external_links` link service containers to services managed outside this Compose application.
`external_links` define the name of an existing service to retrieve by platform lookup mechanism.
An alias can be specified in the form `SERVICE:ALIAS`.

```yaml
external_links:
 - redis
 - database:mysql
 - database:postgresql
```

### extra_hosts

`extra_hosts` add hostname mappings on container network interface configuration (`/etc/hosts`). 
Values MUST set hostname and IP address for additional hosts in for form `HOSTNAME:IP`.

```yml
extra_hosts:
 - "somehost:162.242.195.82"
 - "otherhost:50.31.209.229"
```

Compose implementation MUST create matching entry with the ip address and hostname in container's 
`/etc/hosts`:

```
162.242.195.82  somehost
50.31.209.229   otherhost
```

### healthcheck

`healthcheck` declare a check that's run to determine whether or not containers for this
service are "healthy". This override 
[HEALTHCHECK Dockerfile instruction](https://docs.docker.com/engine/reference/builder/#healthcheck)
set by service Docker image.


```yml
healthcheck:
  test: ["CMD", "curl", "-f", "http://localhost"]
  interval: 1m30s
  timeout: 10s
  retries: 3
  start_period: 40s
```

`interval`, `timeout` and `start_period` are specified as durations in the form `[value unit]+`. 
The supported units are `us`, `ms`, `s`, `m` and `h`.

`test` define the command the Compose implementation will run to check container health. It can be 
either a string or a list. If it's a list, the first item must be either `NONE`, `CMD` or `CMD-SHELL`. 
If it's a string, it's equivalent to specifying `CMD-SHELL` followed by that string.

```yml
# Hit the local web app
test: ["CMD", "curl", "-f", "http://localhost"]
```

Using `CMD-SHELL` will run command configured as a string within container shell (`/bin/sh`). 
Both forms below are equivalent:

```yml
test: ["CMD-SHELL", "curl -f http://localhost || exit 1"]
```

```yml
test: curl -f https://localhost || exit 1
```

`NONE` disable healthcheck, and is mostly usefull to override Healthcheck set by image and disable
it. Alternatively Healthcheck set by the image can be disabled by setting `disable: true`:

```yml
healthcheck:
  disable: true
```

### image

`image` specify the image to start the container from. Can either be a repository/tag, a digested 
reference or a partial image ID.

```yml
    image: redis
    image: redis:5
    image: redis@sha356:0ed5d5928d4737458944eb604cc8509e245c3e19d02ad83935398bc4b991aac7
    image: library/redis
    image: docker.io/library/redis
    image: my_private.registry:5000/redis
    image: a4bc65fd
```    

If the image does not exist on platform, Compose implementation MUST attempt to pull it. Compose 
Implementation with Build support MAY offer alternative option for end-user to control precedence of
pull over building image from source, but pulling image MUST be the default behaviour.

`image` MAY be omitted from a Compose file as long as a `build` section is declared. Compose Implementation 
without Build support MUST fail when `image` is missing from Compose file.

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

`isolation` specify a container’s isolation technology. Supported values are platform-specific.

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

Compose Implementation MUST create container with canonical labels:

- `com.docker.compose.project` set on all resources created by Docmpose implementation to the user project name
- `com.docker.compose.service` set on service containers with service name as defined in the Compose file
- `com.docker.compose.network` set on networks with network name as defined in the Compose file
- `com.docker.compose.volume` set on volumes with volume name as defined in the Compose file

`com.docker.compose` label prefix is reserved. Such labels MUST NOT be overriden if specified by Compose file. 
An attempt to do so SHOULD result in an error.

### links

`links` define a network link to containers in another service. Either specify both the service name and
a link alias (`SERVICE:ALIAS`), or just the service name.

```yml
web:
  links:
   - db
   - db:database
   - redis
```

Containers for the linked service are reachable at a hostname identical to the alias, or the service name 
if no alias was specified.

Links are not required to enable services to communicate - when no specific network configuration is set, 
any service MUST be able to reach any other service at that service’s name on `default` network. If services
do declare networks they are attached to, `links` won't override network configuration and services not
attached to a shared network won't be able to communicate. Compose Implementation MAY NOT warn user about this
configuration mismatch.

Links also express implicit dependency between services in the same way as
[depends_on](#depends_on), so they determine the order of service startup.

### logging

`logging` define the logging configuration for the service.

```yml
logging:
  driver: syslog
  options:
    syslog-address: "tcp://192.168.0.42:123"
```

The `driver` name specifies a logging driver for the service's containers. The default and available values 
are platform specific. Options for the logging driver can be set by `options` as key-value pairs.


### network_mode

`network_mode` set service containers network mode. Available values are platform specific, but Compose 
implementation MUST support :

- `none` which disable networking on container
- `host` which give container raw access to host's network interface
- specific syntax `service:name` 

```yml
    network_mode: "host"
    network_mode: "none"
    network_mode: "service:[service name]"
```

### networks

`networks` do define the Networks service container are attached to, referencing entries under the
[top-level `networks` key](#networks).

```yml
services:
  some-service:
    networks:
     - some-network
     - other-network
```

#### aliases

`aliases` declare alternative hostnames for this service on the network. Other containers on the same 
network can use either the service name or this alias to connect to one of the service's containers.

Since `aliases` is network-scoped, the same service can have different aliases on different networks.

> **Note**: A network-wide alias can be shared by multiple containers, and even by multiple services. 
If it is, then exactly which container the name resolves to is not guaranteed.

The general format is shown here.

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

In the example below, service `frontend` will be able to reach `backed` service at
the hostname `backed` or `database` on the `back-tier` network, and service `monitoring` 
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

`pid` sets the PID mode for container created by Compose implemetnation.  
Supported values are platform specific

### ports

Expose container ports.

> **Note**: Port mapping is incompatible with `network_mode: host`

#### Short syntax

Shor syntax is a coma separated string to set host IP, host port and container port.
Host IP if not set will bind to all network interfaces. port can be either a single 
value or a range.

Either specify both ports (`HOST:CONTAINER`), or just the container
port (an ephemeral host port is chosen).

`HOST:CONTAINER` SHOULD always be specified as a (quoted) string, to avoid conflicts 
with [yaml base-60 float](https://yaml.org/type/float.html).

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

#### Long syntax

The long form syntax allows the configuration of additional fields that can't be
expressed in the short form.

- `target`: the port inside the container
- `published`: the publicly exposed port
- `protocol`: the port protocol (`tcp` or `udp`)
- `mode`: `host` for publishing a host port on each node, or `ingress` for a port to be load balanced.

```yaml
ports:
  - target: 80
    published: 8080
    protocol: tcp
    mode: host
```

### restart

`restart` do define the policy platform will apply on container termination.

- `no` is the default restart policy, and it does not restart a container under any circumstance. 
- `always` policy always restarts container until removal. 
- `on-failure` policy restarts a container if the exit code indicates an on-failure error.
- `unless-stopped` policy restarts a container without consideratoin for exit code, but will stop 
restarting when container is stopped by explicit user command.

```yml
    restart: "no"
    restart: always
    restart: on-failure
    restart: unless-stopped
```

### secrets

`secrets` grant access to sensitive data defined by [secrets](secrets) on a per-service basis. Two 
different syntax variants are supported.

> **Note**: The secret MUST exist or be
> [defined in the top-level `secrets` configuration](#secrets)
> of this Compose file.


#### Short syntax

The short syntax variant only specifies the secret name. This grants the
container access to the secret and mounts it at `/run/secrets/<secret_name>`
within the container. The source name and destination mountpoint are both set
to the secret name.

The following example uses the short syntax to grant the `frontend` service
access to the `server-certificate` secret. The value of `server-certificate` is set 
to the contents of the file `./server.cert`.

```yaml
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
the service's task containers.

- `source`: The name of the secret as it exists in Docker.
- `target`: The name of the file to be mounted in `/run/secrets/` in the
  service's task containers. Defaults to `source` if not specified.
- `uid` and `gid`: The numeric UID or GID that owns the file within
  `/run/secrets/` in the service's task containers. Both default to `0` if not
  specified.
- `mode`: The permissions for the file to be mounted in `/run/secrets/`
  in the service's task containers, in octal notation. For instance, `0444`
  represents world-readable. The default in Docker 1.13.1 is `0000`, but is
  be `0444` in newer versions. Secrets cannot be writable because they are mounted
  in a temporary filesystem, so if you set the writable bit, it is ignored. The
  executable bit can be set. If you aren't familiar with UNIX file permission
  modes, you may find this
  [permissions calculator](http://permissions-calculator.org/){: target="_blank" class="_" }
  useful.

The following example sets name of the `server-certificate` to `server.crt` within the
container, sets the mode to `0440` (group-readable) and sets the user and group
to `103`. The value of `server-certificate` is provided by platform by Lookup and not
directly managed by Compose implementation.

```yml
version: "3"
services:
  frontend:
    image: awesome/webapp
    secrets:
      - source: server-certificate
        target: server.cert
        uid: '103'
        gid: '103'
        mode: 0440
secrets:
  server-certificate:
    external: true
```

You can grant a service access to multiple secrets and you can mix long and
short syntax. Defining a secret does not imply granting a service access to it.

### security_opt

`security_opt` override the default labeling scheme for each container.

```yml
security_opt:
  - label:user:USER
  - label:role:ROLE
```

### stop_grace_period

`stop_grace_period` specify how long to wait when attempting to stop a container if it doesn't
handle SIGTERM (or whatever stop signal has been specified with
[`stop_signal`](#stopsignal)), before sending SIGKILL. Specified
as a [duration](#specifying-durations).

```yml
    stop_grace_period: 1s
    stop_grace_period: 1m30s
```    

Default value is 10 seconds for the container to exit before sending SIGKILL.

### stop_signal

`stop_signal` sets an alternative signal to stop the container. By default, container is stopped by
Compose Implementation by sending SIGTERM. Setting an alternative signal using `stop_signal` causes
 Compose Implementation to send that signal instead.

```yml
stop_signal: SIGUSR1
```

### sysctls

`sysctls` define kernel parameters to set in the container. `sysctls` can use either an array or a map.

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

`tmpfs` mount a temporary file system inside the container. Can be a single value or a list.

```yml
tmpfs: /run
```

```yml
tmpfs:
  - /run
  - /tmp
```

### ulimits

`ulimits` override the default ulimits for a container. You can either specify a single limit as an integer or soft/hard limits as a mapping.

```yml
ulimits:
  nproc: 65535
  nofile:
    soft: 20000
    hard: 40000
```

### userns_mode

`userns_mode` allow to set user namespace for service. Supported values are platform specific and might depend on platform configuration

```yml
userns_mode: "host"
```

### volumes

`volumes` define Mount host paths or named volumes into service containers.

You can mount a host path as part of a definition for a single service, and
there is no need to define it in the top level `volumes` key.

But, if you want to reuse a volume across multiple services, then define a named
volume in the [top-level `volumes` key](#volumes), as the containers backing a service 
can be deployed on distinct nodes, and this may be a different node each time the service is updated.

This example shows a named volume (`db-data`) being used by the `backed` service,
and a bind mount defined for a single service 

```yaml
version: "{{ site.compose_file_v3 }}"
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

Short syntax uses a single string with coma separated values to specify a volume mount
(`VOLUME:CONTAINER_PATH`), or an access mode (`VOLUME:CONTAINER:ACCESS_MODE`).

`ACCESS_MODE` can be set read-only by `ro` or read and write by `rw` (default)

You can mount a relative path on the host, that expands relative to
the directory of the Compose configuration file being used. Relative paths
should always begin with `.` or `..`.

#### Long syntax

The long form syntax allows the configuration of additional fields that can't be
expressed in the short form.

- `type`: the mount type `volume`, `bind`, `tmpfs` or `npipe`
- `source`: the source of the mount, a path on the host for a bind mount, or the
  name of a volume defined in the
  [top-level `volumes` key](#volumes). Not applicable for a tmpfs mount.
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

`domainname` allows to declare a custom domain name to use for service container. MUST be a valid RFC 1123 hostname.

### hostname

`hostname` allows to declare a custom host name to use for service container. MUST be a valid RFC 1123 hostname.

### ipc

`ipc` allows to configure the IPC isolation mode set by service container.

### mac_address

`mac_address` allows to set a MAC address for service container.

### privileged

`privileged` configure service container to run with elevated privileged. Support and actual impacts are platform-specific.

### read_only

`read_only` configure service container to be created with a read-only filesystem.

### shm_size

`shm_size` allows to configure the size of the shared memory (`/dev/shm`) allowed by service container.

Value express a byte value as a string in `{amount}{byte unit} :

```
    2b
    1024kb
    2048k
    300m
    1gb
```

The supported units are `b`, `k`, `m` and `g`, and their alternative notation `kb`,
`mb` and `gb`.

### stdin_open

`stdin_open` configure sevice container to run with an alocated stdin

### tty

`tty` configure sevice container to run with a TTY.

### user

`user` override the user set to run container process  by image (i.e. Dockerfile `USER`)

### working_dir

`working_dir` override the working directory set to run container process by image (i.e. Dockerfile `WORKDIR`)




## Networks

Networks are communication channels between services managed by the platform. The networking model exposed to a service is limited to a simple IP connection with target services and external resources, while the Network definition allows to fine-tune the actual implementation provided by the platform.

*TODO* describe configuration attributes 

## Volumes

Volumes are persistent data stored implemented by the platform. The Compose specification offers a neutral abstraction for services to mount volumes, and configuration parameters to allocate them on infrastructure.


`volumes` section allows to configure named volumes that can be reused across multiple services, and are
easily retrieved and inspected using the docker command line or API. Here's an example of a two-service setup where a database's data directory is shared with another service as a volume so that it can be periodically backed up:

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

An entry under the top-level `volumes` key can be empty, in which case it uses the default configuration by the platform. Optionally, you can configure it with the following keys:

### driver

Specify which volume driver should be used for this volume. Default and available values are platform specific. If the driver is not available, the Compose implementation MUST return an error and stop application deployment.

```yml
driver: foobar
```

### driver_opts

Specify a list of options as key-value pairs to pass to the driver for this volume. Those options are driver-dependent. Consult the driver's documentation for more information. 

```yml
volumes:
  example:
    driver_opts:
      type: "nfs"
      o: "addr=10.40.0.199,nolock,soft,rw"
      device: ":/docker/example"
```

### external

If set to `true`, specifies that this volume has been created on platform outside Compose control. Compose implementation MUST NOT attempt to create it, and MUST raises an error if it doesn't exist.

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

### name

`name` set a custom name for this volume. The name field can be used to reference volumes that contain special 
characters. The name is used as is and will **not** be scoped with the stack name.

```yaml
version: "3"
volumes:
  data:
    name: 'my-app-data'
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



## Configs

Configs allow services to adapt their behaviour without the need to rebuild a Docker image. Configs are comparable to Volumes from a service point of view as they are mounted into service's containers filesystem. The actual implementation detail to get configuration provided by the platform can be set from the Configuration definition. 

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

```yaml
configs:
  http_config:
    file: ./httpd.conf
```

Alternatively, `http_config` can be declared as external, doing so Compose implementation will lookup `server-certificate` to expose configuration data to relevant services.

```yaml
configs:
  http_config:
    external: true
```

External configs lookup can also use a distinct key by specifying a `name`. The following
example modifies the previous one to lookup for config using a parameter `HTTP_CONFIG_KEY`. Doing
so the actual lookup key will be set at deployment time by [interpolation](#interpolation) of
variables, but exposed to containers as hard-coded ID `http_config`.

```yaml
configs:
  http_config:
    external: true
    name: '${HTTP_CONFIG_KEY}'
```

Compose file need to explicitly grant access to the configs to relevant services in the application.



## Secrets

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

```yaml
secrets:
  server-certificate:
    file: ./server.cert
```

Alternatively, `server-certificate` can be declared as external, doing so Compose implementation will lookup `server-certificate` to expose secret to relevant services.

```yaml
secrets:
  server-certificate:
    external: true
```

External secrets lookup can also use a distinct key by specifying a `name`. The following
example modifies the previous one to lookup for secret using a parameter `CERTIFICATE_KEY`. Doing
so the actual lookup key will be set at deployment time by [interpolation](#interpolation) of
variables, but exposed to containers as hard-coded ID `server-certificate`.

```yaml
secrets:
  server-certificate:
    external: true
    name: '${CERTIFICATE_KEY}'
```

Compose file need to explicitly grant access to the secrets to relevant services in the application.
