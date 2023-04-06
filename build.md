# The Compose Specification - Build support
{:.no_toc}

*Note:* Build is an OPTIONAL part of the Compose Specification

* ToC
{:toc}

## Introduction

Compose specification is a platform-neutral way to define multi-container applications. A Compose implementation
focusing on development use-case to run application on local machine will obviously also support (re)building
application from sources. The Compose Build specification allows to define the build process within a Compose file
in a portable way.

## Definitions

Compose Specification is extended to support an OPTIONAL `build` subsection on services. This section define the
build requirements for service container image. Only a subset of Compose file services MAY define such a Build
subsection, others being created based on `Image` attribute. When a Build subsection is present for a service, it
is *valid* for a Compose file to miss an `Image` attribute for corresponding service, as Compose implementation
can build image from source.

Build can be either specified as a single string defining a context path, or as a detailed build definition.

In the former case, the whole path is used as a Docker context to execute a docker build, looking for a canonical
`Dockerfile` at context root. Context path can be absolute or relative, and if so relative path MUST be resolved
from Compose file parent folder. As an absolute path prevent the Compose file to be portable, Compose implementation
SHOULD warn user accordingly.

In the later case, build arguments can be specified, including an alternate `Dockerfile` location. This one can be
absolute or relative path. If Dockerfile path is relative, it MUST be resolved from context path.  As an absolute
path prevent the Compose file to be portable, Compose implementation SHOULD warn user if an absolute alternate
Dockerfile path is used.

## Consistency with Image

When service definition do include both `Image` attribute and a `Build` section, Compose implementation can't
guarantee a pulled image is strictly equivalent to building the same image from sources. Without any explicit
user directives, Compose implementation with Build support MUST first try to pull Image, then build from source
if image was not found on registry. Compose implementation MAY offer options to customize this behaviour by user
request.

## Publishing built images

Compose implementation with Build support SHOULD offer an option to push built images to a registry. Doing so, it
MUST NOT try to push service images without an `Image` attribute. Compose implementation SHOULD warn user about
missing `Image` attribute which prevent image being pushed.

Compose implementation MAY offer a mechanism to compute an `Image` attribute for service when not explicitly
declared in yaml file. In such a case, the resulting Compose configuration is considered to have a valid `Image`
attribute, whenever the actual raw yaml file doesn't explicitly declare one.

## Illustrative sample

The following sample illustrates Compose specification concepts with a concrete sample application. The sample is non-normative.

```yaml
services:
  frontend:
    image: awesome/webapp
    build: ./webapp

  backend:
    image: awesome/database
    build:
      context: backend
      dockerfile: ../backend.Dockerfile

  custom:
    build: ~/custom
```

When used to build service images from source, such a Compose file will create three docker images:

* `awesome/webapp` docker image is built using `webapp` sub-directory within Compose file parent folder as docker build context. Lack of a `Dockerfile` within this folder will throw an error.
* `awesome/database` docker image is built using `backend` sub-directory within Compose file parent folder. `backend.Dockerfile` file is used to define build steps, this file is searched relative to context path, which means for this sample `..` will resolve to Compose file parent folder, so `backend.Dockerfile` is a sibling file.
* a docker image is built using `custom` directory within user's HOME as docker context. Compose implementation warn user about non-portable path used to build image.

On push, both `awesome/webapp` and `awesome/database` docker images are pushed to (default) registry. `custom` service image is skipped as no `Image` attribute is set and user is warned about this missing attribute.

## Build definition

The `build` element define configuration options that are applied by Compose implementations to build Docker image from source.
`build` can be specified either as a string containing a path to the build context or a detailed structure:

Using this string syntax, only the build context can be configured as either:
- a relative path to the Compose file's parent folder. This path MUST be a directory and contain a `Dockerfile`

```yml
services:
  webapp:
    build: ./dir
```

- a git repository URL. Git URLs accept context configuration in their fragment section, separated by a colon (`:`). 
The first part represents the reference that Git will check out, and can be either a branch, a tag, or a remote reference. 
The second part represents a subdirectory inside the repository that will be used as a build context.

```yml
services:
  webapp:
    build: https://github.com/mycompany/example.git#branch_or_tag:subdirectory
```

Alternatively `build` can be an object with fields defined as follow

### context (REQUIRED)

`context` defines either a path to a directory containing a Dockerfile, or a url to a git repository.

When the value supplied is a relative path, it MUST be interpreted as relative to the location of the Compose file.
Compose implementations MUST warn user about absolute path used to define build context as those prevent Compose file
from being portable.

