# Hotel Test API

A RESTful API service for hotel and room management with pricing calculation capabilities.

## 🛠️ Technology Stack

### Language & Framework
- **Language:** Go 1.25.6
- **Web Framework:** Echo v4.15.0
- **Database:** SQLite with GORM ORM v1.31.1
- **Validation:** go-playground/validator v10.30.1
- **Configuration:** Viper v1.21.0

### Key Dependencies
- `gorm.io/gorm` - ORM for database operations
- `gorm.io/driver/sqlite` - SQLite driver
- `github.com/labstack/echo/v4` - HTTP web framework
- `github.com/go-playground/validator/v10` - Request validation
- `github.com/spf13/viper` - Configuration management

## Follow up via github
[Github Hotel Property Service](https://github.com/kracker71/hotel-property-service)

## 🏗️ Architecture

This project follows **Clean Architecture** (also known as **Hexagonal Architecture** or **Ports and Adapters**) principles with clear separation of concerns:

```
┌─────────────────────────────────────────────────────────┐
│                    Presentation Layer                   │
│              (HTTP Handlers & DTOs)                     │
│    internal/transport/http/{handler, dto, router}       │
└───────────────────────┬─────────────────────────────────┘
                        │
┌───────────────────────▼─────────────────────────────────┐
│                    Application Layer                    │
│                 (Business Logic & Use Cases)            │
│                  internal/service                       │
└───────────────────────┬─────────────────────────────────┘
                        │
          ┌─────────────▼─────────────┐
          │                           │
┌─────────▼─────────┐       ┌─────────▼─────────┐
│   Domain Layer    │       │    Ports Layer    │
│  (Entities & BL)  │       │   (Interfaces)    │
│  internal/domain  │       │  internal/port    │
└───────────────────┘       └─────────┬─────────┘
                                      │
                        ┌─────────────▼─────────────┐
                        │     Adapters Layer        │
                        │  (Repository Impls)       │
                        │   internal/adapter        │
                        └─────────────┬─────────────┘
                                      │
                        ┌─────────────▼─────────────┐
                        │  Infrastructure Layer     │
                        │  (Database, External)     │
                        │    internal/infra         │
                        └───────────────────────────┘
```

### Layer Responsibilities

- **Domain Layer**: Contains business entities (`Hotel`, `Room`, `Benefit`, etc.) and core business rules
- **Ports Layer**: Defines interfaces (contracts) for repositories and external services (Hexagonal Architecture ports)
- **Adapters Layer**: Implements port interfaces with concrete implementations (database repositories, mappers)
- **Application Layer**: Implements use cases and orchestrates business logic through services
- **Infrastructure Layer**: Provides foundational infrastructure (database connection, migrations, seeders)
- **Presentation Layer**: HTTP handlers, request/response DTOs, routing, and API contracts

## 📁 Project Structure

```
hotel-property-service
├── backend/
│    ├── cmd/
│    │   └── server/
│    │       └── main.go                 # Application entry point
│    ├── internal/
│    │   ├── config/
│    │   │   └── config.go               # Configuration management
│    │   ├── domain/                     # Business entities
│    │   │   ├── hotel.go
│    │   │   ├── room.go
│    │   │   ├── benefit.go
│    │   │   └── facility.go
│    │   ├── port/                       # Repository interfaces (Ports)
│    │   │   ├── hotel.go
│    │   │   └── room.go
│    │   ├── adapter/                    # Repository implementations (Adapters)
│    │   │   ├── hotel_repository.go
│    │   │   ├── room_repository.go
│    │   │   ├── entity/                 # Database/GORM models
│    │   │   │   ├── hotel.go
│    │   │   │   ├── room.go
│    │   │   │   ├── facility.go
│    │   │   │   ├── benefit.go
│    │   │   └── mapper/                 # Domain ↔ DB model mappers
│    │   │       ├── hotel.go
│    │   │       └── room.go
│    │   ├── service/                    # Business logic layer
│    │   │   ├── hotel.go
│    │   │   ├── room.go
│    │   │   └── pricing.go
│    │   ├── infra/                      # Infrastructure implementations
│    │   │   └── database/
│    │   │       ├── sqlite.go           # DB connection
│    │   │       └── seeder.go           # Data seeding
│    │   └── transport/
│    │       └── http/
│    │           ├── router.go
│    │           ├── handler/            # HTTP handlers
│    │           │   ├── hotel.go
│    │           │   ├── room.go
│    │           │   └── pricing.go
│    │           └── dto/                # Request/Response DTOs
│    │               ├── hoteldto/
│    │               ├── roomdto/
│    │               ├── pricingdto/
│    │               └── mapperdto/         # DTO ↔ Domain mappers
│    ├── docs/                           # Swagger documentation
│    │   ├── docs.go
│    │   ├── swagger.json
│    │   └── swagger.yaml
│    ├── data/                           # SQLite database files
│    ├── config.yaml                     # Application configuration
│    ├── go.mod
│    ├── go.sum
│    ├── Dockerfile
│    └── Makefile
└docker-compose.yaml
```

## 🗄️ Database Schema

### Entity Relationship Diagram

```
┌─────────────────┐          ┌─────────────────────────┐                       
│     Hotel       │          │       Facility          │           
├─────────────────┤          ├─────────────────────────┤
│ hotel_id (PK)   │◄───────► │ facility_id (PK)        │
│ name            │          │ hotel_id (FK/idx)       │
│ address         │          │ name                    │
│ is_active       │          │ description             │
│ created_at      │          │ is_active               │
│ updated_at      │          │ created_at              │
└────────┬────────┘          │ updated_at              │
         │                   └─────────────────────────┘
         │                  
         │                   
         ▼                   
┌─────────────────────────┐         ┌────────────────────────┐
│        Room             │         │     Benefit            │
├─────────────────────────┤         ├────────────────────────┤
│ room_id (PK)            │         │ benefit_id (PK)        │
│ physical_room_id (idx)  │◄───────▶│ physical_room_id (idx) │
│ hotel_id (FK)           │         │ name                   │
│ name                    │         │ description            │
│ description             │         │ is_active              │
│ type                    │         │ created_at             │
│ base_price              │         │ updated_at             │
│ currency                │         └────────────────────────┘
│ cancellation_policy     │
│ is_active               │
│ created_at              │
│ updated_at              │
└─────────────────────────┘         

```

### Tables

#### `Hotel`
| Column     | Type    | Description           |
|------------|---------|----------------------|
| hotel_id   | string  | Primary Key          |
| name       | string  | Hotel name           |
| address    | string  | Hotel address        |
| is_active  | boolean | Active status        |
| created_at | int64   | Creation timestamp   |
| updated_at | int64   | Last update timestamp|

#### `Room`
| Column              | Type    | Description                                    |
|---------------------|---------|------------------------------------------------|
| room_id             | string  | Primary Key                                    |
| physical_room_id    | string  | Physical room identifier (indexed)             |
| hotel_id            | string  | Foreign Key → Hotel                            |
| name                | string  | Room name                                      |
| description         | string  | Room description                               |
| type                | string  | Room type                                      |
| base_price          | int64   | Base price per night (integer, e.g., cents)    |
| currency            | string  | Currency code                                  |
| cancellation_policy | string  | Either `NON_REFUNDABLE` or `FREE_CANCELLATION` |
| is_active           | boolean | Active status                                  |
| created_at          | int64   | Creation timestamp                             |
| updated_at          | int64   | Last update timestamp                          |

#### `Facility`
| Column       | Type    | Description                   |
|--------------|---------|-------------------------------|
| facility_id  | string  | Primary Key                   |
| hotel_id     | string  | Foreign Key → Hotel (indexed) |
| name         | string  | Facility name                 |
| description  | string  | Description                   |
| is_active    | boolean | Active status                 |
| created_at   | int64   | Creation timestamp            |
| updated_at   | int64   | Last update timestamp         |

#### `Benefit`
| Column          | Type    | Description                        |
|-----------------|---------|------------------------------------|
| benefit_id      | string  | Primary Key                        |
| physical_room_id| string  | Foreign Key → Room.physical_room_id (indexed) |
| name            | string  | Benefit name                       |
| description     | string  | Description                        |
| is_active       | boolean | Active status                      |
| created_at      | int64   | Creation timestamp                 |
| updated_at      | int64   | Last update timestamp              |

> Notes: Benefits are associated with a room via `physical_room_id` (one-to-many). Facilities are owned by a hotel (one-to-many). Junction tables for hotel/facility or room/benefit are not used in the current adapter models.

## 📡 API Endpoints

Base URL: `http://localhost:3000/api/v1`

### 📖 Swagger Documentation

Interactive API documentation is available via Swagger UI:

**URL:** `http://localhost:3000/swagger/index.html`

The Swagger UI provides:
- Interactive API testing
- Complete request/response schemas
- Validation rules
- Example values
- Try-it-out functionality

To regenerate Swagger documentation after making changes:
```bash
cd backend
swag init -g cmd/server/main.go -o docs
```

### Hotel Endpoints

#### 1. Get All Hotels
```http
GET /api/v1/hotels
```

**Response:** `200 OK`
```json
{
  "hotels": [
    {
      "id": "uuid-string",
      "name": "Grand Hotel",
      "address": "123 Main St, City",
      "facilities": [
        {
          "id": "facility-uuid",
          "name": "Swimming Pool",
          "description": "Olympic-sized pool"
        }
      ]
    }
  ]
}
```

---

#### 2. Get Hotel by ID
```http
GET /api/v1/hotel/:hotelID
```

**Path Parameters:**
- `hotelID` (string, required): Hotel UUID

**Response:** `200 OK`
```json
{
  "hotel": {
    "id": "uuid-string",
    "name": "Grand Hotel",
    "address": "123 Main St, City",
    "facilities": [
      {
        "id": "facility-uuid",
        "name": "Swimming Pool",
        "description": "Olympic-sized pool"
      }
    ]
  }
}
```

**Error Responses:**
- `400 Bad Request`: Invalid or missing hotelID
- `500 Internal Server Error`: Server error

---

### Room Endpoints

#### 3. Get All Rooms by Hotel
```http
GET /api/v1/hotel/:hotelID/rooms
```

**Path Parameters:**
- `hotelID` (string, required): Hotel UUID

**Response:** `200 OK`
```json
{
  "rooms": [
    {
      "id": "room-uuid",
      "hotelID": "hotel-uuid",
      "name": "Deluxe Suite",
      "type": "Suite",
      "description": "Spacious room with city view",
      "basePrice": 150.00,
      "currency": "USD",
      "cancellationPolicy": "FREE_CANCELLATION",
      "benefits": [
        {
          "id": "benefit-uuid",
          "name": "Free Breakfast",
          "description": "Continental breakfast included"
        }
      ]
    }
  ]
}
```

**Error Responses:**
- `400 Bad Request`: Invalid or missing hotelID
- `500 Internal Server Error`: Server error

---

#### 4. Get Room Details
```http
GET /api/v1/hotel/:hotelID/room/:roomID
```

**Path Parameters:**
- `hotelID` (string, required): Hotel UUID
- `roomID` (string, required): Room UUID

**Validation Rules:**
- `hotelID`: Required, must be valid UUID v4
- `roomID`: Required, must be valid UUID v4

**Response:** `200 OK`
```json
{
  "room": {
    "id": "room-uuid",
    "hotelID": "hotel-uuid",
    "name": "Deluxe Suite",
    "type": "Suite",
    "description": "Spacious room with city view",
    "basePrice": 150.00,
    "currency": "USD",
    "cancellationPolicy": "Flexible",
    "benefits": [
      {
        "id": "benefit-uuid",
        "name": "Free Breakfast",
        "description": "Continental breakfast included"
      }
    ]
  }
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body or validation failed
- `500 Internal Server Error`: Server error

---

### Pricing Endpoints

#### 5. Calculate Room Pricing
```http
POST /api/v1/pricing
```

**Request Body:**
```json
{
  "hotelID": "hotel-uuid",
  "roomID": "room-uuid",
  "nights": 3
}
```

**Validation Rules:**
- `hotelID`: Required, must be valid UUID v4
- `roomID`: Required, must be valid UUID v4
- `nights`: Required, minimum value 1

**Response:** `200 OK`
```json
{
  "finalPrice": 450.00
}
```

**Pricing Calculation Formula:**
```
Final Price = Base Price × Number of Nights × Cancellation Policy (1.2 if cancellation policy is `FREE_CANCELLATION`)
```

**Error Responses:**
- `400 Bad Request`: Invalid request body or validation failed
- `500 Internal Server Error`: Server error or hotel/room mismatch

---

## 🚀 Getting Started

### Prerequisites
- Go 1.25.6 or higher
- Docker & Docker Compose (optional)
- Swag CLI tool (for regenerating API docs) (optional, already generated)

### [Optional] Install Swag CLI -- generated swagger.json is included in .zip file

```bash 
go install github.com/swaggo/swag/cmd/swag@latest 
```


### Configuration

Edit `config.yaml`:
```yaml
server:
  port: 3000

database:
  driver: sqlite
  dsn: ./data/hotel.db
  migration: true
  seeding: true
```

### Run Locally

```bash
# Go to backend directory
cd /backend

# Install dependencies
make init

# Run the application
make run
```

The API will be available at:
- API Base: `http://localhost:3000/api/v1`
- Swagger UI: `http://localhost:3000/swagger/index.html`

### Run with Docker

```bash
# Build and run
docker-compose up --build

# Stop
docker-compose down
```

### Build

```bash
# Using Go
cd /backend
go build -o hotel-api cmd/server/main.go

# Using Makefile (if available)
cd /backend
make build
```

## 📝 Business Logic

### Room Price Calculation

The pricing service calculates room prices based on:
1. **Base Price**: The nightly rate of the room
2. **Number of Nights**: Duration of stay
3. **Cancellation Policy Ratio**: Multiplier by 1.2 if `FREE_CANCELLATION`

**Validation:**
- Ensures the hotel and room IDs match
- Only active rooms can be priced
- Minimum 1 night stay required

## 🔒 Validation

The API uses `go-playground/validator` for request validation:
- UUID format validation for IDs
- Required field validation
- Minimum value checks for numeric fields

## 📦 Response Format

All successful responses return JSON with appropriate HTTP status codes:
- `200 OK`: Successful GET requests
- `400 Bad Request`: Validation errors or invalid input
- `500 Internal Server Error`: Server-side errors

Error responses follow this structure:
```json
{
  "message": "Error message description"
}
```

## 🧪 Development

### Project Conventions
- **Hexagonal Architecture** (Ports & Adapters) for clean separation
- **Dependency Injection** for better testability and flexibility
- **Repository Pattern** via ports for data access abstraction
- **DTO Pattern** for API contracts and data transfer
- **Mapper Pattern** for converting between layers (domain ↔ db, domain ↔ dto)

### Adding New Features
1. Define domain entities in `internal/domain/`
2. Create port (interface) in `internal/port/`
3. Implement adapter in `internal/adapter/` with:
   - Database models in `internal/adapter/entity/`
   - Mappers in `internal/adapter/mapper/`
4. Add business logic in `internal/service/`
5. Create DTOs in `internal/transport/http/dto/`
6. Create DTO mappers in `internal/transport/http/dto/mapper/`
7. Implement handlers in `internal/transport/http/handler/`
8. Register routes in `internal/transport/http/router.go`
