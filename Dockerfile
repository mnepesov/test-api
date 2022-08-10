FROM golang:1.18-buster AS dependencies

ENV NAME "astronomy"

WORKDIR /app/${NAME}

# Prepare go mod dependencies
COPY go.* ./
RUN go mod download && go mod verify

FROM dependencies AS build

ENV NAME "astronomy"

WORKDIR /app/${NAME}

# Copy application and build it
COPY . .

RUN ["make", "build"]

# final image with binaries
FROM alpine:latest
ARG NAME
ARG LOG_LEVEL
ARG ENV
ENV NAME "astronomy"
ENV LOG_LEVEL ${LOG_LEVEL}
ENV ENV ${ENV}

WORKDIR /app/${NAME}

RUN apk --no-cache add ca-certificates

COPY --from=build /app/${NAME}/*.env ./
COPY --from=build /app/${NAME}/bin/${NAME} ./${NAME}

CMD ./${NAME}  -ll ${LOG_LEVEL} -e ${ENV}
