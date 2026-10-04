## Usage

Sources:

- [README.md](https://github.com/realZachi/pg-jev/blob/8d9598d87d5ff460998d91ec070226176024a841/README.md)
- [CHANGELOG.md](https://github.com/realZachi/pg-jev/blob/8d9598d87d5ff460998d91ec070226176024a841/CHANGELOG.md)
- [sql/jev--0.2.0--0.2.1.sql](https://github.com/realZachi/pg-jev/blob/8d9598d87d5ff460998d91ec070226176024a841/sql/jev--0.2.0--0.2.1.sql)
- [PGXN 0.2.1](https://pgxn.org/dist/jev/0.2.1/)

`jev` filters, ranks and classifies rows using natural-language conditions evaluated by a configured Jev-compatible API. It requires `plpython3u`, superuser installation; TypeSafe-hosted endpoints also require an API key. The upstream documented PostgreSQL range is 14–17; no preload is needed.

### Query rows

```sql
CREATE EXTENSION jev CASCADE;
SET jev.api_key = 'your-key';
CREATE TABLE jev_demo (id integer, body text);
INSERT INTO jev_demo VALUES (1, 'The customer requests a refund');
SELECT id, jev_prob(jev_demo, 'the customer requests a refund')
FROM jev_demo;
SELECT jev_stats();
```

`jev()` returns a boolean predicate; `jev_prob()` returns a probability. `jev_choice()` classifies among options, and `jev_score()` evaluates ordered levels. Cache and connection pools are session-local; `jev_cache_clear()` clears session state.

### Service and data boundaries

Row contents are transmitted to the configured API. Set `jev.api_url` for the intended service and use `jev.max_rows_per_statement` and `jev.max_chars_per_statement` to cap work. API latency and charges depend on the service and data volume. The API key must be handled as a credential.

PL/Python runs with server operating-system privileges. Hosts that withhold superuser access or PL/Python cannot run this extension. Local package tests use the upstream mock API; they do not validate the remote model's judgment quality.

### 0.2.1 Endpoints

In 0.2.1, `jev.api_key` is required only for `*.typesafe.ai` hosts. For a locally operated or other compatible endpoint, unset keys omit the Authorization header; a configured key is still sent. Set `jev.api_url` to the intended trusted service. Install the matching files and run `ALTER EXTENSION jev UPDATE` for 0.2.0 databases. Model accuracy and data-handling policy still depend on the chosen service.
