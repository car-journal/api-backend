dev:
	nodemon --exec go run cmd/main.go --signal SIGTERM

build:
	go build -o car-journal cmd/main.go

goose-create: # name=
	goose create $(name) sql

create-seeder: # name=
	touch database/seeder/${name}.sql
	echo "INSERT INTO table_name (column1, column2, column3) VALUES\n(value1, value2, value3)" > database/seeder/$(name).sql
	echo "ON CONFLICT (id) DO NOTHING;" >> database/seeder/$(name).sql

run-seeder: # user= name= 
	psql -U ${user} -d car_journal -f database/seeder/${name}.sql

run-all-seeders: # user= 
	for file in database/seeder/*.sql; do \
		psql -U ${user} -d car_journal -f $$file; \
	done
