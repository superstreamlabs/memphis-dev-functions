# Add timestamp
## Description
Add timestamp adds the current time to an event in the `timestamp` field.

The format of the timestamp will be in `YYYY-MM-DD HH:MM:SS OFFSET TIMEZONE`.
## Supported event formats
JSON 
## Inputs:

N/A

## Test event 

### Inputs

N/A

### Event:

```json
{
    "id": 1,
}
```

## Output to the test event

### Modified Event:
```json
{
    "id": 1,
    "timestamp": "2023-12-22 15:30:50 -0600 CST"
}
```