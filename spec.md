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

```yaml
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

```yaml
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

```yaml
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

```yaml
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

```yaml
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

```yaml
credential_spec:
  file: my-credential-spec.json
```

When using `registry:`, the credential spec is read from the Windows registry on
the daemon's host. A registry value with the given name must be located in:

    HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Virtualization\Containers\CredentialSpecs

The following example load the credential spec from a value named `my-credential-spec`
in the registry:

```yaml
credential_spec:
  registry: my-credential-spec
```  

#### Example gMSA configuration
When configuring a gMSA credential spec for a service, you only need
to specify a credential spec with `config`, as shown in the following example:
```
version: "3.8"
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

- Compose implementation MUST start services in dependency order. In the following
  example, `db` and `redis` are started before `web`. 

- Compose implementation MUST stop services in dependency order. In the following
  example, `web` is stopped before `db` and `redis`.

Simple example:

```yaml
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

```yaml
devices:
  - "/dev/ttyUSB0:/dev/ttyUSB0"
```

### dns

`dns` define custom DNS servers to set on container network interface configuration. Can be a single value or a list.

```yaml
dns: 8.8.8.8
```

```yaml
dns:
  - 8.8.8.8
  - 9.9.9.9
```

### dns_search

`dns` define custom DNS search domainsto set on container network interface configuration. Can be a single value or a list.

```yaml
dns_search: example.com
``` 

```yaml
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

```yaml
entrypoint: /code/entrypoint.sh
```

The entrypoint can also be a list, in a manner similar to
[Dockerfile](https://docs.docker.com/engine/reference/builder/#cmd):

```yaml
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


```yaml
env_file: .env
```

`env_file` can also be a list. The files in the list are processed from the top down. 
For the same variable specified in file `a.env` and assigned a different value in file 
`b.env`, if `b.env` is listed below (after), then the value from `b.env` stands. 

```yaml
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
