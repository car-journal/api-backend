# Implementation Plan: Maintenance Management

## Overview

Implement two new domains — `maintenancecategory` and `maintenanceentry` — following the existing layered architecture used by `fuelentry` and `odometerentry`. Each domain gets a model, payload types, const map, repository, service, HTTP controllers, route registration, and container wiring. The `maintenanceentry` service reuses the odometer upsert pattern from `fuelentryservice`.

## Tasks

- [x] 1. Add `MaintenanceEntry` model and finalize `MaintenanceCategory` model
  - [x] 1.1 Create `internal/model/maintenanceentry_model.go` with the `MaintenanceEntry` struct
    - Define `MaintenanceEntryTableName = "maintenance_entries"` constant
    - Embed `BaseModel`; add fields: `CarID`, `OdometerEntryID`, `CategoryID` (all `uuid.UUID`), `Brand`, `Name` (string), `Price` (float64), `PerformedAt` (time.Time), `Notes` (*string)
    - _Requirements: 6.4_
  - [x] 1.2 Update `internal/model/maintenance_category.go` to add `MaintenanceCategoryTableName` constant
    - Add `const MaintenanceCategoryTableName = "maintenance_categories"` to match the pattern of other models
    - _Requirements: 1.1_

- [x] 2. Implement `maintenancecategory` domain — const, payload, repository
  - [x] 2.1 Create `internal/domain/maintenancecategory/const/maintenancecategory_const.go`
    - Define `ValidSorts` map: `"name"`, `"created_at"`, `"updated_at"`
    - _Requirements: 2.4_
  - [x] 2.2 Create `internal/domain/maintenancecategory/payload/maintenancecategory_payload.go`
    - Define `CreatePayload` with `Name string` (validate:"required") and `Description *string`
    - Define `ListPayload` with `Name string`, `Description string`, `Sorts []string`, `PageParams pagination.Page`
    - Define `UpdatePayload` with `ID string`, `Name *string`, `Description nullable.String`
    - _Requirements: 1.1, 1.3, 2.1, 2.2, 2.3, 4.1, 4.3, 4.4_
  - [x] 2.3 Create `internal/domain/maintenancecategory/repository/maintenancecategory_repository.go`
    - Define `Interface` with `Save`, `List`, `FindByID`, `Delete`
    - `Save`: upsert on conflict `id`, updating `name`, `description`, `updated_at`
    - `List`: select with `COUNT(*) OVER() AS affected_records`, apply `ILike` for `Name`/`Description` filters, `GenerateOrderQuery` for sorts, `GeneratePaginationQuery` for page params
    - `FindByID`: query by `id`, return nil (not error) on `gorm.ErrRecordNotFound`
    - `Delete`: soft delete via GORM `Delete`
    - _Requirements: 1.2, 2.2, 2.3, 2.4, 2.5, 2.6, 5.2_
  - [ ]* 2.4 Write property test for `MaintenanceCategoryRepository.Save` — upsert idempotence
    - **Property 2: Category upsert idempotence**
    - **Validates: Requirements 1.2**
  - [ ]* 2.5 Write property test for `MaintenanceCategoryRepository.List` — pagination invariant
    - **Property 4: Category list pagination invariant**
    - **Validates: Requirements 2.1, 2.5, 2.6**
  - [ ]* 2.6 Write property test for `MaintenanceCategoryRepository.List` — name filter correctness
    - **Property 5: Category name filter correctness**
    - **Validates: Requirements 2.2**
  - [ ]* 2.7 Write property test for `MaintenanceCategoryRepository.List` — description filter correctness
    - **Property 6: Category description filter correctness**
    - **Validates: Requirements 2.3**
  - [ ]* 2.8 Write property test for `MaintenanceCategoryRepository.Delete` — soft delete
    - **Property 9: Category soft delete**
    - **Validates: Requirements 5.1, 5.2**

