# Adding A Severity Field

This function's goal is to check a field, in this case, `time_since_last_produce` and add a json field `severity` that could be used to check how well a given service is behaving. 

## Example Use Case Definition

A user has a service which they cannot change, but has data in their messages that is a quantative measure which could possibly be used to check if a service is being overloaded or not. 

With Memphis Functions, the user could simply attach a function to the station (or to a separate station which a copy of every kth message might be sent to in order to save on Lambda compute time). This function could check that qualitive measure and change the message that is being produced to warn that the service is being overloaded. Or, even better, the Function could directly alert a monitoring tool of the health of the producing service. 

This could be used to monitor the health of the system and to make sure that latency or other measures is optimal.

## Input

This function requires the input to be in JSON format. Additionally, the following inputs must be supplied on function attachment: `field`, `cutoff`, `high`, and `low`.

The `field` will be field in the message that is checked to see if it is higher than the cutoff. The message must have a numerical value for this field. `high` and `low` are the labels for the severity that is given to the new field `severity` that will be added to label the message severity depending on if the `field` is higher or lower than the cutoff.

## Output

The given message will be modified to have an added `severity` field which denotes the how severe some condition is.

 
BG
