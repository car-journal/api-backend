# Design Document: Maintenance Management

## Overview

This document describes the technical design for the maintenance management feature in the Car Journal backend. The feature introduces two new domains — **maintenance categories** and **maintenance entries** — following the existing domain-driven structure under `internal/domain/`. Both domains mirror the patterns already established by `fuelentry` and `odometerentry`, including GORM-based repositories with upsert-on-conflict, service-layer orchestration, and HTTP controllers wired through the existing container and routing infrastructure.

The maintenance entry domain reuses the odometer upsert pattern from `fuelentryservice`: before persisting a new entry, the service checks whether an `OdometerEntry` with the given reading already exists via `OdometerEntryRepository.FindByReading`. If it does, the existing ID is reused; otherwise a new `OdometerEntry` is created first.

---

## Architecture

The feature follows the same layered architecture used throughout the codebase:

```
HTTP Request
    │
    ▼
Controller (transport/http/controller/internals/)
    │  parses body/query params, calls service, writes response
    ▼
Service (internal/domain/{domain}/service/)
    │  orchestrates business logic, calls repositories
    ▼
Repository (internal/domain/{domain}/repository/)
    │  executes GORM queries against PostgreSQL
    ▼
Model (internal/model/)
    │  GORM struct with BaseModel (UUID PK, soft-delete, AffectedRecords)
    ▼
PostgreSQL (existing migrations already applied)
```

The two new domains are:

- `internal/domain/maintenancecategory/` — CRUD for `maintenance_categories`
- `internal/domain/maintenanceentry/` — CRUD for `maintenance_entries`, with odometer upsert logic

> **Note:** The existing directory `internal/domain/maintenancecategory/` and `internal/domain/maintenanceentriy/` (typo in current tree) are empty placeholders. The implementation will populate `maintenancecategory` and create `maintenanceentry` (corrected spelling).

---

## Components and Interfaces

### Maintenance Category Domain

**Package layout** (mirrors `internal/domain/car/`):

```
internal/domain/maintenancecategory/
  const/    maintenancecategory_const.go   — valid sort map
  payload/  maintenancecategory_payload.go — CreatePayload, ListPayload, UpdatePayload
  repository/ maintenancecategory_repository.go
  service/    maintenancecategory_service.go
```

**Repository interface:**

```go
type Interface interface {
    Save(ctx context.Context, model internalmodel.MaintenanceCategory) error
    List(ctx context.Context, payload maintenancecategorypayload.ListPayload) ([]*internalmodel.MaintenanceCategory, error)
    FindByID(ctx context.Context, id string) (*internalmodel.MaintenanceCategory, error)
    Delete(ctx context.Context, id string) error
}
```

**Service interface:**

```go
type Interface interface {
    Create(ctx context.Context, payload maintenancecategorypayload.CreatePayload) error
    List(ctx context.Context, payload maintenancecategorypayload.ListPayload) ([]*internalmodel.MaintenanceCategory, error)
    FindByID(ctx context.Context, id string) (*internalmodel.MaintenanceCategory, error)
    Update(ctx context.Context, payload maintenancecategorypayload.UpdatePayload) error
    Delete(ctx context.Context, id string) error
}
```

---

### Maintenance Entry Domain

**Package layout** (mirrors `internal/domain/fuelentry/`):

```
internal/domain/maintenanceentry/
  const/    maintenanceentry_const.go   — valid sort map
  payload/  maintenanceentry_payload.go — CreatePayload, ListPayload, UpdatePayload
  repository/ maintenanceentry_repository.go
  service/    maintenanceentry_service.go
```

**Repository interface:**

```go
type Interface interface {
    Save(ctx context.Context, model internalmodel.MaintenanceEntry) error
    List(ctx context.Context, payload maintenanceentrypayload.ListPayload) ([]*internalmodel.MaintenanceEntry, error)
    FindByID(ctx context.Context, id string) (*internalmodel.MaintenanceEntry, error)
    Delete(ctx context.Context, id string) error
}
```

**Service interface:**

