# Golang Bootcamp

## Table of Contents  

- [Introduction](#introduction)
- [Deliverables](#Deliverables)

## Introduction

This project should be a whole Golang REST API which must incluide:
- An endpoint for reading from an external API
  - Write the information in a CSV file
- An endpoint for reading the CSV
  - Display the information as a JSON
- An endpoint for reading the CSV concurrently with some criteria (details below)
- Unit testing for the principal logic
- Follow conventions, best practices
- Clean architecture
- Go routines usage


## Deliverables

## First deliverable
  - Run the application
  - Use next endpoint to get a pokemon from the csv file:
    ```
      'http://localhost:8080/get/pokemons/:pokemonID'
    ```
    - If the pokemon exists you should get a JSON like this:
      ```json
        {
          "id": 2,
          "name": " test 2",
          "base_experience": "456"
        }
      ```
    - If not you should get an error message.

## Second deliverable
  - Pending

## Final deliverable
  - Pending