- [x] 3. Implement `maintenancecategory` service
  - [x] 3.1 Create `internal/domain/maintenancecategory/service/maintenancecategory_service.go`
    - Define `Interface` with `Create`, `List`, `FindByID`, `Update`, `Delete`
    - `Create`: validate `Name` non-empty (return `errortype.InvalidInput` if empty), generate UUID, call `repository.Save`
    - `List`: delegate to `repository.List`
    - `FindByID`: call `repository.FindByID`; if nil, return `httperror.New(errortype.RecordNotFound, ...)`
    - `Update`: call `FindByID` first (propagate not-found), apply non-nil fields from `UpdatePayload`, call `repository.Save`
    - `Delete`: delegate to `repository.Delete`
    - _Requirements: 1.1, 1.3, 2.1, 3.1, 3.2, 4.1, 4.2, 4.3, 4.4, 5.1_
  - [ ]* 3.2 Write property test for `MaintenanceCategoryService.Create` — create–read round trip
    - **Property 1: Category create–read round trip**
    - **Validates: Requirements 1.1, 3.1**
  - [ ]* 3.3 Write property test for `MaintenanceCategoryService.Create` — empty/whitespace name rejected
    - **Property 3: Empty or whitespace name is rejected**
    - **Validates: Requirements 1.3**
  - [ ]* 3.4 Write property test for `MaintenanceCategoryService.FindByID` — not-found for unknown ID
    - **Property 7: Category not-found for unknown ID**
    - **Validates: Requirements 3.2, 4.2**
  - [ ]* 3.5 Write property test for `MaintenanceCategoryService.Update` — update round trip
    - **Property 8: Category update round trip**
    - **Validates: Requirements 4.1, 4.3, 4.4**

- [x] 4. Checkpoint — maintenance category domain
  - Ensure all tests pass, ask the user if questions arise.

- [x] 5. Implement `maintenancecategory` HTTP controllers and routes
  - [x] 5.1 Create `transport/http/controller/internals/maintenancecategory/create.go`
    - Parse body into `maintenancecategorypayload.CreatePayload` via `bodyparser.Parse`
    - Wrap in `database.Run`; call `maintenancecategoryservice.Create`
    - Respond with `parser.JSON`; return HTTP 201 on success
    - _Requirements: 1.4, 11.1, 11.3, 11.4, 11.5_
  - [x] 5.2 Create `transport/http/controller/internals/maintenancecategory/list.go`
    - Parse query params into `maintenancecategorypayload.ListPayload` (name, description, sorts, page)
    - Wrap in `database.Run`; call `maintenancecategoryservice.List`; build `filter.PageResponse`
    - _Requirements: 2.1, 11.1_
  - [x] 5.3 Create `transport/http/controller/internals/maintenancecategory/find_by_id.go`
    - Extract `:id` from path via `mux.Vars`; call `maintenancecategoryservice.FindByID`
    - Return HTTP 200 with category data or propagate error
    - _Requirements: 3.3, 11.1, 11.4_
  - [x] 5.4 Create `transport/http/controller/internals/maintenancecategory/update.go`
    - Parse body into `maintenancecategorypayload.UpdatePayload`; extract `:id` from path
    - Wrap in `database.Run`; set `body.ID`; call `maintenancecategoryservice.Update`
    - _Requirements: 4.5, 11.1_
  - [x] 5.5 Create `transport/http/controller/internals/maintenancecategory/delete.go`
    - Extract `:id` from path; wrap in `database.Run`; call `maintenancecategoryservice.Delete`
    - _Requirements: 5.3, 11.1_
  - [x] 5.6 Create `transport/http/route/internals/internal_maintenancecategory_routes.go`
    - Register `POST /maintenance-categories`, `GET /maintenance-categories`, `GET /maintenance-categories/{id}`, `PUT /maintenance-categories/{id}`, `DELETE /maintenance-categories/{id}`
    - Use `routeconst.MaintenanceCategoryPrefix`
    - _Requirements: 11.1_
  - [x] 5.7 Register `MaintenanceCategory(internalRoutes, app)` in `transport/http/route/internals/internal_routes.go`
    - _Requirements: 11.1_

