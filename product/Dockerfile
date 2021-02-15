FROM golang:1.15-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /out/product .
RUN ls

FROM alpine
WORKDIR /app
COPY --from=build /out/product /app/
EXPOSE 8080
ENTRYPOINT ./product
