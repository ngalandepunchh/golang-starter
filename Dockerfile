FROM scratch

# add the dir to the container

ADD . /src/cadence-service

WORKDIR /src/cadence-service

RUN ls

# get dependencies

RUN dep ensure -v

# run the tests

RUN go test ./...

# Coverage - sonar


WORKDIR /src/cadence-service/cmd/cadence-service

# Build
RUN go build -o main .

# expose ports if needed
EXPOSE 8080

CMD ["main"]
