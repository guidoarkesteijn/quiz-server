############################
# STEP 1 build executable binary
############################
FROM golang:alpine as builder
# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates
# Create appuser
RUN adduser -D -g '' appuser
COPY . $GOPATH/src/guidoarkesteijn/quiz-server/
WORKDIR $GOPATH/src/guidoarkesteijn/quiz-server/
# Fetch dependencies.
# Using go get.
# TODO find a way to get all my dependencies in one go command.
RUN go get github.com/twinj/uuid 
RUN go get github.com/golang/protobuf/proto
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/project-quiz/quiz-go-model
# Using go mod.
# RUN go mod download
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/quiz-server
############################
# STEP 2 build a small image
############################
FROM scratch
# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builder /go/bin/quiz-server /go/bin/quiz-server
# Use an unprivileged user.
USER appuser
# Run the quiz server binary.
ENTRYPOINT ["/go/bin/quiz-server"]