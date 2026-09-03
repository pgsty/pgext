## Usage

Sources:

- [Official documentation](https://github.com/brunoenten/pg_oxr/blob/6df5bd5515c2e088a03d44a0357d3ae129dfe11e/README.md)
- [Extension control file](https://github.com/brunoenten/pg_oxr/blob/6df5bd5515c2e088a03d44a0357d3ae129dfe11e/oxr.control)
- [Official repository](https://github.com/brunoenten/pg_oxr)

`oxr` OpenExchangeRates API client and currency-rate cache inside PostgreSQL.

### Enablement

Install the files for the intended server, then create `oxr` in the target database:

```sql
CREATE EXTENSION oxr CASCADE;
```

The reviewed control or official workflow requires `http`. `CASCADE` only succeeds when those extension files are already installed on the server.

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Store your API key (replace 'your_api_key' with your actual key)
SELECT oxr.set_config('api_key', 'your_api_key');
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `oxr.set_config` | FUNCTION | Callable function from the reviewed install surface. |
| `oxr.get_config` | FUNCTION | Callable function from the reviewed install surface. |
| `oxr.get_historical_rate` | FUNCTION | Callable function from the reviewed install surface. |
| `oxr.get_latest_rate` | FUNCTION | Callable function from the reviewed install surface. |
| `oxr.historical_rates` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `oxr.set_app_id` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 13, 14, 15, 16, 17, 18; do not infer unlisted majors.
- The extension fixes or creates schema objects under `oxr`; include them in privilege and backup review.
- Catalog lifecycle is abandoned; test upgrades, dump/restore, and server compatibility before production use.
