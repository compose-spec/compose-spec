# Compose Develop Specification 

> **Note:** 
>
> Develop is an optional part of the Compose Specification

## Introduction

Compose focuses on the development use-case of running applications on a local machine. It also supports some development hooks to improve the velocity of your local workflow, also known as your "inner loop". This document defines how Compose behaves to efficiently assist the developer.

This section defines the development constraints and workflows set by Compose. Only a subset of
Compose file services may require a `develop` subsection.

## Illustrative example

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

## Attributes

The `develop` subsection defines configuration options that are applied by Compose to assist you during development of a service with optimized workflows.

### watch

The `watch` attribute defines a list of rules that control automatic service updates based on local file changes. `watch` is a sequence, each individual item in the sequence defines a rule to be applied by 
Compose to monitor source code for changes. For more information, see [Use Compose Watch](https://docs.docker.com/compose/file-watch/).

#### action

`action` defines the action to take when changes are detected. If `action` is set to:

- `rebuild`, Compose rebuilds the service image based on the `build` section and recreates the service with the updated image.
- `restart`, Compose restarts the service container. [![Compose v2.32.0](https://img.shields.io/badge/compose-v2.32.0-blue?style=flat-square)](https://github.com/docker/compose/releases/v2.32.0)
- `sync`, Compose keeps the existing service container(s) running, but synchronizes source files with container content according to the `target` attribute.
- `sync+restart`, Compose synchronizes source files with container content according to the `target` attribute, and then restarts the container.
- `sync+exec`, Compose synchronizes source files with container content according to the `target` attribute, and then executes a command inside the container. [![Compose v2.32.0](https://img.shields.io/badge/compose-v2.32.0-blue?style=flat-square)](https://github.com/docker/compose/releases/v2.32.0)


#### exec

`exec` is only relevant when `action` is set to `sync+exec`. Comparable to [service hooks](05-services.md#post_start), `exec` is used to defined command to be ran inside container:

- `command`: The command to run after the container has started. This attribute is required.
- `user`: The user to run the command. If not set, the command is run with the same user as the main service command.
- `privileged`: Lets the command run with privileged access.
- `working_dir`: The working directory in which to run the command. If not set, it is run in the same working directory as the main service command.
- `environment`: Sets the environment variables to run the command. The command inherits the `environment` set for the service, this section lets you to append or override values.

```yaml
services:
  frontend:
    image: ...
    develop:
      watch: 
        # sync content then run command to reload service without interruption
        - path: ./etc/config
          action: sync+exec
          target: /etc/config/
          exec:
            command: app reload
```

#### ignore

The `ignore` attribute can be used to define a list of patterns for paths to be ignored. Any updated file
that matches a pattern, or belongs to a folder that matches a pattern, won't trigger services to be re-created. 
The syntax is the same as `.dockerignore` file: 

- `*` matches 0 or more characters in a file name. 
- `?` matches a single character in file name. 
- `*/*` matches two nested folders with arbitrary names
- `**` matches an arbitrary number of nested folders

If the build context includes a `.dockerignore` file, the patterns in this file is loaded as implicit content
for the `ignores` file, and values set in the Compose model are appended.

#### path

`path` attribute defines the path to source code (relative to the project directory) to monitor for changes. Updates to any file
inside the path, which doesn't match any `ignore` rule, triggers the configured action.

#### target

`target` attribute only applies when `action` is configured for `sync`. Files within `path` with changes are synchronized
with container filesystem, so that the latter is always running with up-to-date content.

