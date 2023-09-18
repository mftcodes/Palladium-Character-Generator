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
  go mod init pfcg
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
* Return to the root directory and install
  ```
  go install pfcg
  ```
  * You should be able to run the program now with the command
    ```
    pfcg
    ```
  * If the command is not found, try this first
    ```
    export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
    ```
  * Alternatively, you can run `go build .` and then run the binary that is created
