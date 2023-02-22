## Extension

Special extension fields can be of any format as long as their name starts with the `x-` character sequence. They can be used
within any structure in a Compose file. This is the sole exception for Compose implementations to silently ignore unrecognized field.

```yml
x-custom:
  foo:
    - bar
    - zot

services:
  webapp:
    image: awesome/webapp
    x-foo: bar
```

The contents of such fields are unspecified by Compose specification, and can be used to enable custom features. Compose implementation to encounter an unknown extension field MUST NOT fail, but COULD warn about unknown field.

For platform extensions, it is highly recommended to prefix extension by platform/vendor name, the same way browsers add
support for [custom CSS features](https://www.w3.org/TR/2011/REC-CSS2-20110607/syndata.html#vendor-keywords)

```yml
service:
  backend:
    deploy:
      placement:
        x-aws-role: "arn:aws:iam::XXXXXXXXXXXX:role/foo"
        x-aws-region: "eu-west-3"
        x-azure-region: "france-central"
```

### Informative Historical Notes

This section is informative. At the time of writing, the following prefixes are known to exist:

| prefix     | vendor/organization |
| ---------- | ------------------- |
| docker     | Docker              |
| kubernetes | Kubernetes          |

### Using extensions as fragments

With the support for extension fields, Compose file can be written as follows to improve readability of reused fragments:

```yml
x-logging: &default-logging
  options:
    max-size: "12m"
    max-file: "5"
  driver: json-file

services:
  frontend:
    image: awesome/webapp
    logging: *default-logging
  backend:
    image: awesome/database
    logging: *default-logging
```

### specifying byte values

Value express a byte value as a string in `{amount}{byte unit}` format:
The supported units are `b` (bytes), `k` or `kb` (kilo bytes), `m` or `mb` (mega bytes) and `g` or `gb` (giga bytes).

```
    2b
    1024kb
    2048k
    300m
    1gb
```

### specifying durations

Value express a duration as a string in the in the form of `{value}{unit}`.
The supported units are `us` (microseconds), `ms` (milliseconds), `s` (seconds), `m` (minutes) and `h` (hours).
Value can can combine multiple values and using without separator.

```
  10ms
  40s
  1m30s
  1h5m30s20ms
```