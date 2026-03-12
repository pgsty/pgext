

## Usage

> [pglogical_ticker: Have an accurate view on pglogical replication delay](https://github.com/enova/pglogical_ticker)

A background worker that periodically updates ticker tables to measure pglogical replication lag from the provider's standpoint.

### Enabling

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'pglogical,pglogical_ticker'
pglogical_ticker.database = 'mydb'
pglogical_ticker.naptime = 10      -- tick interval in seconds (default 10)
```

Install on both provider and all subscribers:

```sql
CREATE EXTENSION pglogical_ticker;
```

### Deploy Ticker Tables

Run on the **provider** only (propagated to subscribers via pglogical):

```sql
-- Deploy ticker tables (one per replication set)
SELECT pglogical_ticker.deploy_ticker_tables();

-- Add ticker tables to replication
SELECT pglogical_ticker.add_ticker_tables_to_replication();
```

For cascading replication:

```sql
SELECT pglogical_ticker.deploy_ticker_tables('my_cascaded_set_name');
SELECT pglogical_ticker.add_ticker_tables_to_replication('my_cascaded_set_name');
```

### Manual Tick

```sql
SELECT pglogical_ticker.tick();
```

### Launching the Ticker

The ticker auto-launches on server start if configured in `shared_preload_libraries`. Otherwise:

```sql
SELECT pglogical_ticker.launch();

-- Or, only launch if replication set tables exist
SELECT pglogical_ticker.launch_if_repset_tables();
```

### Viewing Replication Delay

On the **provider**:

```sql
SELECT * FROM pglogical_ticker.all_repset_tickers();
```

On a **subscriber**:

```sql
SELECT * FROM pglogical_ticker.all_subscription_tickers();
```

### Configuration

- `pglogical_ticker.database` - Database to run the ticker in
- `pglogical_ticker.naptime` - Tick interval in seconds (default 10)
- `pglogical_ticker.restart_time` - Seconds before auto-restart (default 10, -1 to disable)
