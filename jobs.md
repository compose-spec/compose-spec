# Jobs top-level element

A job is a container that runs to completion. Unlike [services](05-services.md) which are long-running
processes that can be scaled and deployed, jobs execute a task and then exit. Jobs are useful for
database migrations, batch processing, scheduled maintenance tasks, or any one-shot operation.

A Compose file may declare a `jobs` top-level element as a map whose keys are string representations of job names,
and whose values are job definitions. A job definition contains the configuration that is applied to each
job container.

A job can always be triggered manually using a `run` command, regardless of whether it also has
automated triggers configured.

A job definition supports all attributes defined in the [Container Specification](container_spec.md),
as well as the following job-specific attribute.

## triggers

`triggers` defines the conditions under which a job is executed automatically.

```yaml
jobs:
  db-migration:
    image: myapp:latest
    command: python manage.py migrate

  cleanup:
    image: busybox
    command: sh -c 'find /data -mtime +30 -delete'
    volumes:
      - data:/data
    triggers:
      schedule: "0 3 * * *"

  backup:
    image: backup-tool
    command: /backup.sh
    triggers:
      schedule: "0 0 * * 0"
```

### schedule

`schedule` defines a crontab expression that determines when the job runs automatically.
The format follows the standard crontab syntax:

```
 ┌───────────── minute (0–59)
 │ ┌───────────── hour (0–23)
 │ │ ┌───────────── day of the month (1–31)
 │ │ │ ┌───────────── month (1–12)
 │ │ │ │ ┌───────────── day of the week (0–6, Sunday to Saturday)
 │ │ │ │ │
 * * * * *
```

```yaml
jobs:
  hourly-report:
    image: reporter
    command: ./generate-report.sh
    triggers:
      schedule: "0 * * * *"
```
