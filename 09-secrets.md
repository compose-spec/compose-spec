## Secrets top-level element

Secrets provides a more secure way of getting sensitive information in to your application's services, so you don't have to rely on using environment variables. If you’re injecting passwords and API keys as environment variables, you risk unintentional information exposure. Environment variables are often available to all processes, and it can be difficult to track access. They can also be printed in logs when debugging errors without your knowledge. Using secrets mitigates these risks.

The top-level `secrets` element defines or references sensitive data that can be granted to your services. The source of the secret is set with either `file` or `environment`.

- `file`: The secret is created with the contents of the file at the specified path.
- `environment`: The secret is created with the value of an environment variable.
- `external`: If set to true, `external` specifies that this secret has already been created. Compose does not attempt to create it, and if it does not exist, an error occurs.
- `name`: The name of the secret object in Docker. This field can be used to
  reference secrets that contain special characters. The name is used as is
  and will **not** be scoped with the project name.

### Examples

#### Example 1

`server-certificate` secret is created as `<project_name>_server-certificate` when the application is deployed, by registering the content of the `server.cert` as a platform secret.

```yml
secrets:
  server-certificate:
    file: ./server.cert
```

#### Example 2 

`token` secret  is created as `<project_name>_token` when the application is deployed,
by registering content of the `OAUTH_TOKEN` environment variable as a platform secret.

```yml
secrets:
  token:
    environment: "OAUTH_TOKEN"
```

#### Example 3 

When `server-certificate` is declared as external, Compose looks up the `server-certificate` secret to expose to relevant services.

```yml
secrets:
  server-certificate:
    external: true
```

External secrets lookup can also use a distinct key by specifying a `name`. The following
example extends the previous example and has Compose looking for a secret using the parameter `CERTIFICATE_KEY`. The actual lookup key is set at deployment time by [interpolation](12-interpolation.md) of
variables, but exposed to containers as hard-coded ID `server-certificate`.

```yml
secrets:
  server-certificate:
    external: true
    name: "${CERTIFICATE_KEY}"
```

If `external` is set to `true`, all other attributes apart from `name` are irrelevant. If Compose detects any other attribute, it SHOULD reject the compose file as invalid.

The Compose file needs to explicitly grant access to the secrets to relevant services in the application.

## Additional resources:

- [Using secrets in Compose](https://docs.docker.com/compose/use-secrets/)
- [Secrets attribute for services top-level element](05-services.md#secrets)