```go
type Interface interface {
    Create(ctx context.Context, payload maintenanceentrypayload.CreatePayload) error
    List(ctx context.Context, payload maintenanceentrypayload.ListPayload) ([]*internalmodel.MaintenanceEntry, error)
    FindByID(ctx context.Context, id string) (*internalmodel.MaintenanceEntry, error)
    Update(ctx context.Context, payload maintenanceentrypayload.UpdatePayload) error
    Delete(ctx context.Context, id string) error
}
```

The `MaintenanceEntryService` depends on:
- `MaintenanceCategoryRepository` — to verify `CategoryID` exists
- `MaintenanceEntryRepository` — for persistence
- `OdometerEntryRepository` — for the odometer upsert pattern
- `uuid.UUIDInterface` — for ID generation

---

### HTTP Controllers

One file per action, under:

```
transport/http/controller/internals/maintenancecategory/
  create.go
  list.go
  find_by_id.go
  update.go
  delete.go

transport/http/controller/internals/maintenanceentry/
  create.go
  list.go
  find_by_id.go
  update.go
  delete.go
```

All controllers follow the existing pattern:
1. Parse body/query params via `bodyparser.Parse` or `filter.ParsePage`
2. Wrap execution in `database.Run` (transaction)
3. Extract auth user via `authz.GetAuthUser` where needed
4. Delegate to service
5. Write response via `parser.JSON`

---

### Routing

Two new route files under `transport/http/route/internals/`:

```
internal_maintenancecategory_routes.go
internal_maintenanceentry_routes.go
```

Both are registered in `internal_routes.go` alongside the existing domains. New prefix constants are added to `transport/http/route/const/prefix.go`:

```go
MaintenanceCategoryPrefix = "/maintenance-categories"
MaintenanceEntryPrefix    = "/maintenance-entries"
```

---

### Container Wiring

`transport/container/repository.go` gains:
- `MaintenanceCategory maintenancecategoryrepository.Interface`
- `MaintenanceEntry    maintenanceentryrepository.Interface`

`transport/container/service.go` gains:
- `MaintenanceCategory maintenancecategoryservice.Interface`
- `MaintenanceEntry    maintenanceentryservice.Interface`

---

## Data Models

### `MaintenanceCategory` (existing model, already in `internal/model/maintenance_category.go`)

```go
type MaintenanceCategory struct {
    BaseModel
    Name        string  `json:"name"`
    Description *string `json:"description"`
}
```

Table: `maintenance_categories` (migration already applied)

| Column        | Type         | Notes                    |
|---------------|--------------|--------------------------|
| id            | UUID PK      | from BaseModel           |
| name          | VARCHAR(255) | NOT NULL                 |
| description   | VARCHAR(255) | nullable                 |
| created_at    | TIMESTAMPTZ  |                          |
| updated_at    | TIMESTAMPTZ  |                          |
| deleted_at    | TIMESTAMPTZ  | soft delete via GORM     |

---

### `MaintenanceEntry` (new model, to be added to `internal/model/`)

```go
const MaintenanceEntryTableName = "maintenance_entries"

type MaintenanceEntry struct {
    BaseModel

    CarID           uuid.UUID `json:"car_id"`
    OdometerEntryID uuid.UUID `json:"odometer_entry_id"`
    CategoryID      uuid.UUID `json:"category_id"`
    Brand           string    `json:"brand"`
    Name            string    `json:"name"`
    Price           float64   `json:"price"`
    PerformedAt     time.Time `json:"performed_at"`
    Notes           *string   `json:"notes"`
}
```

Table: `maintenance_entries` (migration already applied)

