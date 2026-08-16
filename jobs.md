# Jobs top-level element

A job is a container that runs to completion. Unlike [services](05-services.md) which are long-running
processes that can be scaled and deployed, jobs execute a task and then exit. Jobs are useful for
database migrations, batch processing, scheduled maintenance tasks, or any one-shot operation.

A Compose file may declare a `jobs` top-level element as a map whose keys are string representations of job names,
and whose values are job definitions. A job definition contains the configuration that is applied to each
job container.

A job can always be triggered manually using a `run` command, regardless of whether it also has
automated triggers configured.

A job definition supports all attributes defined in the [Container Specification](container_spec.md)
and the [Workload Specification](workload_spec.md), as well as the following job-specific attributes.

A job can depend on services or on other jobs using `depends_on`, but a service cannot
depend on a job: the "run this job before my service starts" pattern is not what jobs
are for.

## profiles

`profiles` defines a list of named profiles for the job to be active under, with the same
semantics as for [services](15-profiles.md). If unassigned, the job is always active: its
automated triggers are armed and it can be triggered manually. If assigned, the job is
ignored — triggers are not armed — unless one of its profiles is activated, or the job is
explicitly targeted by a `run` command. In that case its profile is added to the set of
active profiles.

```yaml
jobs:
  seed-fixtures:
    image: myapp:latest
    command: python manage.py seed
    profiles:
      - debug
    triggers:
      manual: true
```

## triggers

`triggers` defines the conditions under which a job is executed. A job must declare
exactly one trigger.

```yaml
jobs:
  db-migration:
    image: myapp:latest
    command: python manage.py migrate
    triggers:
      manual: true

  cleanup:
    image: busybox
    command: sh -c 'find /data -mtime +30 -delete'
    volumes:
      - data:/data
    triggers:
      schedule:
        - "0 3 * * *"

  backup:
    image: backup-tool
    command: /backup.sh
    triggers:
      schedule:
        - "0 0 * * 0"
```

### manual

`manual` declares that the job has no automated trigger: it only runs when explicitly
requested using a `run` command. As any job can be triggered manually regardless of its
automated triggers, `manual: true` makes this intent explicit for jobs that are only
run on demand.

```yaml
jobs:
  db-migration:
    image: myapp:latest
    command: python manage.py migrate
    triggers:
      manual: true
```

### schedule

`schedule` defines when the job runs automatically, as a list of schedules. As with
`volumes`, each entry uses either a short syntax — a plain crontab expression — or a
long syntax — a schedule object.

Crontab expressions follow the standard crontab syntax:

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
      schedule:
        - "0 * * * *"
```

A plain crontab entry is equivalent to a schedule object declaring only `cron`:
Compose implementations MUST canonicalize it to the object form.

A schedule object supports the following attributes:

- `cron` (required): the crontab expression.
- `timezone`: the timezone used to evaluate the cron expression (e.g. `Europe/Paris`).
  When not set, the platform's local timezone is used.
- `concurrency`: the policy applied when the schedule fires while a previous run is
  still in progress — `forbid` (the default) prevents the new run, `queue` runs it
  once the current one completes.
- `missed_fires`: the policy applied to fires missed while the platform was
  unavailable — `one` (the default) runs a single catch-up, `skip` ignores them.

```yaml
jobs:
  backup:
    image: backup-tool
    command: /backup.sh
    triggers:
      schedule:
        - cron: "0 3 * * *"
          timezone: "Europe/Paris"
          concurrency: forbid
          missed_fires: one
        # weekly full backup, plain crontab entry
        - "0 1 * * 0"
```
