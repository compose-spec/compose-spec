## Networks top-level element

Networks are the layer that allow services to communicate with each other. The networking model exposed to a service
is limited to a simple IP connection with target services and external resources, while the Network definition allows
fine-tuning the actual implementation provided by the platform.

Networks can be created by specifying the network name under a top-level `networks` section.
Services can connect to networks by specifying the network name under the service [`networks`](05-services.md#networks) subsection

In the following example, at runtime, networks `front-tier` and `back-tier` will be created and the `frontend` service
connected to the `front-tier` network and the `back-tier` network.

```yml
services:
  frontend:
    image: awesome/webapp
    networks:
      - front-tier
      - back-tier

networks:
  front-tier:
  back-tier:
```

### driver

`driver` specifies which driver should be used for this network. Compose implementations MUST return an error if the
driver is not available on the platform.

```yml
driver: overlay
```

Default and available values are platform specific. Compose specification MUST support the following specific drivers:
`none` and `host`

- `host` use the host's networking stack
- `none` disable networking

#### host or none

The syntax for using built-in networks such as `host` and `none` is different, as such networks implicitly exists outside
the scope of the Compose implementation. To use them one MUST define an external network with the name `host` or `none` and
an alias that the Compose implementation can use (`hostnet` or `nonet` in the following examples), then grant the service
access to that network using its alias.

```yml
services:
  web:
    networks:
      hostnet: {}

networks:
  hostnet:
    external: true
    name: host
```

```yml
services:
  web:
    ...
    networks:
      nonet: {}

networks:
  nonet:
    external: true
    name: none
```

### driver_opts

`driver_opts` specifies a list of options as key-value pairs to pass to the driver for this network. These options are
driver-dependent - consult the driver's documentation for more information. Optional.

```yml
driver_opts:
  foo: "bar"
  baz: 1
```

### attachable

If `attachable` is set to `true`, then standalone containers SHOULD be able attach to this network, in addition to services.
If a standalone container attaches to the network, it can communicate with services and other standalone containers
that are also attached to the network.

```yml
networks:
  mynet1:
    driver: overlay
    attachable: true
```

### enable_ipv6

`enable_ipv6` enable IPv6 networking on this network.

### ipam

`ipam` specifies custom a IPAM configuration. This is an object with several properties, each of which is optional:

- `driver`: Custom IPAM driver, instead of the default.
- `config`: A list with zero or more configuration elements, each containing:
  - `subnet`: Subnet in CIDR format that represents a network segment
  - `ip_range`: Range of IPs from which to allocate container IPs
  - `gateway`: IPv4 or IPv6 gateway for the master subnet
  - `aux_addresses`: Auxiliary IPv4 or IPv6 addresses used by Network driver, as a mapping from hostname to IP
- `options`: Driver-specific options as a key-value mapping.

A full example:

```yml
ipam:
  driver: default
  config:
    - subnet: 172.28.0.0/16
      ip_range: 172.28.5.0/24
      gateway: 172.28.5.254
      aux_addresses:
        host1: 172.28.1.5
        host2: 172.28.1.6
        host3: 172.28.1.7
  options:
    foo: bar
    baz: "0"
```

### internal

By default, Compose implementations MUST provides external connectivity to networks. `internal` when set to `true` allow to
create an externally isolated network.

### labels

Add metadata to containers using Labels. Can use either an array or a dictionary.

Users SHOULD use reverse-DNS notation to prevent labels from conflicting with those used by other software.

```yml
labels:
  com.example.description: "Financial transaction network"
  com.example.department: "Finance"
  com.example.label-with-empty-value: ""
```

```yml
labels:
  - "com.example.description=Financial transaction network"
  - "com.example.department=Finance"
  - "com.example.label-with-empty-value"
```

Compose implementations MUST set `com.docker.compose.project` and `com.docker.compose.network` labels.

### external

If set to `true`, `external` specifies that this network’s lifecycle is maintained outside of that of the application.
Compose Implementations SHOULD NOT attempt to create these networks, and raises an error if one doesn't exist.

If `external` is set to `true` and network configuration has other but `name` attributes set, considering resource is
not managed by compose lifecycle, Compose Implementations SHOULD reject a Compose file as invalid.

In the example below, `proxy` is the gateway to the outside world. Instead of attempting to create a network, Compose
implementations SHOULD interrogate the platform for an existing network simply called `outside` and connect the
`proxy` service's containers to it.

```yml

services:
  proxy:
    image: awesome/proxy
    networks:
      - outside
      - default
  app:
    image: awesome/app
    networks:
      - default

networks:
  outside:
    external: true
```

### name

`name` sets a custom name for this network. The name field can be used to reference networks which contain special characters.
The name is used as is and will **not** be scoped with the project name.

```yml
networks:
  network1:
    name: my-app-net
```

It can also be used in conjunction with the `external` property to define the platform network that the Compose implementation
should retrieve, typically by using a parameter so the Compose file doesn't need to hard-code runtime specific values:

```yml
networks:
  network1:
    external: true
    name: "${NETWORK_ID}"
```
