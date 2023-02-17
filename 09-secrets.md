## Secrets top-level element

Secrets are a flavour of Configs focussing on sensitive data, with specific constraint for this usage. As the platform implementation may significantly differ from Configs, dedicated Secrets section allows to configure the related resources.

The top-level `secrets` declaration defines or references sensitive data that can be granted to the services in this
application. The source of the secret is either `file` or `external`.

- `file`: The secret is created with the contents of the file at the specified path.
- `environment`: The secret is created with the value of an environment variable.
- `external`: If set to true, specifies that this secret has already been created. Compose implementation does
  not attempt to create it, and if it does not exist, an error occurs.
- `name`: The name of the secret object in Docker. This field can be used to
  reference secrets that contain special characters. The name is used as is
  and will **not** be scoped with the project name.

In this example, `server-certificate` secret is created as `<project_name>_server-certificate` when the application is deployed,
by registering content of the `server.cert` as a platform secret.

```yml
secrets:
  server-certificate:
    file: ./server.cert
```

In this example, `token` secret  is created as `<project_name>_token` when the application is deployed,
by registering content of the `OAUTH_TOKEN` environment variable as a platform secret.

```yml
secrets:
  token:
    environment: "OAUTH_TOKEN"
```

Alternatively, `server-certificate` can be declared as external, doing so Compose implementation will lookup `server-certificate` to expose secret to relevant services.

```yml
secrets:
  server-certificate:
    external: true
```

External secrets lookup can also use a distinct key by specifying a `name`. The following
example modifies the previous one to look up for secret using a parameter `CERTIFICATE_KEY`. Doing
so the actual lookup key will be set at deployment time by [interpolation](12-interpolation.md) of
variables, but exposed to containers as hard-coded ID `server-certificate`.

```yml
secrets:
  server-certificate:
    external: true
    name: "${CERTIFICATE_KEY}"
```

If `external` is set to `true` and secret configuration has other but `name` attributes set, considering resource is
not managed by compose lifecycle, Compose Implementations SHOULD reject a Compose file as invalid.

Compose file need to explicitly grant access to the secrets to relevant services in the application.

