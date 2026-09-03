## Usage

Sources:

- [Official Advanced Storage Pack documentation](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/)
- [Official configuration guide](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/configuring/)
- [Official usage guide](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/using/)

`refdata` is an EDB table access method for mostly static reference tables that are read frequently and modified rarely.

### Enablement

Install Advanced Storage Pack 1.7.0, preload `refdata`, restart, and create the extension:

```ini
shared_preload_libraries = 'refdata'
```

```sql
CREATE EXTENSION refdata;
```

### Create Reference Data

```sql
CREATE TABLE market_symbol (
  symbol_id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  symbol    text NOT NULL,
  name      text NOT NULL
) USING refdata;

CREATE TABLE trade (
  symbol_id integer NOT NULL REFERENCES market_symbol(symbol_id),
  traded_at timestamptz NOT NULL,
  price     float8 NOT NULL
);
```

Foreign-key readers avoid row locks in the reference table. Direct modification of a Refdata table instead takes a table-level `ExclusiveLock`.

### Concurrency Boundary

Use Refdata only when changes are infrequent and can tolerate serialization. A write blocks concurrent modifications to the table and to referencing tables, so bulk reference updates require an explicit maintenance plan. Test DDL, vacuuming, foreign keys, logical replication, backup/restore, and failover with the exact EDB/PostgreSQL version before migrating an existing heap table.

