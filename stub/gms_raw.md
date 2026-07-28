## Usage

Sources:

- [Official upstream README](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/README)
- [Official extension control file (gms_raw.control)](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/gms_raw/gms_raw.control)
- [Official extension SQL (gms_raw--1.0.sql)](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/gms_raw/gms_raw--1.0.sql)

`gms_raw` — process raw type data for PL/SQL applications. Use it when porting or emulating the corresponding database API. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION gms_raw;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gms_raw.bit_and(r1 in raw, r2 in raw)` is an extension function and returns `raw`.
- `gms_raw.bit_complement(r1 in raw)` is an extension function and returns `raw`.
- `gms_raw.bit_or(r1 in raw, r2 in raw)` is an extension function and returns `raw`.
- `gms_raw.bit_xor(r1 in raw, r2 in raw)` is an extension function and returns `raw`.
- `gms_raw.cast_from_binary_double(n in binary_double, endianess in integer default 1)` is an extension function and returns `raw`.
- `gms_raw.cast_from_binary_float(n in float, endianess in integer default 1)` is an extension function and returns `raw`.
- `gms_raw.cast_from_binary_integer(n in bigint, endianess in integer default 1)` is an extension function and returns `raw`.
- `gms_raw.cast_from_number(n in number)` is an extension function and returns `raw`.
- `gms_raw.cast_to_binary_double(r in raw, endianess in integer default 1)` is an extension function and returns `binary_double`.
- `gms_raw.cast_to_binary_float(r in raw, endianess in integer default 1)` is an extension function and returns `float4`.
- `gms_raw.cast_to_binary_integer(r in raw, endianess in integer default 1)` is an extension function and returns `binary_integer`.
- `gms_raw.cast_to_number(r in raw)` is an extension function and returns `number`.
- `gms_raw.cast_to_nvarchar2(r in raw)` is an extension function and returns `nvarchar2`.
- `gms_raw.cast_to_raw(c in varchar2)` is an extension function and returns `raw`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
