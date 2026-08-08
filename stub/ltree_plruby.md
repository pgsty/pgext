## Usage

Sources:

- [PL/Ruby v2.5.0 README](https://github.com/commandprompt/plruby/blob/v2.5.0/README.md)
- [ltree_plruby v1.0 control file](https://github.com/commandprompt/plruby/blob/v2.5.0/ltree_plruby/ltree_plruby.control)
- [ltree_plruby v1.0 extension SQL](https://github.com/commandprompt/plruby/blob/v2.5.0/ltree_plruby/ltree_plruby--1.0.sql)

`ltree_plruby` installs a PostgreSQL transform between `ltree` paths and Ruby arrays for the `plruby` language. An `ltree` argument becomes an array of label strings, and an array of valid labels can be returned directly as an `ltree` value.

### Install and Use the Transform

```sql
CREATE EXTENSION ltree;
CREATE EXTENSION plruby;
CREATE EXTENSION ltree_plruby;

CREATE FUNCTION ruby_append_label(ltree, text)
RETURNS ltree
LANGUAGE plruby
TRANSFORM FOR TYPE ltree
AS $$
  path = args[0]
  path << args[1]
  path
$$;

SELECT ruby_append_label('Top.Science'::ltree, 'Astronomy');
```

The transform is used only by functions that declare `TRANSFORM FOR TYPE ltree`.

### Objects and Caveats

- `ltree_to_plruby(internal)` implements SQL-to-Ruby conversion.
- `plruby_to_ltree(internal)` implements Ruby-to-SQL conversion.
- The extension version is `1.0`, it requires both `ltree` and `plruby`, and it is relocatable.
- Every returned array element must be a valid `ltree` label. PostgreSQL rejects invalid characters, empty labels, or paths that exceed `ltree` limits.
- PL/Ruby remains untrusted. Installing this transform does not sandbox Ruby code or reduce the privileges required to create PL/Ruby functions.
