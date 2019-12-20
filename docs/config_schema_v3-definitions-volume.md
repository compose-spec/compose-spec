# Untitled Schema Schema

```txt
undefined#/definitions/volume
```




| Abstract            | Extensible | Status         | Identifiable | Custom Properties | Additional Properties | Access Restrictions | Defined In                                                                  |
| :------------------ | ---------- | -------------- | ------------ | :---------------- | --------------------- | ------------------- | --------------------------------------------------------------------------- |
| Can be instantiated | No         | Unknown status | No           | Forbidden         | Forbidden             | none                | [config_schema_v3.9.json\*](config_schema_v3.9.json "open original schema") |

## volume Type

`object` ([Details](config_schema_v3-definitions-volume.md))

# undefined Properties

| Property                    | Type          | Required | Nullable       | Defined by                                                                                                                              |
| :-------------------------- | ------------- | -------- | -------------- | :-------------------------------------------------------------------------------------------------------------------------------------- |
| [name](#name)               | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-volume-properties-name.md "undefined#/definitions/volume/properties/name")               |
| [driver](#driver)           | `string`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-volume-properties-driver.md "undefined#/definitions/volume/properties/driver")           |
| [driver_opts](#driver_opts) | `object`      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-volume-properties-driver_opts.md "undefined#/definitions/volume/properties/driver_opts") |
| [external](#external)       | Multiple      | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-volume-properties-external.md "undefined#/definitions/volume/properties/external")       |
| [labels](#labels)           | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-volume-properties-labels.md "undefined#/definitions/volume/properties/labels")           |
| `^x-`                       | Not specified | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-volume-patternproperties-x-.md "undefined#/definitions/volume/patternProperties/^x-")    |

## name




`name`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-volume-properties-name.md "undefined#/definitions/volume/properties/name")

### name Type

`string`

## driver




`driver`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-volume-properties-driver.md "undefined#/definitions/volume/properties/driver")

### driver Type

`string`

## driver_opts




`driver_opts`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-volume-properties-driver_opts.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-volume-properties-driver_opts.md "undefined#/definitions/volume/properties/driver_opts")

### driver_opts Type

`object` ([Details](config_schema_v3-definitions-volume-properties-driver_opts.md))

## external




`external`

-   is optional
-   Type: any of the folllowing: `boolean` or `object` ([Details](config_schema_v3-definitions-volume-properties-external.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-volume-properties-external.md "undefined#/definitions/volume/properties/external")

### external Type

any of the folllowing: `boolean` or `object` ([Details](config_schema_v3-definitions-volume-properties-external.md))

## labels




`labels`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-volume-properties-labels.md "undefined#/definitions/volume/properties/labels")

### labels Type

unknown

## Pattern: `^x-`




`^x-`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-volume-patternproperties-x-.md "undefined#/definitions/volume/patternProperties/^x-")

### ^x- Type

unknown
