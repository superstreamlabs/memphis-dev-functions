# Sparcify Messages
## Description
Sparcify messages removes the given keys in the events based on the keys given in the `keys` input. 
## Supported event formats
JSON 
## Inputs:
Input name | Description | Type
|---|---|---|
| keys | A comma separated list of values which denotes what fields to remove from the event | string |
## Test event 

### Inputs
Input name | Value
|---|---|
| keys | key1,key2,key3,key4 

### Event:

```json
{
    "id": 1,
    "key1": "Some data",
    "key2": "Some more data",
    "key3": "Even more data",
    "key4": "Much more data",
    "key5": "This data is saved"
}
```

## Output to the test event

### Modified Event:
```json
{
    "id": 1,
    "key5": "This data is saved"
}
```