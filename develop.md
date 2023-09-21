# The Compose Specification - Development support
{:.no_toc}

*Note:* Development support is an OPTIONAL part of the Compose Specification.

* ToC
{:toc}

## Introduction

Compose specification is a platform-neutral way to define multi-container applications. Compose focusing on development
use-case to run application on local machine also supports some development hooks to improve developers velocity on their
local workflow (also known as "inner loop"). This document defines how Compose behaves to efficiently assist the developer.

## Definitions

Compose Specification supports an OPTIONAL `develop` subsection on services. This section defines the development
constraints and workflows to be set by Compose to assist the developer working on service codebase. Only a subset of
Compose file services may define such a `develop` subsection.

## Illustrative sample

The following sample illustrates Compose specification concepts with a concrete sample application:

```yaml
services:
  frontend:
    image: example/webapp
    build: ./webapp
    develop:
      watch: 
        # sync static content
        - path: ./webapp/html
          action: sync
          target: /var/www
          ignore:
            - node_modules/

  backend:
    image: example/backend
    build: ./backend
    develop:
      watch: 
        # rebuild image and recreate service
        - path: ./backend/src
          action: rebuild
```

## develop

The `develop` element defines configuration options that are applied by Compose to assist developer during development of
a service with optimized workflows.

### watch

The `watch` attribute defines event sources and strategy to be adopted to update a running service as Compose detects
changes in source code. `watch` is a sequence, each individual item in the sequence defines a rule to be applied by 
Compose to monitor source code for changes.

##### action

`action` define the action to take place as changes have been detected:

- `rebuild` strategy will rebuild service image based on the `build` section and recreate the service with updated image.
- `sync` strategy keep existing service container(s) running, but synchronize source files with container content according to `target` attribute. 


##### ignore

The `ignore` attribute can be used to define a list of patterns for paths to be ignored. Any updated file
that matches a pattern, or belongs to a folder that matches a pattern, won't trigger service to be re-created. 
The syntax is the same as `.dockerignore` file: 

- `*` matches 0 or more characters in a file name. 
- `?` matches a single character in file name. 
- `*/*` matches two nested folders with arbitrary names
- `**` matches an arbitrary number of nested folders

If the build context includes a `.dockerignore` file, the patterns in this file are loaded as implicit content
for the `ignores` file, and values set in the Compose model are appended.

##### path

`path` attribute defines the path to source code (relative to the project directory) to monitor for changes. Updates to any file
inside path, which doesn't match any `ignore` rule, will trigger configured action.

##### target

`target` attribute only applies when `action` is configured for `sync`. Files within `path` with changes will be synchronized
with container filesystem, so that the latter is always running with up-to-date content.

