## Configs top-level element

Configs allow services to adapt their behaviour without the need to rebuild a Docker image. Configs are comparable to Volumes from a service point of view as they are mounted into service's container's filesystem. The actual implementation detail to get the configuration provided by the platform can be set from the Configuration definition.

When granted access to a config, the config content is mounted as a file in the container. The location of the mount point within the container defaults to `/<config-name>` in Linux containers and `C:\<config-name>` in Windows containers.

By default, the config must be owned by the user running the container command but can be overridden by service configuration.
By default, the config must have world-readable permissions (mode 0444), unless service is configured to override this.

Services can only access configs when explicitly granted by a [`configs`](05-services.md#configs) subsection.

The top-level `configs` declaration defines or references
configuration data that can be granted to the services in this
application. The source of the config is either `file` or `external`.

- `file`: The config is created with the contents of the file at the specified path.
- `external`: If set to true, specifies that this config has already been created. Compose does not
  attempt to create it, and if it does not exist, an error occurs.
- `name`: The name of the config object on Platform to lookup. This field can be used to
  reference configs that contain special characters. The name is used as is
  and will **not** be scoped with the project name.


In this example, `<project_name>_http_config` is created when the application is deployed,
by registering the content of the `httpd.conf` as configuration data.

```yml
configs:
  http_config:
    file: ./httpd.conf
```

Alternatively, `http_config` can be declared as external. Compose will lookup `http_config` to expose configuration data to relevant services.

```yml
configs:
  http_config:
    external: true
```

External configs lookup can also use a distinct key by specifying a `name`. The following
example modifies the previous one to lookup for config using a parameter `HTTP_CONFIG_KEY`. The
the actual lookup key will be set at deployment time by [interpolation](12-interpolation.md) of
variables, but exposed to containers as hard-coded ID `http_config`.

```yml
configs:
  http_config:
    external: true
    name: "${HTTP_CONFIG_KEY}"
```

If `external` is set to `true` and the configuration has other but `name` attributes set, considering resource is
not managed by compose lifecycle, Compose rejects the Compose file as invalid.

Compose file need to explicitly grant access to the configs to relevant services in the application.

