## Fragments

With Docker Compose you can use built-in [YAML](http://www.yaml.org/spec/1.2/spec.html#id2765878) features to make your Compose file neater and more efficient. Anchors and aliases let you create re-usable blocks. This is useful if you start to find common configurations that span multiple services. Having re-usable blocks minimizes potential mistakes.

Anchors are created using the `&` sign. The sign is followed by an alias name. You can use this alias with the `*` sign later to reference the value following the anchor. Make sure there is no space between the `&` and the `*` characters and the following alias name.

### Single-line anchor example

```yml
volumes:
  db-data: &default-volume
    driver: default
  metrics: *default-volume
```

In the example above, a `default-volume` anchor is created based on the `db-data` volume. It is later reused by the alias `*default-volume` to define the `metrics` volume. 

Anchor resolution MUST take place before [variables interpolation](12-interpolation.md), so variables can't be used to set anchors or aliases.

### Multi-line anchor example

```yml
services:
  first:
    image: my-image:latest
    environment: &env
      - CONFIG_KEY
      - EXAMPLE_KEY
      - DEMO_VAR
  second:
    image: another-image:latest
    environment: *env
```

If you have an anchor that you want to use in more than one service, use it in conjunction with an [extension](11-extension.md) to make your Compose file easier to maintain.

### Extend anchor values examples

You may want to extend the anchor to add additional values or partially override values. You can do this by using the
[YAML merge type](http://yaml.org/type/merge.html). 

#### Example 1

In the following example, `metrics` volume specification uses alias
to avoid repetition but overrides `name` attribute:

```yml

services:
  backend:
    image: awesome/database
    volumes:
      - db-data
      - metrics
volumes:
  db-data: &default-volume
    driver: default
    name: "data"
  metrics:
    <<: *default-volume
    name: "metrics"
```

#### Example 2

```yml
services:
  first:
    image: my-image:latest
    environment: &env
      - CONFIG_KEY
      - EXAMPLE_KEY
      - DEMO_VAR
  second:
    image: another-image:latest
    environment:
      <<: *env
      - AN_EXTRA_KEY
      - SECOND_SPECIFIC_KEY
```

The `second` service now pulls in the base environment configuration from the `env` anchor and adds two additional configuration items.
