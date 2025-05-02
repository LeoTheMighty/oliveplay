FROM rust:1.76 as builder

WORKDIR /usr/src/app
COPY . .
RUN cargo build --release --bin controller

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y libssl-dev ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /usr/src/app/target/release/controller /usr/local/bin/controller

CMD ["controller"] 