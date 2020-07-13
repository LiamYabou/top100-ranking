# About this project
Top100 is a project that presents the top 100 categories in this world. At present, you can browser the top 100 products of the best sellers on Amazon with http://staging.amazon.top100.technology.

The explanation of the architecture can be found [here](https://github.com/LiamYabou/top100-scrapy/wiki/Architecture).

## Top100 Ranking
The top100-ranking is a microservice of the Top100 project that provisions a server of the RPC to handle the request of fetching dedicated list of rankings.

# Devlopment
## Dependencies
- golang 1.14
- rabbitmq 3.8
- postgresql 12

## Environment Variables
We use [direnv](https://direnv.net/) to streamline the loads of the env variables in the project.
```
export ENV=development
export APP_NAME=top100-ranking-staging
export DB_NAME=top100_development
export DB_USER=postgres
export DB_PASSWORD=
export DB_PORT=5432
export DB_HOST=localhost
export SSL_MODE=disable
export APP_URI=.../top100-ranking
export CLOUDAMQP_URL=amqp://guest:guest@localhost:5672
export TEST_DB_DSN=postgres://postgres:@localhost:5432/top100_test?sslmode=disable
export MAX_POOL_CONNECTIONS=25
export MIN_POOL_CONNECTIONS=5
export GOROUTINE_CONCURRENCY=25
```

## Log Monitor
```
tail -f logs/development.log
```

## RPC Server
You can run the following command to launch the server of the RPC.
```
bin/launch_rpc_server
```

## Testing
> If you have some inquires, please ask for help from the `make help` command first.
```
make test
```

# Contributing
If you have any suggestions or any issues you discovered, you can contact me via hello@mengliu.dev or commit a new `pull request`. I appreciate your help!
