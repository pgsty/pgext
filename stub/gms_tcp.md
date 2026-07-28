## Usage

Sources:

- [Official upstream README](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/README)
- [Official extension control file (gms_tcp.control)](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/gms_tcp/gms_tcp.control)
- [Official extension SQL (gms_tcp--1.0--1.1.sql)](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/gms_tcp/gms_tcp--1.0--1.1.sql)

`gms_tcp` — provides TCP/IP client=side access functionality in PL/SQL. Use it when porting or emulating the corresponding database API. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION gms_tcp;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gms_tcp.available(c in gms_tcp.connection, timeout in int default 0)` is an extension function and returns `integer`.
- `gms_tcp.available_real(c in gms_tcp.connection, timeout in int)` is an extension function and returns `integer`.
- `gms_tcp.close_all_connections()` is an extension function and returns `void`.
- `gms_tcp.close_connection(c in gms_tcp.connection)` is an extension function and returns `void`.
- `gms_tcp.connection_in(cstring)` is an extension function and returns `gms_tcp`.
- `gms_tcp.connection_out(gms_tcp.connection)` is an extension function and returns `cstring`.
- `gms_tcp.crlf()` is an extension function and returns `varchar2`.
- `gms_tcp.flush(c in gms_tcp.connection)` is an extension function and returns `void`.
- `gms_tcp.get_line(c in gms_tcp.connection, remove_crlf in boolean default false, peek in boolean default false)` is an extension function and returns `text`.
- `gms_tcp.get_line_real(c in gms_tcp.connection, remove_crlf in boolean, peek in boolean, ch_charset in boolean default false)` is an extension function and returns `text`.
- `gms_tcp.get_raw(c in gms_tcp.connection, len in integer default 1, peek in boolean default false)` is an extension function and returns `raw`.
- `gms_tcp.get_raw_real(c in gms_tcp.connection, len in integer, peek in boolean)` is an extension function and returns `raw`.
- `gms_tcp.get_text(c in gms_tcp.connection, len in integer default 1, peek in boolean default false)` is an extension function and returns `text`.
- `gms_tcp.get_text_real(c in gms_tcp.connection, len in integer, peek in boolean, ch_charset in boolean default false)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.1`.
- The control file marks the extension as non-relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
