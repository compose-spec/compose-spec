# Untitled object in undefined Schema

```txt
undefined#/definitions/deployment/properties/restart_policy
```




| Abstract            | Extensible | Status         | Identifiable | Custom Properties | Additional Properties | Access Restrictions | Defined In                                                                  |
| :------------------ | ---------- | -------------- | ------------ | :---------------- | --------------------- | ------------------- | --------------------------------------------------------------------------- |
| Can be instantiated | No         | Unknown status | No           | Forbidden         | Forbidden             | none                | [config_schema_v3.9.json\*](config_schema_v3.9.json "open original schema") |

## restart_policy Type

`object` ([Details](config_schema_v3-definitions-deployment-properties-restart_policy.md))

# undefined Properties

| Property                      | Type      | Required | Nullable       | Defined by                                                                                                                                                                                            |
| :---------------------------- | --------- | -------- | -------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [condition](#condition)       | `string`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-restart_policy-properties-condition.md "undefined#/definitions/deployment/properties/restart_policy/properties/condition")       |
| [delay](#delay)               | `string`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-restart_policy-properties-delay.md "undefined#/definitions/deployment/properties/restart_policy/properties/delay")               |
| [max_attempts](#max_attempts) | `integer` | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-restart_policy-properties-max_attempts.md "undefined#/definitions/deployment/properties/restart_policy/properties/max_attempts") |
| [window](#window)             | `string`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-restart_policy-properties-window.md "undefined#/definitions/deployment/properties/restart_policy/properties/window")             |

## condition




`condition`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-restart_policy-properties-condition.md "undefined#/definitions/deployment/properties/restart_policy/properties/condition")

### condition Type

`string`

## delay




`delay`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-restart_policy-properties-delay.md "undefined#/definitions/deployment/properties/restart_policy/properties/delay")

### delay Type

`string`

### delay Constraints

**duration**: the string must be a duration string, according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339 "check the specification")

## max_attempts




`max_attempts`

-   is optional
-   Type: `integer`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-restart_policy-properties-max_attempts.md "undefined#/definitions/deployment/properties/restart_policy/properties/max_attempts")

### max_attempts Type

`integer`

## window




`window`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-restart_policy-properties-window.md "undefined#/definitions/deployment/properties/restart_policy/properties/window")

### window Type

`string`

### window Constraints

**duration**: the string must be a duration string, according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339 "check the specification")