| Column            | Type         | Notes                                          |
|-------------------|--------------|------------------------------------------------|
| id                | UUID PK      | from BaseModel                                 |
| car_id            | UUID         | NOT NULL                                       |
| odometer_entry_id | UUID         | FK → odometer_entries(id) ON DELETE CASCADE    |
| category_id       | UUID         | FK → maintenance_categories(id) ON DELETE CASCADE |
| brand             | VARCHAR(255) | NOT NULL                                       |
| name              | VARCHAR(255) | NOT NULL                                       |
| price             | FLOAT        | NOT NULL                                       |
| performed_at      | TIMESTAMPTZ  | NOT NULL, DEFAULT NOW()                        |
| notes             | VARCHAR(255) | nullable                                       |
| created_at        | TIMESTAMPTZ  |                                                |
| updated_at        | TIMESTAMPTZ  |                                                |
| deleted_at        | TIMESTAMPTZ  | soft delete via GORM                           |

---

### Payload Types

**`maintenancecategorypayload`**

```go
type CreatePayload struct {
    Name        string  `json:"name"        validate:"required"`
    Description *string `json:"description"`
}

type ListPayload struct {
    Name        string       `json:"name"`
    Description string       `json:"description"`
    Sorts       []string     `json:"sorts"`
    PageParams  pagination.Page
}

type UpdatePayload struct {
    ID          string          `json:"id"`
    Name        *string         `json:"name"`
    Description nullable.String `json:"description"`
}
```

**`maintenanceentrypayload`**

```go
type CreatePayload struct {
    CarID           string    `json:"car_id"            validate:"required"`
    CategoryID      string    `json:"category_id"       validate:"required"`
    OdometerReading float64   `json:"odometer_reading"  validate:"required"`
    ReadingUnit     string    `json:"reading_unit"      validate:"required"`
    Brand           string    `json:"brand"             validate:"required"`
    Name            string    `json:"name"              validate:"required"`
    Price           float64   `json:"price"             validate:"required"`
    PerformedAt     time.Time `json:"performed_at"      validate:"required"`
    Notes           *string   `json:"notes"`
}

type ListPayload struct {
    CategoryIDs []string     `json:"category_ids"`
    Sorts       []string     `json:"sorts"`
    PageParams  pagination.Page
}

type UpdatePayload struct {
    ID              string          `json:"id"`
    CategoryID      *string         `json:"category_id"`
    OdometerReading *float64        `json:"odometer_reading"`
    ReadingUnit     *string         `json:"reading_unit"     validate:"required_with=OdometerReading"`
    Brand           *string         `json:"brand"`
    Name            *string         `json:"name"`
    Price           *float64        `json:"price"`
    PerformedAt     nullable.Time   `json:"performed_at"`
    Notes           nullable.String `json:"notes"`
}
```

---

### Sort Maps (in `const/` packages)

**`maintenancecategoryconst`**

```go
var ValidSorts = map[string]string{
    "name":       "name",
    "created_at": "created_at",
    "updated_at": "updated_at",
}
```

**`maintenanceentryconst`**

```go
var ValidSorts = map[string]string{
    "brand":        "brand",
    "name":         "name",
    "price":        "price",
    "performed_at": "performed_at",
    "created_at":   "created_at",
    "updated_at":   "updated_at",
}
```

---

## Correctness Properties

*A property is a characteristic or behavior that should hold true across all valid executions of a system — essentially, a formal statement about what the system should do. Properties serve as the bridge between human-readable specifications and machine-verifiable correctness guarantees.*

### Property 1: Category create–read round trip

*For any* valid `CreatePayload` (non-empty `Name`, optional `Description`), creating a `MaintenanceCategory` and then retrieving it by the returned ID should yield a record whose `Name` and `Description` match the payload.

**Validates: Requirements 1.1, 3.1**

---

### Property 2: Category upsert idempotence

*For any* `MaintenanceCategory` record, saving it a second time with the same `id` but different field values should result in the stored record reflecting the second write's values (not the first).

**Validates: Requirements 1.2**

---

### Property 3: Empty or whitespace name is rejected

*For any* string composed entirely of whitespace characters (including the empty string), attempting to create a `MaintenanceCategory` with that string as `Name` should return a validation error and leave the category count unchanged.

**Validates: Requirements 1.3**

---

### Property 4: Category list pagination invariant

*For any* set of N inserted `MaintenanceCategory` records and any valid `PageParams` (limit L, page P), the returned list length should be `min(L, max(0, N - (P-1)*L))` and `AffectedRecords` should equal N.

