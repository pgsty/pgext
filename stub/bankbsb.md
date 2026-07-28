## Usage

Sources:

- [Official upstream README](https://github.com/timwmillard/bankbsb/blob/fa62a91aa3c3845527e29449fd2feb328427f5c4/README.md)
- [Official extension control file (bankbsb.control)](https://github.com/timwmillard/bankbsb/blob/fa62a91aa3c3845527e29449fd2feb328427f5c4/bankbsb.control)
- [Official extension SQL (bankbsb--0.0.1.sql)](https://github.com/timwmillard/bankbsb/blob/fa62a91aa3c3845527e29449fd2feb328427f5c4/bankbsb--0.0.1.sql)

`bankbsb` — A PostgreSQL Extension to add an Australian banking BSB number type. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION bankbsb;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `bankbsb_cmp(bsb, bsb)` is an extension function and returns `integer`.
- `bankbsb_eq(bsb, bsb)` is an extension function and returns `boolean`.
- `bankbsb_ge(bsb, bsb)` is an extension function and returns `boolean`.
- `bankbsb_gt(bsb, bsb)` is an extension function and returns `boolean`.
- `bankbsb_in(cstring)` is an extension function and returns `bsb`.
- `bankbsb_le(bsb, bsb)` is an extension function and returns `boolean`.
- `bankbsb_lt(bsb, bsb)` is an extension function and returns `boolean`.
- `bankbsb_ne(bsb, bsb)` is an extension function and returns `boolean`.
- `bankbsb_out(bsb)` is an extension function and returns `cstring`.
- `bsb` is an extension-defined type.
- `bankbsb_ops` is an extension-defined operator class.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
