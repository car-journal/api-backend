# Requirements Document

## Introduction

This feature adds maintenance tracking to the Car Journal backend. It introduces two new domains: **maintenance categories** (a reference table for classifying maintenance work) and **maintenance entries** (individual maintenance records tied to a car and an odometer reading). The feature follows the existing domain-driven structure under `internal/domain/` and reuses the odometer entry upsert pattern already established in the fuel entry domain.

## Glossary

- **MaintenanceCategory**: A reference record that classifies a type of vehicle maintenance (e.g., "Powertrain", "Chassis & Handling").
- **MaintenanceEntry**: A record of a specific maintenance event performed on a car, linked to a category and an odometer reading.
- **OdometerEntry**: An existing model that records a car's odometer reading at a point in time.
- **OdometerEntryRepository**: The existing repository for `odometer_entries`, which exposes `FindByReading`.
- **MaintenanceCategoryRepository**: The repository responsible for persistence operations on `maintenance_categories`.
- **MaintenanceCategoryService**: The service layer that orchestrates business logic for maintenance categories.
- **MaintenanceEntryRepository**: The repository responsible for persistence operations on `maintenance_entries`.
- **MaintenanceEntryService**: The service layer that orchestrates business logic for maintenance entries.
- **Controller**: An HTTP handler that parses requests, delegates to a service, and writes responses.
- **Sorts**: A slice of strings in `"field:direction"` format (e.g., `"name:ASC"`) used to order query results.
- **PageParams**: A `pagination.Page` struct carrying `Limit`, `Page`, and `Offset` for paginated queries.

---

## Requirements

### Requirement 1: Maintenance Category — Create

**User Story:** As an API consumer, I want to create a maintenance category, so that I can classify maintenance entries under a named category.

#### Acceptance Criteria

1. WHEN a valid `CreatePayload` (containing `Name` and optional `Description`) is provided, THE MaintenanceCategoryService SHALL persist a new `MaintenanceCategory` record via MaintenanceCategoryRepository.
2. THE MaintenanceCategoryRepository SHALL use an upsert (ON CONFLICT on `id`) when saving a `MaintenanceCategory`.
3. IF `Name` is empty in the `CreatePayload`, THEN THE MaintenanceCategoryService SHALL return a validation error.
4. WHEN a `MaintenanceCategory` is successfully created, THE Controller SHALL return HTTP 201 with the created record's ID.

---

### Requirement 2: Maintenance Category — List

**User Story:** As an API consumer, I want to list maintenance categories with optional filtering and pagination, so that I can browse available categories.

#### Acceptance Criteria

1. WHEN a `ListPayload` is provided, THE MaintenanceCategoryService SHALL return a paginated list of `MaintenanceCategory` records.
2. WHERE `Name` is non-empty in `ListPayload`, THE MaintenanceCategoryRepository SHALL apply a case-insensitive partial match filter on the `name` column.
3. WHERE `Description` is non-empty in `ListPayload`, THE MaintenanceCategoryRepository SHALL apply a case-insensitive partial match filter on the `description` column.
4. WHERE `Sorts` is non-empty in `ListPayload`, THE MaintenanceCategoryRepository SHALL apply the sort order using `database.GenerateOrderQuery` with the valid sort map for the domain.
5. THE MaintenanceCategoryRepository SHALL include `COUNT(*) OVER()` as `affected_records` in list queries to support pagination metadata.
6. WHEN `PageParams` is provided in `ListPayload`, THE MaintenanceCategoryRepository SHALL apply `LIMIT` and `OFFSET` accordingly.

---

### Requirement 3: Maintenance Category — FindByID

**User Story:** As an API consumer, I want to retrieve a single maintenance category by its ID, so that I can view its details.

#### Acceptance Criteria

1. WHEN a valid category ID is provided, THE MaintenanceCategoryService SHALL return the matching `MaintenanceCategory` record.
2. IF no `MaintenanceCategory` exists for the given ID, THEN THE MaintenanceCategoryService SHALL return a `RecordNotFound` error.
3. WHEN the record is found, THE Controller SHALL return HTTP 200 with the category data.

---

### Requirement 4: Maintenance Category — Update

**User Story:** As an API consumer, I want to update a maintenance category, so that I can correct or refine its name and description.

#### Acceptance Criteria

1. WHEN a valid `UpdatePayload` (containing `ID` and at least one updatable field) is provided, THE MaintenanceCategoryService SHALL fetch the existing record, apply the changes, and persist via MaintenanceCategoryRepository.
2. IF no `MaintenanceCategory` exists for the given `ID` in `UpdatePayload`, THEN THE MaintenanceCategoryService SHALL return a `RecordNotFound` error.
3. WHERE `Name` is non-nil in `UpdatePayload`, THE MaintenanceCategoryService SHALL update the `Name` field on the model.
4. WHERE `Description` is non-nil in `UpdatePayload`, THE MaintenanceCategoryService SHALL update the `Description` field on the model.
5. WHEN the update is successful, THE Controller SHALL return HTTP 200.

---

### Requirement 5: Maintenance Category — Delete

**User Story:** As an API consumer, I want to delete a maintenance category, so that I can remove obsolete categories.

#### Acceptance Criteria

