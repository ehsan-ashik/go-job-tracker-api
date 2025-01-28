sudo gzip -d --force ./db_backups/last/job-tracker-latest.sql.gz

docker exec -i db psql -U postgres job-tracker < ./db_backups/last/job-tracker-latest.sql