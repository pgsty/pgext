## Usage

Sources:

- [Official documentation](https://github.com/USC-InfoLab/fft_approximate/blob/f330f53d42e0e4fecb4dd1ccd29e480fc5407d72/README.md)
- [Extension control file](https://github.com/USC-InfoLab/fft_approximate/blob/f330f53d42e0e4fecb4dd1ccd29e480fc5407d72/fft_approximate.control)
- [Official repository](https://github.com/USC-InfoLab/fft_approximate)

`fft_approximate` FFT-based approximate aggregate range queries over time-series data.

### Enablement

Install the files for the intended server, then create `fft_approximate` in the target database:

```sql
CREATE EXTENSION fft_approximate;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT approximate_avg(c, 0, 10, 10 ORDER BY k) FROM coeffs;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `re` | FUNCTION | Callable function from the reviewed install surface. |
| `approximate_avg` | AGGREGATE | Aggregate exposed by the extension. |
| `complex` | TYPE | User-facing data type created by the extension. |
| `im` | FUNCTION | Callable function from the reviewed install surface. |
| `complex_add` | FUNCTION | Callable function from the reviewed install surface. |
| `complex_avg` | FUNCTION | Callable function from the reviewed install surface. |
| `complex_avg_accum` | FUNCTION | Callable function from the reviewed install surface. |
| `complex_in` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
