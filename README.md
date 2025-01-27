# Learning Platform API

This project is a collaborative effort by three contributors to create a learning platform API using Golang.

## Getting Started

### Prerequisites

Ensure you have the following installed on your machine:
- [Go](https://golang.org/dl/) (version 1.16 or later)
- [Git](https://git-scm.com/)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/OucheneMohamedNourElIslem658/learn_oo.git
    cd learning-platform-api
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

### Running the Server

To start the Go server, run the following command:
```sh
go run main.go
```

The server should now be running on `http://localhost:8080`.

### Prerequisites

Ensure you have the following installed on your machine:
- [Docker](https://www.docker.com/get-started)

### Running the Server Using Docker

You can run the application using Docker. Follow these steps:

1. Pull the Docker image:
    ```sh
    docker pull fethi279/learn_oo:latest
    ```

2. Run the container:
    ```sh
    docker run -d -p 8080:8000 fethi279/learn_oo:latest
    ```

The server should now be running on `http://localhost:8080`.

### Contributors

- Seffih Fadi: [GitHub Profile](https://github.com/seffihfadi)

- Fethi Boukourou: [GitHub Profile](https://github.com/bkrfethi)
- Alaa Eddine Saouchi: [GitHub Profile](https://github.com/alaasao)

You can find the web app ui code in this repository link [here](https://github.com/seffihfadi/learnoo).

### Acknowledgments

Special thanks to all contributors and the open-source community.

### Features

This learning platform includes the following features:
- Quizzes to test knowledge and understanding
- Certification upon successful completion of courses
- Subscription plans for accessing premium content
- User-friendly interface for easy navigation
- Progress tracking to monitor learning journey
- Interactive content to enhance learning experience
- Regular updates with new courses and materials
- Community support and discussion forums
- Personalized learning paths based on user preferences
- Secure and scalable architecture for reliable performance
- Integration with third-party tools and services
- Analytics and reporting for insights into learning progress
- Multi-language support for a global audience
