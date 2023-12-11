import json
from memphis import create_function

def handler(event, context):
    return create_function(event, event_handler = event_handler, as_dict=True)

def event_handler(msg_payload, msg_headers, inputs):
    msg_payload[inputs["field_to_ingest"]] = "Hello from Memphis!"

    return msg_payload, msg_headers
