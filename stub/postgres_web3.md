## Usage

Sources:

- [Project README](https://github.com/Yen/postgres_web3/blob/85e356422b38390d68796483068f9c1913f6a7f3/README.md)
- [Extension control file](https://github.com/Yen/postgres_web3/blob/85e356422b38390d68796483068f9c1913f6a7f3/postgres_web3.control)
- [Version 1.0 SQL definitions](https://github.com/Yen/postgres_web3/blob/85e356422b38390d68796483068f9c1913f6a7f3/postgres_web3--1.0.sql)

`postgres_web3` 1.0 supplies fixed-width `int128`, `uint128`, `int256`, `uint256`, `hex160`, and `hex256` types for blockchain-sized integers, addresses, and hashes. It defines arithmetic, comparison and bitwise operators, casts, B-tree and hash operator classes, and numeric aggregates.

### Model Web3 values

```sql
CREATE EXTENSION postgres_web3;

CREATE TABLE chain_transfer (
    tx_hash hex256 PRIMARY KEY,
    from_address hex160 NOT NULL,
    amount uint256 NOT NULL
);

INSERT INTO chain_transfer (tx_hash, from_address, amount)
VALUES ('01'::hex256, '88ff'::hex160,
        '340282366920938463463374607431768211456'::uint256);

SELECT amount + '1'::uint256
FROM chain_transfer
WHERE from_address = '88ff'::hex160;
```

Integer input uses base-10 text. Hex input is case-insensitive, contains no `0x` prefix, and is emitted as lowercase text. Arithmetic is range checked, so operations that do not fit the destination type raise an error.

### Build and precision constraints

The reviewed README requires C23 `_BitInt`, Clang 16, and a specially built PostgreSQL without LLVM support, and recommends its prebuilt Linux container instead of a source build. These constraints are historical and narrow; verify compiler, architecture, PostgreSQL ABI, backup, restore, replication, and client-driver behavior before storing durable production values.

The `avg` aggregates use a double-precision internal sum and the upstream documentation warns that results are inaccurate for large sums. Use an exact alternative when averages affect balances or accounting. Casts marked for assignment can still fail on negative or out-of-range values, and fixed-width hex output is not an Ethereum checksum-address implementation.
