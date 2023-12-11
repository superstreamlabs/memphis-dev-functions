const { memphis } = require("memphis-functions");

exports.handler = async (event) => {
    return await memphis.createFunction(event, eventHandler, asJson = true);
};

function eventHandler(payload, headers, inputs) {
    payload[inputs.field_to_ingest] = "Hello from Memphis!";

    return {
        processedMessage: payload,
        processedHeaders: headers
    };
}
