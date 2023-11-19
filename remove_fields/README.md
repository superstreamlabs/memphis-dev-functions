# Example Function: Sparcify-Messages

This function's goal is to remove unnnecessary fields from a JSON entry.

## Example Use Case Definition

A service could be streaming data to Memphis and also logging that data somewhere else. The user could have a lighter, down-stream service that doesn't require all the data and seems to be struggling with meeting throughput requiremenets. 

A Memphis Function could be used to lighten the messages before they reach the consumer. 

## Input

A JSON message and inputs which say which keys to remove.

Here is an example inputs:
```json
{
    "remove1": "key_to_remove",
    "remove2": "key_to_remove",
    "remove3": "key_to_remove",
}
```

The keys are ignored.

## Output

The given JSON message with the field specified in the inputs removed.