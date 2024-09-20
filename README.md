## Initialization
```bash
go mod init fx-golang-server
```

## Docker
- Remove containers not in the file docker-compose.dev.yml
  ```
  docker-compose -f docker-compose.dev.yml up --build -d --remove-orphans 
  ```
- Run kafka containers
  ```
  docker-compose -f docker-compose.kafka.yml up --build -d
  ```
- Run elastic containers
  ```
  docker-compose -f docker-compose.elastic.yml up --build -d
  ```

### Elastic search
- [Dashboard](http://localhost:5601/)
```
{
  "query": {
    "match_all": {}
  }
}
{
  "query": {
    "wildcard": {
      "phone": "*191954*"
    }
  }
}
{
    "query": {
        "bool": {
            "should": [
                {
                    "wildcard": {
                        "phone": "*191*"
                    }
                },
                {
                    "wildcard": {
                        "email": "*191*"
                    }
                }
            ]
        }
    }
}
```

## Swagger
```bash
# Install
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

# Add swag to PATH
export PATH=$PATH:$(go env GOPATH)/bin
source ~/.zshrc
# or
source ~/.bash_profile

# use swag
swag init -g cmd/api/main.go --output docs
```

## Database
```sql
SELECT * FROM pg_stat_activity;
SELECT * FROM pg_stat_database;

SHOW shared_buffers;
SHOW work_mem;

SELECT
  datname,
  usename,
  count(*) AS connections
FROM
  pg_stat_activity
GROUP BY
  datname, usename;
 
SELECT count(*) FROM pg_stat_activity;
SHOW max_connections;

SHOW autovacuum;
```

## Containerization

```bash
# Build an image on local
docker build --build-arg TELEGRAM_TOKEN=$(grep TELEGRAM_TOKEN .env | cut -d '=' -f2) \
             --build-arg TELEGRAM_CHAT_ID=$(grep TELEGRAM_CHAT_ID .env | cut -d '=' -f2) \
            -t fx-golang-server .

# Run container
docker run -d -p 8080:8080 --name fx-golang-server-container fx-golang-server

# Start container
docker start fx-golang-server-container
```

## Deployment

```bash
gcloud builds submit --config cloudbuild.yaml --project $PROJECT_ID
```
