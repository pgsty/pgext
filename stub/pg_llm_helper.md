## Usage

Sources:

- [Official README](https://github.com/GeoffMontee/pg_llm_helper/blob/feae47af77a9cb735400b66b0b39f046d60cf924/README.md)
- [Extension SQL](https://github.com/GeoffMontee/pg_llm_helper/blob/feae47af77a9cb735400b66b0b39f046d60cf924/pg_llm_helper--1.0.sql)
- [Shared-memory implementation](https://github.com/GeoffMontee/pg_llm_helper/blob/feae47af77a9cb735400b66b0b39f046d60cf924/pg_llm_helper.c)
- [Extension control file](https://github.com/GeoffMontee/pg_llm_helper/blob/feae47af77a9cb735400b66b0b39f046d60cf924/pg_llm_helper.control)

`pg_llm_helper` 1.0 installs a postmaster-wide error hook and a 100-entry shared-memory ring containing failed query text, error text, SQLSTATE, severity, backend PID, and timestamp. It can optionally send the current backend's last captured error to an LLM through pgai. Because the history crosses sessions and databases and can contain credentials or personal data, it must be restricted like a cluster-wide diagnostic log.

### Preload and Install

Add the library to `shared_preload_libraries` and restart; loading it later cannot allocate the shared memory or register the hooks.

```conf
shared_preload_libraries = 'pg_llm_helper'
```

Create the extension in a controlled administration database and immediately remove default access:

```sql
CREATE EXTENSION pg_llm_helper;

REVOKE EXECUTE ON FUNCTION get_last_error() FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION get_error_history(integer) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION clear_error_history() FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION llm_help_last_error() FROM PUBLIC;

SELECT * FROM get_last_error();
SELECT * FROM get_error_history(10);
```

The upstream README calls `vector` a required dependency, but the reviewed control file declares no `requires` entry and the installed error-capture SQL does not use vector objects. The LLM helper does require a compatible `ai.openai_chat_complete` function at call time; pgai is not declared as an extension dependency. Validate the exact packaged dependency set rather than relying on `CREATE EXTENSION ... CASCADE`.

### User-Facing Functions

- `get_last_error()`: find the newest ring entry whose backend PID equals the current PID.
- `get_error_history(integer)`: copy up to 100 populated slots across all backends; array traversal is not guaranteed to be newest-first.
- `clear_error_history()`: clear the entire cluster-wide ring, not only the caller's entries.
- `llm_help_last_error()`: format the last error and its query, call the hard-coded `gpt-4o-mini` model through pgai, and return the response.

The compile-time limits are 100 errors, 8192 bytes of query text, and 1024 bytes of error text. Entries contain neither database nor user identity. PID reuse can cause a new backend to find an older process's entry until that slot is overwritten.

### Privacy and Reliability Boundaries

The hook captures `ERROR` and higher severities from `debug_query_string`. Queries and error messages can contain literals, API keys, customer data, object names, and security-policy details. Any role allowed to call `get_error_history` can read other sessions' captured text, and any role allowed to call `clear_error_history` can erase evidence. Limit execution to audited administrators, align log retention and disclosure policy, and test multi-tenant and PID-reuse behavior.

Calling `llm_help_last_error` sends captured query and error text to an external provider synchronously. Apply data classification, redaction, egress, residency, retention, cost, timeout, and provider-key controls before granting it. The returned suggestion is untrusted diagnostic text, never authorization or executable remediation.

The reviewed project targets PostgreSQL 17 and uses global logging and shared-memory hooks. Test hook chaining with every other preloaded module, startup/restart behavior, buffer contention, truncation, crashes, upgrades, and performance under error storms. Do not assume the ring is an audit log: it is small, overwrite-prone, globally clearable, and lacks durable ordering or identity fields.
