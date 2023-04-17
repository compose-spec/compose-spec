## Merge and override

Compose implementations SHOULD allow users to define a Compose application model through multiple Compose files. 
When doing so, a Compose implementation MUST follow the rules declared in this section to merge Compose files.

### Mapping

A YAML `mapping` gets merged by adding missing entries and merging the conflicting ones.

Merging the following example YAML trees:
```yaml
services:
  foo:
    key1: value1
    key2: value2
```    

```yaml
services:
  foo:
    key2: VALUE
    key3: value3
```

MUST result in a Compose application model equivalent to the YAML tree:

```yaml
services:
  foo:
    key1: value1
    key2: VALUE
    key3: value3
```

### Sequence

A YAML `sequence` is merged by appending values from the overriding Compose file to the previous one.

Merging the following example YAML trees:
```yaml
services:
  foo:
    DNS:
      - 1.1.1.1
```    

```yaml
services:
  foo:
    DNS: 
      - 8.8.8.8
```

MUST result in a Compose application model equivalent to the YAML tree:

```yaml
services:
  foo:
    DNS:
      - 1.1.1.1
      - 8.8.8.8
```

## Exceptions

There are exceptions to those rules:

### Shell commands

Service's [command](#command), [entrypoint](#entrypoint) and [healthcheck](#healthcheck) `test`: 
For usability, the value MUST be overridden by the latest Compose file, and not appended.

Merging the following example YAML trees:
```yaml
services:
  foo:
    command: ["echo", "foo"]
```    

```yaml
services:
  foo:
    command: ["echo", "bar"]
```

MUST result in a Compose application model equivalent to the YAML tree:

```yaml
services:
  foo:
    DNS:
      command: ["echo", "bar"]
```


### Unique resources

Applies to service [ports](#ports), [volumes](#volumes), [secrets](#secrets) and [configs](#configs).
While these types are modeled in a Compose file as a sequence, they have special uniqueness requirements:

| attribute   | unique key               |
|-------------|--------------------------|
| volumes     |  target                  |
| secrets     |  source                  |
| configs     |  source                  |
| ports       |  {ip, target, published, protocol}   |

While merging Compose files, a Compose implementation MUST append new entries that do not violate a uniqueness constraint and merge entries that share a unique key.

Merging the following example YAML trees:
```yaml
services:
  foo:
    volumes:
      - foo:/work
```    

```yaml
services:
  foo:
    volumes:
      - bar:/work
```

MUST result in a Compose application model equivalent to the YAML tree:

```yaml
services:
  foo:
    volumes:
      - bar:/work
```

### Reset value

In addition to the previously described mechanism, an override Compose file can also be used to remove elements from application model.
For this purpose, custom YAML tag `!reset` can be set to override value set by the overriden Compose file, and replace with default
value or `null` on target attribute.

Merging the following example YAML trees:
```yaml
services:
  foo:
    build:
      context: /path
```    

```yaml
services:
  foo:
    build: !reset
```

MUST result in a Compose application model equivalent to the YAML tree:

```yaml
services:
  foo:
    build: null
```
