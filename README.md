![image](https://github.com/memphisdev/memphis-dev-functions/assets/70286779/bf9e6907-d060-44db-abbb-42efa4b630b0)

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
  
# Welcome to Memphis Functions
Memphis Functions Is A Faster, Easier, And Dev-First Platform For<br>
Developing Highly-Customizable Event-driven Features and Pipelines By Creating or Employing Serverless Functions.<br>
Memphis Functions has the capability to seamlessly integrate with various streaming platforms and other sources, such as Kafka and S3, enhancing their functionality.
<br><br>
<strong>This repository hosts public functions created by Memphis.dev and the community,<br>accessible across all Memphis.dev deployments and the Cloud.</strong>

## How to attach a new function
1. [Create](https://cloud.memphis.dev) a Memphis.dev Account
2. Create a new station
3. Attach a new function
4. Produce some messages

## How to develop a new private function
A function comprises code files (based on [Memphis template](https://github.com/memphisdev/memphis-dev-academy/tree/master/memphis-functions)) and a `memphis.yaml` file contained within a unified directory.<br>
The directory ought to be included in a Git repository that's linked with Memphis.<br>
Here is a brief hirarchy diagram of how a compatible function file tree should be constructed: <img src="https://github.com/memphisdev/memphis-dev-functions/assets/70286779/57591907-ce74-454c-a9e3-f7348db88c48" width="50%" />
<br>

**Step-by-step Guide:**
1. Clone or create a new repository (At the moment, support is exclusively available for GitHub.)
2. Within the repository, establish a fresh directory and initialize it to your chosen programming language.
```bash
mkdir my-function && cd my-function && npm init -y
```
3. Write your code inside the `eventHandler` block.<br>Incoming events will be accumulated and dispatched to a function collectively in a batch, therefore the wrapper.
```js
export const handler = async (event) => {
    return await createFunction(event, eventHandler);
};

/**
 * https://github.com/memphisdev/memphis.js/tree/functions_wrapper#creating-a-memphis-function
 * @param {Uint8Array} payload 
 * @param {Object} headers 
 * @param {Object} inputs 
 * @returns {Object} 
 */
function eventHandler(payload, headers, inputs) {
    // Handle event here

    // A short example of converting the payload to a json object and returning it as an Uint8Array
    const decodedPayload = payload.toString('utf-8');
    const asJson = JSON.parse(decodedPayload);

    return {
        processedMessage: Buffer.from(JSON.stringify(asJson), 'utf-8'),
        processedHeaders: headers
    };
}
```
Messages will return to the Memphis Station in a batch as well.<br>
4. Connect the designated repository with your Memphis account<br>
5. `my-function` should be available through the Functions main page or a station

## How to develop a new public function
1. Fork https://github.com/memphisdev/memphis-dev-functions
2. Add your function's directory, including memphis.yaml file
3. Create a PR
4. The addition of the new function will take place following a thorough review and subsequent approval
5. Get your swag pack! 😍
