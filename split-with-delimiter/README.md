# Split a string
## Description
Split a string key in a JSON object using custom delimiter
## Supported event formats
JSON 
## Inputs:
Input name | Description | Type
|---|---|---|
| key_to_split | The name of the key that should be splitted | string |
| delimiter | The delimiter to use | string |
## Example

### Test Event:

```json
{
    "active": true,
    "aggregate_usage": null,
    "amount": null,
    "amount_decimal": null,
    "api_version": "2022-11-15"
}
```

### Inputs
Input name | Value
|---|---|
| key_to_split | api_version
| delimiter | -

### Output
```json
{
    "active": true,
    "aggregate_usage": null,
    "amount": null,
    "amount_decimal": null,
    "api_version": ["2022","11","15"]
}
```