## Usage

Sources:

- [Official PostgreSQL-IBAN README](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/31ee56f66b9df44e4fca083d83add588630af7a3/README.md)
- [Version 1.1 extension SQL](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/31ee56f66b9df44e4fca083d83add588630af7a3/iban--1.1.sql)
- [IBAN type and validator implementation](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/31ee56f66b9df44e4fca083d83add588630af7a3/iban.cpp)

`iban` adds an `iban` data type and `iban_validate(text)` for syntactic validation of International Bank Account Numbers. The input routine checks the country-specific BBAN shape and ISO 7064 mod-97 checksum, making the type useful for rejecting malformed values at the database boundary.

### Store validated values

```sql
CREATE EXTENSION iban;

CREATE TABLE beneficiaries (
    beneficiary_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    account_iban   iban NOT NULL
);

INSERT INTO beneficiaries (account_iban)
VALUES ('GB82WEST12345698765432');

SELECT account_iban, iban_validate(account_iban::text)
FROM beneficiaries;
```

Invalid text is rejected when cast or assigned to `iban`. To validate without raising an input error, call the boolean helper first:

```sql
SELECT iban_validate('GB82WEST12345698765432'); -- true
SELECT iban_validate('GB00INVALID');             -- false
```

The extension supplies implicit casts from `iban` to `text` and `varchar`, plus an assignment cast to `bpchar`.

### Validation boundaries

Validation uppercases a temporary copy, but the input function stores the original representation. Lowercase letters can therefore validate while remaining lowercase on output. Spaces and visual separators are not normalized; canonicalize application input before storing it if a single display form is required.

The country and BBAN rules are compiled into this release and may lag the current IBAN registry. A successful result confirms only the encoded structure and checksum—not that the account exists, is open, belongs to a named party, or is reachable by a payment network. The binary receive function does not repeat text-input validation, so only accept trusted PostgreSQL binary-protocol encoders for this custom type. Test the exact build and registry coverage required by the application before enforcing it as a long-lived domain rule.
