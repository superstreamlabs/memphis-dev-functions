# Extract Email
## Description
Extract Email searches the event field given by the input `email` for emails and sets the event field given by the input `out` as a list of the emails found. 
## Supported event formats
JSON 
## Inputs:
Input name | Description | Type
|---|---|---|
| email | The field that contains emails to be extracted | string |
| out | The event field that is used to store the emails extracted | string |
## Test event 

### Inputs
Input name | Value
|---|---|
| email | message 
| out | emails 

### Event:

```json
{
    "id": 1,
    "message": "This is an example message. Please use john.doe@email.com or john.doe2@email.com to contact me back"
}
```

## Output to the test event

### Modified Event:
```json
{
    "id": 1,
    "message": "This is an example message. Please use john.doe@email.com or john.doe2@email.com to contact me back",
    "emails": ["john.doe@email.com", "john.doe2@email.com"]
}
```