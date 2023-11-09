FROM golang:1.18-bullseye AS builder

COPY . /dineflow-menu-services
WORKDIR /dineflow-menu-services
RUN go mod tidy
RUN go build

# Create a new stage for the final image
FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y ca-certificates wget
RUN wget -O /usr/local/bin/wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh

COPY --from=builder /dineflow-menu-services/dineflow-menu-services /app/dineflow-menu-services

EXPOSE 8090

CMD ["wait-for-it.sh", "menu_db:3306", "--", "/app/dineflow-menu-services"]