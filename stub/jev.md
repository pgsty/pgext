## Usage

Sources:

- [PGXN 0.2.0](https://pgxn.org/dist/jev/0.2.0/)

`jev` filters, ranks and classifies rows using natural-language conditions evaluated by a TypeSafe API. It requires `plpython3u`, superuser installation and an API key. The upstream documented PostgreSQL range is 14–17; no preload is needed.

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
