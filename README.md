# Golang server app template
Template for basic Golang server using MUX router

## Includes configs for

- [x] Rotating Logs
- [x] Hot reload with [air](https://github.com/cosmtrek/air)
- [x] Mysql Database
- [x] Docker
- [x] Send email
- [ ] Redis
- [ ] Rabbit MQ
- [ ] Metrics monitoring endpoint for [Prometheus](https://prometheus.io/)

## Setting up
1. Clone repos
2. Open go.mod, change the module name
3. Run `go mod download`
4. Copy `.env.example` to `.env` and configure with your details\
5. Install `air` for live reload. Install from [here](https://github.com/cosmtrek/air).

## Sample routes

| Endpoint | Action  |
|:---------|:-------|
| ``/``         | Home page with html|
| ``/api/v1``      | Responds with json format|

