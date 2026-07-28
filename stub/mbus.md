## Usage

Sources:

- [Official upstream README](https://github.com/plumqqz/mbus4/blob/fb315dd9b5f585798e92a75ac8e09713bc8f39a3/readme.rus)
- [Official extension control file (mbus.control)](https://github.com/plumqqz/mbus4/blob/fb315dd9b5f585798e92a75ac8e09713bc8f39a3/mbus.control)
- [Official extension SQL (mbus--1.1.sql)](https://github.com/plumqqz/mbus4/blob/fb315dd9b5f585798e92a75ac8e09713bc8f39a3/mbus--1.1.sql)

`mbus` — simple message bus for pg. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION mbus;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `can_post(qname text)` is an extension function and returns `boolean`.
- `clear_tempq()` is an extension function and returns `void`.
- `consume` is an extension function.
- `consume_temp(tqname text)` is an extension function and returns `SETOF`.
- `create_consumer` is an extension function.
- `create_run_function(qname text)` is an extension function and returns `void`.
- `create_temporary_consumer` is an extension function.
- `create_temporary_queue()` is an extension function and returns `text`.
- `create_trigger` is an extension function.
- `create_view(qname text, cname text default 'default', sname text default 'public', viewname text default null)` is an extension function and returns `void`.
- `create_view_prop(qname text, cname text, sname text, viewname text, with_delay boolean default false, with_expire boolean default false)` is an extension function and returns `void`.
- `drop_consumer(cname text, qname text)` is an extension function and returns `void`.
- `drop_queue(qname text)` is an extension function and returns `void`.
- `drop_trigger(src text, dst text)` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `1.1`.
- Install the confirmed extension dependencies first: `hstore`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
