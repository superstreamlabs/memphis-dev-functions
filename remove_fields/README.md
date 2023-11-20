# Sparcify-Messages

This function's goal is to remove unnnecessary fields from a JSON entry.

## Example Use Case Definition

A service could be streaming data to Memphis and also logging that data somewhere else. The user could have a lighter, down-stream service that doesn't require all the data and seems to be struggling with meeting throughput requiremenets. 

A Memphis Function could be used to lighten the messages before they reach the consumer. 

## Input

A JSON message and an input `keys` which is a comma separated list of keys to remove from the JSON object.

Here is an example value for the `keys` inputs: `key1, key2, key3`.

## Output

The given JSON message with the field specified in the inputs removed.
