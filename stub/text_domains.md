## Usage

Sources:

- [Official README](https://github.com/cahutton/text_domains/blob/a4e1aa7463e07ffdf1ce9f2d380f054bf1af1b5c/README.md)
- [Official extension SQL](https://github.com/cahutton/text_domains/blob/a4e1aa7463e07ffdf1ce9f2d380f054bf1af1b5c/text_domains--1.0.sql)
- [Official control file](https://github.com/cahutton/text_domains/blob/a4e1aa7463e07ffdf1ce9f2d380f054bf1af1b5c/text_domains.control)

`text_domains` version `1.0` defines 68 reusable PostgreSQL domains for nonempty text and fixed-shape alphanumeric, alphabetic, uppercase, lowercase, and digit strings. These domains centralize simple validation rules while retaining normal `text` storage and operators.

### Core Workflow

Install the extension and use the domains directly in table definitions:

```sql
CREATE EXTENSION text_domains;

CREATE TABLE country_codes (
  code alpha2 NOT NULL PRIMARY KEY,
  display_name nonempty_text NOT NULL,
  numeric_code digit3
);

INSERT INTO country_codes VALUES ('US', 'United States', '840');
```

A value that does not satisfy the domain's check constraint is rejected at assignment time.

### Domain Families

- `nonempty_text` rejects the empty string.
- `alnum` and `nonempty_alnum` accept zero-or-more and one-or-more alphanumeric characters; `alnum1` through selected fixed widths such as `alnum10`, `alnum16`, and `alnum22` require exactly that many characters.
- `alpha`, `upper`, `lower`, and `digit` provide equivalent POSIX character-class families, including nonempty and fixed-width variants.
- All patterned domains use `COLLATE "C"` and `SIMILAR TO` checks in the extension SQL.

### Operational Caveats

Domain checks accept NULL unless the table column separately declares `NOT NULL`; even `nonempty_text` does not make a column non-null. The available fixed widths are an explicit, non-contiguous set in the SQL file, so inspect it rather than assuming every length exists. Character-class and collation behavior should be tested with the deployment encoding and representative non-ASCII input. Domain constraints are intentionally simple and do not normalize case, trim whitespace, or provide application-specific formats. Changing a shared domain constraint affects every dependent column and may require validating existing data first.
