# bwk-dash

![Bulwark Home Node Dashboard](/client/src/img/screenshot.png?raw=true "Bulwark Home Node Dashboard")

# Updating README

## Requirements
The follow items are required by the system:
- golang 1.9+
- bulwark node already setup with rpc information


## Configuration
Before running the dashboard please configure the environment variables located in the ```.env``` file in the root of the project folder.

*Note: these can also be passed via the command line.  Ex. ```DASH_PORT=8888 go run cmd/bwk-dash/*.go```*


## Start (Testing)
To start the system please execute the following commands during development:
- ```go run cmd/bwk-dash/*.go``` - this will start the api server on the configured port in ```.env```.
- ```cd client && yarn run start``` - will start the development web server on port 8080.

The development or testing flow varies from production, ```webpack-dev-server``` is used for hot-reload of UI during development.

Go to [http://localhost:8080](http://localhost:8080)


## Build
Once the system is completed the only requirement will be to run the binary and specify the location of the website. 
The website will be built using ```webpack``` and will be served by the api. 

To build the binary and install it into the ```$GOBIN``` or ```$GOPATH/bin``` folder please run the build script ```./script/build.sh```.  Once finished test to make sure it was installed in a reachable path by running ```which bwk-dash```, you should get the absolute path of the binary back on success.


## Structure
The following scruture is being used:
- ```client/``` - the location of the UI, all yarn/npm commands should be ran from this folder.
- ```cmd/``` - location of golang binaries.
- ```handler/``` - api handlers
- ```rpc/``` - rpc interface for bulwark node for api.
- ```script/``` - has the build script and other helpful scripts to come.
- ```sys/``` - system related functions for api.
- ```vendor/``` - the location of 3rd party packages.


## TODO
- Finish cron binary to query at intervals and store in sqlite3 database.
- Change API handler to query database for information.
  - Maybe setup cache at equal interval to cron binary.
- Add middleware for database connection.
