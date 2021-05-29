FROM golang:1.12.0-alpine3.9

RUN mkdir /app

# ADD : project contents를 /app 아래로 copy
ADD . /app 

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]
