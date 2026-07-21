## Usage

Sources:

- [pgfr_analyze v2.29.2 README](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/pgfr_analyze/README.md)
- [pgfr_analyze control file](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/pgfr_analyze/extension.control)
- [pg_flight_recorder v2.29.2 reference](https://github.com/dventimisupabase/pg_flight_recorder/blob/v2.29.2/REFERENCE.md)
- [v2.29.2 release notes](https://github.com/dventimisupabase/pg_flight_recorder/releases/tag/v2.29.2)

pgfr_analyze is the read-oriented analysis half of PostgreSQL Flight Recorder. It interprets history captured by pgfr_record to compare periods, summarize waits, assemble incident timelines, detect regressions, and estimate blast radius. Install it after pgfr_record and use it for diagnosis rather than collection.

### Install the Analysis Layer

    CREATE EXTENSION pg_cron;
    CREATE EXTENSION pgfr_record;
    CREATE EXTENSION pgfr_analyze;
    SELECT pgfr_record.enable();

pgfr_analyze depends on recorder objects and has no useful history until pgfr_record has collected samples. It does not require its own background worker.

### Start an Incident Investigation

Summarize what happened around a known timestamp:

    SELECT *
    FROM pgfr_analyze.what_happened_at(
      now() - interval '15 minutes'
    );

Build an ordered incident view over a period:

    SELECT *
    FROM pgfr_analyze.incident_timeline(
      now() - interval '30 minutes',
      now()
    );

Compare a suspect period with a baseline using compare, then narrow the result with wait_summary or anomaly_report. Always inspect the installed function signatures with psql or pg_get_function_arguments because optional parameters can change between releases.

### Analysis Index

- compare: contrast metrics across baseline and incident windows.
- wait_summary: aggregate sampled wait events.
- report and anomaly_report: produce broad or anomaly-focused summaries.
- what_happened_at: inspect observations around one timestamp.
- incident_timeline: order notable events over a range.
- detect_regressions and detect_query_storms: flag worsening behavior or query bursts.
- blast_radius: identify affected sessions or workloads.
- capacity_summary and capacity_report: summarize capacity-related signals.
- configuration analysis: relate relevant setting changes to the period.

### Interpretation Workflow

1. Confirm pgfr_record.health_check() and the available sample interval.
2. Choose explicit baseline and incident windows with comparable workload.
3. Use compare and wait_summary to locate the dominant change.
4. Correlate activity, locks, replication, vacuum, and configuration evidence.
5. Validate hypotheses against PostgreSQL logs, query plans, application telemetry, and host metrics.

### Caveats

- Results are observations and heuristics, not proof of causality. Sparse sampling can miss short events.
- Counter resets, extension upgrades, restarts, retention gaps, and workload seasonality can distort comparisons.
- Some findings are richer when pg_stat_statements is enabled and sufficiently sized.
- Access to analysis output can expose query text and operational metadata; restrict privileges accordingly.
- v2.29.2 mainly improves managed-service scheduling behavior in pgfr_record. It does not replace the need to verify that collection jobs actually ran.
