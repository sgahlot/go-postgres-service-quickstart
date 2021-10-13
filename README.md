449225
# go-service-quickstart
Sample microservice code, modeled after Quarkus quickstart apps, for PostgreSQL. It uses `binding client`
to get the connection string instead of making the connection string itself.

It allows insertion, retrieval and deletion (TBD) operations on Fruit list.

_This app is using `go-kit` and `gorilla` libraries to provide the REST endpoints._

## Run the app
_If running against local PostgreSQL, please follow the instructions given at
[Install PostgreSQL](#install-postgresql) to install PostgreSQL_

_If running against remote PostgreSQL, please update the bindings (host/username/password etc.)
in `${PWD}/test-bindings/bindings` directory to point to remote db_
 
  * Using Bindings
    * `make run`
  * Not using Bindings
    * `SERVICE_BINDING_ROOT="" DB_URL="<POSTGRES_DB_URL>" make run`

_[Operations supported](#operations-supported)_


## Build and run executable
Use following command to build and run the executable:

* Mac
  * Using bindings:
    * `TARGET_OS=darwin make run_binary`
  * Not using bindings
    * `TARGET_OS=darwin SERVICE_BINDING_ROOT="" DB_URL="<POSTGRES_DB_URL>" make run_binary`

* Linux 
  * Using bindings:
    * `make run_binary`
  * Not using bindings:
    * `SERVICE_BINDING_ROOT="" DB_URL="<POSTGRES_DB_URL>" make run_binary`

_[Operations supported](#operations-supported)_


## Build Docker image for the service
Use following commands to build and run the app in Docker container:

* Build Docker image
  * `make build_image`
* Run the app in container (will build the image if not already built)
  * _Change the `test-bindings/bindings/host` value to `postgresql`_
  * `SERVICE_BINDING_ROOT=/bindings make start_container`
* Stop the container (will also remove the container)
  * `make stop_container`

_[Operations supported](#operations-supported)_


**Alternate method to manually build and run the container:**

* Build the executable by running following command:
  * `make build_binary`
* Run following command from root directory of the project to build docker image of the service:
  ```
  docker build -t go-postgres-quickstart:0.0.1-SNAPSHOT -f resources/docker/go/Dockerfile .
  ```
  _The image will be named `go-postgres-quickstart:0.0.1-SNAPSHOT`_
* Create and run a container using the image created in previous step:
  ```
  docker run --name go-postgres-fruit-app -d -p 8080:8080 --rm
      -e SERVICE_BINDING_ROOT=<BINDING_ROOT_DIR>
      -e DB_NAME=<DB_NAME_CONTAINING_FRUIT_COLLECTION>
      go-postgres-quickstart:0.0.1-SNAPSHOT
  ```
  _This will create a container named `go-postgres-fruit-app` listening on port 8080._  

## Tools used to perform API calls
You can either use `curl` or `httpie` to invoke the API from command line.
UI support is still a TODO. Examples provide use `httpie` tool

## Operations supported
### Insert fruit

`http POST localhost:8080/api/v1/fruits name="SOME FRUIT NAME" description="SOME DESCRIPTION"`

_above will gets translated to a POST call using name and description as JSON payload_

### Retrieve all fruits

`http http://localhost:8080/api/v1/fruits\?name\="ALL"`

_above will retrieve all the fruits from database_


## Environment variables used by the service (and binding client)
* `SERVICE_BINDING_ROOT`

  _Specfies the binding root containing a separate file for each value that's used by
   binding client to make the connection string_
* `DB_URL`

  _Database URL in case above property is NOT defined or want to run this service without binding client_
* `DB_NAME`

  _Name of the database from which the fruit collection is to be retrieved and used_


## Install PostgreSQL
One can install PostgreSQL in Docker, to run the app locally.

### Start PostgreSQL

`make start_postgresql`

### Initialize db
_Above should fetch the latest image of PostgreSQL and run a container named postgres-test-db locally
The database should be initialized, but in case it is not, please run following commands, in a
terminal, to initialize the local Postgres DB:_

`docker exec -it postgres-test-db bash`

Once you're inside the postgres container, run following command:

`psql -h 127.0.0.1 -U postgres`

For password, please provide the POSTGRES_PASSWORD value from `resources/docker/postgres/.env` file

Paste the contents of `resources/docker/postgres/db-init/init.sql` in the postgres prompt

To verify that the "fruit" relation is created:

`\l` -> should list all the databases

`\c test` -> should select `test` database

To verify that the `fruit` collection is also created with 3 documents in it:

`db.fruit.find`

To exit the postgresql shell as well as postgresql docker container:
```
exit
exit 
```

### To stop PostgreSQL:

`make stop_postgresql`
