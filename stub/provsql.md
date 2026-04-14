## Usage

- Source: [GitHub repo](https://github.com/PierreSenellart/provsql), [project docs](https://provsql.org/docs/), [getting started guide](https://provsql.org/docs/user/getting-provsql.html)
- ProvSQL adds m-semiring provenance and uncertainty management to PostgreSQL, with support for probabilities, Shapley values, and semiring evaluation.

```sql
shared_preload_libraries = 'provsql'
CREATE EXTENSION provsql CASCADE;
```

The upstream quick install also calls out these prerequisites: PostgreSQL 10 or later, a C++17 compiler, PostgreSQL headers, `uuid-ossp`, and Boost libraries.

## Core Workflow

ProvSQL is loaded through `shared_preload_libraries`, then installed with `CREATE EXTENSION provsql CASCADE;`.

Typical use cases include:

- evaluating provenance over different semirings
- computing probabilities and expected values
- computing game-theoretic contributions such as Shapley values
- using the built-in compiled semirings for common cases

## Notes

The project homepage and documentation live at [provsql.org](https://provsql.org/). The README links to the user guide for the full installation and testing workflow.
