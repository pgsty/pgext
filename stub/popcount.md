## Usage

Sources:

- [Official README](https://github.com/eschmar/postgresql-popcount/blob/34fadc73163f22ac7b3f567ebded20d80703a597/readme.md)
- [Extension SQL for version 1.0.0](https://github.com/eschmar/postgresql-popcount/blob/34fadc73163f22ac7b3f567ebded20d80703a597/popcount--1.0.0.sql)
- [PGXN distribution page](https://pgxn.org/dist/popcount/)

`popcount` provides several C implementations for counting set bits in PostgreSQL `bit(n)` values. It was designed to compare lookup-table, Hamming-weight, compiler-intrinsic, and assembly strategies. PostgreSQL 14 and later already provide built-in `bit_count`, which upstream recommends instead for modern servers.

### Core Workflow

On PostgreSQL releases before 14, install the extension and choose an implementation only after validating it on the target CPU and bit widths:

```sql
CREATE EXTENSION popcount;

SELECT popcount(B'101101');
SELECT popcount64(B'101101');
SELECT popcountAsm64(B'101101');
```

Every function is immutable and strict, accepts a varying-length `bit` value, and returns `integer`. Null input therefore returns null without calling the C function.

### Available Implementations

- `popcount` uses an 8-bit lookup-table strategy.
- `popcount32` and `popcount64` use 32-bit and 64-bit Hamming-weight algorithms.
- `popcountAsm` and `popcountAsm64` use 32-bit and 64-bit intrinsic/instruction-oriented paths.
- `popcount256` uses an unrolled assembly-oriented implementation.

The upstream benchmark favored `popcountAsm64` on its tested hardware, but that result is not a portability or performance guarantee. CPU instruction availability, compiler flags, alignment, bit width, PostgreSQL version, and workload can change both safety and speed.

On PostgreSQL 14 or later, prefer the built-in operation:

```sql
SELECT bit_count(B'101101');
```

### Operational Notes

Version 1.0.0 is relocatable and declares no preload or extension dependency. Because the extension includes architecture-specific optimized code, run its installation checks on every target build and deployment architecture; do not copy a binary between incompatible CPUs.

The extension adds six global function names, so schema placement and search-path behavior matter. Benchmark with representative `bit(n)` widths and compare correctness across all implementations before choosing one. Migrating to built-in `bit_count` removes a native extension dependency and should be tested for return-type and plan differences in application queries.
