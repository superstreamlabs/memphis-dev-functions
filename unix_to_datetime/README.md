# Unix time to Date time

This function's goal is to take a message that logged a time in POSIX time and convert it to a date time format.

## Example Use Case Definition

Some downstream service that our messages will be loaded in requires time be in a date time format for it to work. This function could run before messages were ingested into that service so that the time was in the right format.

## Input

A JSON message and inputs which describe the field containing the POSIX timestamp and where the translated timestamp should go. 

These inputs will be given as `timestamp` and `out`. The value `timestamp` contains will be the field where the POSIX timestamp is and the `out` value will be the field the translated timestamp will be put into. 

## Output

A JSON message which has the given `out` key's value set to a date time representation of that time. Make `out` the same as `timestamp` to modify the object in place.
