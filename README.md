# A Simplified Microblogging Backend

This project is a backend implementation for a simplified microblogging platform, similar to Twitter, as part of the Ualá technical challenge.  It allows users to post tweets, follow other users, and view timelines.

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [Flow Diagrams](#flow-diagrams)
    - [Hexagonal Architecture](#hexagonal-architecture)
    - [User Interaction Flow](#user-interaction-flow)
    - [Timeline Generation](#timeline-generation)
- [API Endpoints](#api-endpoints)
- [Technologies Used](#technologies-used)
- [Project Structure](#project-structure)
- [Infrastructure Implementation](#infrastructure-implementation)
    - [Current Development Implementation](#current-development-implementation)
    - [Production-Ready Implementation](#production-ready-implementation)
    - [Infrastructure Comparison Diagram](#infrastructure-comparison-diagram)
- [Setup and Running](#setup-and-running)
    - [Prerequisites](#prerequisites)
    - [Running Locally (In-Memory Version)](#running-locally-in-memory-version)
    - [Running with Docker Compose](#running-with-docker-compose)
- [Logging and Monitoring](#logging-and-monitoring)
    - [Components](#components)
    - [Features](#features)
    - [Accessing Logs and Metrics](#accessing-logs-and-metrics)
- [Testing](#testing)

## Features

* **Post Tweets**: Users can publish short messages (up to 280 characters).
* **Follow Users**: Users can follow other users to see their tweets.
* **View Timeline**: Users can view a timeline of tweets from users they follow, optimized for reads.

## Architecture

The system uses a **Hexagonal Architecture (Ports and Adapters)** to ensure a clean separation of concerns.

### High-Level Architecture Overview

This project implements a Hexagonal Architecture (also known as Ports and Adapters) pattern, which separates the core business logic from external concerns. The architecture consists of three main layers:

1. **Domain Layer**: Contains the core business entities, logic, and interfaces (ports)
2. **Application Layer**: Orchestrates the domain objects to fulfill use cases
3. **Infrastructure Layer**: Implements the interfaces defined in the domain layer (adapters)

## Flow Diagrams

### Hexagonal Architecture

```mermaid
flowchart TB
    subgraph External World
        REST[REST API]
        DB[Database]
    end

    subgraph Adapters/Infrastructure
        REST_Adapter[REST Handlers]
        Repo_Adapter[Repository Implementations]
    end

    subgraph Application
        Use_Cases[Use Cases]
    end

    subgraph Domain
        Entities[Domain Entities]
        Ports[Ports/Interfaces]
    end

    REST --> REST_Adapter
    REST_Adapter --> Use_Cases
    Use_Cases --> Ports
    Ports --> Repo_Adapter
    Repo_Adapter --> DB
    Use_Cases --> Entities
    Entities --> Ports
```

### User Interaction Flow

```mermaid
sequenceDiagram
    participant Client
    participant REST as REST API
    participant App as Application Layer
    participant Domain as Domain Layer
    participant Repo as Repository
    participant Events as Event Bus

    %% Post Tweet Flow
    Client->>REST: POST /v1/tweets
    REST->>App: Create Tweet
    App->>Domain: Create Tweet Entity
    Domain-->>App: Return Tweet
    App->>Repo: Save Tweet
    App->>Events: Publish TweetCreated Event
    REST-->>Client: Response

    %% Follow User Flow
    Client->>REST: POST /v1/users/follow
    REST->>App: Follow User
    App->>Repo: Update Followee Relationship
    App->>Events: Publish UserFollowed Event
    REST-->>Client: Response

    %% Get Timeline Flow
    Client->>REST: GET /v1/users/timeline
    REST->>App: Get Timeline
    App->>Repo: Fetch Timeline
    Repo-->>App: Return Timeline
    App-->>REST: Return Timeline
    REST-->>Client: Timeline Response
```

### Timeline Generation

```mermaid
flowchart TD
    A[Tweet Created] --> B{Event Bus}
    B --> C[Timeline Updater Handler]
    C --> D[Get Followers]
    D --> E[For Each Follower]
    E --> F[Get Follower's Timeline]
    F --> G[Add Tweet to Timeline]
    G --> H[Update Timeline Repository]

    I[User Followed] --> B
    B --> J[Timeline Updater Handler]
    J --> K[Get Followee's Tweets]
    K --> L[Add Tweets to Follower's Timeline]
    L --> M[Update Timeline Repository]
```

## API Endpoints

* `POST /v1/tweets`
    * Header: `X-User-ID: <author_user_id>`
    * Body: `{"tweetId": "string", "content": "string"}`
* `POST /v1/users/follow`
    * Header: `X-User-ID: <follower_user_id>`
    * Body: `{"followeeId": "string"}`
* `POST /v1/users/unfollow`
    * Header: `X-User-ID: <follower_user_id>`
    * Body: `{"unFolloweeId": "string"}`
* `GET /v1/users/timeline`
    * Header: `X-User-ID: <user_id>`
* `POST /v1/users`
    * Body: `{"userId": "string"}`

## Technologies Used

* **Language**: Golang
* **Containerization**: Docker
* **Logging & Monitoring**: Grafana, Loki, Promtail, Prometheus, Node-exporter
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

## Infrastructure Implementation

### Current Development Implementation

This project currently uses in-memory implementations for all infrastructure components, making it easy to run locally without external dependencies:

* **In-Memory Event Bus**: Events (like tweet creation and user following) are published and handled within the application's memory
* **In-Memory Repositories**: All data (users, tweets, followers, timelines) is stored in memory and lost when the application restarts
* **Timeline Update Mechanism**: When events occur (new tweet, new follow), a handler updates the relevant timelines in memory

This approach is perfect for development and testing but not suitable for production use.

### Production-Ready Implementation

For a production environment, the following infrastructure would replace the in-memory implementations:

* **Event Bus**: RabbitMQ or Amazon SQS would handle event publishing and subscription
* **Databases**:
  * **MySQL**: For persistent storage of users, tweets, and follower relationships
  * **Redis**: For caching user timelines, providing fast read access
* **Timeline Updates**: Could be implemented either through:
  * Event-driven approach (current implementation): Events trigger handlers that update the Redis cache
  * Database triggers: MySQL triggers could update the Redis cache when database changes occur

### Infrastructure Comparison Diagram

```mermaid
flowchart TB
    subgraph "Development Environment"
        IME[In-Memory Event Bus]
        IMR[In-Memory Repositories]
        IMT[In-Memory Timeline]

        IME --> IMT
        IMR --> IMT
    end

    subgraph "Production Environment"
        subgraph "Event Bus"
            RMQ[RabbitMQ/SQS]
        end

        subgraph "Persistent Storage"
            SQL[MySQL Database]
            RC[Redis Cache]
        end

        RMQ --> |Timeline Update Handler| RC
        SQL --> |Store/Retrieve Data| APP
        RC --> |Fast Timeline Reads| APP
        APP --> |Publish Events| RMQ
        APP --> |Write Data| SQL
    end

    APP[Application]
```

## Setup and Running

### Prerequisites

* Go (version 1.23.9)
* Docker
* Docker Compose

### Running Locally (In-Memory Version)

```bash
make run
```

### Running with Docker Compose

To run the application with Docker Compose, which includes Grafana and Loki for log monitoring:

```bash
make docker-compose-up
```

This will start the following services:
- Microblog application (http://localhost:8080)
- Grafana (http://localhost:3000) - Use admin/admin for login
- Loki
- Promtail (for log collection)
- Prometheus (http://localhost:9090)
- Node-exporter (for system metrics collection)

To stop all services:

```bash
make docker-compose-down
```

## Logging and Monitoring

This project includes a comprehensive logging and monitoring setup using Grafana and Loki:

### Components

- **Loki**: A horizontally-scalable, highly-available log aggregation system
- **Promtail**: An agent that ships the contents of local logs to Loki
- **Prometheus**: A monitoring system and time series database for metrics collection
- **Node-exporter**: An exporter for hardware and OS metrics for Prometheus
- **Grafana**: A visualization and analytics platform for monitoring and observability

### Accessing Logs and Metrics

1. Start the services with `make docker-compose-up`
2. Open Grafana at http://localhost:3000 (login with admin/admin)
3. Navigate to the "Microblog Logs" dashboard to view and search logs in real-time
4. Navigate to the "System Metrics" dashboard to view CPU, memory, disk, and network metrics
5. Navigate to the "Application Metrics" dashboard to view HTTP request rates and durations
6. You can also access Prometheus directly at http://localhost:9090 to query metrics
