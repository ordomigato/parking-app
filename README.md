# Setup Instructions
## Docker

### step 1
Copy contents of `example.env` to `.env` and update values appropriately. Don't forget to update postgres service in `docker-compose.yml`

### Step 2
Have docker running and run `docker-compose up postgres`

### Step 3
Run `docker-compose up server`. This will also migrate the db on initial load.

#### Migration Notes
Basic command to create new migration table (in `./server` director):
`migrate create -ext sql -dir store/migrations/ -seq table_name`
