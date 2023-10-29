import memphis
import json

# Lambda handler required by Lambda. 
# We return in the hander unlike in Go. 
# Make sure to pass both the context and event to the created
# function
def lambda_handler(event, context):
    my_funct = create_function(user_func = flatten_wrapper)
    
    return my_funct(event, context)

def flatten_wrapper(message_payload):
    payload = json.loads(message_payload)
    out_dict = {}
    flatten(out_dict, "", payload)

    return json.dumps(out_dict)

def flatten(out_dict: dict, parent_key, value):
    if isinstance(value, dict):
        for key, item in value.items():
            flatten(out_dict, key, item)
    elif isinstance(value, list):
        for index, item in enumerate(value):
            flatten(out_dict, f"{parent_key}_{index}", item)
    else:
        out_dict[parent_key] = value
        

# Identical to the memphis call, exported here for the 
# example until the repo gets updated
def create_function(
        user_func: callable, 
    ) -> None:
        def lambda_handler(event, context):
            import json

            processed_events = {}
            processed_events["successfullMessages"] = []
            processed_events["errorMessages"] = []
            for message in event.messages:
                try:
                    processed_message = user_func(message["payload"])
                    
                    processed_events["successfullMessages"].append({
                        "headers": message["headers"],
                        "payload": processed_message
                    })

                except Exception as e:
                    processed_events["errorMessages"].append({
                        "headers": message["headers"],
                        "payload": message["payload"],
                        "error": str(e)  
                    })
            
            try:
                return json.dumps(processed_events).encode('utf-8')
            except Exception as e:
                return f"Returned message types from user function are not able to be converted into JSON: {e}"

        return lambda_handler