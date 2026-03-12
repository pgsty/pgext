


## Usage

> [plluau: Lua as an untrusted procedural language](https://github.com/pllua/pllua)

`plluau` is the untrusted variant of `pllua`, allowing Lua functions to access the filesystem, load arbitrary modules, and perform operations restricted in the trusted version.

```sql
CREATE EXTENSION plluau;
```

### Create Functions

```sql
CREATE FUNCTION read_file(path text) RETURNS text LANGUAGE plluau AS $$
  local f = io.open(path, "r")
  if f then
    local content = f:read("*a")
    f:close()
    return content
  end
  return nil
$$;
```

### Differences from pllua (Trusted)

| Feature | pllua (trusted) | plluau (untrusted) |
|---------|----------------|--------------------|
| File I/O | Restricted | Full access |
| Module loading | Whitelisted only | Unrestricted |
| OS access | No | Yes |
| Suitable for | User-defined functions | Admin/superuser functions |

### Same API as pllua

`plluau` shares the same SPI interface, trigger support, set-returning functions, and data type handling as `pllua`. All SPI functions (`spi.execute`, `spi.prepare`, `spi.rows`), coroutine-based set returns, and trigger functions work identically.

```sql
CREATE FUNCTION run_command(cmd text) RETURNS text LANGUAGE plluau AS $$
  local handle = io.popen(cmd)
  local result = handle:read("*a")
  handle:close()
  return result
$$;
```

### Initialization

Configure via GUC (superuser only):

```sql
SET pllua.on_untrusted_init = 'myvar = {}';
```

Only superusers can create `plluau` functions due to the unrestricted access it provides.
