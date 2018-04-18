# bwk-dash

## Requirements
The follow items are required by the system:
- golang 1.9+
- bulwark node already setup with rpc information

## Configuration
Before running the dashboard please configure the environment variables located in the ```.env``` file in the root of the project folder.

*Note: these can also be passed via the command line.  Ex. DASH_PORT=8888 go run cmd/bwk-dash/*.go*

```
# The donation address that will be converted QR code.
DASH_DONATION_ADDRESS=TESTADDRESSHERE

# What port to run the api on.
DASH_PORT=3000

# The location of the node.
DASH_RPC_ADDR=localhost

# The rpc port for the node.
DASH_RPC_PORT=52541

# The rpc user for the node.
DASH_RPC_USER=rpcbwk

# The rpc password for the rpc user.
DASH_RPC_PASS=rpcbwkpassword!1

# Where is the website located.
DASH_WEBSITE=./client/dist
```

## Start (Testing)
*Note: build/install scripts coming soon*

To start the system please execute the following commands during development:
- ```go run cmd/bwk-dash/*.go``` - this will start the api server on the configured port in ```.env```.
- ```cd client && yarn run start``` - will start the development web server on port 8080.

The development or testing flow varies from production, ```webpack-dev-server``` is used for hot-reload of UI during development.

Go to [http://localhost:8080](http://localhost:8080)

## Build
Once the system is completed the only requirement will be to run the binary and specify the location of the website. 
The website will be built using ```webpack``` and will be served by the api. 

## Structure
The following scruture is being used:
- ```client/``` - the location of the UI, all yarn/npm commands should be ran from this folder.
- ```cmd/``` - location of golang binaries.
- ```handler/``` - api handlers
- ```rpc/``` - rpc interface for bulwark node for api.
- ```sys/``` - system related functions for api.
- ```vendor/``` - the location of 3rd party packages.
