# Palladium-Fantasy-Character-Generator
## Objective
Learning the basics and syntax of the Go programming language by creating a console app that builds characters for the Palladium table top RPG. 

## Getting Started
* Clone project
* `cd` into project root directory
  * `cd` into `docker`
  * run `docker-compose up` to bring up the database
  * use your favorite method to create db tables and insert data from the `docker/db-dump/initDb.sql` file
* Work through the build process below
* Run the program

## Build Process
* Beginning at the repo root, initialize the project
  ```
  go mod init PALLADIUM_FCG
  go mod tidy
  ```
  * This will create the `go.mod` and `go.sum` files
* Navigate to each sub-directory and build each module
  ```
  cd builder/
  go build .

  cd ../dbservice
  go build .
  ```
  * continue that until you have run a build in all sub-directories
  * the `build` command should finish without any errors
* Return to the root directory and build
  ```
  go build .
  ```
  * you should now have the binary file `PALLADIUM_FCG`
  * run it with `./PALLADIUM_FCG`

