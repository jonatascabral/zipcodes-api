FROM golang

COPY ./ /api

WORKDIR /api

ENTRYPOINT ["go", "run", "/api/main.go"]