## Usage

Sources:

- [Official README at the catalog revision](https://github.com/ab1smo/pg_normalize_email/blob/709ccbe4cc2733adeb629bc4873d3f3afda868c4/README.md)
- [Extension SQL at the catalog revision](https://github.com/ab1smo/pg_normalize_email/blob/709ccbe4cc2733adeb629bc4873d3f3afda868c4/sql/pg_normalize_email--1.0.9.sql)
- [PGXN distribution page](https://pgxn.org/dist/pg_normalize_email/)

`pg_normalize_email` 1.0.9 provides one immutable normalization function for selected email providers. It lowercases addresses, maps one provider domain, and removes provider-specific dots or plus tags. Use it only when deliberate account-deduplication policy accepts the risk of treating distinct input strings as one identity.

### Core Workflow

```sql
CREATE EXTENSION pg_normalize_email;

SELECT normalize_email('First.Last+news@googlemail.com');
-- firstlast@gmail.com

SELECT normalize_email('User+tag@outlook.com');
-- user@outlook.com

CREATE INDEX account_email_normalized_idx
ON account (normalize_email(email));
```

### Function Behavior

- `normalize_email(text)` requires exactly one at-sign and raises an exception for NULL or malformed splitting.
- All input is lowercased, and googlemail.com is rewritten to gmail.com.
- For gmail.com and live.com, dots in the local part and the first plus suffix are removed.
- For ya.ru, hotmail.com, and outlook.com, only the first plus suffix is removed. Other domains are only lowercased.

### Caveats

- The function is not a complete RFC email validator: it does not trim whitespace, validate nonempty parts, normalize internationalized domains, or account for quoted local parts.
- Provider behavior can change, and provider-specific normalization may merge addresses that an application considers distinct. Keep the original address and treat normalized uniqueness as an explicit product decision.
- Because the function is declared immutable, changing its rules later requires rebuilding dependent expression indexes and reviewing stored derived values.
