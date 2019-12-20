# Untitled object in undefined Schema

```txt
undefined#/definitions/healthcheck
```




| Abstract            | Extensible | Status         | Identifiable | Custom Properties | Additional Properties | Access Restrictions | Defined In                                                                  |
| :------------------ | ---------- | -------------- | ------------ | :---------------- | --------------------- | ------------------- | --------------------------------------------------------------------------- |
| Can be instantiated | No         | Unknown status | No           | Forbidden         | Forbidden             | none                | [config_schema_v3.9.json\*](config_schema_v3.9.json "open original schema") |

## healthcheck Type

`object` ([Details](config_schema_v3-definitions-healthcheck.md))

# undefined Properties

| Property                      | Type      | Required | Nullable       | Defined by                                                                                                                                          |
| :---------------------------- | --------- | -------- | -------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------- |
| [disable](#disable)           | `boolean` | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-healthcheck-properties-disable.md "undefined#/definitions/healthcheck/properties/disable")           |
| [interval](#interval)         | `string`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-healthcheck-properties-interval.md "undefined#/definitions/healthcheck/properties/interval")         |
| [retries](#retries)           | `number`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-healthcheck-properties-retries.md "undefined#/definitions/healthcheck/properties/retries")           |
| [test](#test)                 | Merged    | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-healthcheck-properties-test.md "undefined#/definitions/healthcheck/properties/test")                 |
| [timeout](#timeout)           | `string`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-healthcheck-properties-timeout.md "undefined#/definitions/healthcheck/properties/timeout")           |
| [start_period](#start_period) | `string`  | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-healthcheck-properties-start_period.md "undefined#/definitions/healthcheck/properties/start_period") |

## disable




`disable`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-healthcheck-properties-disable.md "undefined#/definitions/healthcheck/properties/disable")

### disable Type

`boolean`

## interval




`interval`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-healthcheck-properties-interval.md "undefined#/definitions/healthcheck/properties/interval")

### interval Type

`string`

### interval Constraints

**duration**: the string must be a duration string, according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339 "check the specification")

## retries




`retries`

-   is optional
-   Type: `number`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-healthcheck-properties-retries.md "undefined#/definitions/healthcheck/properties/retries")

### retries Type

`number`

## test




`test`

-   is optional
-   Type: merged type ([Details](config_schema_v3-definitions-healthcheck-properties-test.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-healthcheck-properties-test.md "undefined#/definitions/healthcheck/properties/test")

### test Type

merged type ([Details](config_schema_v3-definitions-healthcheck-properties-test.md))

one (and only one) of

-   [Untitled string in undefined](config_schema_v3-definitions-healthcheck-properties-test-oneof-0.md "check type definition")
-   [Untitled array in undefined](config_schema_v3-definitions-healthcheck-properties-test-oneof-1.md "check type definition")

## timeout




`timeout`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-healthcheck-properties-timeout.md "undefined#/definitions/healthcheck/properties/timeout")

### timeout Type

`string`

### timeout Constraints

**duration**: the string must be a duration string, according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339 "check the specification")

## start_period




`start_period`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-healthcheck-properties-start_period.md "undefined#/definitions/healthcheck/properties/start_period")

### start_period Type

`string`

### start_period Constraints

**duration**: the string must be a duration string, according to [RFC 3339, section 5.6](https://tools.ietf.org/html/rfc3339 "check the specification")
