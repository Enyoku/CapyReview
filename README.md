# CapyReview

***CapyReview** is an **API Gateway** written in Go for a moddular microservice-based platform.
It acts as the single entry point for routing HTTP requests to various backend services.

---

## Architecture

- **API Gateway (`./`)** - Main entry point, forwards requests to microservices.
- **Auth Service (`auth/`)** - Handles registration, login, andd authentication tokens
- **Content Service (`content/`)** - Responsible for managing data related to **movies**, **serials**, **games**.
- **Review Service (`review/`)** - Responsible for managing data related to users **reviews**.

---

## Setup Instructions

### Requirements
- Go 1.20+
- Git
- (Optional) Docker & Docker Compose

### Clone the Project

```bash
git clone https://github.com/Enyoku/CapyReview.git
cd CapyReview
```

### Build Services

```bash
# Build Auth Service
cd auth
go build -o auth_service
cd ..

# Build Content Service
cd content
go build -o content_service
cd ..

# Build API Gateway
cd capyReview
go build -o capy_gateway
cd ..

# Build Review Service
cd review
go build -o review_service
cd ..
```

### Run Services (separatly)

```bash
# Terminal 1
cd auth
./auth_service

# Terminal 2
cd content
./content_service

# Terminal 3
cd capyReview
./capy_gateway

# Terminal 4
cd review
./review_service
```

---

## Authentication

Handled by the `auth` service.
On login, a JWT token is generated and stored in the client`s **HTTP-only cookies** for security.

You **do not** need to manually include the token in headers.

This setup protects the token form JavaScript access and ensures the session flows naturally through requests.

## License

## Roadmap

- [ ] Integrate Review Service (in progress)
- [ ] Docker Compose support
- [ ] OpenAPI/Swagger documentation
- [ ] HTTPS support

