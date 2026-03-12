

## Usage

> [pg_fact_loader: build fact tables with Postgres](https://github.com/enova/pg_fact_loader)

Build and maintain fact tables using queue-based change data capture. Processes audit/change log tables to incrementally update fact tables.

### Enabling

```sql
CREATE EXTENSION pg_fact_loader;
```

Optionally with pglogical for replica-based setup:

```sql
CREATE EXTENSION pglogical;
CREATE EXTENSION pglogical_ticker;
CREATE EXTENSION pg_fact_loader;
```

### Workflow

1. **Replicate source tables** to a reporting database (via pglogical or other means)
2. **Create audit/change log tables** on the OLTP system for source tables
3. **Create a fact table** structure for aggregated data
4. **Create a merge function** that takes a key ID and returns one row of the fact table
5. **Configure pg_fact_loader** to wire queue tables to fact tables
6. **Backfill** the fact table initially
7. **Schedule** the worker to process changes continuously

### Configuration Tables

```sql
-- Register a fact table
INSERT INTO fact_loader.fact_tables (fact_table_relid, fact_table_agg_proid, ...)
VALUES ('public.customers_fact'::regclass, 'customers_fact_merge'::regproc, ...);

-- Register queue (audit) tables
INSERT INTO fact_loader.queue_tables (queue_table_relid, queue_of_base_table_relid, ...)
VALUES ('audit.customers_audit'::regclass, 'public.customers'::regclass, ...);

-- Connect queue tables to fact tables with merge functions
INSERT INTO fact_loader.queue_table_deps
    (fact_table_id, queue_table_id, insert_merge_proid, update_merge_proid, delete_merge_proid)
VALUES (1, 1, 'customers_fact_merge'::regproc, 'customers_fact_merge'::regproc, 'customers_fact_merge'::regproc);

-- Define how to retrieve the key from queue entries
INSERT INTO fact_loader.key_retrieval_sequences
    (queue_table_dep_id, return_columns, is_fact_key)
VALUES (1, '{customer_id}', true);
```

### Running the Worker

```sql
-- Process pending changes
SELECT fact_loader.worker();

-- Schedule this to run periodically (e.g., every few seconds via pg_cron)
```

### Initial Backfill

```sql
-- Run the merge function for every existing row
SELECT customers_fact_merge(customer_id) FROM customers;
```

### Adding Batch ID Fields

```sql
SELECT fact_loader.add_batch_id_fields();
```

### Key Features

- Queue-based incremental fact table updates
- Supports insert, update, and delete events
- Handles multi-level key retrieval (joins through multiple tables)
- Fact table dependency chains (child facts updated after parent)
- Checks replication lag before processing (when used with pglogical)
