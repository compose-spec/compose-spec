## Version top-level element

Top-level `version` property is defined by the specification for backward compatibility but is only informative.

A Compose implementation SHOULD NOT use this version to select an exact schema to validate the Compose file, but
prefer the most recent schema at the time it has been designed.

Compose implementations SHOULD validate whether they can fully parse the Compose file. If some fields are unknown, typically
because the Compose file was written with fields defined by a newer version of the specification, Compose implementations
SHOULD warn the user. Compose implementations MAY offer options to ignore unknown fields (as defined by ["loose"](01-status.md#requirements-and-optional-attributes) mode).

## Name top-level element

Top-level `name` property is defined by the specification as project name to be used if user doesn't set one explicitly. 
Compose implementations MUST offer a way for user to override this name, and SHOULD define a mechanism to compute a
default project name, to be used if the top-level `name` element is not set.

Whenever project name is defined by top-level `name` or by some custom mechanism, it MUST be exposed for 
[interpolation](12-interpolation.md) and environment variable resolution as `COMPOSE_PROJECT_NAME`

```yml
services:
  foo:
    image: busybox
    environment:
      - COMPOSE_PROJECT_NAME
    command: echo "I'm running ${COMPOSE_PROJECT_NAME}"
```
