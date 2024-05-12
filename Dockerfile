FROM golang:alpine

# Install system dependencies including 'make'
RUN apk update && apk add --no-cache gcc libc-dev make

# Copying app to docker and making it as working directory
RUN mkdir /app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. They will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY ./ /app

#Creating a non root user
RUN adduser -D user
USER user