# Setup Instructions
## Docker

### step 1
Copy contents of `example.env` to `.env` and update values appropriately.

### step 2
Don't forget to update environemnt variables in postgres service in `docker-compose.yml`
```
    POSTGRES_USER: user
    POSTGRES_PASSWORD: password
    POSTGRES_DB: dbname
```

### Step 3
Have docker running and run `docker-compose up postgres`

### Step 4
Run `docker-compose up server`. This will also migrate the db on initial load.

#### NOTES
Authentication flow heavily based on https://codevoweb.com/golang-gorm-fiber-jwt-authentication/
