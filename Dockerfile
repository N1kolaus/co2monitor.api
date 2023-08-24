# Verwende das offizielle Go-Image als Basis
FROM golang:1.21

# Setze das Arbeitsverzeichnis innerhalb des Containers
WORKDIR /app

# Kopiere die Go-Anwendungsquelle in das Arbeitsverzeichnis
COPY . .

# Kompiliere die Go-Anwendung
RUN go build -o main .

# Exponiere den Port, auf dem die Anwendung lauschen wird
EXPOSE 8080

# Starte die Go-Anwendung
CMD ["./main"]
