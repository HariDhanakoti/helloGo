# helloGo

# Run in local dev mode

go run main.go

# Build and run the helloGo application

docker build -t gowebapp .

docker run -it --rm -p 8080:8080 --name web gowebapp