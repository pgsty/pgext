## Usage

Sources:

- [Official source](https://github.com/determinate/pg_systemone/blob/bdccd600eb458eb30018db6ebc6d1ee6ae828c78/README.md)
- [Official source](https://github.com/determinate/pg_systemone/blob/bdccd600eb458eb30018db6ebc6d1ee6ae828c78/pg_systemone.control)
- [Official source](https://github.com/determinate/pg_systemone/blob/bdccd600eb458eb30018db6ebc6d1ee6ae828c78/pg_systemone--1.0.sql)
- [Official source](https://github.com/determinate/pg_systemone/blob/bdccd600eb458eb30018db6ebc6d1ee6ae828c78/src/pg_systemone.c)

`pg_systemone` exposes SystemOne decision-model inference through C SQL functions. It sends requests to a separate HTTP service; it does not embed the model in PostgreSQL. Upstream documents PostgreSQL 14–17 and requires libcurl plus a compatible server.

### Core Workflow

```sql
CREATE EXTENSION pg_systemone;
SET systemone.api_url = 'http://127.0.0.1:9931/v1/systemone';
SELECT systemone_prob('The shipment arrived damaged.', 'Does this need customer support?');
```

`systemone_choice(state, instructions, criteria)` returns a categorical choice. `systemone_prob(state, question)` returns a probability, `systemone_score(state, question, levels)` scores ordered labels, and `systemone_match_score(job_desc, resume)` compares a job description and resume. These are model responses, not deterministic truth.

### Configuration and Access

`systemone.api_key` supplies an optional bearer token. `systemone.batch_size` is currently a hint for future batching, not a guarantee of batched requests. Legacy `jev_*` functions and `jev.*` settings are aliases; this is a distinct C extension from the Python-based Jev extension, but their SQL names can collide.

Installation requires a superuser; no preload is documented. The extension is relocatable. HTTP calls run synchronously in the backend and transmit supplied text. Endpoint and credential settings are user-settable, and the SQL script leaves function execution public: restrict function privileges and network egress before exposing it to untrusted roles. The README's license badge is not backed by a license file in this revision; the catalog leaves the license unverified.
