## Initialization
```bash
go mod init fx-golang-server
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
