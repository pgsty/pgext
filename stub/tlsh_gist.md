## Usage

Sources:

- [Official upstream README](https://github.com/jinyyu/tlsh_gist/blob/8b028591e961a9ca8d561b34b02817c15f01bf86/README.MD)
- [Official extension control file (tlsh_gist.control)](https://github.com/jinyyu/tlsh_gist/blob/8b028591e961a9ca8d561b34b02817c15f01bf86/tlsh_gist.control)
- [Official extension SQL (tlsh_gist--1.0.sql)](https://github.com/jinyyu/tlsh_gist/blob/8b028591e961a9ca8d561b34b02817c15f01bf86/tlsh_gist--1.0.sql)

`tlsh_gist` — tlsh_gist --------- A PostgreSQL plugin for tlsh hash, whitch is a fuzzy matching program and library, Similar files will have similar hash values which allows for the detection of similar objects by comparing their hash values TLSH. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION tlsh_gist;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gtlsh_in(cstring)` is an extension function and returns `gtlsh`.
- `gtlsh_out(gtlsh)` is an extension function and returns `cstring`.
- `tlsh_compress(internal)` is an extension function and returns `internal`.
- `tlsh_consistent(internal, tlsh, smallint, oid, internal)` is an extension function.
- `tlsh_decompress(internal)` is an extension function and returns `internal`.
- `tlsh_dist(tlsh,tlsh)` is an extension function and returns `int4`.
- `tlsh_distance(internal, tlsh, smallint, oid, internal)` is an extension function and returns `float8`.
- `tlsh_equal(tlsh, tlsh)` is an extension function.
- `tlsh_in(cstring)` is an extension function and returns `tlsh`.
- `tlsh_mean(tlsh, tlsh)` is an extension function and returns `tlsh`.
- `tlsh_out(tlsh)` is an extension function and returns `cstring`.
- `tlsh_penalty(internal, internal, internal)` is an extension function and returns `internal`.
- `tlsh_picksplit(internal, internal)` is an extension function and returns `internal`.
- `tlsh_same(gtlsh, gtlsh, internal)` is an extension function and returns `internal`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
