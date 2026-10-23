## Usage

Sources:

- [README.md](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/README.md)
- [hdr_histogram--1.0.0.sql](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/hdr_histogram--1.0.0.sql)
- [hdr_histogram.c](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/hdr_histogram.c)
- [META.json](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/META.json)
- [.github/workflows/ci.yml](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/.github/workflows/ci.yml)
- [hdr_histogram.control](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/hdr_histogram.control)
- [Official 1.0.0 release archive](https://github.com/vovcke/pg_hdr_histogram/releases/download/v1.0.0/hdr_histogram-1.0.0.tar.gz)

`hdr_histogram` stores compressed HDR histograms and computes percentiles with configurable precision. `hdr_add` aggregates numeric samples; `hdr_merge` combines stored histograms. PostgreSQL 14–18 appear in upstream CI. The trusted, relocatable extension requires zlib and can be installed by a role with CREATE privilege on the database.

The official 1.0.0 release archive used for Pigsty packages installs control/SQL version 1.0; the Git tag's control file declares 1.0.0. Check the installed extension version rather than inferring it from the package version.

### Core Workflow

```sql
CREATE EXTENSION hdr_histogram;
SELECT hdr_percentile(hdr_add(1.0), 50);
SET hdr_histogram.merge_mode = 'lossless';
SELECT hdr_percentile(hdr_add(latency::double precision), 99)
FROM (VALUES (100), (200), (300), (500)) AS samples(latency);
CREATE TABLE latency_histograms (recorded_at timestamptz, hist hdr_histogram);
SELECT hdr_percentile(hdr_merge(hist), 99) FROM latency_histograms;
```

### Operational Boundaries

Samples are converted to int64, so choose one unit consistently. The default highest value is 1e9 with three significant figures; hdr_histogram.lowest, hdr_histogram.highest and hdr_histogram.sig_figs configure new histograms. Nonfinite and out-of-range values error. `hdr_percentile` accepts a percentile from 0 to 100; `hdr_min`, `hdr_max`, `hdr_mean`, `hdr_stddev`, `hdr_total_count`, `hdr_sig_figs`, `hdr_lowest` and `hdr_highest` inspect a histogram. The default merge_mode lossless rejects incompatible merges; strict requires identical parameters, left can discard values outside the accumulator range, and union expands range/precision. The first histogram function call loads the library; the example calls it before applying session GUCs. No shared preload is required for core use. Percentiles are bounded-error estimates, not exact order statistics.