```yml
build:
  context: ./dir
```

```yml
services:
  webapp:
    build: https://github.com/mycompany/webapp.git
```

### dockerfile

`dockerfile` allows to set an alternate Dockerfile. A relative path MUST be resolved from the build context.
Compose implementations MUST warn user about absolute path used to define Dockerfile as those prevent Compose file
from being portable. When set, `dockerfile_inline` attribute is not allowed and a Compose Implementation SHOULD 
reject any Compose file having both set.

```yml
build:
  context: .
  dockerfile: webapp.Dockerfile
```

### dockerfile_inline

`dockerfile_inline` allows to define Dockerfile content as inlined string in a Compose file. When set, `dockerfile` 
attribute is not allowed  and a Compose Implementation SHOULD reject any Compose file having both set.

Use of YAML multi-line string syntax is recommended to define Dockerfile content:

```yml
build:
  context: .
  dockerfile_inline: |
    FROM baseimage
    RUN some command
```


### args

`args` define build arguments, i.e. Dockerfile `ARG` values.

Using following Dockerfile:

```Dockerfile
ARG GIT_COMMIT
RUN echo "Based on commit: $GIT_COMMIT"
```

`args` can be set in Compose file under the `build` key to define `GIT_COMMIT`. `args` can be set a mapping or a list:

```yml
build:
  context: .
  args:
    GIT_COMMIT: cdc3b19
```

```yml
build:
  context: .
  args:
    - GIT_COMMIT=cdc3b19
```

Value can be omitted when specifying a build argument, in which case its value at build time MUST be obtained by user interaction,
otherwise build arg won't be set when building the Docker image.

```yml
args:
  - GIT_COMMIT
```

### ssh

`ssh` defines SSH authentications that the image builder SHOULD use during image build (e.g., cloning private repository)

