## Usage

Sources:

- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/pg/README.md)
- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/pg/pg_mcp/Cargo.toml)
- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/pg/pg_mcp/pg_mcp.control)
- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/pg/pg_mcp/sql/pg_mcp--0.3.0.sql)
- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/pg/pg_mcp/src/session.rs)
- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/pg/pg_mcp/src/lib.rs)
- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/docs/pg-mcp-install.md)
- [Official source](https://github.com/noizu-labs-ml/elixir-mcp-lib/blob/f3401ad75ba2299648a8a0a7f80c559a9fe51ae9/LICENSE)

`pg_mcp` 0.3.0 is the Rust PostgreSQL extension in the Noizu MCP repository. It supplies `mcp_fdw`, MCP call functions and generated SQL objects. The Cargo feature matrix covers PostgreSQL 16–18 with pgrx 0.19.2. It requires superuser installation, creates its own `mcp` schema, is not relocatable, and does not require shared preload.

### Connect to an MCP server

Provision a reachable MCP Streamable HTTP service and the referenced secret table before using this example. The secret lookup runs with the caller's database privileges; configure its expected token row and grants according to the exact source contract.

```sql
CREATE EXTENSION pg_mcp;
CREATE SERVER engine FOREIGN DATA WRAPPER mcp_fdw
  OPTIONS (url 'https://engine.internal/mcp', timeout_ms '15000');
CREATE USER MAPPING FOR CURRENT_USER SERVER engine
  OPTIONS (token_secret 'mcp_secrets.engine');
CREATE SCHEMA mcp_tools;
SELECT mcp.import('engine', 'mcp_tools');
SELECT * FROM mcp.generate_functions('engine', 'mcp_tools');
SELECT mcp.refresh('engine');
```

The control file requires a superuser despite a looser statement in the operator runbook. Do not separately create the extension-owned `mcp` schema. The tracked installation SQL includes development-test wrappers; install SQL generated for the same production library rather than pairing that test-enabled snapshot with an arbitrary binary.

### SQL surface

`mcp.call_tool` returns JSON and `mcp.call_tool_text` returns text. `mcp.get_prompt`, `mcp.read_resource` and `mcp.complete` expose the corresponding MCP calls. In generic mode, imported catalog/read-through tables expose tools and resources; insertion into `tool_calls` invokes one external tool per row. Generated per-tool tables, functions and views are recorded in `mcp.generated`; regeneration replaces only recorded objects with RESTRICT, so dependent user objects can block it.

The optional engine layout uses `all_upstreams` and `per_upstream_schema`. Tool availability follows the remote principal, not merely local table privileges. `mcp.refresh` discards the current backend's session and catalog cache; it does not automatically regenerate SQL objects.

### Privileges and side effects

Use per-role USER MAPPING entries and deliberate foreign-server, schema and object grants. Shared PUBLIC mappings collapse remote identities. Prefer secret references over literal tokens in migrations. Non-loopback endpoints require HTTPS, and anonymous authentication is restricted to loopback. Timeouts bound exchanges; remote errors and unavailable services can fail SQL calls.

Remote tool actions are external side effects and cannot be assumed to roll back with a PostgreSQL transaction. Generated SELECT wrappers can invoke tools; keep the read-only annotation gate and review any override. The protocol's `sql/*` family is experimental. Preserve externally managed secrets and audit objects during upgrades and restores, and use the supplied SQL upgrade chain with matching library files.
