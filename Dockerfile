FROM golang:alpine

WORKDIR /menusplanningapp
COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
# RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go build -o ./bin/api ./cmd/api 

ENV PATH="/go/bin:${PATH}"
RUN chmod +x /menusplanningapp/entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]
CMD ["/menusplanningapp/bin/api"]

EXPOSE 8080 40000

# psql -d menusplanningapp_db -U menusplanningapp_user

