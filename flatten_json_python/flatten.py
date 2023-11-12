from memphis.functions import create_function
import json

"""
Python requires data getting returned, so make sure to return create_function(...)
Make sure to pass both the context and event to the created function
"""
def handler(event, context):
    return create_function(event=event, user_func = flatten_wrapper)

def flatten_wrapper(message_payload, headers):
    payload = json.loads(message_payload)
    out_dict = {}
    flatten(out_dict, "", payload)

    # Convert the dict object to a JSON string, and then encode that as bytes
    # !Important! Bytes object must be encoded with utf-8!
    # Return the headers as well, modify them if needed.
    headers['modified'] = True # Adding a random header for the example...
    return bytes(json.dumps(out_dict), encoding='utf-8'), headers

"""
A basic example of flattening a JSON object recursively
"""
def flatten(out_dict: dict, parent_key, value):
    if isinstance(value, dict):
        for key, item in value.items():
            flatten(out_dict, key, item)
    elif isinstance(value, list):
        for index, item in enumerate(value):
            flatten(out_dict, f"{parent_key}_{index}", item)
    else:
        out_dict[parent_key] = value
        
