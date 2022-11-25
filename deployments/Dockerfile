FROM golang:1.18-alpine AS build

ARG GITLAB_USER
ARG GITLAB_PASSWORD

RUN apk update
RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN echo "https://${GITLAB_USER}:${GITLAB_PASSWORD}@gitlab.com" > ~/.git-credentials
RUN git config --global credential.helper store
RUN echo "machine gitlab.com login ${GITLAB_USER} password ${GITLAB_PASSWORD}" >> ~/.netrc

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o server ./cmd/server/main.go

FROM gcr.io/distroless/static AS final

WORKDIR /app

COPY --from=build /app/server .

EXPOSE 8080

ENTRYPOINT ["/app/server"]