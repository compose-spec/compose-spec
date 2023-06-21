## Version top-level element

Top-level `version` property is defined by the specification for backward compatibility but is only informative.

Compose doesn't use this version to select an exact schema to validate the Compose file, but
prefer the most recent schema at the time it has been designed.

Compose validates whether it can fully parse the Compose file. If some fields are unknown, typically
because the Compose file was written with fields defined by a newer version of the specification, you'll receive a warning message. Compose offers options to ignore unknown fields (as defined by ["loose"](01-status.md#requirements-and-optional-attributes) mode).

## Name top-level element

Top-level `name` property is defined by the specification as project name to be used if user doesn't set one explicitly.
Compose offers a way for you to override this name, and sets a
default project name to be used if the top-level `name` element is not set.

Whenever project name is defined by top-level `name` or by some custom mechanism, it must be exposed for
[interpolation](12-interpolation.md) and environment variable resolution as `COMPOSE_PROJECT_NAME`

```yml
services:
  foo:
    image: busybox
    environment:
      - COMPOSE_PROJECT_NAME
    command: echo "I'm running ${COMPOSE_PROJECT_NAME}"
```
