FROM golang:1.16-alpine

WORKDIR /weact-backend

COPY . .

ENV GOPROXY="https://goproxy.cn,direct"

RUN go mod download

RUN go build -o main.go

EXPOSE 8000

CMD [ "./main" ]