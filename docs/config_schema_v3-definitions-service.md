# Untitled object in undefined Schema

```txt
undefined#/definitions/service
```




| Abstract            | Extensible | Status         | Identifiable | Custom Properties | Additional Properties | Access Restrictions | Defined In                                                                  |
| :------------------ | ---------- | -------------- | ------------ | :---------------- | --------------------- | ------------------- | --------------------------------------------------------------------------- |
| Can be instantiated | No         | Unknown status | No           | Forbidden         | Forbidden             | none                | [config_schema_v3.9.json\*](config_schema_v3.9.json "open original schema") |

## service Type

`object` ([Details](config_schema_v3-definitions-service.md))

# undefined Properties

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

## deploy




`deploy`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-deploy.md "undefined#/definitions/service/properties/deploy")

### deploy Type

unknown

## build




`build`

-   is optional
-   Type: merged type ([Details](config_schema_v3-definitions-service-properties-build.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-build.md "undefined#/definitions/service/properties/build")

### build Type

merged type ([Details](config_schema_v3-definitions-service-properties-build.md))

one (and only one) of

-   [Untitled string in undefined](config_schema_v3-definitions-service-properties-build-oneof-0.md "check type definition")
-   [Untitled object in undefined](config_schema_v3-definitions-service-properties-build-oneof-1.md "check type definition")

## cap_add




`cap_add`

-   is optional
-   Type: `string[]`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-cap_add.md "undefined#/definitions/service/properties/cap_add")

### cap_add Type

`string[]`

### cap_add Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

## cap_drop




`cap_drop`

-   is optional
-   Type: `string[]`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-cap_drop.md "undefined#/definitions/service/properties/cap_drop")

### cap_drop Type

`string[]`

### cap_drop Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

## cgroup_parent




`cgroup_parent`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-cgroup_parent.md "undefined#/definitions/service/properties/cgroup_parent")

### cgroup_parent Type

`string`

## command




`command`

-   is optional
-   Type: merged type ([Details](config_schema_v3-definitions-service-properties-command.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-command.md "undefined#/definitions/service/properties/command")

### command Type

merged type ([Details](config_schema_v3-definitions-service-properties-command.md))

one (and only one) of

-   [Untitled string in undefined](config_schema_v3-definitions-service-properties-command-oneof-0.md "check type definition")
-   [Untitled array in undefined](config_schema_v3-definitions-service-properties-command-oneof-1.md "check type definition")

## configs




`configs`

-   is optional
-   Type: an array of merged types ([Details](config_schema_v3-definitions-service-properties-configs-items.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-configs.md "undefined#/definitions/service/properties/configs")

### configs Type

an array of merged types ([Details](config_schema_v3-definitions-service-properties-configs-items.md))

## container_name




`container_name`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-container_name.md "undefined#/definitions/service/properties/container_name")

### container_name Type

`string`

## credential_spec




`credential_spec`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-service-properties-credential_spec.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-credential_spec.md "undefined#/definitions/service/properties/credential_spec")

### credential_spec Type

`object` ([Details](config_schema_v3-definitions-service-properties-credential_spec.md))

## depends_on




`depends_on`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-depends_on.md "undefined#/definitions/service/properties/depends_on")

### depends_on Type

unknown

## devices




`devices`

-   is optional
-   Type: `string[]`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-devices.md "undefined#/definitions/service/properties/devices")

### devices Type

`string[]`

### devices Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

## dns




`dns`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-dns.md "undefined#/definitions/service/properties/dns")

### dns Type

unknown

## dns_search




`dns_search`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-dns_search.md "undefined#/definitions/service/properties/dns_search")

### dns_search Type

unknown

## domainname




`domainname`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-domainname.md "undefined#/definitions/service/properties/domainname")

### domainname Type

`string`

## entrypoint




`entrypoint`

-   is optional
-   Type: merged type ([Details](config_schema_v3-definitions-service-properties-entrypoint.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-entrypoint.md "undefined#/definitions/service/properties/entrypoint")

### entrypoint Type

merged type ([Details](config_schema_v3-definitions-service-properties-entrypoint.md))

one (and only one) of

-   [Untitled string in undefined](config_schema_v3-definitions-service-properties-entrypoint-oneof-0.md "check type definition")
-   [Untitled array in undefined](config_schema_v3-definitions-service-properties-entrypoint-oneof-1.md "check type definition")

## env_file




`env_file`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-env_file.md "undefined#/definitions/service/properties/env_file")

### env_file Type

unknown

## environment




`environment`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-environment.md "undefined#/definitions/service/properties/environment")

### environment Type

unknown

## expose




`expose`

-   is optional
-   Type: an array of the following:`string` or `number` ([Details](config_schema_v3-definitions-service-properties-expose-items.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-expose.md "undefined#/definitions/service/properties/expose")

### expose Type

an array of the following:`string` or `number` ([Details](config_schema_v3-definitions-service-properties-expose-items.md))

### expose Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

## external_links




`external_links`

-   is optional
-   Type: `string[]`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-external_links.md "undefined#/definitions/service/properties/external_links")

### external_links Type

`string[]`

### external_links Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

## extra_hosts




`extra_hosts`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-extra_hosts.md "undefined#/definitions/service/properties/extra_hosts")

### extra_hosts Type

unknown

## healthcheck




`healthcheck`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-healthcheck.md "undefined#/definitions/service/properties/healthcheck")

### healthcheck Type

unknown

## hostname




`hostname`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-hostname.md "undefined#/definitions/service/properties/hostname")

### hostname Type

`string`

