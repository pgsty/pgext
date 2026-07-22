## Usage

Sources:

- [Official README](https://github.com/petere/pgvihash/blob/70e637321ab07746ecaebf6c47385b06c84c050c/README)
- [Official SQL definition](https://github.com/petere/pgvihash/blob/70e637321ab07746ecaebf6c47385b06c84c050c/vihash.sql)
- [Official C implementation](https://github.com/petere/pgvihash/blob/70e637321ab07746ecaebf6c47385b06c84c050c/vihash.c)
- [Official PGXN distribution](https://pgxn.org/dist/pgvihash/)

`vihash` provides one stable 32-bit text hash, `vihashtext(text)`, copied from PostgreSQL 8.3's `hashtext` implementation. Its purpose is preserving sharding/partition-routing decisions while different PostgreSQL major versions coexist during a migration.

### Core Workflow

Install the extension on every participating database and verify identical results before changing routing code:

```sql
CREATE EXTENSION vihash;

SELECT vihashtext('tenant-42');

SELECT key, vihashtext(key) AS routing_hash
FROM (VALUES ('tenant-1'), ('tenant-2')) AS v(key);
```

In PL/Proxy-era code, replace `RUN ON hashtext(arg)` with `RUN ON vihashtext(arg)` only as part of a planned redistribution/compatibility migration.

### Compatibility Contract

`vihashtext(text) RETURNS integer` hashes the raw database-encoding bytes and is intended to remain version-independent. Its output matches built-in `hashtext` from PostgreSQL 8.0–8.3. It does not match the changed built-in implementation introduced in PostgreSQL 8.4.

### Caveats

If data is already distributed with PostgreSQL 8.4-or-later `hashtext`, switching to `vihashtext` changes routes and requires redistribution. Test a fixed corpus on old and new nodes before cutover, keep database encoding identical, and handle the signed 32-bit result consistently when mapping it to shard numbers.

This is a migration aid, not a cryptographic or collision-resistant hash. Upstream recommends broader stable-hash libraries for new designs, and the repository is archived. Never use the result for authentication, integrity, uniqueness, or adversarial partition keys without collision handling.
