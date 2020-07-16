# a mysql database custom metrics reporter for newrelic

- Create an .env file with all configurations needed (see .env-example)
- Create a configuration file with your metrics (see config-example.yaml)
- `make docker`
- `make run`

## Example run

```bash
docker run -i -t --rm -e NR_API_KEY=xxx -e NR_ACCOUNT_ID=1234567 -e DATABASE_URL="root:pwd@tcp(127.0.0.1:3306)/information_schema" -v /path/to/my-config.yml:/config.yaml kununulabs/newrelic-mysql-reporter:latest
```
