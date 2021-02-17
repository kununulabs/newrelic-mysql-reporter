# A MySQL Database Custom Metrics Reporter for New Relic

A small tool to send single-value custom metrics (events), defined as queries, to newrelic.

## Running the example
```bash
docker build -t kununu/newrelic-mysql-reporter:local .
docker run -d --name newrelic-mysql-reporter-example-db -e MYSQL_ROOT_PASSWORD=example -p 3306:3306 mysql:5.7
sleep 30
cat >.env <<ENV
NEW_RELIC_ACCOUNT_ID=1234567
NEW_RELIC_INSIGHTS_INSERT_KEY=xxx
NEW_RELIC_REGION=eu
DATABASE_URL=root:example@tcp(172.17.0.1:3306)/information_schema
ENV
docker run -it --rm --env-file .env -v ${PWD}/yaml/example.yaml:/config.yaml kununu/newrelic-mysql-reporter:local
docker rm -f newrelic-mysql-reporter-example-db
```

## Use the official docker image
https://hub.docker.com/r/kununulabs/newrelic-mysql-reporter

```bash
docker pull kununulabs/newrelic-mysql-reporter:latest
```
