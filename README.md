# Palladium-Character-Generator
## Objective
Learning the basics and syntax of the Go programming language by creating a console app that builds characters for the Palladium table top RPG. 

## Getting Started
* Clone project
* `cd` into project root directory
  * `cd` into `docker`
  * run `docker-compose up` to bring up the database
  * use your favorite method to create db tables and insert data from the `docker/db-dump/initDb.sql` file
* `cd` back to root
  * init this directory however you like
    * e.g. (`go mod init pcg`)
  * now you will need to go into each module directory and init those directories
    * e.g. `go mod init pcg/character`
    * you will also have to run some commands for replacing modules to local directories
      * e.g. `go mod edit -replace pcg/character=../character`
  * run `go mod tidy` in all directories and sub directories
* now you shuld be able to `cd` back to the root and run with `go run .` 
  * If you have erros read the output carefully, and along the directions above you, should be able to repeat steps above and fix the errors.