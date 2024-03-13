# Project: Script Management Service

## Overview

This project provides a web service for managing scripts. It allows adding new scripts, retrieving a list of all scripts, and fetching a specific script by its ID.

## Project Description

The project implements a web service that handles HTTP requests for script management. It utilizes the `mux` router to define request handlers. Scripts are stored in a database, and access to them is provided via a REST API. Additionally, the `Command` structure is defined to represent scripts.

## Endpoints

1. **Add Script**
   - Method: `POST`
   - Path: `/scripts`
   - Adds a new script to the database.

2. **Get All Scripts**
   - Method: `GET`
   - Path: `/scripts`
   - Retrieves a list of all scripts.

3. **Get Script by ID**
   - Method: `GET`
   - Path: `/scripts/{id}`
   - Retrieves information about a specific script by its ID.

## Running the Project

1. **Add your postgres nickname and password into .env file**

2. **run "psql -U postgres -c "CREATE DATABASE script;" in the terminal"**

3. **run "go run ./cmd"**
