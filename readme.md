# Microservice to get Whois information
App provides  whois information about given domain.
It's small Golang webserver with only one endpoint:

```GET /whois/plain/{domain}```

Endpoint uses whois package https://github.com/rfc1036/whois 

##Get started
1. Download repo.
2. Run docker.
3. Use app.

##How to run docker container
```docker build . -t whois-go```

```docker run -it -d -p 9091:9091 --name=whois-go whois-go```

###How to use
Send request to

```curl 0.0.0.0:9091/whois/plain/facebook.com```