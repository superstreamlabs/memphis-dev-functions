# Adding a Severity Field 
## Description
Add severity checks the event field specified by the input `field` against a given `cutoff` and sets the event field `severity` to values given by the inputs `high` if the event field was greater than or equal to the cutoff and `low` if it is lower than the cutoff.
## Supported event formats
JSON 
## Inputs:
Input name | Description | Type
|---|---|---|
| field | The field which will be checked against the cutoff | number |
| cutoff | The cutoff that the field is checked against | number |
| high | The label given to events whose field is greater than or equal to the cutoff | string |
| low | The label given to events whose field is lower than the cutoff | string |
## Test event 

### Inputs
Input name | Value
|---|---|
| field | measurement 
| cutoff | 10 
| high | severe 
| low | normal
### Event:

```json
{
    "id": 1,
    "measurement": 10
}
```

## Output to the test event

### Modified Event:
```json
{
    "id": 1,
    "measurement": 10,
    "severity": "severe"
}
```
