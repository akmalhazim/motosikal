FROM golang:alpine
RUN apk add git

WORKDIR /app
COPY . .
RUN go build .

EXPOSE 3500

CMD ["./motosikal"]
