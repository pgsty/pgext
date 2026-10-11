## Usage

Sources:

- [README.md](https://github.com/joelonsql/pg-catbench/blob/d81740c55a3aad7c599c502d044ba8db6f6e0837/README.md)
- [Makefile](https://github.com/joelonsql/pg-catbench/blob/d81740c55a3aad7c599c502d044ba8db6f6e0837/Makefile)
- [functions/new_benchmark.sql](https://github.com/joelonsql/pg-catbench/blob/d81740c55a3aad7c599c502d044ba8db6f6e0837/functions/new_benchmark.sql)
- [views/vreport.sql](https://github.com/joelonsql/pg-catbench/blob/d81740c55a3aad7c599c502d044ba8db6f6e0837/views/vreport.sql)
- [catbench.control](https://github.com/joelonsql/pg-catbench/blob/d81740c55a3aad7c599c502d044ba8db6f6e0837/catbench.control)

`catbench` stores definitions and results for PostgreSQL catalog-function benchmarks. Manual benchmark execution exists; automatic commit monitoring is still a project goal.

### Core Workflow

```sql
CREATE EXTENSION timeit;
CREATE EXTENSION catbench;
SELECT * FROM catbench.vreport;
```

### Operational Boundaries

Install `timeit` first. `catbench.new_benchmark()` registers a benchmark definition and `catbench.vreport` exposes reports; upstream companion executables drive collection and execution. The installation SQL is assembled by Makefile. Although the control file declares a module path, this extension Makefile installs SQL only and defines no native module. It needs no preload, is not relocatable, and sets superuser to false without marking the extension trusted. PostgreSQL-major support is not declared.
