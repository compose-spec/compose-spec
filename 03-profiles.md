## Profiles

Profiles let you adjust the Compose application model for various usages and environments. 
Compose SHOULD allow you to define a set of active profiles. The exact mechanism is implementation
specific and MAY include command line flags, environment variables, etc.

The [services top-level element](05-services.md) supports a `profiles` attribute to define a list of named profiles. Services without
a `profiles` attribute set MUST always be enabled. A service MUST be ignored by Compose when none of the listed `profiles` match the active ones, unless the service is
explicitly targeted by a command. In that case its `profiles` MUST be added to the set of active profiles.
All other top-level elements are not affected by `profiles` and are always active.

References to other services, by `links`, `extends`, or shared resource syntax `service:xxx` for example, MUST not
automatically enable a component that would otherwise have been ignored by active profiles. Instead an error MUST be returned.

### Illustrative example

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

In the example above:
- `foo` is the only service that has no `profiles` attribute so is always enabled.
- If the profile `test` is enabled, the application contains the services `bar` and `baz`, which are enabled by the `test` profile, and service `foo`, which is always enabled.
- If the profile `debug` is enabled, the application contains both `foo` and `zot` services, but not `bar` and `baz`, and as such the model is invalid regarding the `depends_on` constraint of `zot`.
- If the profiles `debug` and `test` are enabled, the application contains all services: `foo`, `bar`, `baz` and `zot`.
- If the service `bar` is explicitly told to run, the `test` profile is active even though it was not enabled.
- If the service `baz` is explicitly told to run, the service `baz` and the
  profile `test` will be active and `bar` will be pulled in by the `depends_on` constraint.
- If Compose implementation is executed with `zot` as explicit service to run, again the model will be
  invalid regarding the `depends_on` constraint of `zot`, since `zot` and `bar` have no common `profiles`
  listed.
- If Compose implementation is executed with `zot` as explicit service to run and profile `test` enabled,
  profile `debug` is automatically enabled and service `bar` is pulled in as a dependency starting both
  services `zot` and `bar`.

