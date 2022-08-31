FROM golang:1.15.2-alpine


# RUN mkdir /app/backend

WORKDIR /app/back-api


CMD ["go", "run", "main.go"]