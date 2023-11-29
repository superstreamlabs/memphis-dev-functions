# Geolocation Enrichment
## Description
Takes an IP from the field the `geolocation` input points to and using `https://ip-api.com`, gets geolocation information and places that information into the event with the key as the given `out` input.
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
| geolocation | ip
| out | geolocation 

### Event:

```json
{
    "id": 1, 
    "ip": "156.33.241.5"
}
```

## Output to the test event

### Modified Event:
```json
{
    "id": 1,
    "ip": "156.33.241.5",
    "geolocation": {
        "status": "success",
        "country": "United States",
        "countryCode": "US",
        "region": "DC",
        "regionName": "District of Columbia",
        "city": "Washington",
        "zip": "20003",
        "lat": 38.8842,
        "lon": -76.9941,
        "timezone": "America/New_York",
        "isp": "US Senate",
        "org": "United States Senate",
        "as": "AS3495 US Senate",
        "query": "156.33.241.5"
    }
}
```