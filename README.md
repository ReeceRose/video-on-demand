# Video On Demand (VOD)

## Setup

### Developer Dependencies

1. [Golang](https://go.dev/)
2. [Air](https://github.com/air-verse/air)
3. [Docker](https://docs.docker.com/desktop/)
4. [Docker Compose](https://docs.docker.com/compose/install/)
5. [Make](https://www.gnu.org/software/make/)

```bash
# First time
make dev-build
```

## Development

```bash
# Start all containers, local services are run in "watch" mode and will be recompiled when files change.
make dev
# Stop all containers
make dev-down
```
