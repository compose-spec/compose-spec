# The Compose Specification - Development support
{:.no_toc}

*Note:* Development support is an OPTIONAL part of the Compose Specification.

* ToC
{:toc}

## Introduction

Compose specification is a platform-neutral way to define multi-container applications. A Compose implementation
focusing on development use-case to run application on local machine SHOULD also support some development hooks to improve developers velocity on their local workflow (also known as "inner loop"). The Compose Development specification allows to define how a Compose implementation SHOULD behave to efficiently assist the developer.

## Definitions

Compose Specification is extended to support an OPTIONAL `develop` subsection on services. This section
defines the development constraints and workflows to be set by a Compose Implementation to assist the developer
working on service codebase. Only a subset of Compose file services MAY define such a Development subsection.

## Illustrative sample

The following sample illustrates Compose specification concepts with a concrete sample application. The sample is non-normative.

```yaml
services:
  frontend:
    image: awesome/webapp
    build: ./webapp
    develop:
      watch: 
        ignores:
            - node_modules/
        quiet-period: 5s
        trigger: 
          # synchronize src with container's /app
          strategy: sync
          paths: 
            - src:/app

  backend:
    image: awesome/backend
    build: ./backend
    develop:
      watch: 
        # rebuild image and recreate service
        trigger: build

  db:
    image: awesome/database
    develop:
      watch: 
        trigger:
          strategy: sync
          # get service to reload configuration
          signal: SIGHUP
```

## Developemnt mode definition

The `develop` element defines configuration options that are applied by Compose implementations to assist developer during development of a service.

### watch

The `watch` attribute defines strategy to be adopted to update a running service as Compose Implementation detects 
changes in service source code. "Source code" to watch is identified by the `build.context` section.

#### ignores

The `ignores` attribute can be used to define a list of patterns for paths to be ignored. Any updated file 
that matches a pattern, or belongs to a folder that matches a pattern, won't trigger service to be re-created. The 
syntax is the same as `.dockerignore` file: 

- `*` matches 0 or more characters in a file name. 
- `?` matches a single character in file name. 
- `*/*` matches two nested folders with arbitrary names
- `**` matches an arbitrary number of nested folders

If the build context includes a `.dockerignore` file, the patterns in this file are loaded as implicit content
for the `ignores` file, and values set in the Compose model are appended.

#### quiet_period

If a `quiet_period` is set, Compose Implementations MUST wait at least configured delay before actually refreshing
service image. Using this allows to avoid service image to be re-created many times in a row as multiple files get
updated.

#### trigger

`trigger` define the actions to take place as changes have been detected.

Short syntax can be used to set the strategy name to be adopted by Compose Implementation. This allows to set simple
strategies which do not require additional configuration options:

```yaml
    develop:
      watch: 
        # rebuild image and recreate service
        trigger: build
```        

Long-syntax allow user to pass additional configuration options, specific to selected strategy. `strategy` attribute
is REQUIRED:

- `build` strategy will rebuild service image based on the `build` section and restart the service with updated image.
- `sync` strategy keep existing service contianers running, but can synchronize source files with container content. 

`paths` attribute can be set with `sync` strategy to define local files within service's `build.context` to be synchronized  with
container's filesystem. This attribute uses `source:target` syntax, similar to [spec.md#volumes](service volumes).

`signal` attribute can be set with `sync` strategy to define a signal sent to container after synchronization completed, to force service to
reload source/configuration or restart. Syntax is the same as [spec.md#stop_signal](stop_signal)