**Validates: Requirements 2.1, 2.5, 2.6**

---

### Property 5: Category name filter correctness

*For any* non-empty name filter string and any set of categories, every record returned by `List` should contain the filter string as a case-insensitive substring of its `name`.

**Validates: Requirements 2.2**

---

### Property 6: Category description filter correctness

*For any* non-empty description filter string and any set of categories, every record returned by `List` should contain the filter string as a case-insensitive substring of its `description`.

**Validates: Requirements 2.3**

---

### Property 7: Category not-found for unknown ID

*For any* UUID that was never inserted (or was soft-deleted), calling `FindByID` should return a `RecordNotFound` error.

**Validates: Requirements 3.2, 4.2**

---

### Property 8: Category update round trip

*For any* existing `MaintenanceCategory` and any `UpdatePayload` with non-nil fields, after calling `Update` the record retrieved by `FindByID` should reflect all non-nil fields from the payload.

**Validates: Requirements 4.1, 4.3, 4.4**

---

### Property 9: Category soft delete

*For any* existing `MaintenanceCategory`, after calling `Delete`, a subsequent `FindByID` should return `RecordNotFound`, and a raw unscoped query should still find the record with `deleted_at` set to a non-null value.

**Validates: Requirements 5.1, 5.2**

---

### Property 10: Odometer upsert — reuse existing reading

*For any* `OdometerReading` value for which an `OdometerEntry` already exists, creating a `MaintenanceEntry` with that reading should result in the entry's `odometer_entry_id` equaling the pre-existing `OdometerEntry`'s ID (no new `OdometerEntry` row is created).

**Validates: Requirements 6.1, 6.2, 9.3**

---

### Property 11: Odometer upsert — create new reading

*For any* `OdometerReading` value for which no `OdometerEntry` exists, creating a `MaintenanceEntry` with that reading should result in a new `OdometerEntry` row being created and the entry's `odometer_entry_id` referencing it.

**Validates: Requirements 6.1, 6.3, 9.3**

---

### Property 12: Entry create–read round trip

*For any* valid `CreatePayload`, creating a `MaintenanceEntry` and then retrieving it by the returned ID should yield a record whose fields (`CategoryID`, `Brand`, `Name`, `Price`, `PerformedAt`, `Notes`) match the payload.

**Validates: Requirements 6.4, 8.1**

---

### Property 13: Invalid category ID is rejected

*For any* UUID that does not correspond to an existing `MaintenanceCategory`, attempting to create or update a `MaintenanceEntry` with that `CategoryID` should return a `RecordNotFound` error.

**Validates: Requirements 6.6, 9.4**

---

### Property 14: Entry list pagination invariant

*For any* set of N inserted `MaintenanceEntry` records and any valid `PageParams`, the returned list length and `AffectedRecords` should satisfy the same invariant as Property 4.

**Validates: Requirements 7.1, 7.4, 7.5**

---

### Property 15: Entry category filter correctness

*For any* non-empty `CategoryIDs` filter and any set of entries, every record returned by `List` should have a `category_id` that is a member of the filter set.

**Validates: Requirements 7.2**

---

### Property 16: Entry not-found for unknown ID

*For any* UUID that was never inserted (or was soft-deleted), calling `FindByID` on `MaintenanceEntryService` should return a `RecordNotFound` error.

**Validates: Requirements 8.2, 9.2**

---

### Property 17: Entry update round trip

*For any* existing `MaintenanceEntry` and any `UpdatePayload` with non-nil fields, after calling `Update` the record retrieved by `FindByID` should reflect all non-nil fields from the payload.

**Validates: Requirements 9.1**

---

### Property 18: Entry soft delete

*For any* existing `MaintenanceEntry`, after calling `Delete`, a subsequent `FindByID` should return `RecordNotFound`, and a raw unscoped query should still find the record with `deleted_at` set.

**Validates: Requirements 10.1, 10.2**

---

## Error Handling

All error handling follows the existing `httperror` pattern:

