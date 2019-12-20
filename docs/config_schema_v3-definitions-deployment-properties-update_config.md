# Untitled object in undefined Schema

```txt
undefined#/definitions/deployment/properties/update_config
```




| Abstract            | Extensible | Status         | Identifiable | Custom Properties | Additional Properties | Access Restrictions | Defined In                                                                  |
| :------------------ | ---------- | -------------- | ------------ | :---------------- | --------------------- | ------------------- | --------------------------------------------------------------------------- |
| Can be instantiated | No         | Unknown status | No           | Forbidden         | Forbidden             | none                | [config_schema_v3.9.json\*](config_schema_v3.9.json "open original schema") |

## update_config Type

`object` ([Details](config_schema_v3-definitions-deployment-properties-update_config.md))

# undefined Properties

| Property                                | Type      | Required | Nullable       | Defined by                                                                                                                                                                                                    |
| :-------------------------------------- | --------- | -------- | -------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [parallelism](#parallelism)             | `integer` | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config-properties-parallelism.md "undefined#/definitions/deployment/properties/update_config/properties/parallelism")             |
| [delay](#delay)                         | `string`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config-properties-delay.md "undefined#/definitions/deployment/properties/update_config/properties/delay")                         |
| [failure_action](#failure_action)       | `string`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config-properties-failure_action.md "undefined#/definitions/deployment/properties/update_config/properties/failure_action")       |
| [monitor](#monitor)                     | `string`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config-properties-monitor.md "undefined#/definitions/deployment/properties/update_config/properties/monitor")                     |
| [max_failure_ratio](#max_failure_ratio) | `number`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config-properties-max_failure_ratio.md "undefined#/definitions/deployment/properties/update_config/properties/max_failure_ratio") |
| [order](#order)                         | `string`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config-properties-order.md "undefined#/definitions/deployment/properties/update_config/properties/order")                         |

## parallelism




`parallelism`

-   is optional
-   Type: `integer`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config-properties-parallelism.md "undefined#/definitions/deployment/properties/update_config/properties/parallelism")

### parallelism Type

`integer`

## delay




`delay`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config-properties-delay.md "undefined#/definitions/deployment/properties/update_config/properties/delay")

### delay Type

`string`

### delay Constraints

**duration**: the string must be a duration string, according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339 "check the specification")

## failure_action




`failure_action`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config-properties-failure_action.md "undefined#/definitions/deployment/properties/update_config/properties/failure_action")

### failure_action Type

`string`

## monitor




`monitor`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config-properties-monitor.md "undefined#/definitions/deployment/properties/update_config/properties/monitor")

### monitor Type

`string`

### monitor Constraints

**duration**: the string must be a duration string, according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339 "check the specification")

## max_failure_ratio




`max_failure_ratio`

-   is optional
-   Type: `number`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config-properties-max_failure_ratio.md "undefined#/definitions/deployment/properties/update_config/properties/max_failure_ratio")

### max_failure_ratio Type

`number`

## order




`order`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-update_config-properties-order.md "undefined#/definitions/deployment/properties/update_config/properties/order")

### order Type

`string`

### order Constraints

**enum**: the value of this property must be equal to one of the following values:

| Value           | Explanation |
| :-------------- | ----------- |
| `"start-first"` |             |
| `"stop-first"`  |             |