1. WHEN a valid category ID is provided, THE MaintenanceCategoryService SHALL delegate deletion to MaintenanceCategoryRepository.
2. THE MaintenanceCategoryRepository SHALL perform a soft delete (relying on GORM's `DeletedAt` field) when deleting a `MaintenanceCategory`.
3. WHEN deletion is successful, THE Controller SHALL return HTTP 200.

---

### Requirement 6: Maintenance Entry — Create

**User Story:** As an API consumer, I want to create a maintenance entry for a car, so that I can record a maintenance event at a specific odometer reading.

#### Acceptance Criteria

1. WHEN a valid `CreatePayload` is provided, THE MaintenanceEntryService SHALL check whether an `OdometerEntry` with the given `OdometerReading` already exists by calling `OdometerEntryRepository.FindByReading`.
2. IF an `OdometerEntry` with the given `OdometerReading` already exists, THEN THE MaintenanceEntryService SHALL use the existing `odometer_entry.id` as `maintenance_entry.odometer_entry_id`.
3. IF no `OdometerEntry` with the given `OdometerReading` exists, THEN THE MaintenanceEntryService SHALL create a new `OdometerEntry` and use its newly generated ID as `maintenance_entry.odometer_entry_id`.
4. WHEN the `OdometerEntry` ID is resolved, THE MaintenanceEntryService SHALL persist the new `MaintenanceEntry` via MaintenanceEntryRepository.
5. THE MaintenanceEntryRepository SHALL use an upsert (ON CONFLICT on `id`) when saving a `MaintenanceEntry`.
6. IF the `CategoryID` in `CreatePayload` does not correspond to an existing `MaintenanceCategory`, THEN THE MaintenanceEntryService SHALL return a `RecordNotFound` error.
7. WHEN a `MaintenanceEntry` is successfully created, THE Controller SHALL return HTTP 201 with the created record's ID.

---

### Requirement 7: Maintenance Entry — List

**User Story:** As an API consumer, I want to list maintenance entries with optional filtering and pagination, so that I can review maintenance history.

#### Acceptance Criteria

1. WHEN a `ListPayload` is provided, THE MaintenanceEntryService SHALL return a paginated list of `MaintenanceEntry` records.
2. WHERE `CategoryIDs` is non-empty in `ListPayload`, THE MaintenanceEntryRepository SHALL apply an `IN` filter on the `category_id` column.
3. WHERE `Sorts` is non-empty in `ListPayload`, THE MaintenanceEntryRepository SHALL apply the sort order using `database.GenerateOrderQuery` with the valid sort map for the domain.
4. THE MaintenanceEntryRepository SHALL include `COUNT(*) OVER()` as `affected_records` in list queries to support pagination metadata.
5. WHEN `PageParams` is provided in `ListPayload`, THE MaintenanceEntryRepository SHALL apply `LIMIT` and `OFFSET` accordingly.

---

### Requirement 8: Maintenance Entry — FindByID

**User Story:** As an API consumer, I want to retrieve a single maintenance entry by its ID, so that I can view its full details.

#### Acceptance Criteria

1. WHEN a valid entry ID is provided, THE MaintenanceEntryService SHALL return the matching `MaintenanceEntry` record.
2. IF no `MaintenanceEntry` exists for the given ID, THEN THE MaintenanceEntryService SHALL return a `RecordNotFound` error.
3. WHEN the record is found, THE Controller SHALL return HTTP 200 with the entry data.

---

### Requirement 9: Maintenance Entry — Update

**User Story:** As an API consumer, I want to update a maintenance entry, so that I can correct details of a recorded maintenance event.

#### Acceptance Criteria

1. WHEN a valid `UpdatePayload` is provided, THE MaintenanceEntryService SHALL fetch the existing `MaintenanceEntry`, apply changes, and persist via MaintenanceEntryRepository.
2. IF no `MaintenanceEntry` exists for the given `ID` in `UpdatePayload`, THEN THE MaintenanceEntryService SHALL return a `RecordNotFound` error.
3. WHERE `OdometerReading` is non-nil in `UpdatePayload`, THE MaintenanceEntryService SHALL apply the same odometer upsert logic as in Requirement 6 (check `FindByReading`, create if absent, assign ID).
4. WHERE `CategoryID` is non-nil in `UpdatePayload`, THE MaintenanceEntryService SHALL verify the category exists before updating `category_id`.
5. WHEN the update is successful, THE Controller SHALL return HTTP 200.

---

### Requirement 10: Maintenance Entry — Delete

**User Story:** As an API consumer, I want to delete a maintenance entry, so that I can remove incorrectly recorded events.

#### Acceptance Criteria

1. WHEN a valid entry ID is provided, THE MaintenanceEntryService SHALL delegate deletion to MaintenanceEntryRepository.
2. THE MaintenanceEntryRepository SHALL perform a soft delete when deleting a `MaintenanceEntry`.
3. WHEN deletion is successful, THE Controller SHALL return HTTP 200.

---

### Requirement 11: HTTP Endpoints

**User Story:** As an API consumer, I want RESTful endpoints for maintenance categories and entries, so that I can integrate the feature into client applications.

#### Acceptance Criteria

1. THE Controller SHALL expose the following endpoints for maintenance categories:
   - `POST   /maintenance-categories` — create
   - `GET    /maintenance-categories` — list (with query params for filters, sorts, pagination)
   - `GET    /maintenance-categories/:id` — find by ID
   - `PUT    /maintenance-categories/:id` — update
   - `DELETE /maintenance-categories/:id` — delete
2. THE Controller SHALL expose the following endpoints for maintenance entries:
   - `POST   /maintenance-entries` — create
   - `GET    /maintenance-entries` — list (with query params for filters, sorts, pagination)
   - `GET    /maintenance-entries/:id` — find by ID
   - `PUT    /maintenance-entries/:id` — update
   - `DELETE /maintenance-entries/:id` — delete
3. WHEN a request body fails validation, THE Controller SHALL return HTTP 400 with a descriptive error message.
4. WHEN a service returns a `RecordNotFound` error, THE Controller SHALL return HTTP 404.
5. WHEN an unexpected error occurs, THE Controller SHALL return HTTP 500.