## image




`image`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-image.md "undefined#/definitions/service/properties/image")

### image Type

`string`

## init




`init`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-init.md "undefined#/definitions/service/properties/init")

### init Type

`boolean`

## ipc




`ipc`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-ipc.md "undefined#/definitions/service/properties/ipc")

### ipc Type

`string`

## isolation




`isolation`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-isolation.md "undefined#/definitions/service/properties/isolation")

### isolation Type

`string`

## labels




`labels`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-labels.md "undefined#/definitions/service/properties/labels")

### labels Type

unknown

## links




`links`

-   is optional
-   Type: `string[]`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-links.md "undefined#/definitions/service/properties/links")

### links Type

`string[]`

### links Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

## logging




`logging`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-service-properties-logging.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-logging.md "undefined#/definitions/service/properties/logging")

### logging Type

`object` ([Details](config_schema_v3-definitions-service-properties-logging.md))

## mac_address




`mac_address`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-mac_address.md "undefined#/definitions/service/properties/mac_address")

### mac_address Type

`string`

## network_mode




`network_mode`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-network_mode.md "undefined#/definitions/service/properties/network_mode")

### network_mode Type

`string`

## networks




`networks`

-   is optional
-   Type: merged type ([Details](config_schema_v3-definitions-service-properties-networks.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-networks.md "undefined#/definitions/service/properties/networks")

### networks Type

merged type ([Details](config_schema_v3-definitions-service-properties-networks.md))

one (and only one) of

-   [Untitled Schema](config_schema_v3-definitions-service-properties-networks-oneof-0.md "check type definition")
-   [Untitled object in undefined](config_schema_v3-definitions-service-properties-networks-oneof-1.md "check type definition")

## pid




`pid`

-   is optional
-   Type: `string`
-   can be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-pid.md "undefined#/definitions/service/properties/pid")

### pid Type

`string`

## ports




`ports`

-   is optional
-   Type: an array of merged types ([Details](config_schema_v3-definitions-service-properties-ports-items.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-ports.md "undefined#/definitions/service/properties/ports")

### ports Type

an array of merged types ([Details](config_schema_v3-definitions-service-properties-ports-items.md))

### ports Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

## privileged




`privileged`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-privileged.md "undefined#/definitions/service/properties/privileged")

### privileged Type

`boolean`

## read_only




`read_only`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-read_only.md "undefined#/definitions/service/properties/read_only")

### read_only Type

`boolean`

## restart




`restart`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-restart.md "undefined#/definitions/service/properties/restart")

### restart Type

`string`

## security_opt




`security_opt`

-   is optional
-   Type: `string[]`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-security_opt.md "undefined#/definitions/service/properties/security_opt")

### security_opt Type

`string[]`

### security_opt Constraints

**unique items**: all items in this array must be unique. Duplicates are not allowed.

## shm_size




`shm_size`

-   is optional
-   Type: any of the folllowing: `number` or `string` ([Details](config_schema_v3-definitions-service-properties-shm_size.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-shm_size.md "undefined#/definitions/service/properties/shm_size")

### shm_size Type

any of the folllowing: `number` or `string` ([Details](config_schema_v3-definitions-service-properties-shm_size.md))

## secrets




`secrets`

-   is optional
-   Type: an array of merged types ([Details](config_schema_v3-definitions-service-properties-secrets-items.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-secrets.md "undefined#/definitions/service/properties/secrets")

### secrets Type

an array of merged types ([Details](config_schema_v3-definitions-service-properties-secrets-items.md))

## sysctls




`sysctls`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-sysctls.md "undefined#/definitions/service/properties/sysctls")

### sysctls Type

unknown

## stdin_open




`stdin_open`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-stdin_open.md "undefined#/definitions/service/properties/stdin_open")

### stdin_open Type

`boolean`

## stop_grace_period




`stop_grace_period`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-stop_grace_period.md "undefined#/definitions/service/properties/stop_grace_period")

### stop_grace_period Type

`string`

### stop_grace_period Constraints

**duration**: the string must be a duration string, according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339 "check the specification")

## stop_signal




`stop_signal`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-stop_signal.md "undefined#/definitions/service/properties/stop_signal")

### stop_signal Type

`string`

## tmpfs




`tmpfs`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-tmpfs.md "undefined#/definitions/service/properties/tmpfs")

### tmpfs Type

unknown

## tty




`tty`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-tty.md "undefined#/definitions/service/properties/tty")

### tty Type

`boolean`

## ulimits




`ulimits`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-service-properties-ulimits.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-ulimits.md "undefined#/definitions/service/properties/ulimits")

### ulimits Type

`object` ([Details](config_schema_v3-definitions-service-properties-ulimits.md))

## user




`user`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-user.md "undefined#/definitions/service/properties/user")

### user Type

`string`

## userns_mode




`userns_mode`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-userns_mode.md "undefined#/definitions/service/properties/userns_mode")

### userns_mode Type

`string`

## volumes




`volumes`

-   is optional
-   Type: an array of merged types ([Details](config_schema_v3-definitions-service-properties-volumes-items.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-volumes.md "undefined#/definitions/service/properties/volumes")

### volumes Type

an array of merged types ([Details](config_schema_v3-definitions-service-properties-volumes-items.md))

## working_dir




`working_dir`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-properties-working_dir.md "undefined#/definitions/service/properties/working_dir")

### working_dir Type

`string`

## Pattern: `^x-`




`^x-`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-service-patternproperties-x-.md "undefined#/definitions/service/patternProperties/^x-")

### ^x- Type

unknown
