FROM golang:1.18-bullseye AS builder

COPY . /dineflow-menu-services
WORKDIR /dineflow-menu-services
RUN go mod tidy
RUN go build

FROM golang:1.18-bullseye AS runner
ENV GIN_MODE release

RUN mkdir /app
WORKDIR /app
COPY --from=builder /dineflow-menu-services/dineflow-menu-services /app

EXPOSE 8090

CMD ["/app/dineflow-menu-services"]