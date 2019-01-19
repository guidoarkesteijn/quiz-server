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
COPY . $GOPATH/src/github.com/project-quiz/quiz-server/
WORKDIR $GOPATH/src/github.com/project-quiz/quiz-server/
# Fetch dependencies.
# Using go get.
RUN go get
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

#Environment Variables
ENV DATABASE_IP "127.0.0.1"
ENV DATABASE_PORT "4600"
ENV DATABASE_USERNAME "root"
ENV DATABASE_PASSWORD "password"

# Run the quiz server binary.
ENTRYPOINT ["/go/bin/quiz-server"]