- [x] 6. Implement `maintenanceentry` domain — const, payload, repository
  - [x] 6.1 Create `internal/domain/maintenanceentry/const/maintenanceentry_const.go`
    - Define `ValidSorts` map: `"brand"`, `"name"`, `"price"`, `"performed_at"`, `"created_at"`, `"updated_at"`
    - _Requirements: 7.3_
  - [x] 6.2 Create `internal/domain/maintenanceentry/payload/maintenanceentry_payload.go`
    - Define `CreatePayload` with `CarID`, `CategoryID`, `Brand`, `Name`, `ReadingUnit` (string, required), `OdometerReading`, `Price` (float64, required), `PerformedAt` (time.Time, required), `Notes` (*string)
    - Define `ListPayload` with `CategoryIDs []string`, `Sorts []string`, `PageParams pagination.Page`
    - Define `UpdatePayload` with `ID string`, `CategoryID *string`, `OdometerReading *float64`, `ReadingUnit *string` (validate:"required_with=OdometerReading"), `Brand *string`, `Name *string`, `Price *float64`, `PerformedAt nullable.Time`, `Notes nullable.String`
    - _Requirements: 6.1, 7.1, 7.2, 9.1, 9.3, 9.4_
  - [x] 6.3 Create `internal/domain/maintenanceentry/repository/maintenanceentry_repository.go`
    - Define `Interface` with `Save`, `List`, `FindByID`, `Delete`
    - `Save`: upsert on conflict `id`, updating `car_id`, `odometer_entry_id`, `category_id`, `brand`, `name`, `price`, `performed_at`, `notes`, `updated_at`
    - `List`: select with `COUNT(*) OVER() AS affected_records`, apply `IN` filter for `CategoryIDs`, `GenerateOrderQuery` for sorts, `GeneratePaginationQuery` for page params
    - `FindByID`: query by `id`, return nil on `gorm.ErrRecordNotFound`
    - `Delete`: soft delete via GORM `Delete`
    - _Requirements: 6.5, 7.2, 7.3, 7.4, 7.5, 10.2_
  - [ ]* 6.4 Write property test for `MaintenanceEntryRepository.List` — pagination invariant
    - **Property 14: Entry list pagination invariant**
    - **Validates: Requirements 7.1, 7.4, 7.5**
  - [ ]* 6.5 Write property test for `MaintenanceEntryRepository.List` — category filter correctness
    - **Property 15: Entry category filter correctness**
    - **Validates: Requirements 7.2**
  - [ ]* 6.6 Write property test for `MaintenanceEntryRepository.Delete` — soft delete
    - **Property 18: Entry soft delete**
    - **Validates: Requirements 10.1, 10.2**

- [x] 7. Implement `maintenanceentry` service
  - [x] 7.1 Create `internal/domain/maintenanceentry/service/maintenanceentry_service.go`
    - Define `Interface` with `Create`, `List`, `FindByID`, `Update`, `Delete`
    - Service struct depends on: `maintenancecategoryrepository.Interface`, `maintenanceentryrepository.Interface`, `odometerentryrepository.Interface`, `uuid.UUIDInterface`
    - `Create`:
      1. Verify `CategoryID` exists via `maintenancecategoryrepository.FindByID`; return `RecordNotFound` if nil
      2. Call `odometerentryrepository.FindByReading` with `OdometerReading`
      3. If found, reuse existing `OdometerEntry.ID`; if not, create new `OdometerEntry` via `odometerentryrepository.Save`
      4. Generate new entry UUID; call `maintenanceentryrepository.Save`
    - `List`: delegate to `maintenanceentryrepository.List`
    - `FindByID`: call `maintenanceentryrepository.FindByID`; if nil, return `httperror.New(errortype.RecordNotFound, ...)`
    - `Update`: fetch existing entry (propagate not-found); apply odometer upsert if `OdometerReading` non-nil; verify category if `CategoryID` non-nil; apply remaining non-nil fields; call `maintenanceentryrepository.Save`
    - `Delete`: delegate to `maintenanceentryrepository.Delete`
    - _Requirements: 6.1, 6.2, 6.3, 6.4, 6.6, 7.1, 8.1, 8.2, 9.1, 9.2, 9.3, 9.4, 10.1_
  - [ ]* 7.2 Write property test for `MaintenanceEntryService` — odometer upsert reuse existing reading
    - **Property 10: Odometer upsert — reuse existing reading**
    - **Validates: Requirements 6.1, 6.2, 9.3**
  - [ ]* 7.3 Write property test for `MaintenanceEntryService` — odometer upsert create new reading
    - **Property 11: Odometer upsert — create new reading**
    - **Validates: Requirements 6.1, 6.3, 9.3**
  - [ ]* 7.4 Write property test for `MaintenanceEntryService.Create` — create–read round trip
    - **Property 12: Entry create–read round trip**
    - **Validates: Requirements 6.4, 8.1**
  - [ ]* 7.5 Write property test for `MaintenanceEntryService.Create` — invalid category ID rejected
    - **Property 13: Invalid category ID is rejected**
    - **Validates: Requirements 6.6, 9.4**
  - [ ]* 7.6 Write property test for `MaintenanceEntryService.FindByID` — not-found for unknown ID
    - **Property 16: Entry not-found for unknown ID**
    - **Validates: Requirements 8.2, 9.2**
  - [ ]* 7.7 Write property test for `MaintenanceEntryService.Update` — update round trip
    - **Property 17: Entry update round trip**
    - **Validates: Requirements 9.1**

