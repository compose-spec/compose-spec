## Fragments

It is possible to re-use configuration fragments using [YAML anchors](http://www.yaml.org/spec/1.2/spec.html#id2765878).

```yml
volumes:
  db-data: &default-volume
    driver: default
  metrics: *default-volume
```

In previous sample, an _anchor_ is created as `default-volume` based on `db-data` volume specification. It is later reused by _alias_ `*default-volume` to define `metrics` volume. Same logic can apply to any element in a Compose file. Anchor resolution MUST take place
before [variables interpolation](12-interpolation.md), so variables can't be used to set anchors or aliases.

It is also possible to partially override values set by anchor reference using the
[YAML merge type](http://yaml.org/type/merge.html). In following example, `metrics` volume specification uses alias
to avoid repetition but override `name` attribute:

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
