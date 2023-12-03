[![Github (6)](https://github.com/memphisdev/memphis/assets/107035359/bc2feafc-946c-4569-ab8d-836bc0181890)](https://www.functions.memphis.dev/)
<p align="center">
<a href="https://memphis.dev/discord"><img src="https://img.shields.io/discord/963333392844328961?color=6557ff&label=discord" alt="Discord"></a>
<a href="https://github.com/memphisdev/memphis/issues?q=is%3Aissue+is%3Aclosed"><img src="https://img.shields.io/github/issues-closed/memphisdev/memphis?color=6557ff"></a> 
  <img src="https://img.shields.io/npm/dw/memphis-dev?color=ffc633&label=installations">
<a href="https://github.com/memphisdev/memphis/blob/master/CODE_OF_CONDUCT.md"><img src="https://img.shields.io/badge/Code%20of%20Conduct-v1.0-ff69b4.svg?color=ffc633" alt="Code Of Conduct"></a> 
<img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/memphisdev/memphis?color=61dfc6">
<img src="https://img.shields.io/github/last-commit/memphisdev/memphis?color=61dfc6&label=last%20commit">
</p>

<div align="center">
  
  <img width="200" alt="CNCF Silver Member" src="https://www.cncf.io/wp-content/uploads/2022/07/cncf-white-logo.svg#gh-dark-mode-only">
  
</div>
 
 <b><p align="center">
  <a href="https://memphis.dev/pricing/">Cloud</a> - <a href="https://memphis.dev/docs/">Docs</a> - <a href="https://twitter.com/Memphis_Dev">X</a> - <a href="https://www.youtube.com/channel/UCVdMDLCSxXOqtgrBaRUHKKg">YouTube</a>
</p></b>

<div align="center">

  <h4>

**[Memphis.dev](https://memphis.dev)** is more than a broker. It's a new streaming stack.<br>
Memphis.dev is a highly scalable event streaming and processing engine.<br>

  </h4>
  
</div>

## ![20](https://user-images.githubusercontent.com/70286779/220196529-abb958d2-5c58-4c33-b5e0-40f5446515ad.png) About

Before Memphis came along, handling ingestion and processing of events on a large scale took months to adopt and was a capability reserved for the top 20% of mega-companies. Now, Memphis opens the door for the other 80% to unleash their event and data streaming superpowers quickly, easily, and with great cost-effectiveness.

This repository is responsible for the Memphis Functions Javascript SDK

## Installation

```sh
$ npm install memphis-functions
```

## Importing

For Javascript, you can choose to use the import or required keyword. This library exports a singleton instance of `memphis` with which you can consume and produce messages.

```js
const { memphis } = require('memphis-functions');
```

### Creating a Memphis function
Memphis provides a `createFunction` utility for more easily creating Memphis Functions. 

The user-created `eventHandler` will be called for every message in the given batch of events. The user's `eventHandler` will take in a `payload` as a Uint8Array, `headers` as an object, and `inputs` as an object, and should return an object with keys `{ processedMessage, processedHeaders }`. The returned `processedMessage` should be a Uint8Array, and `processedHeaders` should be an object.

The user function should throw an exception if the message processing has failed. If any exception is thrown (deliberately or by a failed operation), the message will be sent to the dead letter station.

If the returned `processedMessage` and `processedHeaders` are `null`, then the message will be skipped and will not be sent to the station or dead letter station.

> Make sure to encode the `processedMessage` Uint8Array object with utf-8 encoding!

This example function takes the Uint8Array object `payload` and decodes it from base64 encoding so that it may be processed.

```javascript
const { memphis } = require("memphis-functions");

exports.handler = async (event) => {
    return await memphis.createFunction(event, eventHandler);
};

function eventHandler(payload, headers, inputs) {
    const decodedPayload = payload.toString('utf-8');
    const asJson = JSON.parse(decodedPayload);
    asJson.modified = true;

    return {
        processedMessage: Buffer.from(JSON.stringify(asJson), 'utf-8'),
        processedHeaders: headers
    };
}
```

A user created `eventHandler` may also be async:

```javascript
const { memphis } = require("memphis-functions");

exports.handler = async (event) => {
    return await memphis.createFunction(event, eventHandler);
};

async function eventHandler(payload, headers, inputs) {
    const decodedPayload = payload.toString('utf-8');
    const asJson = JSON.parse(decodedPayload);
    asJson.modified = true;

    return {
        processedMessage: Buffer.from(JSON.stringify(asJson), 'utf-8'),
        processedHeaders: headers
    };
}
```

If the user wants to have a message that they would like to validate and send to the dead letter station if the validation fails, then the user can throw an exception. In the following example, the field `check` is a boolean. The following function will send any messages that fail the `check` to the dead letter station.

```javascript
const { memphis } = require("memphis-functions");

exports.handler = async (event) => {
    return await memphis.createFunction(event, eventHandler);
};

async function eventHandler(payload, headers, inputs) {
    const decodedPayload = payload.toString('utf-8');
    const asJson = JSON.parse(decodedPayload);

    if (!asJson.check) {
        throw new Error("Validation Failed!");
    }

    return {
        processedMessage: Buffer.from(JSON.stringify(asJson), 'utf-8'),
        processedHeaders: headers
    };
}
```

If a user would rather just skip the message and not have it be sent to the station or dead letter station, the user could instead return `{ processedMessage: null, processedHeaders: null }`.

```javascript
const { memphis } = require("memphis-functions");

exports.handler = async (event) => {
    return await memphis.createFunction(event, eventHandler);
};

function eventHandler(payload, headers, inputs) {
    const decodedPayload = payload.toString('utf-8');
    const asJson = JSON.parse(decodedPayload);

    if (!asJson.check) {
        return { processedMessage: null, processedHeaders: null };
    }

    return {
        processedMessage: Buffer.from(JSON.stringify(asJson), 'utf-8'),
        processedHeaders: headers
    };
}
```

LLastly, if the user is using another data format like Protocol Buffers, the user may simply decode the `payload` into that format instead of JSON. The following example will be using the [protobufjs](https://github.com/protobufjs/protobuf.js/) package. Assuming we have a .proto definition like this: 

```proto
syntax = "proto3";
package protobuf_example;

message Message{
    string data_field = 1;
}
```

We can decode this and get the data_field out like this:

```javascript
const { memphis } = require("memphis-functions");

exports.handler = async (event) => {
    return await memphis.createFunction(event, eventHandler);
};

function eventHandler(payload, headers, inputs) {
    const root = await protobuf.load("./message.proto");
    const Message = root.lookupType('protobuf_example.Message');

    const my_message = Message.decode(payload)

    // data_field gets translated to dataField by the library...
    // Arbitrarily changing the data field
    my_message.dataField = "My new data"
    
    // .finish() returns a Uint8Array so it may just be returned as the processedMessage
    return {
      "processedMessage" : Message.encode(my_message).finish(),
      "processedHeaders" : headers
    }
}


```
