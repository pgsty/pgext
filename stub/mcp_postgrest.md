## Usage

Sources:

- [Official upstream README](https://github.com/cybertheory/pgmcp/blob/26e25d6beb457ac064b971b9fa856cfdea1ba046/README.md)
- [Official extension control file (mcp_postgrest.control)](https://github.com/cybertheory/pgmcp/blob/26e25d6beb457ac064b971b9fa856cfdea1ba046/mcp_postgrest.control)
- [Official extension SQL (mcp_postgrest--0.1.2.sql)](https://github.com/cybertheory/pgmcp/blob/26e25d6beb457ac064b971b9fa856cfdea1ba046/sql/mcp_postgrest--0.1.2.sql)

`mcp_postgrest` — The mcp_postgrest PostgreSQL extension transforms your database into an AI-powered tool interface compliant with Anthropic’s Model Context Protocol (MCP). Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION mcp_postgrest;

SELECT call_tool('tool_name', '{"arg1": "value"}'::jsonb);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `call_tool(tool_name TEXT, args JSONB)` is an extension function and returns `JSONB`.
- `create_crud_tools_for_table(tablename TEXT)` is an extension function and returns `VOID`.
- `generate_ai_tool(provider TEXT, api_key TEXT, tool_name TEXT, description TEXT, table_names TEXT[])` is an extension function and returns `TEXT`.
- `mcp_on_create_table()` is an extension function and returns `event_trigger`.
- `mcp_tools` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.2`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
