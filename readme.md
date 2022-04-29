# Microservice providing Whois information
This repo contains Docker image which installs Whois client on Alpine Linux distribution and Golang app.
App is a small webserver written in Golang with only one endpoint:

```GET /whois/plain/{domain}```

it provides whois information as a **plain text** about a given domain using whois client https://github.com/rfc1036/whois.

## Get started
1. Download repo.
2. Run docker.
3. Use app.

## How to run docker container
```docker build . -t whois-go```

```docker run -it -d -p 9091:9091 --name=whois-go whois-go```

### How to use
Send request to endpoint like this:

```curl 0.0.0.0:9091/whois/plain/facebook.com```

port - optional GET param to let know Whois client to use it instead of default 43 port

```curl 0.0.0.0:9091/whois/plain/google.ch?port=4343```