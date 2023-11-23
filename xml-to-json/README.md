# XML to JSON
## Description
XML to JSON converts an XML message to a JSON message.

The JSON representation of the message will discard the root element.
## Supported event formats
XML
## Inputs:

N/A

## Test event 

### Inputs

N/A

### Event:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<root>
    <id>1</id>
    <data>My data</data>
</root>
```

## Output to the test event

### Modified Event:
```json
{
    "id": "1", 
    "data": "My data"
}
```