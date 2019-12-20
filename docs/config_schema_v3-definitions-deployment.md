# Untitled Schema Schema

```txt
undefined#/definitions/deployment
```




| Abstract            | Extensible | Status         | Identifiable | Custom Properties | Additional Properties | Access Restrictions | Defined In                                                                  |
| :------------------ | ---------- | -------------- | ------------ | :---------------- | --------------------- | ------------------- | --------------------------------------------------------------------------- |
| Can be instantiated | No         | Unknown status | No           | Forbidden         | Forbidden             | none                | [config_schema_v3.9.json\*](config_schema_v3.9.json "open original schema") |

## deployment Type

`object` ([Details](config_schema_v3-definitions-deployment.md))

# undefined Properties

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

## mode




`mode`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-mode.md "undefined#/definitions/deployment/properties/mode")

### mode Type

`string`

## endpoint_mode




`endpoint_mode`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-endpoint_mode.md "undefined#/definitions/deployment/properties/endpoint_mode")

### endpoint_mode Type

`string`

## replicas




`replicas`

-   is optional
-   Type: `integer`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-replicas.md "undefined#/definitions/deployment/properties/replicas")

### replicas Type

`integer`

## labels




`labels`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-labels.md "undefined#/definitions/deployment/properties/labels")

### labels Type

unknown

## rollback_config




`rollback_config`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-deployment-properties-rollback_config.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-rollback_config.md "undefined#/definitions/deployment/properties/rollback_config")

### rollback_config Type

`object` ([Details](config_schema_v3-definitions-deployment-properties-rollback_config.md))

## update_config




`update_config`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-deployment-properties-update_config.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config.md "undefined#/definitions/deployment/properties/update_config")

### update_config Type

`object` ([Details](config_schema_v3-definitions-deployment-properties-update_config.md))

## resources




`resources`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-deployment-properties-resources.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-resources.md "undefined#/definitions/deployment/properties/resources")

### resources Type

`object` ([Details](config_schema_v3-definitions-deployment-properties-resources.md))

## restart_policy




`restart_policy`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-deployment-properties-restart_policy.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-restart_policy.md "undefined#/definitions/deployment/properties/restart_policy")

### restart_policy Type

`object` ([Details](config_schema_v3-definitions-deployment-properties-restart_policy.md))

## placement




`placement`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-deployment-properties-placement.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-placement.md "undefined#/definitions/deployment/properties/placement")

### placement Type

`object` ([Details](config_schema_v3-definitions-deployment-properties-placement.md))
