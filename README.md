# cloudflare-ddns

This small application will keep a Cloudflare DNS record up to date with the Docker public IP address.

## Usage

### Basic
```
CF_HOST=cloudgenius.co CF_API_KEY=xxxxx -e CF_API_EMAIL=nilesh@cloudgeni.us go run cmd/cloudflare-ddns/main.go
```

## Docker

```
docker run -e CF_HOST=cloudgenius.co -e CF_API_KEY=926e93ebd56638754af6ab32b39fff42e7e89 -e CF_API_EMAIL=nilesh@cloudgeni.us cloudgenius/cloudflare-ddns:latest
```

## Compose v3
```
version: '3'
services:
  dns:
    image: cloudgenius/cloudflare-ddns:latest
    environment:
      - CF_HOST=${DOMAIN}
      - CF_API_KEY=${KEY}
      - CF_API_EMAIL=${MAIL}
    deploy:
      replicas: 1
    healthcheck:
      test: ["CMD", "curl", "-f", "http://${DOMAIN}"]
      interval: 20s
      timeout: 10s
      retries: 3
```


docker build -t cloudgenius/cloudflare-ddns .  
docker push cloudgenius/cloudflare-ddns
docker stack deploy -c ddns.yml ddns
