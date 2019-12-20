# Untitled object in undefined Schema

```txt
undefined
```

The Compose file is a YAML file defining a multi-containers based application.


| Abstract            | Extensible | Status         | Identifiable | Custom Properties | Additional Properties | Access Restrictions | Defined In                                                                |
| :------------------ | ---------- | -------------- | ------------ | :---------------- | --------------------- | ------------------- | ------------------------------------------------------------------------- |
| Can be instantiated | Yes        | Unknown status | No           | Forbidden         | Forbidden             | none                | [config_schema_v3.9.json](config_schema_v3.9.json "open original schema") |

## Untitled object in undefined Type

`object` ([Details](config_schema_v3.md))

# Untitled object in undefined Definitions

## Definitions group service

Reference this group by using

```json
{"$ref":"undefined#/definitions/service"}
```

| Property                                | Type          | Required | Nullable       | Defined by                                                                                                                                            |
| :-------------------------------------- | ------------- | -------- | -------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------- |
| [deploy](#deploy)                       | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-deploy.md "undefined#/definitions/service/properties/deploy")                       |
| [build](#build)                         | Merged        | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-build.md "undefined#/definitions/service/properties/build")                         |
| [cap_add](#cap_add)                     | `array`       | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-cap_add.md "undefined#/definitions/service/properties/cap_add")                     |
| [cap_drop](#cap_drop)                   | `array`       | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-cap_drop.md "undefined#/definitions/service/properties/cap_drop")                   |
| [cgroup_parent](#cgroup_parent)         | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-cgroup_parent.md "undefined#/definitions/service/properties/cgroup_parent")         |
| [command](#command)                     | Merged        | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-command.md "undefined#/definitions/service/properties/command")                     |
| [configs](#configs)                     | `array`       | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-configs.md "undefined#/definitions/service/properties/configs")                     |
| [container_name](#container_name)       | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-container_name.md "undefined#/definitions/service/properties/container_name")       |
| [credential_spec](#credential_spec)     | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-credential_spec.md "undefined#/definitions/service/properties/credential_spec")     |
| [depends_on](#depends_on)               | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-depends_on.md "undefined#/definitions/service/properties/depends_on")               |
| [devices](#devices)                     | `array`       | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-devices.md "undefined#/definitions/service/properties/devices")                     |
| [dns](#dns)                             | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-dns.md "undefined#/definitions/service/properties/dns")                             |
| [dns_search](#dns_search)               | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-dns_search.md "undefined#/definitions/service/properties/dns_search")               |
| [domainname](#domainname)               | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-domainname.md "undefined#/definitions/service/properties/domainname")               |
| [entrypoint](#entrypoint)               | Merged        | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-entrypoint.md "undefined#/definitions/service/properties/entrypoint")               |
| [env_file](#env_file)                   | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-env_file.md "undefined#/definitions/service/properties/env_file")                   |
| [environment](#environment)             | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-environment.md "undefined#/definitions/service/properties/environment")             |
| [expose](#expose)                       | `array`       | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-expose.md "undefined#/definitions/service/properties/expose")                       |
| [external_links](#external_links)       | `array`       | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-external_links.md "undefined#/definitions/service/properties/external_links")       |
| [extra_hosts](#extra_hosts)             | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-extra_hosts.md "undefined#/definitions/service/properties/extra_hosts")             |
| [healthcheck](#healthcheck)             | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-healthcheck.md "undefined#/definitions/service/properties/healthcheck")             |
| [hostname](#hostname)                   | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-hostname.md "undefined#/definitions/service/properties/hostname")                   |
| [image](#image)                         | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-image.md "undefined#/definitions/service/properties/image")                         |
| [init](#init)                           | `boolean`     | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-init.md "undefined#/definitions/service/properties/init")                           |
| [ipc](#ipc)                             | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-ipc.md "undefined#/definitions/service/properties/ipc")                             |
| [isolation](#isolation)                 | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-isolation.md "undefined#/definitions/service/properties/isolation")                 |
| [labels](#labels)                       | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-labels.md "undefined#/definitions/service/properties/labels")                       |
| [links](#links)                         | `array`       | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-links.md "undefined#/definitions/service/properties/links")                         |
| [logging](#logging)                     | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-logging.md "undefined#/definitions/service/properties/logging")                     |
| [mac_address](#mac_address)             | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-mac_address.md "undefined#/definitions/service/properties/mac_address")             |
| [network_mode](#network_mode)           | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-network_mode.md "undefined#/definitions/service/properties/network_mode")           |
| [networks](#networks)                   | Merged        | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-networks.md "undefined#/definitions/service/properties/networks")                   |
| [pid](#pid)                             | `string`      | Optional | can be null    | [Untitled schema](config_schema_v3-definitions-service-properties-pid.md "undefined#/definitions/service/properties/pid")                             |
| [ports](#ports)                         | `array`       | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-ports.md "undefined#/definitions/service/properties/ports")                         |
| [privileged](#privileged)               | `boolean`     | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-privileged.md "undefined#/definitions/service/properties/privileged")               |
| [read_only](#read_only)                 | `boolean`     | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-read_only.md "undefined#/definitions/service/properties/read_only")                 |
| [restart](#restart)                     | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-restart.md "undefined#/definitions/service/properties/restart")                     |
| [security_opt](#security_opt)           | `array`       | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-security_opt.md "undefined#/definitions/service/properties/security_opt")           |
| [shm_size](#shm_size)                   | Multiple      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-shm_size.md "undefined#/definitions/service/properties/shm_size")                   |
| [secrets](#secrets)                     | `array`       | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-secrets.md "undefined#/definitions/service/properties/secrets")                     |
| [sysctls](#sysctls)                     | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-sysctls.md "undefined#/definitions/service/properties/sysctls")                     |
| [stdin_open](#stdin_open)               | `boolean`     | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-stdin_open.md "undefined#/definitions/service/properties/stdin_open")               |
| [stop_grace_period](#stop_grace_period) | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-stop_grace_period.md "undefined#/definitions/service/properties/stop_grace_period") |
| [stop_signal](#stop_signal)             | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-stop_signal.md "undefined#/definitions/service/properties/stop_signal")             |
| [tmpfs](#tmpfs)                         | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-tmpfs.md "undefined#/definitions/service/properties/tmpfs")                         |
| [tty](#tty)                             | `boolean`     | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-tty.md "undefined#/definitions/service/properties/tty")                             |
| [ulimits](#ulimits)                     | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-ulimits.md "undefined#/definitions/service/properties/ulimits")                     |
| [user](#user)                           | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-user.md "undefined#/definitions/service/properties/user")                           |
| [userns_mode](#userns_mode)             | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-userns_mode.md "undefined#/definitions/service/properties/userns_mode")             |
| [volumes](#volumes)                     | `array`       | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-volumes.md "undefined#/definitions/service/properties/volumes")                     |
| [working_dir](#working_dir)             | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-properties-working_dir.md "undefined#/definitions/service/properties/working_dir")             |
| `^x-`                                   | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-service-patternproperties-x-.md "undefined#/definitions/service/patternProperties/^x-")                |

### deploy




`deploy`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-deploy.md "undefined#/definitions/service/properties/deploy")

#### deploy Type

unknown

### build




`build`

-   is optional
-   Type: merged type ([Details](config_schema_v3-definitions-service-properties-build.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-build.md "undefined#/definitions/service/properties/build")

#### build Type

merged type ([Details](config_schema_v3-definitions-service-properties-build.md))

one (and only one) of

-   [Untitled string in undefined](config_schema_v3-definitions-service-properties-build-oneof-0.md "check type definition")
-   [Untitled object in undefined](config_schema_v3-definitions-service-properties-build-oneof-1.md "check type definition")

### cap_add




`cap_add`

-   is optional
-   Type: `string[]`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-cap_add.md "undefined#/definitions/service/properties/cap_add")

#### cap_add Type

`string[]`

#### cap_add Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

### cap_drop




`cap_drop`

-   is optional
-   Type: `string[]`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-cap_drop.md "undefined#/definitions/service/properties/cap_drop")

#### cap_drop Type

`string[]`

#### cap_drop Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

### cgroup_parent




`cgroup_parent`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-cgroup_parent.md "undefined#/definitions/service/properties/cgroup_parent")

#### cgroup_parent Type

`string`

### command




`command`

-   is optional
-   Type: merged type ([Details](config_schema_v3-definitions-service-properties-command.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-command.md "undefined#/definitions/service/properties/command")

#### command Type

merged type ([Details](config_schema_v3-definitions-service-properties-command.md))

one (and only one) of

-   [Untitled string in undefined](config_schema_v3-definitions-service-properties-command-oneof-0.md "check type definition")
-   [Untitled array in undefined](config_schema_v3-definitions-service-properties-command-oneof-1.md "check type definition")

### configs




`configs`

-   is optional
-   Type: an array of merged types ([Details](config_schema_v3-definitions-service-properties-configs-items.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-configs.md "undefined#/definitions/service/properties/configs")

#### configs Type

an array of merged types ([Details](config_schema_v3-definitions-service-properties-configs-items.md))

### container_name




`container_name`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-container_name.md "undefined#/definitions/service/properties/container_name")

#### container_name Type

`string`

### credential_spec




`credential_spec`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-service-properties-credential_spec.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-credential_spec.md "undefined#/definitions/service/properties/credential_spec")

#### credential_spec Type

`object` ([Details](config_schema_v3-definitions-service-properties-credential_spec.md))

### depends_on




`depends_on`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-depends_on.md "undefined#/definitions/service/properties/depends_on")

#### depends_on Type

unknown

### devices




`devices`

-   is optional
-   Type: `string[]`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-devices.md "undefined#/definitions/service/properties/devices")

#### devices Type

`string[]`

#### devices Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

### dns




`dns`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-dns.md "undefined#/definitions/service/properties/dns")

#### dns Type

unknown

### dns_search




`dns_search`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-dns_search.md "undefined#/definitions/service/properties/dns_search")

#### dns_search Type

unknown

### domainname




`domainname`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-domainname.md "undefined#/definitions/service/properties/domainname")

#### domainname Type

`string`

### entrypoint




`entrypoint`

-   is optional
-   Type: merged type ([Details](config_schema_v3-definitions-service-properties-entrypoint.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-entrypoint.md "undefined#/definitions/service/properties/entrypoint")

#### entrypoint Type

merged type ([Details](config_schema_v3-definitions-service-properties-entrypoint.md))

one (and only one) of

-   [Untitled string in undefined](config_schema_v3-definitions-service-properties-entrypoint-oneof-0.md "check type definition")
-   [Untitled array in undefined](config_schema_v3-definitions-service-properties-entrypoint-oneof-1.md "check type definition")

### env_file




`env_file`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-env_file.md "undefined#/definitions/service/properties/env_file")

#### env_file Type

unknown

### environment




`environment`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-environment.md "undefined#/definitions/service/properties/environment")

#### environment Type

unknown

### expose




`expose`

-   is optional
-   Type: an array of the following:`string` or `number` ([Details](config_schema_v3-definitions-service-properties-expose-items.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-expose.md "undefined#/definitions/service/properties/expose")

#### expose Type

an array of the following:`string` or `number` ([Details](config_schema_v3-definitions-service-properties-expose-items.md))

#### expose Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

### external_links




`external_links`

-   is optional
-   Type: `string[]`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-external_links.md "undefined#/definitions/service/properties/external_links")

#### external_links Type

`string[]`

#### external_links Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

### extra_hosts




`extra_hosts`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-extra_hosts.md "undefined#/definitions/service/properties/extra_hosts")

#### extra_hosts Type

unknown

### healthcheck




`healthcheck`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-healthcheck.md "undefined#/definitions/service/properties/healthcheck")

#### healthcheck Type

unknown

### hostname




`hostname`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-hostname.md "undefined#/definitions/service/properties/hostname")

#### hostname Type

`string`

### image




`image`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-image.md "undefined#/definitions/service/properties/image")

#### image Type

`string`

### init




`init`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-init.md "undefined#/definitions/service/properties/init")

#### init Type

`boolean`

### ipc




`ipc`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-ipc.md "undefined#/definitions/service/properties/ipc")

#### ipc Type

`string`

### isolation




`isolation`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-isolation.md "undefined#/definitions/service/properties/isolation")

#### isolation Type

`string`

### labels




`labels`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-labels.md "undefined#/definitions/service/properties/labels")

#### labels Type

unknown

### links




`links`

-   is optional
-   Type: `string[]`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-links.md "undefined#/definitions/service/properties/links")

#### links Type

`string[]`

#### links Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

### logging




`logging`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-service-properties-logging.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-logging.md "undefined#/definitions/service/properties/logging")

#### logging Type

`object` ([Details](config_schema_v3-definitions-service-properties-logging.md))

### mac_address




`mac_address`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-mac_address.md "undefined#/definitions/service/properties/mac_address")

#### mac_address Type

`string`

### network_mode




`network_mode`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-network_mode.md "undefined#/definitions/service/properties/network_mode")

#### network_mode Type

`string`

### networks




`networks`

-   is optional
-   Type: merged type ([Details](config_schema_v3-definitions-service-properties-networks.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-networks.md "undefined#/definitions/service/properties/networks")

#### networks Type

merged type ([Details](config_schema_v3-definitions-service-properties-networks.md))

one (and only one) of

-   [Untitled Schema](config_schema_v3-definitions-service-properties-networks-oneof-0.md "check type definition")
-   [Untitled object in undefined](config_schema_v3-definitions-service-properties-networks-oneof-1.md "check type definition")

### pid




`pid`

-   is optional
-   Type: `string`
-   can be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-pid.md "undefined#/definitions/service/properties/pid")

#### pid Type

`string`

### ports




`ports`

-   is optional
-   Type: an array of merged types ([Details](config_schema_v3-definitions-service-properties-ports-items.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-ports.md "undefined#/definitions/service/properties/ports")

#### ports Type

an array of merged types ([Details](config_schema_v3-definitions-service-properties-ports-items.md))

#### ports Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

### privileged




`privileged`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-privileged.md "undefined#/definitions/service/properties/privileged")

#### privileged Type

`boolean`

### read_only




`read_only`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-read_only.md "undefined#/definitions/service/properties/read_only")

#### read_only Type

`boolean`

### restart




`restart`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-restart.md "undefined#/definitions/service/properties/restart")

#### restart Type

`string`

### security_opt




`security_opt`

-   is optional
-   Type: `string[]`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-security_opt.md "undefined#/definitions/service/properties/security_opt")

#### security_opt Type

`string[]`

#### security_opt Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

### shm_size




`shm_size`

-   is optional
-   Type: any of the folllowing: `number` or `string` ([Details](config_schema_v3-definitions-service-properties-shm_size.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-shm_size.md "undefined#/definitions/service/properties/shm_size")

#### shm_size Type

any of the folllowing: `number` or `string` ([Details](config_schema_v3-definitions-service-properties-shm_size.md))

### secrets




`secrets`

-   is optional
-   Type: an array of merged types ([Details](config_schema_v3-definitions-service-properties-secrets-items.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-secrets.md "undefined#/definitions/service/properties/secrets")

#### secrets Type

an array of merged types ([Details](config_schema_v3-definitions-service-properties-secrets-items.md))

### sysctls




`sysctls`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-sysctls.md "undefined#/definitions/service/properties/sysctls")

#### sysctls Type

unknown

### stdin_open




`stdin_open`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-stdin_open.md "undefined#/definitions/service/properties/stdin_open")

#### stdin_open Type

`boolean`

### stop_grace_period




`stop_grace_period`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-stop_grace_period.md "undefined#/definitions/service/properties/stop_grace_period")

#### stop_grace_period Type

`string`

#### stop_grace_period Constraints

**duration**: the string must be a duration string, according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339 "check the specification")

### stop_signal




`stop_signal`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-stop_signal.md "undefined#/definitions/service/properties/stop_signal")

#### stop_signal Type

`string`

### tmpfs




`tmpfs`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-tmpfs.md "undefined#/definitions/service/properties/tmpfs")

#### tmpfs Type

unknown

### tty




`tty`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-tty.md "undefined#/definitions/service/properties/tty")

#### tty Type

`boolean`

### ulimits




`ulimits`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-service-properties-ulimits.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-ulimits.md "undefined#/definitions/service/properties/ulimits")

#### ulimits Type

`object` ([Details](config_schema_v3-definitions-service-properties-ulimits.md))

### user




`user`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-user.md "undefined#/definitions/service/properties/user")

#### user Type

`string`

### userns_mode




`userns_mode`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-userns_mode.md "undefined#/definitions/service/properties/userns_mode")

#### userns_mode Type

`string`

### volumes




`volumes`

-   is optional
-   Type: an array of merged types ([Details](config_schema_v3-definitions-service-properties-volumes-items.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-volumes.md "undefined#/definitions/service/properties/volumes")

#### volumes Type

an array of merged types ([Details](config_schema_v3-definitions-service-properties-volumes-items.md))

### working_dir




`working_dir`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-working_dir.md "undefined#/definitions/service/properties/working_dir")

#### working_dir Type

`string`

### Pattern: `^x-`




`^x-`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-patternproperties-x-.md "undefined#/definitions/service/patternProperties/^x-")

#### ^x- Type

unknown

## Definitions group healthcheck

Reference this group by using

```json
{"$ref":"undefined#/definitions/healthcheck"}
```

| Property                      | Type      | Required | Nullable       | Defined by                                                                                                                                          |
| :---------------------------- | --------- | -------- | -------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------- |
| [disable](#disable)           | `boolean` | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-healthcheck-properties-disable.md "undefined#/definitions/healthcheck/properties/disable")           |
| [interval](#interval)         | `string`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-healthcheck-properties-interval.md "undefined#/definitions/healthcheck/properties/interval")         |
| [retries](#retries)           | `number`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-healthcheck-properties-retries.md "undefined#/definitions/healthcheck/properties/retries")           |
| [test](#test)                 | Merged    | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-healthcheck-properties-test.md "undefined#/definitions/healthcheck/properties/test")                 |
| [timeout](#timeout)           | `string`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-healthcheck-properties-timeout.md "undefined#/definitions/healthcheck/properties/timeout")           |
| [start_period](#start_period) | `string`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-healthcheck-properties-start_period.md "undefined#/definitions/healthcheck/properties/start_period") |

### disable




`disable`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-healthcheck-properties-disable.md "undefined#/definitions/healthcheck/properties/disable")

#### disable Type

`boolean`

### interval




`interval`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-healthcheck-properties-interval.md "undefined#/definitions/healthcheck/properties/interval")

#### interval Type

`string`

#### interval Constraints

**duration**: the string must be a duration string, according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339 "check the specification")

### retries




`retries`

-   is optional
-   Type: `number`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-healthcheck-properties-retries.md "undefined#/definitions/healthcheck/properties/retries")

#### retries Type

`number`

### test




`test`

-   is optional
-   Type: merged type ([Details](config_schema_v3-definitions-healthcheck-properties-test.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-healthcheck-properties-test.md "undefined#/definitions/healthcheck/properties/test")

#### test Type

merged type ([Details](config_schema_v3-definitions-healthcheck-properties-test.md))

one (and only one) of

-   [Untitled string in undefined](config_schema_v3-definitions-healthcheck-properties-test-oneof-0.md "check type definition")
-   [Untitled array in undefined](config_schema_v3-definitions-healthcheck-properties-test-oneof-1.md "check type definition")

### timeout




`timeout`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-healthcheck-properties-timeout.md "undefined#/definitions/healthcheck/properties/timeout")

#### timeout Type

`string`

#### timeout Constraints

**duration**: the string must be a duration string, according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339 "check the specification")

### start_period




`start_period`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-healthcheck-properties-start_period.md "undefined#/definitions/healthcheck/properties/start_period")

#### start_period Type

`string`

#### start_period Constraints

**duration**: the string must be a duration string, according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339 "check the specification")

## Definitions group deployment

Reference this group by using

```json
{"$ref":"undefined#/definitions/deployment"}
```

| Property                            | Type          | Required | Nullable       | Defined by                                                                                                                                              |
| :---------------------------------- | ------------- | -------- | -------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [mode](#mode)                       | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-mode.md "undefined#/definitions/deployment/properties/mode")                       |
| [endpoint_mode](#endpoint_mode)     | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-endpoint_mode.md "undefined#/definitions/deployment/properties/endpoint_mode")     |
| [replicas](#replicas)               | `integer`     | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-replicas.md "undefined#/definitions/deployment/properties/replicas")               |
| [labels](#labels)                   | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-labels.md "undefined#/definitions/deployment/properties/labels")                   |
| [rollback_config](#rollback_config) | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-rollback_config.md "undefined#/definitions/deployment/properties/rollback_config") |
| [update_config](#update_config)     | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config.md "undefined#/definitions/deployment/properties/update_config")     |
| [resources](#resources)             | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-resources.md "undefined#/definitions/deployment/properties/resources")             |
| [restart_policy](#restart_policy)   | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-restart_policy.md "undefined#/definitions/deployment/properties/restart_policy")   |
| [placement](#placement)             | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-placement.md "undefined#/definitions/deployment/properties/placement")             |

### mode




`mode`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-mode.md "undefined#/definitions/deployment/properties/mode")

#### mode Type

`string`

### endpoint_mode




`endpoint_mode`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-endpoint_mode.md "undefined#/definitions/deployment/properties/endpoint_mode")

#### endpoint_mode Type

`string`

### replicas




`replicas`

-   is optional
-   Type: `integer`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-replicas.md "undefined#/definitions/deployment/properties/replicas")

#### replicas Type

`integer`

### labels




`labels`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-labels.md "undefined#/definitions/deployment/properties/labels")

#### labels Type

unknown

### rollback_config




`rollback_config`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-deployment-properties-rollback_config.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-rollback_config.md "undefined#/definitions/deployment/properties/rollback_config")

#### rollback_config Type

`object` ([Details](config_schema_v3-definitions-deployment-properties-rollback_config.md))

### update_config




`update_config`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-deployment-properties-update_config.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config.md "undefined#/definitions/deployment/properties/update_config")

#### update_config Type

`object` ([Details](config_schema_v3-definitions-deployment-properties-update_config.md))

### resources




`resources`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-deployment-properties-resources.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-resources.md "undefined#/definitions/deployment/properties/resources")

#### resources Type

`object` ([Details](config_schema_v3-definitions-deployment-properties-resources.md))

### restart_policy




`restart_policy`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-deployment-properties-restart_policy.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-restart_policy.md "undefined#/definitions/deployment/properties/restart_policy")

#### restart_policy Type

`object` ([Details](config_schema_v3-definitions-deployment-properties-restart_policy.md))

### placement




`placement`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-deployment-properties-placement.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-placement.md "undefined#/definitions/deployment/properties/placement")

#### placement Type

`object` ([Details](config_schema_v3-definitions-deployment-properties-placement.md))

## Definitions group generic_resources

Reference this group by using

```json
{"$ref":"undefined#/definitions/generic_resources"}
```

| Property | Type | Required | Nullable | Defined by |
| :------- | ---- | -------- | -------- | :--------- |

## Definitions group network

Reference this group by using

```json
{"$ref":"undefined#/definitions/network"}
```

| Property                    | Type          | Required | Nullable       | Defined by                                                                                                                                |
| :-------------------------- | ------------- | -------- | -------------- | :---------------------------------------------------------------------------------------------------------------------------------------- |
| [name](#name)               | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-network-properties-name.md "undefined#/definitions/network/properties/name")               |
| [driver](#driver)           | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-network-properties-driver.md "undefined#/definitions/network/properties/driver")           |
| [driver_opts](#driver_opts) | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-network-properties-driver_opts.md "undefined#/definitions/network/properties/driver_opts") |
| [ipam](#ipam)               | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-network-properties-ipam.md "undefined#/definitions/network/properties/ipam")               |
| [external](#external)       | Multiple      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-network-properties-external.md "undefined#/definitions/network/properties/external")       |
| [internal](#internal)       | `boolean`     | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-network-properties-internal.md "undefined#/definitions/network/properties/internal")       |
| [attachable](#attachable)   | `boolean`     | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-network-properties-attachable.md "undefined#/definitions/network/properties/attachable")   |
| [labels](#labels)           | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-network-properties-labels.md "undefined#/definitions/network/properties/labels")           |
| `^x-`                       | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-network-patternproperties-x-.md "undefined#/definitions/network/patternProperties/^x-")    |

### name




`name`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-name.md "undefined#/definitions/network/properties/name")

#### name Type

`string`

### driver




`driver`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-driver.md "undefined#/definitions/network/properties/driver")

#### driver Type

`string`

### driver_opts




`driver_opts`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-network-properties-driver_opts.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-driver_opts.md "undefined#/definitions/network/properties/driver_opts")

#### driver_opts Type

`object` ([Details](config_schema_v3-definitions-network-properties-driver_opts.md))

### ipam




`ipam`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-network-properties-ipam.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-ipam.md "undefined#/definitions/network/properties/ipam")

#### ipam Type

`object` ([Details](config_schema_v3-definitions-network-properties-ipam.md))

### external




`external`

-   is optional
-   Type: any of the folllowing: `boolean` or `object` ([Details](config_schema_v3-definitions-network-properties-external.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-external.md "undefined#/definitions/network/properties/external")

#### external Type

any of the folllowing: `boolean` or `object` ([Details](config_schema_v3-definitions-network-properties-external.md))

### internal




`internal`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-internal.md "undefined#/definitions/network/properties/internal")

#### internal Type

`boolean`

### attachable




`attachable`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-attachable.md "undefined#/definitions/network/properties/attachable")

#### attachable Type

`boolean`

### labels




`labels`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-labels.md "undefined#/definitions/network/properties/labels")

#### labels Type

unknown

### Pattern: `^x-`




`^x-`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-patternproperties-x-.md "undefined#/definitions/network/patternProperties/^x-")

#### ^x- Type

unknown

## Definitions group volume

Reference this group by using

```json
{"$ref":"undefined#/definitions/volume"}
```

| Property                    | Type          | Required | Nullable       | Defined by                                                                                                                              |
| :-------------------------- | ------------- | -------- | -------------- | :-------------------------------------------------------------------------------------------------------------------------------------- |
| [name](#name)               | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-volume-properties-name.md "undefined#/definitions/volume/properties/name")               |
| [driver](#driver)           | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-volume-properties-driver.md "undefined#/definitions/volume/properties/driver")           |
| [driver_opts](#driver_opts) | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-volume-properties-driver_opts.md "undefined#/definitions/volume/properties/driver_opts") |
| [external](#external)       | Multiple      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-volume-properties-external.md "undefined#/definitions/volume/properties/external")       |
| [labels](#labels)           | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-volume-properties-labels.md "undefined#/definitions/volume/properties/labels")           |
| `^x-`                       | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-volume-patternproperties-x-.md "undefined#/definitions/volume/patternProperties/^x-")    |

### name




`name`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-volume-properties-name.md "undefined#/definitions/volume/properties/name")

#### name Type

`string`

### driver




`driver`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-volume-properties-driver.md "undefined#/definitions/volume/properties/driver")

#### driver Type

`string`

### driver_opts




`driver_opts`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-volume-properties-driver_opts.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-volume-properties-driver_opts.md "undefined#/definitions/volume/properties/driver_opts")

#### driver_opts Type

`object` ([Details](config_schema_v3-definitions-volume-properties-driver_opts.md))

### external




`external`

-   is optional
-   Type: any of the folllowing: `boolean` or `object` ([Details](config_schema_v3-definitions-volume-properties-external.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-volume-properties-external.md "undefined#/definitions/volume/properties/external")

#### external Type

any of the folllowing: `boolean` or `object` ([Details](config_schema_v3-definitions-volume-properties-external.md))

### labels




`labels`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-volume-properties-labels.md "undefined#/definitions/volume/properties/labels")

#### labels Type

unknown

### Pattern: `^x-`




`^x-`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-volume-patternproperties-x-.md "undefined#/definitions/volume/patternProperties/^x-")

#### ^x- Type

unknown

## Definitions group secret

Reference this group by using

```json
{"$ref":"undefined#/definitions/secret"}
```

| Property                            | Type          | Required | Nullable       | Defined by                                                                                                                                      |
| :---------------------------------- | ------------- | -------- | -------------- | :---------------------------------------------------------------------------------------------------------------------------------------------- |
| [name](#name)                       | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-secret-properties-name.md "undefined#/definitions/secret/properties/name")                       |
| [file](#file)                       | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-secret-properties-file.md "undefined#/definitions/secret/properties/file")                       |
| [external](#external)               | Multiple      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-secret-properties-external.md "undefined#/definitions/secret/properties/external")               |
| [labels](#labels)                   | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-secret-properties-labels.md "undefined#/definitions/secret/properties/labels")                   |
| [driver](#driver)                   | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-secret-properties-driver.md "undefined#/definitions/secret/properties/driver")                   |
| [driver_opts](#driver_opts)         | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-secret-properties-driver_opts.md "undefined#/definitions/secret/properties/driver_opts")         |
| [template_driver](#template_driver) | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-secret-properties-template_driver.md "undefined#/definitions/secret/properties/template_driver") |
| `^x-`                               | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-secret-patternproperties-x-.md "undefined#/definitions/secret/patternProperties/^x-")            |

### name




`name`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-secret-properties-name.md "undefined#/definitions/secret/properties/name")

#### name Type

`string`

### file




`file`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-secret-properties-file.md "undefined#/definitions/secret/properties/file")

#### file Type

`string`

### external




`external`

-   is optional
-   Type: any of the folllowing: `boolean` or `object` ([Details](config_schema_v3-definitions-secret-properties-external.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-secret-properties-external.md "undefined#/definitions/secret/properties/external")

#### external Type

any of the folllowing: `boolean` or `object` ([Details](config_schema_v3-definitions-secret-properties-external.md))

### labels




`labels`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-secret-properties-labels.md "undefined#/definitions/secret/properties/labels")

#### labels Type

unknown

### driver




`driver`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-secret-properties-driver.md "undefined#/definitions/secret/properties/driver")

#### driver Type

`string`

### driver_opts




`driver_opts`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-secret-properties-driver_opts.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-secret-properties-driver_opts.md "undefined#/definitions/secret/properties/driver_opts")

#### driver_opts Type

`object` ([Details](config_schema_v3-definitions-secret-properties-driver_opts.md))

### template_driver




`template_driver`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-secret-properties-template_driver.md "undefined#/definitions/secret/properties/template_driver")

#### template_driver Type

`string`

### Pattern: `^x-`




`^x-`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-secret-patternproperties-x-.md "undefined#/definitions/secret/patternProperties/^x-")

#### ^x- Type

unknown

## Definitions group config

Reference this group by using

```json
{"$ref":"undefined#/definitions/config"}
```

| Property                            | Type          | Required | Nullable       | Defined by                                                                                                                                      |
| :---------------------------------- | ------------- | -------- | -------------- | :---------------------------------------------------------------------------------------------------------------------------------------------- |
| [name](#name)                       | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-config-properties-name.md "undefined#/definitions/config/properties/name")                       |
| [file](#file)                       | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-config-properties-file.md "undefined#/definitions/config/properties/file")                       |
| [external](#external)               | Multiple      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-config-properties-external.md "undefined#/definitions/config/properties/external")               |
| [labels](#labels)                   | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-config-properties-labels.md "undefined#/definitions/config/properties/labels")                   |
| [template_driver](#template_driver) | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-config-properties-template_driver.md "undefined#/definitions/config/properties/template_driver") |
| `^x-`                               | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-config-patternproperties-x-.md "undefined#/definitions/config/patternProperties/^x-")            |

### name




`name`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-config-properties-name.md "undefined#/definitions/config/properties/name")

#### name Type

`string`

### file




`file`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-config-properties-file.md "undefined#/definitions/config/properties/file")

#### file Type

`string`

### external




`external`

-   is optional
-   Type: any of the folllowing: `boolean` or `object` ([Details](config_schema_v3-definitions-config-properties-external.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-config-properties-external.md "undefined#/definitions/config/properties/external")

#### external Type

any of the folllowing: `boolean` or `object` ([Details](config_schema_v3-definitions-config-properties-external.md))

### labels




`labels`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-config-properties-labels.md "undefined#/definitions/config/properties/labels")

#### labels Type

unknown

### template_driver




`template_driver`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-config-properties-template_driver.md "undefined#/definitions/config/properties/template_driver")

#### template_driver Type

`string`

### Pattern: `^x-`




`^x-`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-config-patternproperties-x-.md "undefined#/definitions/config/patternProperties/^x-")

#### ^x- Type

unknown

## Definitions group string_or_list

Reference this group by using

```json
{"$ref":"undefined#/definitions/string_or_list"}
```

| Property | Type | Required | Nullable | Defined by |
| :------- | ---- | -------- | -------- | :--------- |

## Definitions group list_of_strings

Reference this group by using

```json
{"$ref":"undefined#/definitions/list_of_strings"}
```

| Property | Type | Required | Nullable | Defined by |
| :------- | ---- | -------- | -------- | :--------- |

## Definitions group list_or_dict

Reference this group by using

```json
{"$ref":"undefined#/definitions/list_or_dict"}
```

| Property | Type | Required | Nullable | Defined by |
| :------- | ---- | -------- | -------- | :--------- |

## Definitions group constraints

Reference this group by using

```json
{"$ref":"undefined#/definitions/constraints"}
```

| Property | Type | Required | Nullable | Defined by |
| :------- | ---- | -------- | -------- | :--------- |

# undefined Properties

| Property              | Type          | Required | Nullable       | Defined by                                                                                     |
| :-------------------- | ------------- | -------- | -------------- | :--------------------------------------------------------------------------------------------- |
| [version](#version)   | `string`      | Required | cannot be null | [Untitled schema](config_schema_v3-properties-version.md "undefined#/properties/version")      |
| [services](#services) | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-properties-services.md "undefined#/properties/services")    |
| [networks](#networks) | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-properties-networks.md "undefined#/properties/networks")    |
| [volumes](#volumes)   | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-properties-volumes.md "undefined#/properties/volumes")      |
| [secrets](#secrets)   | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-properties-secrets.md "undefined#/properties/secrets")      |
| [configs](#configs)   | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-properties-configs.md "undefined#/properties/configs")      |
| `^x-`                 | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-patternproperties-x-.md "undefined#/patternProperties/^x-") |

## version

Version of the Compose specification used. Tools not implementing required version MUST reject the configuration file.


`version`

-   is required
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-properties-version.md "undefined#/properties/version")

### version Type

`string`

## services

Service definition contains configuration that is applied to each container started for that service.


`services`

-   is optional
-   Type: `object` ([Details](config_schema_v3-properties-services.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-properties-services.md "undefined#/properties/services")

### services Type

`object` ([Details](config_schema_v3-properties-services.md))

## networks




`networks`

-   is optional
-   Type: `object` ([Details](config_schema_v3-properties-networks.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-properties-networks.md "undefined#/properties/networks")

### networks Type

`object` ([Details](config_schema_v3-properties-networks.md))

## volumes




`volumes`

-   is optional
-   Type: `object` ([Details](config_schema_v3-properties-volumes.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-properties-volumes.md "undefined#/properties/volumes")

### volumes Type

`object` ([Details](config_schema_v3-properties-volumes.md))

## secrets




`secrets`

-   is optional
-   Type: `object` ([Details](config_schema_v3-properties-secrets.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-properties-secrets.md "undefined#/properties/secrets")

### secrets Type

`object` ([Details](config_schema_v3-properties-secrets.md))

## configs




`configs`

-   is optional
-   Type: `object` ([Details](config_schema_v3-properties-configs.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-properties-configs.md "undefined#/properties/configs")

### configs Type

`object` ([Details](config_schema_v3-properties-configs.md))

## Pattern: `^x-`




`^x-`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-patternproperties-x-.md "undefined#/patternProperties/^x-")

### ^x- Type

unknown
