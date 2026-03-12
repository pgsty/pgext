

## Usage

> [pgmp: multiple precision arithmetic (GMP) for PostgreSQL](https://github.com/dvarrazzo/pgmp/)

The `pgmp` extension integrates the GNU Multiple Precision (GMP) library into PostgreSQL, providing arbitrary-precision integer and rational number types.

```sql
CREATE EXTENSION pgmp;
```

### Data Types

- **`mpz`**: Arbitrary-size integers, limited only by PostgreSQL's 1GB varlena maximum
- **`mpq`**: Arbitrary-precision rational numbers for exact fractional arithmetic

### Basic Usage

```sql
-- Arbitrary precision integers
SELECT '123456789012345678901234567890'::mpz * 2;

-- Exact rational arithmetic (no rounding)
SELECT '1'::mpq / '3'::mpq;  -- 1/3

-- Mixed arithmetic with native PostgreSQL types
SELECT 42::mpz + '100'::mpz;
```

### Operators

Standard arithmetic operators (`+`, `-`, `*`, `/`, `%`) and comparison operators (`=`, `<>`, `<`, `>`, `<=`, `>=`) are supported for both `mpz` and `mpq`.

### Functions

The extension exposes all GMP library functions for these types, including:

- Prime number functions
- Random number generation
- Factorization
- GCD, LCM, and other number theory functions

### Index Support

Both `mpz` and `mpq` support btree and hash indexes for efficient queries and sorting.

Full documentation: https://www.varrazzo.com/pgmp/
