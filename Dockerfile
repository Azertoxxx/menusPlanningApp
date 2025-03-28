FROM golang:alpine

WORKDIR /menusplanningapp
COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go build -o ./bin/api ./cmd/api 

ENV PATH="/go/bin:${PATH}"
RUN chmod +x /menusplanningapp/entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]
CMD ["/menusplanningapp/bin/api"]

EXPOSE 8080

# psql -d menusplanningapp_db -U menusplanningapp_user

# FROM golang:alpine

# # Install necessary dependencies
# RUN apk add --no-cache git gcc g++ musl-dev

# # Install Delve debugger
# RUN go install github.com/go-delve/delve/cmd/dlv@latest

# WORKDIR /menusplanningapp
# COPY . .

# # Install dependencies
# RUN go mod download

# # Build the application with debugging flags
# RUN go build -gcflags="all=-N -l" -o ./bin/api ./cmd/api 

# ENV PATH="/go/bin:${PATH}"
# RUN chmod +x /menusplanningapp/entrypoint.sh

# # Start Delve instead of running the app directly
# CMD ["dlv", "debug", "--headless", "--listen=:40000", "--api-version=2", "--accept-multiclient", "./cmd/api"]

