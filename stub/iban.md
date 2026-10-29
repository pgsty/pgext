## Usage

Sources:

- [Official README](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/6fa399864a71eb681ee2814e86686100ecfe095f/README.md)
- [SQL 1.2](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/6fa399864a71eb681ee2814e86686100ecfe095f/iban--1.2.sql)
- [Upgrade from 1.1](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/6fa399864a71eb681ee2814e86686100ecfe095f/iban--1.1--1.2.sql)
- [Type and validator implementation](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/6fa399864a71eb681ee2814e86686100ecfe095f/iban.cpp)

`iban` SQL 1.2, distributed as upstream 1.2.0, adds a validated International Bank Account Number type and `iban_validate(text)`. Country-specific BBAN structure and a mod-97 checksum reject malformed values; successful validation does not prove account existence, ownership or payment-network reachability. Upstream supports PostgreSQL 14–18.

### Store and Compare Values

```sql
CREATE EXTENSION iban;
CREATE TABLE beneficiaries (
    beneficiary_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    account_iban iban NOT NULL
);
INSERT INTO beneficiaries (account_iban)
VALUES ('gb82 west 1234 5698 7654 32');
SELECT account_iban, iban_validate(account_iban::text)
FROM beneficiaries ORDER BY account_iban;
CREATE INDEX beneficiaries_iban_idx ON beneficiaries (account_iban);
```

Text input is uppercased and plain spaces are removed, so the stored value in this example is GB82WEST12345698765432. Tabs, newlines and other whitespace are not removed. Input longer than 64 bytes is rejected; error output repeats at most 40 input bytes. Binary input is now validated and normalized by the same rules.

### Objects

- `iban_validate(text)` returns a validation result without storing a value.
- `iban` supports casts to text types and comparison operators. `iban_cmp(iban, iban)` provides byte-order comparison.
- Default btree and hash operator classes support ordering, equality, joins, grouping and keys; `iban_hash(iban)` and `iban_hash_extended(iban, bigint)` support hash operations.

### Upgrade and Validation Boundaries

Install matching files, ensure old-library sessions are replaced, then upgrade SQL and normalize rows stored by older versions:

```sql
ALTER EXTENSION iban UPDATE TO '1.2';
UPDATE beneficiaries SET account_iban = account_iban::text::iban;
```

Review existing values and uniqueness collisions before rewriting production data or creating unique indexes: older rows may retain lowercase or noncanonical representations. Country rules are compiled into the extension and may lag registry changes; verify the countries required by the application against this exact release.
