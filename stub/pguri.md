## Usage

Sources:

- [Pinned upstream README](https://github.com/dylex/pguri/blob/6b80c2f61b1de04faa3f2e6c1f768e9a8997db30/README)
- [Pinned version 1.0 installation SQL](https://github.com/dylex/pguri/blob/6b80c2f61b1de04faa3f2e6c1f768e9a8997db30/pguri--1.0.sql)
- [Pinned C parser implementation](https://github.com/dylex/pguri/blob/6b80c2f61b1de04faa3f2e6c1f768e9a8997db30/pguri.c)
- [Pinned extension control file](https://github.com/dylex/pguri/blob/6b80c2f61b1de04faa3f2e6c1f768e9a8997db30/pguri.control)

`pguri` version `1.0` adds parsed `domainname` and `uri` types for storing browsing-history and policy patterns. It separates a URI into scheme, authentication, domain, port, path, query, and fragment fields, provides domain/URI containment, and includes a URI-aware text-search parser and configuration.

### Core Workflow

```sql
CREATE EXTENSION pguri;

SELECT ('https://user@example.com:8443/a/b?q=one#top'::uri).domain;
SELECT 'foo.com'::domainname @> 'www.foo.com'::domainname;
SELECT 'foo.com'::uri @> 'https://www.foo.com/path?q=one'::uri;

SELECT to_tsvector(
  'uri'::regconfig,
  'https://www.example.com/a%20path?q=one#top'
);
```

The URI parser accepts partial values, so a policy-like value such as `foo.com` can contain a more specific URI according to the extension's component and prefix rules.

### Important Objects

- `domainname`: a collatable text-like type with B-tree and hash operator classes, comparison/concatenation operators, `domainname_parts()`, `domainname_parents()`, and containment operators `@>` and `<@`.
- `uri`: a parsed type with fields `domain`, `port`, `path`, `query`, `auth`, `scheme`, and `fragment` plus containment operators `@>` and `<@`.
- `uri(text)` and `text(uri)`: explicit conversion functions for existing text columns; the corresponding casts are not installed by the reviewed SQL.
- Text-search parser, template, dictionary, and configuration `uri`: tokenize URI components and percent-decode selected domain, path, query, and fragment tokens.

### Operational Notes

Containment is application policy logic, not RFC URI equivalence: it compares selected components and uses prefix matching for paths or queries. Test trailing slashes, escaped characters, default ports, case, internationalized names, IPv6 literals, and opaque schemes against application expectations, and retain the original text when lossless round trips matter.

The installation SQL contains an explicit catalog-update workaround that replaces the input/output functions of the already-created composite `uri` type. The C code also depends on PostgreSQL backend text-search and tuple internals. Treat this as version-sensitive legacy code, run round-trip and index tests on the exact PostgreSQL build, and do not assume binary or dump/restore compatibility across major releases without verification.
