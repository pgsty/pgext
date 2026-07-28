## Usage

Sources:

- [Apache MADlib project](https://madlib.apache.org/)
- [MADlib 2.1.0 user guide](https://madlib.apache.org/docs/latest/)
- [Official module index](https://madlib.apache.org/docs/latest/modules.html)

`madlib` is Apache MADlib 2.1.0, an in-database analytics library for PostgreSQL and Greenplum. It provides SQL-callable statistical, graph, matrix, machine-learning, XGBoost, and deep-learning workflows while keeping training data in the database.

### Core Workflow

MADlib is deployed with its `madpack` utility rather than `CREATE EXTENSION`. After installing a release that supports the target database, use a connection scoped to the intended database:

```sh
madpack -p postgres -c analyst@localhost:5432/analytics install
```

Verify the installed schema before starting an algorithm workflow:

```sql
SELECT madlib.version();
```

Choose a module from the official index, create the module's documented input table, run its training or analysis function, and inspect the generated model or result table.

### Major Module Families

- regression, classification, clustering, sampling, and hypothesis tests
- graph algorithms and path analysis
- arrays, matrices, factorization, and sparse vectors
- model preparation, deep learning, and XGBoost
- data transformation and utility functions

### Requirements and Caveats

- MADlib 2.1.0 is the reviewed release. Use the official supported-database and operating-system matrix for the exact build.
- `madpack` creates a substantial schema of functions and support objects and owns install, upgrade, and uninstall behavior; do not substitute `ALTER EXTENSION UPDATE`.
- Individual modules have different Python and native-library dependencies. Validate those dependencies and resource requirements before enabling a module.
- Training functions commonly create model and summary tables. Use dedicated schemas, review generated-object names, and test rollback and cleanup behavior.
