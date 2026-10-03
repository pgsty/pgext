## Usage

Sources:

- [STOMATA 0.1.0 README](https://github.com/CrystallineCore/Stomata/blob/0193b0cdc90a6d9866553f0788a51a462a8fb7ec/README.md)
- [Control](https://github.com/CrystallineCore/Stomata/blob/0193b0cdc90a6d9866553f0788a51a462a8fb7ec/stomata.control)
- [SQL 0.1.0](https://github.com/CrystallineCore/Stomata/blob/0193b0cdc90a6d9866553f0788a51a462a8fb7ec/sql/stomata--0.1.0.sql)

`stomata` 0.1.0 indexes `LIKE` and `ILIKE` using segmented automata. It targets anchored patterns, short literals and underscore wildcards. Upstream labels this first version as testing; it requires PostgreSQL 16 or later and reports testing on 16, 17 and 18.

### Core Workflow

```sql
CREATE EXTENSION stomata;
CREATE TABLE contacts (id bigint, email text);
CREATE INDEX contacts_email_stomata ON contacts USING stomata (email);

SELECT * FROM contacts WHERE email LIKE 'jo%@example.com';
SELECT * FROM contacts WHERE email ILIKE 'JO%';
SELECT * FROM stomata_index_info('contacts_email_stomata');
SELECT stomata_verify('contacts_email_stomata');
```

Install as a superuser; the control file is not trusted. Each index covers one text or varchar column or expression. Expression and partial indexes are supported. Candidate matches are rechecked by the executor.

### Objects and Tuning

- `stomata_index_info`, `stomata_runs`, `stomata_key_stats`: inspect runs, pending rows, sizes and key families.
- `stomata_keys`, `stomata_pattern_keys`, `stomata_pattern_stats`: inspect value or pattern keys.
- `stomata_candidate_pages`, `stomata_estimate`: inspect candidate pages and planner estimates.
- `stomata_verify`: counts visible tuples missing from the index; zero is the expected result.
- `stomata_merge_pending`, `stomata_compact`: flush pending rows or consolidate all runs.
- Index options include `k` (segment width, default 3), `exact` (row-level postings, on), `exact_bigrams` (off), `rollup` (3), `skip_depth` (4), and `pending_limit` (4096 kB). Setting the last option to zero leaves merging to vacuum or explicit maintenance.
- `stomata.estimate_budget` bounds planner estimation work. Rebuild with `REINDEX` after changing index options through `ALTER INDEX`.

### Maintenance and Limits

Inserts can trigger merges and increase statement latency. Vacuum and compaction can rewrite runs and generate substantial WAL; full compaction can temporarily need about twice the live index size. Freed space is reused; rebuild to return it to the operating system.

There are no index-only scans, ordering support, multicolumn indexes or parallel index builds. Negated patterns, regular expressions and similarity operators are not indexed. ASCII case folding reduces precision for non-ASCII case-insensitive searches; nondeterministic collations fall back to scanning every page. Standby page-reuse conflicts can cancel queries. Before 1.0, an upstream disk-format change can require rebuilding indexes.
