# Job Application Tracker - Backend API

A robust REST API service built with Go to manage job applications, resumes, and related resources. This backend system provides a comprehensive solution for tracking job applications with support for resume storage in Azure Blob Storage.

Frontend for the application is available in the following repo: https://github.com/ehsan-ashik/react-admin-job-tracker-frontend

## Features

### Core Functionality
- Complete REST API for job application management
- Resource management for:
  - Jobs
  - Job Descriptions
  - Job Categories
  - Companies
  - Resumes (with PDF storage)
- Advanced querying capabilities:
  - Sorting
  - Pagination
  - Filtering
- Flexible resource creation:
  - Create jobs with associated company and category information
  - Independent CRUD operations for all resources
- Azure Blob Storage integration for resume PDFs
- Multi-part file upload support for resumes

### Technical Features
- Built with Go Fiber for high performance
- GORM integration for database operations
- PostgreSQL database with automatic backups
- Dockerized deployment
- Structured logging
- CORS configuration

## Technical Stack

- **Language**: Go
- **Framework**: Go Fiber
- **ORM**: GORM
- **Database**: PostgreSQL
- **Cloud Storage**: Azure Blob Storage
- **Containerization**: Docker & Docker Compose

## Getting Started

### Prerequisites
- Docker and Docker Compose
- Go 1.23
- PostgreSQL (for local development)
- Azure Storage Account

### Configuration

Create a `.env` file in the root directory:

```env
# Database Configuration
DB_NAME=
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=

# Azure Storage Configuration
AZURE_ACCESS_KEY=
AZURE_ACCOUNT_NAME=
AZURE_CONTAINER_NAME=
```

### Running the Application

#### Using Docker Compose (Recommended)

1. Build and start the services:
   ```bash
   docker-compose up --build
   ```

This will start:
- The API service
- PostgreSQL database
- Database backup service

#### Local Development

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Run the application:
   ```bash
   go run cmd/main.go
   ```

## API Documentation

### Available Endpoints

#### Jobs
- `GET /api/jobs` - List all jobs (supports pagination, sorting, filtering)
- `POST /api/jobs` - Create a new job
- `GET /api/jobs/:id` - Get job details
- `PUT /api/jobs/:id` - Update job
- `DELETE /api/jobs/:id` - Delete job

#### Companies
- `GET /api/companies` - List all companies (supports pagination, sorting, filtering)
- `POST /api/companies` - Create a new company
- `GET /api/companies/:id` - Get company details
- `PUT /api/companies/:id` - Update company
- `DELETE /api/companies/:id` - Delete company

#### Resumes
- `GET /api/resumes` - List all resumes (supports pagination, sorting, filtering)
- `POST /api/resumes` - Create a new resume (multipart form)
- `GET /api/resumes/:id` - Get resume details
- `PUT /api/resumes/:id` - Update resume (multipart form)
- `DELETE /api/resumes/:id` - Delete resume

#### Job Categories
- `GET /api/categories` - List all categories (multipart form)
- `POST /api/categories` - Create a new category
- `GET /api/categories/:id` - Get category details
- `PUT /api/categories/:id` - Update category
- `DELETE /api/categories/:id` - Delete category

#### Job Descriptions
- `GET /api/categories` - List all descriptions (multipart form)
- `POST /api/categories` - Create a new job description
- `GET /api/categories/:id` - Get job description details
- `PUT /api/categories/:id` - Update job description
- `DELETE /api/categories/:id` - Delete job description

## License
This project is licensed under the MIT License - see the LICENSE file for details.