`ssh` property syntax can be either:
* `default` - let the builder connect to the ssh-agent.
* `ID=path` - a key/value definition of an ID and the associated path. Can be either a [PEM](https://en.wikipedia.org/wiki/Privacy-Enhanced_Mail) file, or path to ssh-agent socket

Simple `default` sample
```yaml
build:
  context: .
  ssh: 
    - default   # mount the default ssh agent
```
or 
```yaml
build:
  context: .
  ssh: ["default"]   # mount the default ssh agent
```

Using a custom id `myproject` with path to a local SSH key:
```yaml
build:
  context: .
  ssh: 
    - myproject=~/.ssh/myproject.pem
```
Image builder can then rely on this to mount SSH key during build.
For illustration, [BuildKit extended syntax](https://github.com/compose-spec/compose-spec/pull/234/%5Bmoby/buildkit@master/frontend/dockerfile/docs/syntax.md#run---mounttypessh%5D(https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/syntax.md#run---mounttypessh)) can be used to mount ssh key set by ID and access a secured resource:

`RUN --mount=type=ssh,id=myproject git clone ...`

### cache_from

`cache_from` defines a list of sources the Image builder SHOULD use for cache resolution.

Cache location syntax MUST follow the global format `[NAME|type=TYPE[,KEY=VALUE]]`. Simple `NAME` is actually a shortcut notation for `type=registry,ref=NAME`.

Compose Builder implementations MAY support custom types, the Compose Specification defines canonical types which MUST be supported:

- `registry` to retrieve build cache from an OCI image set by key `ref`


```yml
build:
  context: .
  cache_from:
    - alpine:latest
    - type=local,src=path/to/cache
    - type=gha
```

Unsupported caches MUST be ignored and not prevent user from building image.

### cache_to

`cache_to` defines a list of export locations to be used to share build cache with future builds.

```yml
build:
  context: .
  cache_to: 
   - user/app:cache
   - type=local,dest=path/to/cache
```

Cache target is defined using the same `type=TYPE[,KEY=VALUE]` syntax defined by [`cache_from`](#cache_from). 

Unsupported cache target MUST be ignored and not prevent user from building image.

### additional_contexts

`additional_contexts` defines a list of named contexts the image builder SHOULD use during image build.

`additional_contexts` can be a mapping or a list:

```yml
build:
  context: .
  additional_contexts:
    - resources=/path/to/resources
    - app=docker-image://my-app:latest
    - source=https://github.com/myuser/project.git
```

```yml
build:
  context: .
  additional_contexts:
    resources: /path/to/resources
    app: docker-image://my-app:latest
    source: https://github.com/myuser/project.git
```

When used as a list, the syntax should follow the `NAME=VALUE` format, where `VALUE` is a string. Validation beyond that is the responsibility of the image builder (and is builder specific).

The Compose implementation SHOULD warn the user if the image builder does not support additional contexts and MAY list the unused contexts.

Illustrative examples of how this is used in Buildx can be found [here](https://github.com/docker/buildx/blob/master/docs/reference/buildx_build.md#-additional-build-contexts---build-context).

### extra_hosts

`extra_hosts` adds hostname mappings at build-time. Use the same syntax as [extra_hosts](05-services.md#extra_hosts).

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

### isolation

`isolation` specifies a build’s container isolation technology. Like [isolation](05-services.md#isolation) supported values
are platform-specific.

### privileged

`privileged` configures the service image to build with elevated privileges. Support and actual impacts are platform-specific.

```yml
build:
  context: .
  privileged: true
```

### labels

`labels` add metadata to the resulting image. `labels` can be set either as an array or a map.

reverse-DNS notation SHOULD be used to prevent labels from conflicting with those used by other software.

```yml
build:
  context: .
  labels:
    com.example.description: "Accounting webapp"
    com.example.department: "Finance"
    com.example.label-with-empty-value: ""
```

```yml
build:
  context: .
  labels:
    - "com.example.description=Accounting webapp"
    - "com.example.department=Finance"
    - "com.example.label-with-empty-value"
```

### no_cache

`no_cache` disables image builder cache and enforce a full rebuild from source for all image layers. This only
applies to layers declared in the Dockerfile, referenced images COULD be retrieved from local image store whenever tag
has been updated on registry (see [pull](#pull)).

### pull

`pull` require the image builder to pull referenced images (`FROM` Dockerfile directive), even if those are already 
available in the local image store.

### shm_size

`shm_size` set the size of the shared memory (`/dev/shm` partition on Linux) allocated for building Docker image. Specify
as an integer value representing the number of bytes or as a string expressing a [byte value](11-extension.md#specifying-byte-values).

```yml
build:
  context: .
  shm_size: '2gb'
```

```yaml
build:
  context: .
  shm_size: 10000000
```

### target

`target` defines the stage to build as defined inside a multi-stage `Dockerfile`.

```yml
build:
  context: .
  target: prod
```

### secrets
`secrets` grants access to sensitive data defined by [secrets](05-services.md#secrets) on a per-service build basis. Two
different syntax variants are supported: the short syntax and the long syntax.

Compose implementations MUST report an error if the secret isn't defined in the
[`secrets`](09-secrets.md) section of this Compose file.

#### Short syntax

The short syntax variant only specifies the secret name. This grants the
container access to the secret and mounts it as read-only to `/run/secrets/<secret_name>`
within the container. The source name and destination mountpoint are both set
to the secret name.

The following example uses the short syntax to grant the build of the `frontend` service
access to the `server-certificate` secret. The value of `server-certificate` is set
to the contents of the file `./server.cert`.

```yml
services:
  frontend:
    build: 
      context: .
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
services:
  frontend:
    build:
      context: .
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

Service builds MAY be granted access to multiple secrets. Long and short syntax for secrets MAY be used in the
same Compose file. Defining a secret in the top-level `secrets` MUST NOT imply granting any service build access to it.
Such grant must be explicit within service specification as [secrets](05-services.md#secrets) service element.

### tags

`tags` defines a list of tag mappings that MUST be associated to the build image. This list comes in addition of 
the `image` [property defined in the service section](05-services.md#image)

```yml
tags:
  - "myimage:mytag"
  - "registry/username/myrepos:my-other-tag"
```

### platforms

`platforms` defines a list of target [platforms](05-services.md#platform).

```yml
build:
  context: "."
  platforms:
    - "linux/amd64"
    - "linux/arm64"
```

When the `platforms` attribute is omitted, Compose implementations MUST include the service's platform
in the list of the default build target platforms.

When the `platforms` attribute is defined, Compose implementations SHOULD require it includes the service's 
platform, otherwise users won't be able to run images they built.

Compose implementations SHOULD report an error in the following cases:
* when the list contains multiple platforms but the implementation is incapable of storing multi-platform images
* when the list contains an unsupported platform
```yml
build:
  context: "."
  platforms:
    - "linux/amd64"
    - "unsupported/unsupported"
```
* when the list is non-empty and does not contain the service's platform
```yml
services:
  frontend:
    platform: "linux/amd64"
    build:
      context: "."
      platforms:
        - "linux/arm64"
```

## Implementations

* [docker-compose](https://docs.docker.com/compose)
* [buildX bake](https://docs.docker.com/buildx/working-with-buildx/)

