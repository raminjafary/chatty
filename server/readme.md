
### Download postgres image
docker pull postgres:15-alpine

### install golang-migrate cli
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

### craete users migration
migrate create -ext sql -dir db/migrations add_users_table

### craete rooms migration
migrate create -ext sql -dir db/migrations add_rooms_table



