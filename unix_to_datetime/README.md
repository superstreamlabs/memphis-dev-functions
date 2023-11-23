# Unix time to date time
## Description
Unix time to date time takes a given POSIX time and converts it to a more human readable date time format. 
## Supported event formats
JSON 
## Inputs:
Input name | Description | Type
|---|---|---|
| timestamp | The field that contians the POSIX time | string |
| out | The field where the date time format will be stored | string |
## Test event 

### Inputs
Input name | Value
|---|---|
| timestamp | posix_time
| out | date_time

### Event:

```json
{
    "id": 1,
    "posix_time": "1700003853"
}
```

## Output to the test event

### Modified Event:
```json
{
    "id": 1,
    "posix_time": "1700003853",
    "date_time": "2023-11-14 17:17:33 -0600 CST"
}
```