# Untitled Schema Schema

```txt
undefined#/definitions/network
```




| Abstract            | Extensible | Status         | Identifiable | Custom Properties | Additional Properties | Access Restrictions | Defined In                                                                  |
| :------------------ | ---------- | -------------- | ------------ | :---------------- | --------------------- | ------------------- | --------------------------------------------------------------------------- |
| Can be instantiated | No         | Unknown status | No           | Forbidden         | Forbidden             | none                | [config_schema_v3.9.json\*](config_schema_v3.9.json "open original schema") |

## network Type

`object` ([Details](config_schema_v3-definitions-network.md))

# undefined Properties

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

## name




`name`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-name.md "undefined#/definitions/network/properties/name")

### name Type

`string`

## driver




`driver`

-   is optional
-   Type: `string`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-driver.md "undefined#/definitions/network/properties/driver")

### driver Type

`string`

## driver_opts




`driver_opts`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-network-properties-driver_opts.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-driver_opts.md "undefined#/definitions/network/properties/driver_opts")

### driver_opts Type

`object` ([Details](config_schema_v3-definitions-network-properties-driver_opts.md))

## ipam




`ipam`

-   is optional
-   Type: `object` ([Details](config_schema_v3-definitions-network-properties-ipam.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-ipam.md "undefined#/definitions/network/properties/ipam")

### ipam Type

`object` ([Details](config_schema_v3-definitions-network-properties-ipam.md))

## external




`external`

-   is optional
-   Type: any of the folllowing: `boolean` or `object` ([Details](config_schema_v3-definitions-network-properties-external.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-external.md "undefined#/definitions/network/properties/external")

### external Type

any of the folllowing: `boolean` or `object` ([Details](config_schema_v3-definitions-network-properties-external.md))

## internal




`internal`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-internal.md "undefined#/definitions/network/properties/internal")

### internal Type

`boolean`

## attachable




`attachable`

-   is optional
-   Type: `boolean`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-attachable.md "undefined#/definitions/network/properties/attachable")

### attachable Type

`boolean`

## labels




`labels`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-properties-labels.md "undefined#/definitions/network/properties/labels")

### labels Type

unknown

## Pattern: `^x-`




`^x-`

-   is optional
-   Type: unknown
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-network-patternproperties-x-.md "undefined#/definitions/network/patternProperties/^x-")

### ^x- Type

unknown
