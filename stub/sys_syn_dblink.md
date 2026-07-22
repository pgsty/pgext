## Usage

Sources:

- [Official README](https://github.com/pgstuff/sys_syn_dblink/blob/0c7a18da219a2e71b83ce374e589cfa383a72518/README.md)
- [Official user guide](https://github.com/pgstuff/sys_syn_dblink/blob/0c7a18da219a2e71b83ce374e589cfa383a72518/doc/sys_syn_dblink.adoc)
- [Official extension SQL](https://github.com/pgstuff/sys_syn_dblink/blob/0c7a18da219a2e71b83ce374e589cfa383a72518/sql/sys_syn_dblink.sql)

`sys_syn_dblink` is the receiving-side processor for `sys_syn` queues. It uses a named `dblink` connection to claim batches on a source database, pull changes, transform them, write local target tables, and push processing status back to the source.

### Setup

Install `sys_syn` and configure its input/output queues on the source. On the receiving database, install the declared dependencies and this extension, open the named connection, define input/output/put groups, and generate a process table:

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION dblink;
CREATE EXTENSION sys_syn_dblink;

SELECT dblink_connect('source_syn', 'service=sys_syn_source');

INSERT INTO sys_syn_dblink.in_groups_def
VALUES ('source-cluster', 'in');
INSERT INTO sys_syn_dblink.out_groups_def
VALUES ('source-cluster', 'out');
INSERT INTO sys_syn_dblink.put_groups_def
VALUES ('put');

SELECT sys_syn_dblink.proc_table_create(
    proc_schema     => 'processor_data',
    in_table_id     => 'customer',
    out_group_id    => 'out',
    put_group_id    => 'put',
    dblink_connname => 'source_syn'
);
```

Optional rows in `put_table_transforms` and `put_column_transforms` change table/column mappings. Generated worker functions implement claim, pull, process, and status-push stages. The official guide requires repeated pull/process/push cycles because each call handles a limited batch; claim operations also use an explicit remote transaction.

### Dependencies and Operations

The receiver requires `hstore` and `dblink`; the remote database requires `sys_syn`. Temporal and bitemporal table types additionally require the specified `temporal_tables` version or patch. The reviewed project documents PostgreSQL 9.5 or later, but version `0.0.1` uses extensive generated PL/pgSQL and old APIs, so verify the exact modern-server combination.

Keep the named `dblink` connection open in the session running generated procedures. Protect connection credentials and transformation definitions: they determine remote access and dynamically generated local SQL. Monitor held/retry states, bound batch sizes, serialize claims correctly, and test interrupted transactions so rows are not lost or processed twice. The catalog classifies this version as preview.
