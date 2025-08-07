# Models top-level element

[![Compose v2.38.0](https://img.shields.io/badge/compose-v2.38.0-blue?style=flat-square)](https://github.com/docker/compose/releases/v2.38.0)

The top-level `models` section defines AI models used by the Compose application.

## Example 1

```yaml
services:
  app:
    image: app
    models:
      - ai_model


models:
  ai_model:
    model: ai/model
```

## Example 2

```yaml
services:
  app:
    image: app
    models:
      my_model:
        endpoint_var: MODEL_URL

models:
  my_model:
    model: ai/model
    context_size: 1024
    runtime_flags: 
      - "--a-flag"
      - "--another-flag=42"
```

## Attributes

- `model` (required) is the OCI artifact identifier for model to be pulled and ran by a model runner, which exposes API to application services
- `context_size` defines the context size for the model (tokens)
- `runtime_flags` defines some raw runtime flags to pass to the inference engine


