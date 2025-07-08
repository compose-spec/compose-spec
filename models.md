# Models top-level element

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

## Example 2 (with environment variables)

```yaml
services:
  app:
    image: app
    models:
      my_model:
        endpoint_var: MODEL_URL # app service will receive the model endpoint from this environment variable
        model_var: MODEL_NAME   # app service will receive the model name from this environment variable

models:
  my_model:
    model: ai/model    
```

## Example 3 (with context size and runtime flags)

```yaml
services:
  app:
    image: app
    models:
      my_model:
        endpoint_var: MODEL_URL
        model_var: MODEL_NAME

models:
  my_model:
    model: ai/model
    context_size: 1024
    runtime_flags:
      - "--a-flag"
      - "--another-flag=42"
```

## Example 4 (common runtime configurations)

### Development

```yaml
services:
  app:
    image: app
    models:
      dev_model:
        endpoint_var: DEV_URL
        model_var: DEV_MODEL

models:
  dev_model:
    model: ai/model
    context_size: 4096
    runtime_flags:
      - "--verbose"                       # Set verbosity level to infinity
      - "--verbose-prompt"                # Print a verbose prompt before generation
      - "--log-prefix"                    # Enable prefix in log messages
      - "--log-timestamps"                # Enable timestamps in log messages
      - "--log-colors"                    # Enable colored logging
```

### Conservative with disabled reasoning

```yaml
services:
  app:
    image: app
    models:
      conservative_model:
        endpoint_var: CONSERVATIVE_URL
        model_var: CONSERVATIVE_MODEL

models:
  conservative_model:
    model: ai/model
    context_size: 4096
    runtime_flags:      
      - "--temp"                # Temperature
      - "0.1"
      - "--top-k"               # Top-k sampling
      - "1"
      - "--reasoning-budget"    # Disable reasoning
      - "0"
```

### Creative with high randomness

```yaml
services:
  app:
    image: app
    models:
      creative_model:
        endpoint_var: CREATIVE_URL
        model_var: CREATIVE_MODEL

models:
  creative_model:
    model: ai/model
    context_size: 4096
    runtime_flags:      
      - "--temp"                # Temperature
      - "1"
      - "--top-p"               # Top-p sampling
      - "0.9"
```

### Highly deterministic

```yaml
services:
  app:
    image: app
    models:
      deterministic_model:
        endpoint_var: DET_URL
        model_var: DET_MODEL

models:
  deterministic_model:
    model: ai/model
    context_size: 4096
    runtime_flags:
      - "--temp"                # Temperature
      - "0"
      - "--top-k"               # Top-k sampling
      - "1"
```

### Concurrent processing

```yaml
services:
  app:
    image: app
    models:
      concurrent_model:
        endpoint_var: CONCURRENT_URL
        model_var: CONCURRENT_MODEL

models:
  concurrent_model:
    model: ai/model
    context_size: 2048
    runtime_flags:      
      - "--threads"             # Number of threads to use during generation
      - "8"      
      - "--mlock"               # Lock memory to prevent swapping
```

### Rich vocabulary model

```yaml
services:
  app:
    image: app
    models:
      rich_vocab_model:
        endpoint_var: RICH_VOCAB_URL
        model_var: RICH_VOCAB_MODEL

models:
  rich_vocab_model:
    model: ai/model
    context_size: 4096
    runtime_flags:
      - "--temp"                # Temperature
      - "0.1"
      - "--top-p"               # Top-p sampling
      - "0.9"
```
