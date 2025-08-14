FROM golang:1.24.5-alpine 

WORKDIR /app 

COPY go.mod go.sum ./

RUN apk add --no-cache git 

RUN go mod download 

COPY . .

RUN go build -o app 

EXPOSE 8081

CMD [ "./app" ]
