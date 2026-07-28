## Usage

Sources:

- [Official upstream README](https://github.com/commandprompt/plruby/blob/0720d8e72522c5196db062a1610eb2031a832246/README.md)
- [Official extension control file (hstore_plruby.control)](https://github.com/commandprompt/plruby/blob/0720d8e72522c5196db062a1610eb2031a832246/hstore_plruby/hstore_plruby.control)
- [Official extension SQL (hstore_plruby--1.0.sql)](https://github.com/commandprompt/plruby/blob/0720d8e72522c5196db062a1610eb2031a832246/hstore_plruby/hstore_plruby--1.0.sql)

`hstore_plruby` — transform between hstore and Ruby Hashes. Use it when database code must run in or interoperate with this procedural language. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION hstore_plruby;

CREATE EXTENSION plruby;

CREATE FUNCTION hello(text) RETURNS text LANGUAGE plruby AS $$
    "Hello, #{args[0]}!"
$$;

SELECT hello('world');   -- Hello, world!
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hstore_to_plruby(val internal)` is an extension function and returns `internal`.
- `plruby_to_hstore(val internal)` is an extension function and returns `hstore`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `hstore`, `plruby`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
