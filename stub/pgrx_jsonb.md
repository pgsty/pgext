## Usage

Sources:

- [Official README](https://github.com/jscheid/pgrx_jsonb/blob/777c01c72062b74127e562f179d2810c7f7a7450/README.md)
- [Extension control file](https://github.com/jscheid/pgrx_jsonb/blob/777c01c72062b74127e562f179d2810c7f7a7450/pgrx_jsonb.control)
- [Rust implementation](https://github.com/jscheid/pgrx_jsonb/blob/777c01c72062b74127e562f179d2810c7f7a7450/src/lib.rs)

pgrx_jsonb is a proof of concept for iterating PostgreSQL's internal JSONB representation from Rust/pgrx. Upstream explicitly says not to use it for anything important. Its three demonstration functions are useful only when studying extension internals.

### Core Workflow

Install it only in a disposable development database and compare the demonstration paths:

```sql
CREATE EXTENSION pgrx_jsonb;

SELECT jsonb_to_text('{"answer": 42}'::jsonb);
SELECT jsonb_test('{"answer": 42}'::jsonb);
SELECT jsonb_test2('{"answer": 42}'::jsonb);
```

The first function walks the low-level representation and serializes it, while the other two report whether the top-level value is an object through different code paths.

### Important Objects

- `jsonb_to_text` converts a JSONB value to text through the prototype iterator.
- `jsonb_test` checks for an object through the custom low-level iterator.
- `jsonb_test2` performs the comparison using pgrx's higher-level JSONB wrapper.

### Safety and Compatibility Notes

The iterator relies on PostgreSQL internal JSONB APIs and contains assertions, unwraps, panics, and unimplemented branches for some token types. Unsupported data can raise an error or terminate the backend process. The cataloged version is 0.0.0 and the project has no stable compatibility promise. Use PostgreSQL's built-in JSONB functions or maintained pgrx APIs in real extensions.
