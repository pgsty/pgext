## Usage

Sources:

- [Official README](https://github.com/choplin/plmruby/blob/5fd693cfa3fd44eb1f1178bbd94d184c399fd4a6/README.md)
- [Official control file](https://github.com/choplin/plmruby/blob/5fd693cfa3fd44eb1f1178bbd94d184c399fd4a6/plmruby.control)
- [Official language definition](https://github.com/choplin/plmruby/blob/5fd693cfa3fd44eb1f1178bbd94d184c399fd4a6/plmruby--0.0.1.sql)
- [Official scalar examples](https://github.com/choplin/plmruby/blob/5fd693cfa3fd44eb1f1178bbd94d184c399fd4a6/sql/plmruby.sql)
- [Official runtime build configuration](https://github.com/choplin/plmruby/blob/5fd693cfa3fd44eb1f1178bbd94d184c399fd4a6/build_config.rb)

`plmruby` embeds mruby as a PostgreSQL procedural language for scalar, set-returning, inline, and trigger functions. The reviewed `0.0.1` tree is an unfinished 2015 implementation; its README leaves the quick-start and set-returning documentation incomplete.

### Core Workflow

Only a superuser can install the extension. Named PostgreSQL arguments are exposed under the same names in mruby code.

```sql
CREATE EXTENSION plmruby;
REVOKE USAGE ON LANGUAGE plmruby FROM PUBLIC;

CREATE FUNCTION ruby_add(a integer, b integer)
RETURNS integer
LANGUAGE plmruby
IMMUTABLE STRICT
AS $$
  a + b
$$;

SELECT ruby_add(2, 3);
```

Unnamed arguments use `_1`, `_2`, and so on. Variadic arguments can be named normally. A `DO` block can use the language's inline handler.

### Data and Trigger Model

The implementation converts PostgreSQL numeric values, booleans, strings, dates and timestamps, JSON, arrays, and records into corresponding mruby values. Unsupported types pass through their PostgreSQL text input/output functions, which can lose type-specific semantics.

Row-trigger code receives `new`, `old`, `tg_name`, `tg_when`, `tg_level`, `tg_op`, `tg_relid`, `tg_table_name`, `tg_table_schema`, and `tg_argv`. For a row-level before trigger, returning `:skip` cancels the row and returning a tuple-like hash replaces it; return values are ignored for statement-level and after triggers.

### Trust and Compatibility Boundaries

The installation SQL declares `plmruby` as a trusted language, but the embedded runtime build includes file I/O, directory, environment, process, and system-information mruby gems. Treat it as capable of host-level actions inside a database backend. Revoke public language usage and allow function creation only to roles trusted at the operating-system boundary.

The language has no validator, and the upstream source predates modern supported PostgreSQL releases. Test compilation, type conversion, errors, triggers, and set-returning functions on the exact server build. Do not assume that contemporary Ruby libraries or CRuby behavior are available in the embedded mruby runtime.
