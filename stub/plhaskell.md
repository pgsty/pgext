## Usage

Sources:

- [Project README](https://github.com/ed-o-saurus/PLHaskell/blob/32805cf338d2088697144f29420a1c31b2efddff/README.md)
- [Extension control file](https://github.com/ed-o-saurus/PLHaskell/blob/32805cf338d2088697144f29420a1c31b2efddff/src/plhaskell.control)
- [Version 5.0 SQL language definitions](https://github.com/ed-o-saurus/PLHaskell/blob/32805cf338d2088697144f29420a1c31b2efddff/src/plhaskell--5.0.sql)
- [Usage and query API](https://github.com/ed-o-saurus/PLHaskell/blob/32805cf338d2088697144f29420a1c31b2efddff/docs/Usage.md)
- [Trust model](https://github.com/ed-o-saurus/PLHaskell/blob/32805cf338d2088697144f29420a1c31b2efddff/docs/Trust.md)
- [Known bugs](https://github.com/ed-o-saurus/PLHaskell/blob/32805cf338d2088697144f29420a1c31b2efddff/docs/Known_Bugs.md)

`plhaskell` 5.0 embeds Haskell as a PostgreSQL procedural language. Installation creates trusted `plhaskell` for typed `PGm` actions and untrusted `plhaskellu` for `IO`; only superusers may create functions in the untrusted variant. The database encoding must be UTF-8.

### Define a trusted function

```sql
CREATE EXTENSION plhaskell;

CREATE FUNCTION add_haskell(int, int)
RETURNS int
IMMUTABLE
AS $$
  import Data.Int (Int32)
  import PGutils (PGm)

  add_haskell :: Maybe Int32 -> Maybe Int32 -> PGm (Maybe Int32)
  add_haskell Nothing Nothing = return Nothing
  add_haskell (Just a) (Just b) = return $ Just $ a + b
  add_haskell a Nothing = return a
  add_haskell Nothing b = return b
$$
LANGUAGE plhaskell;

SELECT add_haskell(20, 22);
```

The extension also supports composites, arrays, timestamps, advisory locks, parameterized SPI queries, sets, and anonymous blocks. Function volatility controls whether internal queries are read-only; still review every SQL query and imported module.

### Runtime and trust boundary

The trusted-language claim relies on the extension's Haskell type and module restrictions. It also invokes the GHC API and dynamically loaded Haskell runtime inside PostgreSQL backends, so granting ordinary roles permission to create functions materially expands the database execution surface. Independently review the trusted module set, package database, filesystem and SELinux policy, statement timeouts, role privileges, and denial-of-service behavior. Reserve `plhaskellu` for tightly controlled superuser code.

The default Haskell runtime memory limit is 128 MB; `plhaskell.max_memory = 0` disables it and is unsafe for multi-tenant use. Upstream documents failures with some `lc_messages` values, severe slowdown when a set-returning function is called multiple times, and platform runpath issues. Builds are intended for Linux and tested only on x86-64. Pin matching PostgreSQL, GHC, hint, and runtime libraries, then test backend crashes, errors, upgrades, replication, and dump and restore before production deployment.