- [x] 8. Checkpoint — maintenance entry domain
  - Ensure all tests pass, ask the user if questions arise.

- [x] 9. Implement `maintenanceentry` HTTP controllers and routes
  - [x] 9.1 Create `transport/http/controller/internals/maintenanceentry/create.go`
    - Parse body into `maintenanceentrypayload.CreatePayload` via `bodyparser.Parse`
    - Wrap in `database.Run`; call `maintenanceentryservice.Create`
    - Return HTTP 201 on success
    - _Requirements: 6.7, 11.2, 11.3, 11.4, 11.5_
  - [x] 9.2 Create `transport/http/controller/internals/maintenanceentry/list.go`
    - Parse query params into `maintenanceentrypayload.ListPayload` (category_ids, sorts, page)
    - Wrap in `database.Run`; call `maintenanceentryservice.List`; build `filter.PageResponse`
    - _Requirements: 7.1, 11.2_
  - [x] 9.3 Create `transport/http/controller/internals/maintenanceentry/find_by_id.go`
    - Extract `:id` from path; call `maintenanceentryservice.FindByID`
    - Return HTTP 200 with entry data or propagate error
    - _Requirements: 8.3, 11.2, 11.4_
  - [x] 9.4 Create `transport/http/controller/internals/maintenanceentry/update.go`
    - Parse body into `maintenanceentrypayload.UpdatePayload`; extract `:id` from path
    - Wrap in `database.Run`; set `body.ID`; call `maintenanceentryservice.Update`
    - _Requirements: 9.5, 11.2_
  - [x] 9.5 Create `transport/http/controller/internals/maintenanceentry/delete.go`
    - Extract `:id` from path; wrap in `database.Run`; call `maintenanceentryservice.Delete`
    - _Requirements: 10.3, 11.2_
  - [x] 9.6 Create `transport/http/route/internals/internal_maintenanceentry_routes.go`
    - Register `POST /maintenance-entries`, `GET /maintenance-entries`, `GET /maintenance-entries/{id}`, `PUT /maintenance-entries/{id}`, `DELETE /maintenance-entries/{id}`
    - Use `routeconst.MaintenanceEntryPrefix`
    - _Requirements: 11.2_
  - [x] 9.7 Register `MaintenanceEntry(internalRoutes, app)` in `transport/http/route/internals/internal_routes.go`
    - _Requirements: 11.2_

- [x] 10. Wire repositories and services into the container
  - [x] 10.1 Update `transport/container/repository.go`
    - Add `MaintenanceCategory maintenancecategoryrepository.Interface` and `MaintenanceEntry maintenanceentryrepository.Interface` fields to `RepositoryContainer`
    - Instantiate both in `CreateRepositoryContainer`
    - _Requirements: 1.1, 6.4_
  - [x] 10.2 Update `transport/container/service.go`
    - Add `MaintenanceCategory maintenancecategoryservice.Interface` and `MaintenanceEntry maintenanceentryservice.Interface` fields to `ServiceContainer`
    - Instantiate both in `CreateServiceContainer`, passing the correct repository dependencies
    - `MaintenanceEntryService` receives: `repoContainer.MaintenanceCategory`, `repoContainer.MaintenanceEntry`, `repoContainer.OdometerEntry`, `clientContainer.UUID`
    - _Requirements: 1.1, 6.1_

- [x] 11. Final checkpoint — build and full test pass
  - Ensure all tests pass, ask the user if questions arise.

## Notes

- Tasks marked with `*` are optional and can be skipped for a faster MVP
- Property-based tests use `pgregory.net/rapid`; each runs a minimum of 100 iterations
- Repository-level property tests require a real PostgreSQL test database; service-level tests mock repository interfaces
- Each property test file should include the annotation comment: `// Feature: maintenance-management, Property N: <property text>`
- The existing `internal/domain/maintenancecategory/` directory is an empty placeholder — populate it in place
- The existing `internal/domain/maintenanceentriy/` directory (typo) should be left alone; create the correctly-spelled `internal/domain/maintenanceentry/` as a new directory
