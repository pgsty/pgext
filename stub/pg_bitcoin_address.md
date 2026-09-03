## Usage

Sources:

- [Official documentation](https://github.com/whitslack/pg_bitcoin_address/blob/1ab081b6a55cc4bfc8b01c63c0637fcb4d8ae6b6/README.md)
- [Extension control file](https://github.com/whitslack/pg_bitcoin_address/blob/1ab081b6a55cc4bfc8b01c63c0637fcb4d8ae6b6/pg_bitcoin_address.control)
- [Official repository](https://github.com/whitslack/pg_bitcoin_address)

`pg_bitcoin_address` Bitcoin address types plus Base58Check and Bech32/Bech32m codecs.

### Enablement

Install the files for the intended server, then create `pg_bitcoin_address` in the target database:

```sql
CREATE EXTENSION pg_bitcoin_address;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
=> SELECT pg_column_size('1BitcoinEaterAddressDontSendf59kuE'::text);
pg_column_size | 38

=> SELECT pg_column_size('1BitcoinEaterAddressDontSendf59kuE'::bitcoin_address);
pg_column_size | 26

=> SELECT pg_column_size('bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4'::text);
pg_column_size | 46

=> SELECT pg_column_size('bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4'::bitcoin_address);
pg_column_size | 26
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `bitcoin_address` | FUNCTION | Callable function from the reviewed install surface. |
| `program` | FUNCTION | Callable function from the reviewed install surface. |
| `base58check` | TYPE | User-facing data type created by the extension. |
| `version` | FUNCTION | Callable function from the reviewed install surface. |
| `hrp` | FUNCTION | Callable function from the reviewed install surface. |
| `is_blinding` | FUNCTION | Callable function from the reviewed install surface. |
| `is_liquidtestnet` | FUNCTION | Callable function from the reviewed install surface. |
| `is_liquidv1` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
