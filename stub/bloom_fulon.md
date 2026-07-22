## Usage

Sources:

- [Official extension control file](https://github.com/matheus-consoli/boom-index-pgrx/blob/7025b8186d770eee48ecff55f683c50355d98ff4/bloom_fulon.control)
- [Official Rust implementation](https://github.com/matheus-consoli/boom-index-pgrx/blob/7025b8186d770eee48ecff55f683c50355d98ff4/src/lib.rs)
- [Official Rust package manifest](https://github.com/matheus-consoli/boom-index-pgrx/blob/7025b8186d770eee48ecff55f683c50355d98ff4/Cargo.toml)

`bloom_fulon` is an unfinished `pgrx` example that registers a PostgreSQL index access method. Version `0.0.0` is useful for studying access-method callbacks, but it is not a functional Bloom index and must not hold production data.

### Core Workflow

Creating the extension registers the `bloom_fulon` access method and a default `int4_ops` operator class for integer equality:

```sql
CREATE EXTENSION bloom_fulon;

CREATE TABLE bloom_demo (id integer);
CREATE INDEX bloom_demo_idx ON bloom_demo USING bloom_fulon (id)
  WITH (size = 10);
```

The `size` relation option defaults to 10 and accepts values from -10 through 100. It is only a demonstration option; the implementation does not use it to build a real filter.

### Operational Boundary

The build callback reports zero indexed tuples, insertion returns without storing an entry, and vacuum callbacks perform no maintenance. Both `amgettuple` and `amgetbitmap` are absent, so the access method cannot return matching rows. Its cost estimate is fixed at `1e10` to discourage planner use.

The control file requires superuser installation and is non-relocatable. The reviewed manifest defines features for PostgreSQL 11 through 16, but that build matrix does not make the implementation complete. Use `bloom_fulon` only in a disposable development database for access-method experiments.
