
FROM golang:1.16.5

WORKDIR /app

COPY . .

CMD [ "go", "run", "main.go" ]

EXPOSE 3000