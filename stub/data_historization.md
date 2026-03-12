


## Usage

> [data_historization: Track data changes in partitioned log tables](https://github.com/rodo/postgresql-data-historization)

PL/pgSQL scripts to historize data changes into partitioned tables.

### Initialize Historization

Set up the necessary objects for a table (no data collected yet):

```sql
SELECT historize_table_init('public', 'my_table');
```

### Start Historization

Install triggers to begin collecting changes into the `_log` table:

```sql
SELECT historize_table_start('public', 'my_table');
```

### Stop Historization

Remove triggers and stop collecting changes:

```sql
SELECT historize_table_stop('public', 'my_table');
```

### Reset Historization

Remove cron entries and columns created on the source table:

```sql
SELECT historize_table_reset('public', 'my_table');
```

### Clean Historization

Remove the log table entirely:

```sql
SELECT historize_table_clean('public', 'my_table');
```

### Partition Management

Create and drop partitions manually:

```sql
SELECT historize_create_partition('public', 'my_table_log', 0);
SELECT historize_drop_partition('public', 'my_table_log', 0);
```

Automate with `pg_cron`:

```sql
SELECT cron.schedule_in_database(
    'create-partitions', '00 08 * * *',
    $$SELECT historize_create_partition('my_table', generate_series(1, 4))$$,
    'my_database'
);
```
