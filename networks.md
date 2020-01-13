# The Compose Specification - networks support
version: 3.9

### Definition

Networks are the layer that allow services to communicate with each other.
Networks can be created by specifying the network name under a top level `networks` section.
Services can connect to networks by specifying the network name under the service subsection

In the following example during runtime, networks `front-tier` and `back-tier` will be created and the `frontend` service 
connect to the`front-tier` network and the `back-tier` network

```yaml
version: "3"
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

#### aliases

Aliases (alternative hostnames) for this service on the network. Other containers on the same network can use either the service name or this alias to connect to one of the service's containers.

Since `aliases` is network-scoped, the same service can have different aliases on different networks.

> **Note**: A network-wide alias can be shared by multiple containers, and even by multiple services. If it is, then exactly which container the name resolves to is not guaranteed.

The general format is shown here.

```yaml
services:
  some-service:
    networks:
      some-network:
        aliases:
         - alias1
         - alias3
      other-network:
        aliases:
         - alias2
```

In the example below, three services are provided (`web`, `worker`, and `db`),
along with two networks (`new` and `legacy`). The `db` service is reachable at
the hostname `db` or `database` on the `new` network, and at `db` or `mysql` on
the `legacy` network.

```yaml
version: "3"

services:
  web:
    image: "nginx:alpine"
    networks:
      - new

  worker:
    image: "my-worker-image:latest"
    networks:
      - legacy

  db:
    image: mysql
    networks:
      new:
        aliases:
          - database
      legacy:
        aliases:
          - mysql

networks:
  new:
  legacy:
```

#### ipv4_address, ipv6_address

Specify a static IP address for containers for this service when joining the network.

The corresponding network configuration in the top-level networks section must have an
`ipam` block with subnet configurations covering each static address.

An example:

```yaml
version: "3"

services:
  app:
    image: nginx:alpine
    networks:
      app_net:
        ipv4_address: 172.16.238.10
        ipv6_address: 2001:3984:3989::10

networks:
  app_net:
    ipam:
      driver: default
      config:
        - subnet: "172.16.238.0/24"
        - subnet: "2001:3984:3989::/64"
```

## Implementation

* [Networking in Docker Compose](https://docs.docker.com/compose/networking/)
