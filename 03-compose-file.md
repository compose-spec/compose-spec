## Compose file

The Compose file is a [YAML](http://yaml.org/) file defining
[version](04-version-and-name.md) (DEPRECATED),
[services](05-services.md) (REQUIRED),
[networks](06-networks.md),
[volumes](07-volumes.md),
[configs](08-configs.md) and
[secrets](09-secrets.md).
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

### Profiles

Profiles allow to adjust the Compose application model for various usages and environments. A Compose
implementation SHOULD allow the user to define a set of active profiles. The exact mechanism is implementation
specific and MAY include command line flags, environment variables, etc.

The Services top-level element supports a `profiles` attribute to define a list of named profiles. Services without
a `profiles` attribute set MUST always be enabled. A service MUST be ignored by the Compose
implementation when none of the listed `profiles` match the active ones, unless the service is
explicitly targeted by a command. In that case its `profiles` MUST be added to the set of active profiles.
All other top-level elements are not affected by `profiles` and are always active.

References to other services (by `links`, `extends` or shared resource syntax `service:xxx`) MUST not
automatically enable a component that would otherwise have been ignored by active profiles. Instead the
Compose implementation MUST return an error.

#### Illustrative example

```yaml
services:
  foo:
    image: foo
  bar:
    image: bar
    profiles:
      - test
  baz:
    image: baz
    depends_on:
      - bar
    profiles:
      - test
  zot:
    image: zot
    depends_on:
      - bar
    profiles:
      - debug
```

- Compose application model parsed with no profile enabled only contains the `foo` service.
- If profile `test` is enabled, model contains the services `bar` and `baz` which are enabled by the
  `test` profile and service `foo` which is always enabled.
- If profile `debug` is enabled, model contains both `foo` and `zot` services, but not `bar` and `baz`
  and as such the model is invalid regarding the `depends_on` constraint of `zot`.
- If profiles `debug` and `test` are enabled, model contains all services: `foo`, `bar`, `baz` and `zot`.
- If Compose implementation is executed with `bar` as explicit service to run, it and the `test` profile
  will be active even if `test` profile is not enabled _by the user_.
- If Compose implementation is executed with `baz` as explicit service to run, the service `baz` and the
  profile `test` will be active and `bar` will be pulled in by the `depends_on` constraint.
- If Compose implementation is executed with `zot` as explicit service to run, again the model will be
  invalid regarding the `depends_on` constraint of `zot` since `zot` and `bar` have no common `profiles`
  listed.
- If Compose implementation is executed with `zot` as explicit service to run and profile `test` enabled,
  profile `debug` is automatically enabled and service `bar` is pulled in as a dependency starting both
  services `zot` and `bar`.
