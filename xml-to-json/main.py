import xmltodict
import json
from memphis import create_function

def handler(event, context):
    return create_function(event, event_handler)

def event_handler(msg_payload, msg_headers, inputs):
    payload = str(msg_payload, encoding='utf-8')
    as_dict = xmltodict.parse(payload, encoding='utf-8')

    return bytes(json.dumps(as_dict['root']), encoding='utf-8'), msg_headers