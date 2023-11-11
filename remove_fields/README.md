# Example Function: Sparcify-Messages

This function's goal is to remove unnnecessary fields from a JSON entry.

## Example Use Case Definition

A user has multiple services which are producing to a given station, and each of the services produce messages in the same nested JSON format. For down-stream usage from the consumer, it would help if the JSON was flattened.

The user could attach a Memphis Function to the station so that by making one change, and changing nothing in the producers, the down-stream consumers may consume the messages already flattened.