## Usage

Sources:

- [README.md](https://github.com/darthunix/tessera/blob/f4fab15667db06ff444dff57331668a6b2e0bc1b/README.md)
- [bridge.md](https://github.com/darthunix/tessera/blob/f4fab15667db06ff444dff57331668a6b2e0bc1b/docs/bridge.md)
- [limitations.md](https://github.com/darthunix/tessera/blob/f4fab15667db06ff444dff57331668a6b2e0bc1b/docs/limitations.md)
- [tessera.control](https://github.com/darthunix/tessera/blob/f4fab15667db06ff444dff57331668a6b2e0bc1b/bridge/tessera.control)
- [tessera--0.1.sql](https://github.com/darthunix/tessera/blob/f4fab15667db06ff444dff57331668a6b2e0bc1b/bridge/tessera--0.1.sql)
- [LICENSE](https://github.com/darthunix/tessera/blob/f4fab15667db06ff444dff57331668a6b2e0bc1b/LICENSE)

`tessera` is the bridge for a batch executor over ordinary PostgreSQL heap tables. The project supplies C modules and Rust kernels for supported analytical plan nodes. SQL extension version 0.1 loads the bridge; it exposes a C rendezvous API rather than SQL functions. Upstream targets PostgreSQL 18, 19 and development master.

### Enable the Modules

Install matching bridge, node and kernel libraries. For regular use, preload them in this order and restart PostgreSQL:

```ini
shared_preload_libraries = 'tessera, tessera_nodes, tessera_kernels'
```

Preserve existing library entries. Create the extension as a superuser and compare plans on representative data:

```sql
CREATE EXTENSION tessera;
CREATE TABLE tessera_demo (x integer);
INSERT INTO tessera_demo SELECT generate_series(1, 100000);
ANALYZE tessera_demo;
EXPLAIN ANALYZE SELECT count(*), sum(x) FROM tessera_demo WHERE x > 1000;
SET tessera.enable = off;
EXPLAIN ANALYZE SELECT count(*), sum(x) FROM tessera_demo WHERE x > 1000;
```

The bridge must load before the node and kernel modules. Session preload in the same order is also supported for development. CREATE EXTENSION only loads the bridge in the creating backend; its catalog presence does not load it in later connections. A fresh connection must preload or explicitly load the libraries.

### Plans and Limits

Supported nodes include heap/index scans, batch filters and aggregates, hash joins, sorting and set operations. Unsupported nodes remain PostgreSQL nodes in a mixed plan. Window functions, foreign tables, other table access methods, and scans that write or lock rows remain on the ordinary executor. Inspect EXPLAIN for actual selection; no speedup is guaranteed.

Visibility follows the PostgreSQL snapshot. Hashes and sorts use work_mem and hash_mem_multiplier and can spill under temp_file_limit. Cancellation is checked in batch loops. Rebuild modules when changing PostgreSQL installations and replace loaded postmaster/session libraries when upgrading. The project is under active development; validate result equivalence and workload behavior before use.
