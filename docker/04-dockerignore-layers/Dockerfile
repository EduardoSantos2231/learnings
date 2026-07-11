FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
  curl && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY ["./index.html", "."] 

CMD [ "cat", "index.html" ]
