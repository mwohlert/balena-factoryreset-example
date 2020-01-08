# Balena FactoryReset Example

Example to test the FactoryReset functionality of the balena-supervisor in conjunction with deleting all wireless connections

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

 * balenaOS compatible device running a development balenaOS or a device connected to the balena cloud.
 * wireless connection


### Installing

If you are running the development balenaOS push the application to the device locally:

```
balena push x.x.x.x --logs --source .
```

If you are using the balena cloud, follow the release instructions.

## Usage

In order to trigger a factory reset do a post request to the `factoryReset` route:

```
curl -X POST http://deviceIPAdress:8080/factoryReset
```

dd