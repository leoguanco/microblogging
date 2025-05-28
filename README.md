# A Simplified Microblogging Backend

This project is a backend implementation for a simplified microblogging platform, similar to Twitter, as part of the Ualá technical challenge.  It allows users to post tweets, follow other users, and view timelines.

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [API Endpoints](#api-endpoints)
- [Technologies Used](#technologies-used)
- [Project Structure](#project-structure)
- [Setup and Running](#setup-and-running)
    - [Prerequisites](#prerequisites)
    - [Running Locally (In-Memory Version)](#running-locally-in-memory-version)
- [Testing](#testing)

## Features

* **Post Tweets**: Users can publish short messages (up to 280 characters).
* **Follow Users**: Users can follow other users to see their tweets.
* **View Timeline**: Users can view a timeline of tweets from users they follow, optimized for reads.

## Architecture

The system uses a **Hexagonal Architecture (Ports and Adapters)** to ensure a clean separation of concerns.

### High-Level Architecture Overview

## API Endpoints

* `POST /v1/tweets`
    * Header: `X-User-ID: <author_user_id>`
    * Body: `{"content": "string"}`
* `POST /v1/users/{user_to_follow_id}/follow`
    * Header: `X-User-ID: <follower_user_id>`
* `POST /v1/users/{user_to_unfollow_id}/unfollow`
    * Header: `X-User-ID: <follower_user_id>`
* `GET /v1/users/{user_id}/timeline`

(More details in the architecture document)

## Technologies Used

* **Language**: Golang
* **Containerization**: Docker
* **(Development)**: In-memory data stores.

## Project Structure

This project follows a clean and modular structure based on hexagonal architecture principles:

* `/cmd/api`: Entry point for the application, contains main.go and server initialization
* `/internal`: Core application code isn't meant to be imported by external projects
  * `/domain`: Contains core business logic, entities, and domain interfaces (ports)
  * `/application`: Application services that orchestrate domain objects to fulfill use cases
  * `/infrastructure`: Implementations of the domain interfaces (adapters) including repositories and external services
* `/pkg`: Reusable packages that could be imported by other projects
* `Dockerfile`: Container definition for production deployment
* `Makefile`: Common commands for building, testing, and running the application

This structure separates business logic from technical implementations, making the codebase more maintainable and testable.

## Setup and Running

### Prerequisites

* Go (version 1.23.9)
* Docker
* Docker Compose

### Running Locally (In-Memory Version)

```bash
cd microblogging
make run