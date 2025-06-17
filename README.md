# api-backend

## Create migration
### goose create ${table_name} ${file_extension}
Example
```
goose create users sql
goose create users go
```

## Create seeder
### make create-seeder name=${table_name}
Example
```
make create-seeder name=users
```

## Run Seeder
### make run-seeder name=${table_name}
Example
```
make run-seeder name=${users}
```
### make run-all-seeders