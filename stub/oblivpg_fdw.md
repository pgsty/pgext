## Usage

Sources:

- [Official README](https://github.com/rogerioacp/oblivpg_ftw/blob/5a26f0ffa02242aa187d22bde8ca8dc3bc05dde0/README.md)
- [Extension SQL](https://github.com/rogerioacp/oblivpg_ftw/blob/5a26f0ffa02242aa187d22bde8ca8dc3bc05dde0/oblivpg_fdw--1.0.sql)
- [FDW implementation](https://github.com/rogerioacp/oblivpg_ftw/blob/5a26f0ffa02242aa187d22bde8ca8dc3bc05dde0/oblivpg_fdw.c)

oblivpg_fdw is a research prototype for oblivious table access through Intel SGX and a Secure Oblivious Engine using Path ORAM or Forest ORAM. It is coupled to custom PostgreSQL source-tree builds, enclave libraries, shared memory, mirror tables, and manually supplied object identifiers. It is not a general-purpose or production-ready FDW.

### Core Workflow

Creating the extension registers the wrapper and a server named `obliv`, and also creates the global mapping table `obl_ftw`. The README's test flow then manually records the OIDs of a foreign table, mirror table, and mirror index before opening and initializing the enclave:

```sql
CREATE EXTENSION oblivpg_fdw;

INSERT INTO obl_ftw (
  ftw_table_oid, mirror_table_oid, mirror_index_oid,
  ftw_table_nblocks, ftw_index_nblocks, init
) VALUES (16399, 16392, 16395, 100, 100, false);

SELECT open_enclave();
SELECT init_soe(16399);
```

The numeric OIDs are only placeholders here; upstream requires resolving them from the exact test schema. Its sample foreign-table definition is commented out in the installation SQL, so there is no portable table contract to copy blindly.

### Important Objects

- `oblivpg_fdw` is the wrapper and `obliv` is the server created at installation.
- `obl_ftw` stores the foreign-table, mirror-table, and mirror-index OIDs plus block counts.
- `open_enclave` and `close_enclave` control the SGX enclave lifecycle.
- `init_soe` initializes the Secure Oblivious Engine for a mapping.
- `load_blocks`, `attach_shmem`, and `set_nextterm` expose low-level test and data-loading hooks.

### Security and Compatibility Notes

The build requires Intel SGX SDK and external SOE/ORAM libraries, with compile-time modes that can bypass the enclave or inject dummy queries. Hard-coded OIDs are unstable across restore and schema recreation. Installing the extension drops any existing server `obliv` and foreign table `ftw_users`, then creates global objects. The prototype does not provide a security proof, modern compatibility claim, key-management procedure, or recovery workflow; use it only to reproduce the original research in an isolated environment.
