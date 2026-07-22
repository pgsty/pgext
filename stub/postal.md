## Usage

Sources:

- [Official README](https://github.com/pramsey/pgsql-postal/blob/105a6ca855dbf4d33b160ca50d00c3c76a7e2061/README.md)
- [Extension SQL](https://github.com/pramsey/pgsql-postal/blob/105a6ca855dbf4d33b160ca50d00c3c76a7e2061/postal--1.0.sql)
- [libpostal wrapper implementation](https://github.com/pramsey/pgsql-postal/blob/105a6ca855dbf4d33b160ca50d00c3c76a7e2061/postal.c)

postal exposes libpostal address normalization and parsing inside PostgreSQL. It can expand an address into normalized alternatives or label address components as JSONB. It is useful for search preparation and data cleanup, but normalization is probabilistic preprocessing rather than proof that an address exists or is deliverable.

### Core Workflow

After the server administrator installs libpostal and its model data, create the extension and call the two functions:

```sql
CREATE EXTENSION postal;

SELECT unnest(postal_normalize('412 first ave, victoria, bc'));

SELECT postal_parse('412 first ave, victoria, bc');
```

Store the original input alongside derived values. Choose one normalized form according to application rules rather than assuming the first alternative is canonical.

### Important Objects

- `postal_normalize(text)` returns a `text[]` containing libpostal's normalized alternatives.
- `postal_parse(text)` returns a `jsonb` object whose keys label components such as house number, road, city, or state.

### Operational Notes

libpostal has a noticeable first-call initialization delay and loads large model data into every backend that invokes it; upstream observed roughly 1 GB per active backend. Limit concurrency or isolate normalization in a controlled worker pool. Results depend on the installed libpostal version and data model, so upgrades can change derived output. The checked wrapper targets PostgreSQL 9.4-era JSONB and has no Windows installation path; test modern server compatibility and representative international addresses.
