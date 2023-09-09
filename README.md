# Quantia: Banking Backend

## Introduction

Quantia is a robust backend architecture for a banking application, designed to automate and streamline banking operations. It is developed using the Go programming language, specifically version 1.21, along with Go SDK 1.21.1.

## Features

- User registration and login
- Account operations (perform deposits, withdrawals and view balance)
- Ability to transfer money between accounts

## Components

Quantia is built leveraging several advanced components:

- Go programming language
- Robust database layer for storing critical user and transactional data
- An intuitive REST API layer for seamless interfacing between the client and server

## Pre-requisites

To successfully build and run Quantia, you need:

- Go programming language version 1.21 installed
- Go SDK 1.21.1 installed
- A configured and running database

## Setting Up

To set up Quantia on your local machine, follow these steps:

1.  Clone the repo to your local machine
    ```bash
    git clone https://github.com/yourusername/quantia.git
    ```
2.  Move to the project directory
    ```bash
    cd quantia
    ```
3.  Install the necessary Go packages
    ```bash
    go get
    ```
4.  Start the server
    ```bash
    go run .
    ```

## API endpoints

Please refer to `docs/api.docs.md` for detailed information about the provided endpoints and their usage.

## Contribution

We warmly welcome any and all contributions. If you would like to contribute to the project, please fork the repository, make your changes, and submit a pull request.

## Licence

Quantia is licensed to the end-user under the terms of the MIT license.

## Feedback

Please do not hesitate to provide feedback or [report an issue](https://github.com/quabynah-bilson/quantia-go/issues). We also welcome feature 
requests.