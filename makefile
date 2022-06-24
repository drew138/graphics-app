# https://dev.to/techschoolguru/how-to-write-run-database-migration-in-golang-5h6g


# https://adityarama1210.medium.com/simple-golang-api-uploader-using-google-cloud-storage-3d5e45df74a5



migrations:
	./scripts/makemigrations.sh

migrateup:
	./scripts/migrateup.sh "postgres://postgres:local-password@localhost:5432/graphics?sslmode=disable"

migratedown:
	./scripts/migratedown.sh "postgres://postgres:local-password@localhost:5432/graphics?sslmode=disable"