| Condition | Error Type | HTTP Status |
|---|---|---|
| `Name` is empty in `CreatePayload` | `errortype.InvalidInput` | 400 |
| Request body fails `validate` tags | `errortype.InvalidInput` | 400 |
| Record not found by ID | `errortype.RecordNotFound` | 404 |
| `CategoryID` does not exist | `errortype.RecordNotFound` | 404 |
| Unexpected DB or runtime error | `errortype.InternalServer` | 500 |

The `bodyparser.Parse` function handles struct validation via `go-playground/validator` and returns an `InvalidInput` error automatically. Services return `httperror.New(errortype.RecordNotFound, ...)` for missing records, matching the pattern in `fuelentryservice` and `carservice`. The `parser.JSON` writer in each controller reads the error type from `httperror.Interface` and sets the appropriate HTTP status code.

---

## Testing Strategy

### Unit / Example Tests

Each service method gets example-based tests covering:
- Happy path (valid input → expected output)
- `RecordNotFound` path (unknown ID → error)
- Validation error path (empty `Name` → error)
- Category-not-found path for entry create/update

Each controller gets example-based tests covering:
- HTTP 201 on successful create (with ID in response)
- HTTP 200 on successful list, find-by-id, update, delete
- HTTP 400 on invalid request body
- HTTP 404 on `RecordNotFound` from service
- HTTP 500 on unexpected service error

### Property-Based Tests

Property-based testing is appropriate here because both domains involve pure data transformation logic (create/read/update/delete with filtering and pagination), and input variation (random names, descriptions, readings, IDs) meaningfully exercises edge cases.

**Library:** [`pgregory.net/rapid`](https://github.com/flyingmutant/rapid) — idiomatic Go PBT library with no external dependencies.

**Configuration:** Each property test runs a minimum of **100 iterations**.

**Tag format:** Each property test is annotated with a comment:
```
// Feature: maintenance-management, Property N: <property text>
```

**Properties to implement as property-based tests** (one test per property):

| Property | Test file | What varies |
|---|---|---|
| 1 — Category create–read round trip | `maintenancecategory_service_test.go` | Name, Description |
| 2 — Category upsert idempotence | `maintenancecategory_repository_test.go` | Name, Description on second save |
| 3 — Empty/whitespace name rejected | `maintenancecategory_service_test.go` | Whitespace string variants |
| 4 — Category list pagination invariant | `maintenancecategory_repository_test.go` | N records, limit, page |
| 5 — Category name filter | `maintenancecategory_repository_test.go` | Name values, filter substring |
| 6 — Category description filter | `maintenancecategory_repository_test.go` | Description values, filter substring |
| 7 — Category not-found | `maintenancecategory_service_test.go` | Random UUIDs |
| 8 — Category update round trip | `maintenancecategory_service_test.go` | UpdatePayload field combinations |
| 9 — Category soft delete | `maintenancecategory_repository_test.go` | Any valid category |
| 10 — Odometer upsert reuse | `maintenanceentry_service_test.go` | OdometerReading values |
| 11 — Odometer upsert create | `maintenanceentry_service_test.go` | OdometerReading values |
| 12 — Entry create–read round trip | `maintenanceentry_service_test.go` | All CreatePayload fields |
| 13 — Invalid category ID rejected | `maintenanceentry_service_test.go` | Random category UUIDs |
| 14 — Entry list pagination invariant | `maintenanceentry_repository_test.go` | N records, limit, page |
| 15 — Entry category filter | `maintenanceentry_repository_test.go` | CategoryIDs filter sets |
| 16 — Entry not-found | `maintenanceentry_service_test.go` | Random UUIDs |
| 17 — Entry update round trip | `maintenanceentry_service_test.go` | UpdatePayload field combinations |
| 18 — Entry soft delete | `maintenanceentry_repository_test.go` | Any valid entry |

Repository-level property tests use a real PostgreSQL test database (via `testcontainers-go` or a local test DB). Service-level property tests mock the repository interfaces to keep tests fast and deterministic.
