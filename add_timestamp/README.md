# Add timestamp

This function's goal is to add a new field, `timestamp` that contains the time when the Function received the message. 

## Example Use Case Definition

A user might want to sort events by the time they arrived when visualizing the resulting data. To display this data in sorted order, a timestamp needs to be added but our producing service does not add one. A Memphis Function could be used to add this timestamp to the messages so that they would be able to be visualized the way the user wants.

## Input

A JSON message.

## Output

The given JSON message with the added `timestamp` entry added that contains a date time formatted timestamp.
