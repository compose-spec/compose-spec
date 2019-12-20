# Untitled object in undefined Schema

```txt
undefined#/definitions/deployment/properties/placement
```




| Abstract            | Extensible | Status         | Identifiable | Custom Properties | Additional Properties | Access Restrictions | Defined In                                                                  |
| :------------------ | ---------- | -------------- | ------------ | :---------------- | --------------------- | ------------------- | --------------------------------------------------------------------------- |
| Can be instantiated | No         | Unknown status | No           | Forbidden         | Forbidden             | none                | [config_schema_v3.9.json\*](config_schema_v3.9.json "open original schema") |

## placement Type

`object` ([Details](config_schema_v3-definitions-deployment-properties-placement.md))

# undefined Properties

| Property                                        | Type      | Required | Nullable       | Defined by                                                                                                                                                                                                    |
| :---------------------------------------------- | --------- | -------- | -------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [constraints](#constraints)                     | `array`   | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-placement-properties-constraints.md "undefined#/definitions/deployment/properties/placement/properties/constraints")                     |
| [preferences](#preferences)                     | `array`   | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-placement-properties-preferences.md "undefined#/definitions/deployment/properties/placement/properties/preferences")                     |
| [max_replicas_per_node](#max_replicas_per_node) | `integer` | Optional | cannot be null | [Untitled schema](config_schema_v3-definitions-deployment-properties-placement-properties-max_replicas_per_node.md "undefined#/definitions/deployment/properties/placement/properties/max_replicas_per_node") |

## constraints




`constraints`

-   is optional
-   Type: `string[]`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-placement-properties-constraints.md "undefined#/definitions/deployment/properties/placement/properties/constraints")

### constraints Type

`string[]`

## preferences




`preferences`

-   is optional
-   Type: `object[]` ([Details](config_schema_v3-definitions-deployment-properties-placement-properties-preferences-items.md))
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-placement-properties-preferences.md "undefined#/definitions/deployment/properties/placement/properties/preferences")

### preferences Type

`object[]` ([Details](config_schema_v3-definitions-deployment-properties-placement-properties-preferences-items.md))

## max_replicas_per_node




`max_replicas_per_node`

-   is optional
-   Type: `integer`
-   cannot be null
-   defined in: [Untitled schema](config_schema_v3-definitions-deployment-properties-placement-properties-max_replicas_per_node.md "undefined#/definitions/deployment/properties/placement/properties/max_replicas_per_node")

### max_replicas_per_node Type

`integer`
