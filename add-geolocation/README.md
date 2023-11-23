# Geolocation Enrichment
## Description
Takes a geolocation from the `geolocation` input and places it into the event with the key as the given `out` input.
## Supported event formats
JSON 
## Inputs:
Input name | Description | Type
|---|---|---|
| geolocation | The geolocation of the event | string |
| out | The event fieldname which gets set as the geolocation | string |
## Test event 

### Inputs
Input name | Value
|---|---|
| geolocation | 127.0.0.1 
| out | geolocation 

### Event:

```json
{
    "id": 1
}
```

## Output to the test event

### Modified Event:
```json
{
    "id": 1,
    "geolocation": "127.0.0.1"
}
```