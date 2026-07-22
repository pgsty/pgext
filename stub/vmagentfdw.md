## Usage

Sources:

- [Official README](https://github.com/alesharik/vmagent-fdw/blob/1ef48ceb1e9fa6d27fc22a668d9385ffb834caea/README.md)
- [Official FDW implementation](https://github.com/alesharik/vmagent-fdw/blob/1ef48ceb1e9fa6d27fc22a668d9385ffb834caea/src/fdw.rs)
- [Official HTTP client implementation](https://github.com/alesharik/vmagent-fdw/blob/1ef48ceb1e9fa6d27fc22a668d9385ffb834caea/src/client.rs)

`vmagentfdw` is a read-only Rust foreign data wrapper that exposes vmagent's active and dropped scrape targets as PostgreSQL rows. It reads the vmagent `/api/v1/targets` endpoint synchronously during a foreign-table scan, so remote latency and availability directly affect the querying backend.

### Core Workflow

Install the extension, register the handler and validator, and point a foreign server at a trusted vmagent endpoint:

```sql
CREATE EXTENSION vmagentfdw;

CREATE FOREIGN DATA WRAPPER vmagent_wrapper
  HANDLER vmagent_fdw_handler
  VALIDATOR vmagent_fdw_validator;

CREATE SERVER local_vmagent
  FOREIGN DATA WRAPPER vmagent_wrapper
  OPTIONS (address 'http://127.0.0.1:8429');
```

Define only the columns needed by queries. Their names drive the wrapper's value mapping:

```sql
CREATE FOREIGN TABLE vmagent_targets (
  state text,
  health text,
  last_samples_scraped bigint,
  last_scrape_duration double precision,
  last_scrape timestamp,
  last_error text,
  scrape_url text,
  scrape_pool text,
  labels jsonb,
  discovered_labels jsonb,
  labels_job text
)
SERVER local_vmagent;

SELECT state, health, scrape_url, labels_job
FROM vmagent_targets;
```

### Column Mapping

- `state` is synthesized as `active` or `dropped` from the response section containing the target.
- `health`, `last_samples_scraped`, `last_scrape_duration`, `last_scrape`, `last_error`, `scrape_url`, and `scrape_pool` map the corresponding target fields.
- `labels` and `discovered_labels` return the complete label maps as JSONB.
- A column named `labels_<key>` projects one entry from the target's label map. Unrecognized columns return NULL.

### Operational Caveats

The implementation fetches all active and dropped targets for each scan and does not push PostgreSQL filters, ordering, or limits into the HTTP request. It implements no write callbacks. The reviewed source declares builds for PostgreSQL `14`, `15`, and `16`, but compatibility still depends on the installed binary. Restrict who may create or query servers, use a bounded and trusted endpoint, and enforce suitable statement/network timeouts. The documented options provide only the endpoint `address`; authentication and TLS policy must therefore be handled by the selected URL and surrounding network controls.
