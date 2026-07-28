## Usage

Sources:

- [Official upstream README](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/endpoint/README.md)
- [Official extension control file (endpoint.control)](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/endpoint/endpoint.control)

`endpoint` — This extension does not enforce any security constraints on it's own. Use it when an application needs this specific database capability. Upstream describes it as a work in progress.

### Core Workflow

```sql
CREATE EXTENSION endpoint;

select endpoint.request(
    '0.3',                              -- version
	'GET',                              -- method
	'/endpoint/0.3/row/{meta.row_id}',  -- path
	'{"key": "val"}',                   -- query string as JSON
	'{"key": "val"}'                    -- post args as JSON
);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.5.0`.
- Install the confirmed extension dependencies first: `meta`.
- The control file marks the extension as non-relocatable.
- Upstream describes the project as a work in progress.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
