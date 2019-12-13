# Schema

```

```

The Compose file is a YAML file defining a multi-containers based application.

| Abstract            | Extensible | Status       | Identifiable | Custom Properties | Additional Properties | Defined In |
| ------------------- | ---------- | ------------ | ------------ | ----------------- | --------------------- | ---------- |
| Can be instantiated | Yes        | Experimental | No           | Forbidden         | Forbidden             |            |

# Definitions

| Property                                | Type       | Group                       |
| --------------------------------------- | ---------- | --------------------------- |
| [attachable](#attachable)               | `boolean`  | `#/definitions/network`     |
| [build](#build)                         | complex    | `#/definitions/service`     |
| [cap_add](#cap_add)                     | `string[]` | `#/definitions/service`     |
| [cap_drop](#cap_drop)                   | `string[]` | `#/definitions/service`     |
| [cgroup_parent](#cgroup_parent)         | `string`   | `#/definitions/service`     |
| [command](#command)                     | complex    | `#/definitions/service`     |
| [container_name](#container_name)       | `string`   | `#/definitions/service`     |
| [credential_spec](#credential_spec)     | `object`   | `#/definitions/service`     |
| [depends_on](#depends_on)               | reference  | `#/definitions/service`     |
| [deploy](#deploy)                       | reference  | `#/definitions/service`     |
| [devices](#devices)                     | `string[]` | `#/definitions/service`     |
| [disable](#disable)                     | `boolean`  | `#/definitions/healthcheck` |
| [dns](#dns)                             | reference  | `#/definitions/service`     |
| [dns_search](#dns_search)               | reference  | `#/definitions/service`     |
| [domainname](#domainname)               | `string`   | `#/definitions/service`     |
| [driver](#driver)                       | `string`   | `#/definitions/secret`      |
| [driver_opts](#driver_opts)             | `object`   | `#/definitions/secret`      |
| [endpoint_mode](#endpoint_mode)         | `string`   | `#/definitions/deployment`  |
| [entrypoint](#entrypoint)               | complex    | `#/definitions/service`     |
| [env_file](#env_file)                   | reference  | `#/definitions/service`     |
| [environment](#environment)             | reference  | `#/definitions/service`     |
| [expose](#expose)                       | multiple   | `#/definitions/service`     |
| [external](#external)                   | multiple   | `#/definitions/config`      |
| [external_links](#external_links)       | `string[]` | `#/definitions/service`     |
| [extra_hosts](#extra_hosts)             | reference  | `#/definitions/service`     |
| [file](#file)                           | `string`   | `#/definitions/config`      |
| [healthcheck](#healthcheck)             | reference  | `#/definitions/service`     |
| [hostname](#hostname)                   | `string`   | `#/definitions/service`     |
| [image](#image)                         | `string`   | `#/definitions/service`     |
| [init](#init)                           | `boolean`  | `#/definitions/service`     |
| [internal](#internal)                   | `boolean`  | `#/definitions/network`     |
| [interval](#interval)                   | `string`   | `#/definitions/healthcheck` |
| [ipam](#ipam)                           | `object`   | `#/definitions/network`     |
| [ipc](#ipc)                             | `string`   | `#/definitions/service`     |
| [isolation](#isolation)                 | `string`   | `#/definitions/service`     |
| [labels](#labels)                       | reference  | `#/definitions/config`      |
| [links](#links)                         | `string[]` | `#/definitions/service`     |
| [logging](#logging)                     | `object`   | `#/definitions/service`     |
| [mac_address](#mac_address)             | `string`   | `#/definitions/service`     |
| [mode](#mode)                           | `string`   | `#/definitions/deployment`  |
| [name](#name)                           | `string`   | `#/definitions/config`      |
| [network_mode](#network_mode)           | `string`   | `#/definitions/service`     |
| [pid](#pid)                             | `string`   | `#/definitions/service`     |
| [placement](#placement)                 | `object`   | `#/definitions/deployment`  |
| [ports](#ports)                         | `array`    | `#/definitions/service`     |
| [privileged](#privileged)               | `boolean`  | `#/definitions/service`     |
| [read_only](#read_only)                 | `boolean`  | `#/definitions/service`     |
| [replicas](#replicas)                   | `integer`  | `#/definitions/deployment`  |
| [resources](#resources)                 | `object`   | `#/definitions/deployment`  |
| [restart](#restart)                     | `string`   | `#/definitions/service`     |
| [restart_policy](#restart_policy)       | `object`   | `#/definitions/deployment`  |
| [retries](#retries)                     | `number`   | `#/definitions/healthcheck` |
| [rollback_config](#rollback_config)     | `object`   | `#/definitions/deployment`  |
| [security_opt](#security_opt)           | `string[]` | `#/definitions/service`     |
| [shm_size](#shm_size)                   | multiple   | `#/definitions/service`     |
| [start_period](#start_period)           | `string`   | `#/definitions/healthcheck` |
| [stdin_open](#stdin_open)               | `boolean`  | `#/definitions/service`     |
| [stop_grace_period](#stop_grace_period) | `string`   | `#/definitions/service`     |
| [stop_signal](#stop_signal)             | `string`   | `#/definitions/service`     |
| [sysctls](#sysctls)                     | reference  | `#/definitions/service`     |
| [template_driver](#template_driver)     | `string`   | `#/definitions/config`      |
| [test](#test)                           | complex    | `#/definitions/healthcheck` |
| [timeout](#timeout)                     | `string`   | `#/definitions/healthcheck` |
| [tmpfs](#tmpfs)                         | reference  | `#/definitions/service`     |
| [tty](#tty)                             | `boolean`  | `#/definitions/service`     |
| [ulimits](#ulimits)                     | `object`   | `#/definitions/service`     |
| [update_config](#update_config)         | `object`   | `#/definitions/deployment`  |
| [user](#user)                           | `string`   | `#/definitions/service`     |
| [userns_mode](#userns_mode)             | `string`   | `#/definitions/service`     |
| [working_dir](#working_dir)             | `string`   | `#/definitions/service`     |

## attachable

`attachable`

- is optional
- type: `boolean`
- defined in this schema

### attachable Type

`boolean`

## build

`build`

- is optional
- type: complex
- defined in this schema

### build Type

**One** of the following _conditions_ need to be fulfilled.

#### Condition 1

`string`

#### Condition 2

`object` with following properties:

| Property     | Type           | Required |
| ------------ | -------------- | -------- |
| `args`       |                | Optional |
| `cache_from` |                | Optional |
| `context`    | string         | Optional |
| `dockerfile` | string         | Optional |
| `labels`     |                | Optional |
| `network`    | string         | Optional |
| `shm_size`   | integer,string | Optional |
| `target`     | string         | Optional |

#### args

`args`

- is optional
- type: reference

##### args Type

- []() – `#/definitions/list_or_dict`

#### cache_from

`cache_from`

- is optional
- type: reference

##### cache_from Type

- []() – `#/definitions/list_of_strings`

#### context

`context`

- is optional
- type: `string`

##### context Type

`string`

#### dockerfile

`dockerfile`

- is optional
- type: `string`

##### dockerfile Type

`string`

#### labels

`labels`

- is optional
- type: reference

##### labels Type

- []() – `#/definitions/list_or_dict`

#### network

`network`

- is optional
- type: `string`

##### network Type

`string`

#### shm_size

`shm_size`

- is optional
- type: multiple

##### shm_size Type

Unknown type `integer,string`.

```json
{
  "type": ["integer", "string"],
  "simpletype": "multiple"
}
```

#### target

`target`

- is optional
- type: `string`

##### target Type

`string`

## cap_add

`cap_add`

- is optional
- type: `string[]`
- defined in this schema

### cap_add Type

Array type: `string[]`

All items must be of the type: `string`

## cap_drop

`cap_drop`

- is optional
- type: `string[]`
- defined in this schema

### cap_drop Type

Array type: `string[]`

All items must be of the type: `string`

## cgroup_parent

`cgroup_parent`

- is optional
- type: `string`
- defined in this schema

### cgroup_parent Type

`string`

## command

`command`

- is optional
- type: complex
- defined in this schema

### command Type

**One** of the following _conditions_ need to be fulfilled.

#### Condition 1

`string`

#### Condition 2

Array type:

All items must be of the type: `string`

## container_name

`container_name`

- is optional
- type: `string`
- defined in this schema

### container_name Type

`string`

## credential_spec

`credential_spec`

- is optional
- type: `object`
- defined in this schema

### credential_spec Type

`object` with following properties:

| Property   | Type   | Required |
| ---------- | ------ | -------- |
| `config`   | string | Optional |
| `file`     | string | Optional |
| `registry` | string | Optional |

#### config

`config`

- is optional
- type: `string`

##### config Type

`string`

#### file

`file`

- is optional
- type: `string`

##### file Type

`string`

#### registry

`registry`

- is optional
- type: `string`

##### registry Type

`string`

## depends_on

`depends_on`

- is optional
- type: reference
- defined in this schema

### depends_on Type

- []() – `#/definitions/list_of_strings`

## deploy

`deploy`

- is optional
- type: reference
- defined in this schema

### deploy Type

- []() – `#/definitions/deployment`

## devices

`devices`

- is optional
- type: `string[]`
- defined in this schema

### devices Type

Array type: `string[]`

All items must be of the type: `string`

## disable

`disable`

- is optional
- type: `boolean`
- defined in this schema

### disable Type

`boolean`

## dns

`dns`

- is optional
- type: reference
- defined in this schema

### dns Type

- []() – `#/definitions/string_or_list`

## dns_search

`dns_search`

- is optional
- type: reference
- defined in this schema

### dns_search Type

- []() – `#/definitions/string_or_list`

## domainname

`domainname`

- is optional
- type: `string`
- defined in this schema

### domainname Type

`string`

## driver

`driver`

- is optional
- type: `string`
- defined in this schema

### driver Type

`string`

## driver_opts

`driver_opts`

- is optional
- type: `object`
- defined in this schema

### driver_opts Type

`object` with following properties:

| Property | Type | Required |
| -------- | ---- | -------- |


## endpoint_mode

`endpoint_mode`

- is optional
- type: `string`
- defined in this schema

### endpoint_mode Type

`string`

## entrypoint

`entrypoint`

- is optional
- type: complex
- defined in this schema

### entrypoint Type

**One** of the following _conditions_ need to be fulfilled.

#### Condition 1

`string`

#### Condition 2

Array type:

All items must be of the type: `string`

## env_file

`env_file`

- is optional
- type: reference
- defined in this schema

### env_file Type

- []() – `#/definitions/string_or_list`

## environment

`environment`

- is optional
- type: reference
- defined in this schema

### environment Type

- []() – `#/definitions/list_or_dict`

## expose

`expose`

- is optional
- type: multiple
- defined in this schema

### expose Type

Array type: multiple

All items must be of the type: Unknown type `string,number`.

```json
{
  "type": "array",
  "items": {
    "type": ["string", "number"],
    "format": "expose",
    "simpletype": "multiple"
  },
  "uniqueItems": true,
  "definitiongroup": "service",
  "simpletype": "multiple"
}
```

## external

`external`

- is optional
- type: multiple
- defined in this schema

### external Type

Either one of:

- `boolean`
- `object`

## external_links

`external_links`

- is optional
- type: `string[]`
- defined in this schema

### external_links Type

Array type: `string[]`

All items must be of the type: `string`

## extra_hosts

`extra_hosts`

- is optional
- type: reference
- defined in this schema

### extra_hosts Type

- []() – `#/definitions/list_or_dict`

## file

`file`

- is optional
- type: `string`
- defined in this schema

### file Type

`string`

## healthcheck

`healthcheck`

- is optional
- type: reference
- defined in this schema

### healthcheck Type

- []() – `#/definitions/healthcheck`

## hostname

`hostname`

- is optional
- type: `string`
- defined in this schema

### hostname Type

`string`

## image

`image`

- is optional
- type: `string`
- defined in this schema

### image Type

`string`

## init

`init`

- is optional
- type: `boolean`
- defined in this schema

### init Type

`boolean`

## internal

`internal`

- is optional
- type: `boolean`
- defined in this schema

### internal Type

`boolean`

## interval

`interval`

- is optional
- type: `string`
- defined in this schema

### interval Type

`string`

## ipam

`ipam`

- is optional
- type: `object`
- defined in this schema

### ipam Type

`object` with following properties:

| Property | Type   | Required |
| -------- | ------ | -------- |
| `config` | array  | Optional |
| `driver` | string | Optional |

#### config

`config`

- is optional
- type: `object[]`

##### config Type

Array type: `object[]`

All items must be of the type: `object` with following properties:

| Property | Type   | Required |
| -------- | ------ | -------- |
| `subnet` | string | Optional |

#### subnet

`subnet`

- is optional
- type: `string`

##### subnet Type

`string`

#### driver

`driver`

- is optional
- type: `string`

##### driver Type

`string`

## ipc

`ipc`

- is optional
- type: `string`
- defined in this schema

### ipc Type

`string`

## isolation

`isolation`

- is optional
- type: `string`
- defined in this schema

### isolation Type

`string`

## labels

`labels`

- is optional
- type: reference
- defined in this schema

### labels Type

- []() – `#/definitions/list_or_dict`

## links

`links`

- is optional
- type: `string[]`
- defined in this schema

### links Type

Array type: `string[]`

All items must be of the type: `string`

## logging

`logging`

- is optional
- type: `object`
- defined in this schema

### logging Type

`object` with following properties:

| Property  | Type   | Required |
| --------- | ------ | -------- |
| `driver`  | string | Optional |
| `options` | object | Optional |

#### driver

`driver`

- is optional
- type: `string`

##### driver Type

`string`

#### options

`options`

- is optional
- type: `object`

##### options Type

`object` with following properties:

| Property | Type | Required |
| -------- | ---- | -------- |


## mac_address

`mac_address`

- is optional
- type: `string`
- defined in this schema

### mac_address Type

`string`

## mode

`mode`

- is optional
- type: `string`
- defined in this schema

### mode Type

`string`

## name

`name`

- is optional
- type: `string`
- defined in this schema

### name Type

`string`

## network_mode

`network_mode`

- is optional
- type: `string`
- defined in this schema

### network_mode Type

`string`

## pid

`pid`

- is optional
- type: `string`
- defined in this schema

### pid Type

`string`, nullable

## placement

`placement`

- is optional
- type: `object`
- defined in this schema

### placement Type

`object` with following properties:

| Property                | Type    | Required |
| ----------------------- | ------- | -------- |
| `constraints`           | array   | Optional |
| `max_replicas_per_node` | integer | Optional |
| `preferences`           | array   | Optional |

#### constraints

`constraints`

- is optional
- type: `string[]`

##### constraints Type

Array type: `string[]`

All items must be of the type: `string`

#### max_replicas_per_node

`max_replicas_per_node`

- is optional
- type: `integer`

##### max_replicas_per_node Type

`integer`

#### preferences

`preferences`

- is optional
- type: `object[]`

##### preferences Type

Array type: `object[]`

All items must be of the type: `object` with following properties:

| Property | Type   | Required |
| -------- | ------ | -------- |
| `spread` | string | Optional |

#### spread

`spread`

- is optional
- type: `string`

##### spread Type

`string`

## ports

`ports`

- is optional
- type: `array`
- defined in this schema

### ports Type

Array type: `array`

All items must be of the type:

**One** of the following _conditions_ need to be fulfilled.

#### Condition 1

`number`

#### Condition 2

`string`

#### Condition 3

`object` with following properties:

| Property    | Type    | Required |
| ----------- | ------- | -------- |
| `mode`      | string  | Optional |
| `protocol`  | string  | Optional |
| `published` | integer | Optional |
| `target`    | integer | Optional |

#### mode

`mode`

- is optional
- type: `string`

##### mode Type

`string`

#### protocol

`protocol`

- is optional
- type: `string`

##### protocol Type

`string`

#### published

`published`

- is optional
- type: `integer`

##### published Type

`integer`

#### target

`target`

- is optional
- type: `integer`

##### target Type

`integer`

## privileged

`privileged`

- is optional
- type: `boolean`
- defined in this schema

### privileged Type

`boolean`

## read_only

`read_only`

- is optional
- type: `boolean`
- defined in this schema

### read_only Type

`boolean`

## replicas

`replicas`

- is optional
- type: `integer`
- defined in this schema

### replicas Type

`integer`

## resources

`resources`

- is optional
- type: `object`
- defined in this schema

### resources Type

`object` with following properties:

| Property       | Type   | Required |
| -------------- | ------ | -------- |
| `limits`       | object | Optional |
| `reservations` | object | Optional |

#### limits

`limits`

- is optional
- type: `object`

##### limits Type

`object` with following properties:

| Property | Type   | Required |
| -------- | ------ | -------- |
| `cpus`   | string | Optional |
| `memory` | string | Optional |

#### cpus

`cpus`

- is optional
- type: `string`

##### cpus Type

`string`

#### memory

`memory`

- is optional
- type: `string`

##### memory Type

`string`

#### reservations

`reservations`

- is optional
- type: `object`

##### reservations Type

`object` with following properties:

| Property            | Type   | Required |
| ------------------- | ------ | -------- |
| `cpus`              | string | Optional |
| `generic_resources` |        | Optional |
| `memory`            | string | Optional |

#### cpus

`cpus`

- is optional
- type: `string`

##### cpus Type

`string`

#### generic_resources

`generic_resources`

- is optional
- type: reference

##### generic_resources Type

- []() – `#/definitions/generic_resources`

#### memory

`memory`

- is optional
- type: `string`

##### memory Type

`string`

## restart

`restart`

- is optional
- type: `string`
- defined in this schema

### restart Type

`string`

## restart_policy

`restart_policy`

- is optional
- type: `object`
- defined in this schema

### restart_policy Type

`object` with following properties:

| Property       | Type    | Required |
| -------------- | ------- | -------- |
| `condition`    | string  | Optional |
| `delay`        | string  | Optional |
| `max_attempts` | integer | Optional |
| `window`       | string  | Optional |

#### condition

`condition`

- is optional
- type: `string`

##### condition Type

`string`

#### delay

`delay`

- is optional
- type: `string`

##### delay Type

`string`

#### max_attempts

`max_attempts`

- is optional
- type: `integer`

##### max_attempts Type

`integer`

#### window

`window`

- is optional
- type: `string`

##### window Type

`string`

## retries

`retries`

- is optional
- type: `number`
- defined in this schema

### retries Type

`number`

## rollback_config

`rollback_config`

- is optional
- type: `object`
- defined in this schema

### rollback_config Type

`object` with following properties:

| Property            | Type    | Required |
| ------------------- | ------- | -------- |
| `delay`             | string  | Optional |
| `failure_action`    | string  | Optional |
| `max_failure_ratio` | number  | Optional |
| `monitor`           | string  | Optional |
| `order`             | string  | Optional |
| `parallelism`       | integer | Optional |

#### delay

`delay`

- is optional
- type: `string`

##### delay Type

`string`

#### failure_action

`failure_action`

- is optional
- type: `string`

##### failure_action Type

`string`

#### max_failure_ratio

`max_failure_ratio`

- is optional
- type: `number`

##### max_failure_ratio Type

`number`

#### monitor

`monitor`

- is optional
- type: `string`

##### monitor Type

`string`

#### order

`order`

- is optional
- type: `enum`

The value of this property **must** be equal to one of the [known values below](#rollback_config-known-values).

##### order Known Values

| Value         | Description |
| ------------- | ----------- |
| `start-first` |             |
| `stop-first`  |             |

#### parallelism

`parallelism`

- is optional
- type: `integer`

##### parallelism Type

`integer`

## security_opt

`security_opt`

- is optional
- type: `string[]`
- defined in this schema

### security_opt Type

Array type: `string[]`

All items must be of the type: `string`

## shm_size

`shm_size`

- is optional
- type: multiple
- defined in this schema

### shm_size Type

Either one of:

- `number`
- `string`

## start_period

`start_period`

- is optional
- type: `string`
- defined in this schema

### start_period Type

`string`

## stdin_open

`stdin_open`

- is optional
- type: `boolean`
- defined in this schema

### stdin_open Type

`boolean`

## stop_grace_period

`stop_grace_period`

- is optional
- type: `string`
- defined in this schema

### stop_grace_period Type

`string`

## stop_signal

`stop_signal`

- is optional
- type: `string`
- defined in this schema

### stop_signal Type

`string`

## sysctls

`sysctls`

- is optional
- type: reference
- defined in this schema

### sysctls Type

- []() – `#/definitions/list_or_dict`

## template_driver

`template_driver`

- is optional
- type: `string`
- defined in this schema

### template_driver Type

`string`

## test

`test`

- is optional
- type: complex
- defined in this schema

### test Type

**One** of the following _conditions_ need to be fulfilled.

#### Condition 1

`string`

#### Condition 2

Array type:

All items must be of the type: `string`

## timeout

`timeout`

- is optional
- type: `string`
- defined in this schema

### timeout Type

`string`

## tmpfs

`tmpfs`

- is optional
- type: reference
- defined in this schema

### tmpfs Type

- []() – `#/definitions/string_or_list`

## tty

`tty`

- is optional
- type: `boolean`
- defined in this schema

### tty Type

`boolean`

## ulimits

`ulimits`

- is optional
- type: `object`
- defined in this schema

### ulimits Type

`object` with following properties:

| Property | Type | Required |
| -------- | ---- | -------- |


## update_config

`update_config`

- is optional
- type: `object`
- defined in this schema

### update_config Type

`object` with following properties:

| Property            | Type    | Required |
| ------------------- | ------- | -------- |
| `delay`             | string  | Optional |
| `failure_action`    | string  | Optional |
| `max_failure_ratio` | number  | Optional |
| `monitor`           | string  | Optional |
| `order`             | string  | Optional |
| `parallelism`       | integer | Optional |

#### delay

`delay`

- is optional
- type: `string`

##### delay Type

`string`

#### failure_action

`failure_action`

- is optional
- type: `string`

##### failure_action Type

`string`

#### max_failure_ratio

`max_failure_ratio`

- is optional
- type: `number`

##### max_failure_ratio Type

`number`

#### monitor

`monitor`

- is optional
- type: `string`

##### monitor Type

`string`

#### order

`order`

- is optional
- type: `enum`

The value of this property **must** be equal to one of the [known values below](#update_config-known-values).

##### order Known Values

| Value         | Description |
| ------------- | ----------- |
| `start-first` |             |
| `stop-first`  |             |

#### parallelism

`parallelism`

- is optional
- type: `integer`

##### parallelism Type

`integer`

## user

`user`

- is optional
- type: `string`
- defined in this schema

### user Type

`string`

## userns_mode

`userns_mode`

- is optional
- type: `string`
- defined in this schema

### userns_mode Type

`string`

## working_dir

`working_dir`

- is optional
- type: `string`
- defined in this schema

### working_dir Type

`string`

# Properties

| Property              | Type     | Required     | Nullable | Defined by    |
| --------------------- | -------- | ------------ | -------- | ------------- |
| [configs](#configs)   | `object` | Optional     | No       | (this schema) |
| [networks](#networks) | `object` | Optional     | No       | (this schema) |
| [secrets](#secrets)   | `object` | Optional     | No       | (this schema) |
| [services](#services) | `object` | Optional     | No       | (this schema) |
| [version](#version)   | `string` | **Required** | No       | (this schema) |
| [volumes](#volumes)   | `object` | Optional     | No       | (this schema) |
| `^x-`                 | complex  | Pattern      | No       | (this schema) |

## configs

`configs`

- is optional
- type: `object`
- defined in this schema

### configs Type

`object` with following properties:

| Property | Type | Required |
| -------- | ---- | -------- |


## networks

`networks`

- is optional
- type: `object`
- defined in this schema

### networks Type

`object` with following properties:

| Property | Type | Required |
| -------- | ---- | -------- |


## secrets

`secrets`

- is optional
- type: `object`
- defined in this schema

### secrets Type

`object` with following properties:

| Property | Type | Required |
| -------- | ---- | -------- |


## services

Service definition contains configuration that is applied to each container started for that service.

`services`

- is optional
- type: `object`
- defined in this schema

### services Type

`object` with following properties:

| Property | Type | Required |
| -------- | ---- | -------- |


## version

Version of the Compose specification used. Tools not implementing required version MUST reject the configuration file.

`version`

- is **required**
- type: `string`
- defined in this schema

### version Type

`string`

## volumes

`volumes`

- is optional
- type: `object`
- defined in this schema

### volumes Type

`object` with following properties:

| Property | Type | Required |
| -------- | ---- | -------- |


## Pattern: `^x-`

Applies to all properties that match the regular expression `^x-`

`^x-`

- is a property pattern
- type: complex
- defined in this schema

### Pattern ^x- Type

Unknown type ``.

```json
{
  "simpletype": "complex"
}
```
