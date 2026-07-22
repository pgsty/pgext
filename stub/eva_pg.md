## Usage

Sources:

- [Extension control file](https://github.com/eva-ics/eva4/blob/main/eva_pg/eva_pg/eva_pg.control)
- [pgrx implementation](https://github.com/eva-ics/eva4/blob/main/eva_pg/eva_pg/src/lib.rs)
- [Extension package metadata](https://github.com/eva-ics/eva4/blob/main/eva_pg/eva_pg/Cargo.toml)

`eva_pg` 0.0.1 is the PostgreSQL subproject of the EVA ICS monorepo. It exposes one pgrx function, `oid_match`, for matching EVA object identifiers against EVA mask syntax; it does not import the wider platform's service APIs into PostgreSQL.

### Core Workflow

Install the extension as a superuser, then compare an EVA identifier in `kind:path/segments` form with a mask:

```sql
CREATE EXTENSION eva_pg;

SELECT oid_match('sensor:plant/line1/temp', 'sensor:plant/+/temp');
SELECT oid_match('unit:plant/line1/pump', '+:plant/#');
```

The whole mask `*` or `#` matches immediately. Otherwise, the identifier and mask must contain a colon separating kind from path or the function raises an error.

### Matching Rules

- A mask kind of `+` or `?` matches any one kind; any other kind must match exactly.
- A path segment of `+` or `?` matches one segment.
- A path segment of `*` or `#` accepts the remaining path immediately, so later mask segments are ignored.
- Without a remainder wildcard, identifier and mask paths must have the same number of segments.

### Operational Notes

The control file requires superuser installation and is not relocatable. The Rust package currently declares pgrx features for PostgreSQL 13 through 18, with PostgreSQL 16 as its default build feature; build it for the exact target major. The subproject contains no independent SQL regression suite, so test malformed identifiers, empty segments, wildcard placement, and application-specific masks before relying on it in authorization or filtering logic.
