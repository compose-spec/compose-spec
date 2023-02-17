## Volumes top-level element

Volumes are persistent data stores implemented by the platform. The Compose specification offers a neutral abstraction
for services to mount volumes, and configuration parameters to allocate them on infrastructure.

The `volumes` section allows the configuration of named volumes that can be reused across multiple services. Here's
an example of a two-service setup where a database's data directory is shared with another service as a volume named
`db-data` so that it can be periodically backed up:

```yml
services:
  backend:
    image: awesome/database
    volumes:
      - db-data:/etc/data

  backup:
    image: backup-service
    volumes:
      - db-data:/var/lib/backup/data

volumes:
  db-data:
```

An entry under the top-level `volumes` key can be empty, in which case it uses the platform's default configuration for
creating a volume. Optionally, you can configure it with the following keys:

### driver

Specify which volume driver should be used for this volume. Default and available values are platform specific. If the driver is not available, the Compose implementation MUST return an error and stop application deployment.

```yml
driver: foobar
```

### driver_opts

`driver_opts` specifies a list of options as key-value pairs to pass to the driver for this volume. Those options are driver-dependent.

```yml
volumes:
  example:
    driver_opts:
      type: "nfs"
      o: "addr=10.40.0.199,nolock,soft,rw"
      device: ":/docker/example"
```

### external

If set to `true`, `external` specifies that this volume already exist on the platform and its lifecycle is managed outside
of that of the application. Compose implementations MUST NOT attempt to create these volumes, and MUST return an error if they
do not exist.

If `external` is set to `true` and volume configuration has other but `name` attributes set, considering resource is
not managed by compose lifecycle, Compose Implementations SHOULD reject a Compose file as invalid.


In the example below, instead of attempting to create a volume called
`{project_name}_db-data`, Compose looks for an existing volume simply
called `db-data` and mounts it into the `backend` service's containers.

```yml
services:
  backend:
    image: awesome/database
    volumes:
      - db-data:/etc/data

volumes:
  db-data:
    external: true
```

### labels

`labels` are used to add metadata to volumes. You can use either an array or a dictionary.

It's recommended that you use reverse-DNS notation to prevent your labels from
conflicting with those used by other software.

```yml
labels:
  com.example.description: "Database volume"
  com.example.department: "IT/Ops"
  com.example.label-with-empty-value: ""
```

```yml
labels:
  - "com.example.description=Database volume"
  - "com.example.department=IT/Ops"
  - "com.example.label-with-empty-value"
```

Compose implementation MUST set `com.docker.compose.project` and `com.docker.compose.volume` labels.

### name

`name` set a custom name for this volume. The name field can be used to reference volumes that contain special
characters. The name is used as is and will **not** be scoped with the stack name.

```yml
volumes:
  data:
    name: "my-app-data"
```

It can also be used in conjunction with the `external` property. Doing so the name of the volume used to lookup for
actual volume on platform is set separately from the name used to refer to it within the Compose file:

```yml
volumes:
  db-data:
    external:
      name: actual-name-of-volume
```

This make it possible to make this lookup name a parameter of a Compose file, so that the model ID for volume is
hard-coded but the actual volume ID on platform is set at runtime during deployment:

```yml
volumes:
  db-data:
    external:
      name: ${DATABASE_VOLUME}
```

