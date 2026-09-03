## Usage

Sources:

- [Official documentation](https://github.com/AlgoNode/pg_algorand/blob/7bc3552bdb84a684b75dde2c7c8e91aaafa089cc/README.md)
- [Extension control file](https://github.com/AlgoNode/pg_algorand/blob/7bc3552bdb84a684b75dde2c7c8e91aaafa089cc/pg_algorand.control)
- [Official repository](https://github.com/AlgoNode/pg_algorand)

`pg_algorand` Algorand address, application, asset, transaction, and checksum encoding utilities.

### Enablement

Install the files for the intended server, then create `pg_algorand` in the target database:

```sql
CREATE EXTENSION pg_algorand;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT
  COUNT(*) FROM account
WHERE
  addr = AddressTxt2Bin('ALGONODEIBJTET5OSEAXIHDSIEG7C2DOFB2WDYLRZTXN3NXVJ3NJD26L4E');
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `AddressBin2Txt` | FUNCTION | Callable function from the reviewed install surface. |
| `AddressTxt2Bin` | FUNCTION | Callable function from the reviewed install surface. |
| `GetNFDSigNameLSIG` | FUNCTION | Callable function from the reviewed install surface. |
| `GetNFDSigRevAddressBinLSIG` | FUNCTION | Callable function from the reviewed install surface. |
| `GetNFDSigRevAddressLSIG` | FUNCTION | Callable function from the reviewed install surface. |
| `algoaddr` | TYPE | User-facing data type created by the extension. |
| `algoaddr_in` | FUNCTION | Callable function from the reviewed install surface. |
| `algoaddr_out` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
