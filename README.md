# Location API

A Beego v2 REST API for managing locations and their image galleries, backed by PostgreSQL and MinIO (S3-compatible) object storage.

**Repo:** https://github.com/200215-Moynul-Islam/location-api.git

## Features

- List locations with their hero/base image
- List all gallery images for a location
- Stream a location's base image or any gallery image from object storage
- Set any gallery image as a location's new base/hero image (server-side copy in storage)

## Tech Stack

- **Framework:** Beego v2
- **Database:** PostgreSQL (via `beego/client/orm` + `lib/pq`)
- **Object Storage:** MinIO (S3-compatible, via AWS SDK v2)
- **Migrations:** [golang-migrate](https://github.com/golang-migrate/migrate)

## Project Structure

```
.
├── conf/                # app.conf (Beego + Postgres + MinIO config)
├── controllers/         # HTTP handlers
├── dtos/                # Request/response payloads
├── models/               # ORM models (Location, LocationImage)
├── repositories/        # Data access (Postgres + S3/MinIO)
├── services/             # Business logic
├── routers/              # Route definitions
├── migrations/           # SQL migrations
├── utils/                # Config, S3 client, response helpers
├── docker-compose.yml     # Postgres + migration runner
└── main.go
```

## Setup

1. Clone the repo:
   ```bash
   git clone https://github.com/200215-Moynul-Islam/location-api.git
   cd location-api
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Copy the config sample and fill in your values:
   ```bash
   cp conf/app.conf.sample conf/app.conf
   ```
4. Start PostgreSQL and run migrations:
   ```bash
   docker compose up
   ```
5. Set up MinIO (see below).
6. Install the [bee](https://github.com/beego/bee) CLI if you don't have it:
   ```bash
   go install github.com/beego/bee/v2@latest
   ```
7. Run the API with bee (supports hot reload on file changes):
   ```bash
   bee run
   ```
   The server starts on the port set by `httpport` in `conf/app.conf` (`8086` by default).

## Setting up MinIO

`docker-compose.yml` only starts Postgres and runs migrations — MinIO is not included, so it needs to be started separately.

1. **Run a MinIO container:**

   ```bash
   docker run -d --name location-minio \
     -p 9000:9000 -p 9001:9001 \
     -e MINIO_ROOT_USER=minioadmin \
     -e MINIO_ROOT_PASSWORD=minioadmin \
     -v minio_data:/data \
     minio/minio server /data --console-address ":9001"
   ```

   - `9000` is the S3 API port (this is what `MINIO_ENDPOINT` should point to).
   - `9001` is the web console, available at http://localhost:9001.

2. **Log in to the console** at http://localhost:9001 with the root user/password you set above (`minioadmin` / `minioadmin` by default).

3. **Create the bucket** the API expects — its name must match `MINIO_BUCKET` in `conf/app.conf` (`locations` by default):

   - Click **Buckets** in the left sidebar → **Create Bucket**.
   - Enter `locations` as the bucket name → **Create Bucket**.

4. **Match `conf/app.conf` to what you just set up:**

   ```
   MINIO_ENDPOINT=http://localhost:9000
   MINIO_REGION=us-east-1
   MINIO_ACCESS_KEY=minioadmin
   MINIO_SECRET_KEY=minioadmin
   MINIO_BUCKET=locations
   ```

   `MINIO_REGION` isn't enforced by MinIO but is required by the AWS SDK — any value works as long as it's set.

5. **Upload images** for each location into the bucket, and make sure `base_image_key` (in `locations`) and `image_key` (in `location_images`) match the object keys you uploaded them under — the API streams objects by these keys, it doesn't upload them for you.

## Configuration

Set in `conf/app.conf`:

| Key                                     | Description                     |
| --------------------------------------- | ------------------------------- |
| `httpport`                              | Port the API listens on         |
| `MINIO_ENDPOINT`                        | MinIO/S3 endpoint URL           |
| `MINIO_REGION`                          | S3 region                       |
| `MINIO_ACCESS_KEY` / `MINIO_SECRET_KEY` | S3 credentials                  |
| `MINIO_BUCKET`                          | Bucket used for location images |
| `POSTGRES_*`                            | Database connection details     |

## API Endpoints

All routes are prefixed with `/api/v1`.

| Method | Endpoint                    | Description                                                |
| ------ | --------------------------- | ---------------------------------------------------------- |
| GET    | `/health`                   | Health check                                               |
| GET    | `/locations`                | List all locations                                         |
| GET    | `/locations/:id/base-image` | Stream a location's base image                             |
| PUT    | `/locations/:id/base-image` | Set a location's base image from an existing gallery image |
| GET    | `/locations/:id/images`     | List a location's gallery images                           |
| GET    | `/location-images/:id`      | Stream a specific gallery image                            |

### `PUT /locations/:id/base-image`

Request body:

```json
{
  "locationImageId": 12
}
```

## Data Model

- **locations** — `id`, `country`, `state`, `city`, `slug`, `base_image_key`
- **location_images** — `id`, `location_id` (FK → locations, cascade delete), `image_key`

Images are stored in MinIO/S3 by key; `base_image_key` and `image_key` reference object keys in the configured bucket. Setting a base image copies the chosen gallery object onto the location's base image key rather than storing a reference.
