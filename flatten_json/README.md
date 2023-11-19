# Example Function: Flattening JSON

This function's goal is to take a nested JSON entry and flatten it to remove the nesting.

## Example Use Case Definition

A user has multiple services which are producing to a given station, and each of the services produce messages in the same nested JSON format. For down-stream usage from the consumer, it would help if the JSON was flattened.

The user could attach a Memphis Function to the station so that by making one change, and changing nothing in the producers, the down-stream consumers may consume the messages already flattened.

## Input

A JSON message.

## Output

A flattened representation of the JSON message. Array items are given as array_key: array_key_{index}. 