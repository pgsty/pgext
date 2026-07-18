## Usage

Sources:

- [Project README](https://github.com/sycertech/pksuid/blob/1ab87a77bd2a2e4e3e95ab900f11c27f63ce48e4/README.md)
- [Extension control file](https://github.com/sycertech/pksuid/blob/1ab87a77bd2a2e4e3e95ab900f11c27f63ce48e4/pksuid-extension/pksuid.control)
- [PostgreSQL extension implementation](https://github.com/sycertech/pksuid/blob/1ab87a77bd2a2e4e3e95ab900f11c27f63ce48e4/pksuid-extension/src/lib.rs)
- [Pksuid type implementation](https://github.com/sycertech/pksuid/blob/1ab87a77bd2a2e4e3e95ab900f11c27f63ce48e4/pksuid/src/lib.rs)
- [Extension Cargo manifest](https://github.com/sycertech/pksuid/blob/1ab87a77bd2a2e4e3e95ab900f11c27f63ce48e4/pksuid-extension/Cargo.toml)

`pksuid` 0.1.0 provides a PostgreSQL type and generator for prefixed KSUIDs. Values combine an application-readable prefix with a time-sortable, high-entropy KSUID, producing identifiers such as a customer-specific `client_...` value without relying on a central sequence.

### Use prefixed keys

```sql
CREATE EXTENSION pksuid;

CREATE TABLE client (
    id pksuid PRIMARY KEY DEFAULT pksuid_generate('client'),
    name text NOT NULL
);

INSERT INTO client (name) VALUES ('Ada');
SELECT * FROM client ORDER BY id;
```

The source-defined SQL function is `pksuid_generate(text)`, and the data type is `pksuid`. The README's `CREATE EXTENSION prefixed_ksuid` line is stale; the actual control filename and catalog extension name are `pksuid`.

### Comparison semantics

The Rust type's equality and ordering implementations compare only the underlying KSUID, not the prefix. Therefore the prefix is a display/domain label, not part of database equality or sort order. Its parser splits text at `_`; keep prefixes free of underscores and validate prefix policy in the application. The reviewed extension crate is 0.1.0 and exposes pgrx build features for PostgreSQL 11 through 16; build separately for each target major version.
