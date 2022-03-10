FROM golang:rc-alpine

WORKDIR /app
COPY . .
RUN go build .

EXPOSE 3500

CMD ["./motosikal"]
