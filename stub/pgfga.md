## Usage

Sources:

- [pgfga README at the reviewed commit](https://github.com/craigpastro/pgfga/blob/dcdaad718a596d4e0284aafc739db40f273f788b/README.md)
- [pgfga.control at the reviewed commit](https://github.com/craigpastro/pgfga/blob/dcdaad718a596d4e0284aafc739db40f273f788b/pgfga.control)
- [Cargo package metadata](https://github.com/craigpastro/pgfga/blob/dcdaad718a596d4e0284aafc739db40f273f788b/Cargo.toml)
- [SQL API and table definitions](https://github.com/craigpastro/pgfga/blob/dcdaad718a596d4e0284aafc739db40f273f788b/src/lib.rs)
- [Authorization evaluator](https://github.com/craigpastro/pgfga/blob/dcdaad718a596d4e0284aafc739db40f273f788b/src/check.rs)
- [Tuple storage implementation](https://github.com/craigpastro/pgfga/blob/dcdaad718a596d4e0284aafc739db40f273f788b/src/storage.rs)

`pgfga` is an experimental relationship-based access-control engine inspired by Zanzibar. It stores JSON authorization schemas and relationship tuples, then evaluates direct relations and permission rewrites through `pgfga.check`. Management functions include `pgfga.create_schema`, `pgfga.read_schemas`, `pgfga.create_tuple`, `pgfga.read_tuples`, and `pgfga.delete_tuple`.

### Minimal Relationship Check

The following example uses psql's `\gset` to retain the generated schema UUID:

```sql
CREATE EXTENSION pgfga;

SELECT pgfga.create_schema(
  '{"namespaces":{"document":{"relations":{"viewer":[{"namespace":"user"}]}},"user":{}}}'::json
) AS schema_id \gset

SELECT pgfga.create_tuple(
  :'schema_id'::uuid, 'document', 'doc-1', 'viewer', 'user', 'anya', ''
);
SELECT pgfga.check(
  :'schema_id'::uuid, 'document', 'doc-1', 'viewer', 'user', 'anya', ''
);
```

### Caveats

- The repository is archived, and upstream calls version `0.1.0` experimental work in progress. Do not use it as a production authorization boundary.
- Tuple writes verify only that the schema UUID exists; they do not validate namespaces, relations, or subject restrictions against the JSON schema. The tuple table also lacks a schema foreign key and secondary lookup indexes.
- The reviewed intersection and exclusion evaluators suppress some nested check errors and can continue to a true result. That is unsafe fail-open behavior for authorization decisions.
- Read paths materialize matching tuples into Rust vectors, and recursive checks issue repeated SPI queries. Benchmark depth and fan-out only after correcting the security semantics.
