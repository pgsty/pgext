

## Usage

> [pglogical: PostgreSQL Logical Replication](https://github.com/2ndQuadrant/pglogical)

A logical replication system for PostgreSQL using a publish/subscribe model. Requires no triggers or external programs.

### Enabling

Add to `postgresql.conf`:

```ini
wal_level = 'logical'
max_worker_processes = 10
max_replication_slots = 10
max_wal_senders = 10
shared_preload_libraries = 'pglogical'
```

```sql
CREATE EXTENSION pglogical;
```

### Provider (Publisher) Setup

```sql
-- Create a node on the provider
SELECT pglogical.create_node(
    node_name := 'provider1',
    dsn := 'host=providerhost port=5432 dbname=mydb'
);

-- Add all tables in public schema to default replication set
SELECT pglogical.replication_set_add_all_tables('default', ARRAY['public']);

-- Add all sequences in public schema
SELECT pglogical.replication_set_add_all_sequences('default', ARRAY['public']);
```

### Subscriber Setup

```sql
-- Create a node on the subscriber
SELECT pglogical.create_node(
    node_name := 'subscriber1',
    dsn := 'host=subscriberhost port=5432 dbname=mydb'
);

-- Create a subscription to the provider
SELECT pglogical.create_subscription(
    subscription_name := 'subscription1',
    provider_dsn := 'host=providerhost port=5432 dbname=mydb'
);
```

### Replication Set Management

```sql
-- Create a custom replication set
SELECT pglogical.create_replication_set('my_set');

-- Add a specific table
SELECT pglogical.replication_set_add_table('my_set', 'my_table', true);

-- Remove a table
SELECT pglogical.replication_set_remove_table('my_set', 'my_table');
```

### Row and Column Filtering

```sql
-- Row filtering: only replicate rows matching a condition
SELECT pglogical.replication_set_add_table(
    set_name := 'default',
    relation := 'my_table',
    row_filter := 'id > 1000'
);

-- Column filtering: only replicate specific columns
SELECT pglogical.replication_set_add_table(
    set_name := 'default',
    relation := 'my_table',
    columns := '{id, name, updated_at}'
);
```

### Subscription Management

```sql
-- Check subscription status
SELECT * FROM pglogical.show_subscription_status();

-- Drop subscription
SELECT pglogical.drop_subscription('subscription1');
```

### Key Features

- Selective replication (per-table, row filtering, column filtering)
- Replication between different PostgreSQL major versions
- Delayed replication
- No need for superuser on subscriber
