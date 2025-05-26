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
    - [Running with Docker (Production-like with Cassandra/Redis)](#running-with-docker-production-like-with-cassandraredis)
- [Testing](#testing)

## Features

* **Post Tweets**: Users can publish short messages (up to 280 characters).
* **Follow Users**: Users can follow other users to see their tweets.
* **View Timeline**: Users can view a timeline of tweets from users they follow, optimized for reads.

## Architecture

The system uses a **Hexagonal Architecture (Ports and Adapters)**  to ensure a clean separation of concerns.
Refer to the [High-Level Architecture Document](PATH_TO_ARCHITECTURE_DOC_OR_WIKI_PAGE) for more details.

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
* **Databases (Production Target)**: Apache Cassandra, Redis
* **Containerization**: Docker
* **Orchestration (Target)**: Kubernetes
* **(Development)**: In-memory data stores.

## Project Structure

(Briefly describe the Golang project structure as outlined above)

## Setup and Running

### Prerequisites

* Go (version X.Y.Z)
* Docker (if running with Docker)
* Docker Compose (if using it for multi-container setup)

### Running Locally (In-Memory Version)

```bash
git clone <repository_url>
cd ualatter
# (Instructions for setting up any environment variables if needed)
go run cmd/api/main.go