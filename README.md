# Booking Backend Architecture

## Monorepo Root Structure
The repository is organized as a monorepo utilizing npm workspaces to manage dependencies and orchestrate scripts across multiple microservices. The root package.json defines the workspaces encompassing three core services: hotel-service, booking-service, and notification-service. This structure allows for centralized dependency management while keeping the codebases logically separated. The root directory also maintains shared configurations for code formatting tools like Prettier, ensuring consistency across all enclosed microservices. Scripts defined at the root level permit the execution of development servers for individual workspaces from a central location.

## Microservice Specifications

### 1. Hotel Service
**Specification:** This service is responsible for managing hotel and room inventories. It is built on Node.js using the Express framework. It utilizes Sequelize as its Object-Relational Mapper (ORM) to interface with a MySQL database.
**Working and Approach:** The service exposes a RESTful API architecture. It processes requests related to hotel entities and room categories. Database migrations and model synchronization are handled through Sequelize-CLI. Data validation is strictly enforced using Zod schemas to ensure structural integrity of incoming payloads before they reach the controller layer. Application logging is managed by Winston, configured for daily log rotation. The folder structure follows a standard separation of concerns, dividing logic into controllers, routers, models, and repositories.

### 2. Booking Service
**Specification:** This service handles the core transactional logic of the application, managing reservation creation and confirmation. It operates on Node.js and Express, utilizing Prisma ORM with a MariaDB adapter for database operations. It integrates Redis via ioredis and utilizes redlock for distributed locking mechanisms.
**Working and Approach:** The service intercepts incoming booking requests and interfaces with MariaDB for high-performance data manipulation. The architecture separates concerns into controllers for routing, services for business logic, and repositories for direct database access. Prisma's schema definition provides a strongly typed database client, which reduces runtime errors and streamlines query construction.

### 3. Notification Service
**Specification:** This service functions as an asynchronous worker system dedicated to handling outbound communications, primarily email dispatch. It relies on BullMQ backed by Redis for message queuing. Nodemailer is used for email delivery, coupled with Handlebars for compiling dynamic email templates.
**Working and Approach:** To prevent the core booking flows from being blocked by slow network operations, this service decouples notification logic from the main request lifecycle. It listens to Redis queues for incoming jobs (such as booking confirmations). Upon receiving a job, dedicated processors compile the required email template using Handlebars and dispatch the payload via Nodemailer.

## Core Architectural Concepts

### Idempotency
To prevent unintended duplicate operations, specifically in scenarios where network instability might cause a client to retry a booking confirmation, the system implements idempotency keys.
**Approach:** During the initial creation of a booking, the service generates a unique idempotency key using the uuid library. This key is stored in the database within a dedicated IdempotencyKey table and linked via a relation to the newly created booking record. When a client subsequently requests to confirm the booking, they must provide this key. The system verifies the state of the key in the database. If the key indicates that the operation has already been finalized, the system rejects the duplicate request. This guarantees that critical state changes occur exactly once.

### Distributed Locks for Maximum Isolation (Redlock)
To ensure maximum isolation during concurrent operations on shared resources, the booking system employs distributed locking via the Redlock algorithm.
**Approach:** Before a booking can be created, the system must guarantee that multiple concurrent requests do not attempt to book the same hotel resources simultaneously. When a request enters the createBookingService, the system attempts to acquire a Redis lock on a specific resource identifier (e.g., hotel:[hotelId]) for a predefined Time-To-Live (TTL). If the lock is successfully acquired, the process continues, writing to the database and generating the idempotency key. If another request attempts to acquire the lock for the same resource while it is held, the acquisition fails, and the system immediately halts the operation, throwing an error indicating the resource is temporarily locked. This mechanism enforces strict serialization at the application level, preventing race conditions before database interaction occurs.

### ACID Compliant Transactions
For operations requiring strict adherence to Atomicity, Consistency, Isolation, and Durability, the system utilizes database-level transactions combined with explicit row-level locking.
**Approach:** The booking confirmation process involves multiple critical database updates that must either succeed completely or fail entirely without leaving partial data. The service executes these operations within a Prisma $transaction block. To handle isolation at the database level, the transaction begins by executing a raw SQL query: `SELECT * FROM IdempotencyKey WHERE key=${key} FOR UPDATE`. The FOR UPDATE clause instructs the MariaDB engine to place an exclusive lock on the specific row retrieved. Any concurrent transaction attempting to read or modify this row is forced to wait until the current transaction completes. Once the row is secured, the system validates the idempotency state, updates the booking, marks the key as finalized, and commits the transaction. If any failure occurs during this sequence, the transaction is rolled back, ensuring complete data integrity.
