# Go Hello World Project

This project is a simple Go application that serves a "Hola Mundo" message over HTTP. It is dockerized and ready to be run in a container.

## Project Structure

```
go-hello-world
├── src
│   ├── main.go
├── Dockerfile
├── go.mod
└── README.md
```

## Getting Started

### Prerequisites

- Docker installed on your machine.

### Building the Docker Image

To build the Docker image for the application, navigate to the project directory and run the following command:

```
docker build -t go-hello-world .
```

### Running the Docker Container

After building the image, you can run the Docker container with the following command:

```
docker run -p 8080:8080 go-hello-world
```

This command maps port 8080 of the container to port 8080 on your host machine.

### Accessing the Application

Once the container is running, you can access the "Hola Mundo" endpoint by navigating to the following URL in your web browser or using a tool like curl:

```
http://localhost:8080
```

You should see the message "Hola Mundo" displayed in your browser.

## License

This project is licensed under the MIT License.