## Usage

Sources:

- [README](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/README.md)
- [Control file](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/pg_money.control)
- [Cargo.toml](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/Cargo.toml)
- [src/lib.rs](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/src/lib.rs)
- [docs/SECURITY.md](https://github.com/RustedBytes/pg-money/blob/4dfa040a59ef4d53eebd274749ada42ee184be29/docs/SECURITY.md)

`pg_money` provides exact currency-aware amounts, arithmetic, rounding, and allocation on PostgreSQL 14–18. Its types are distinct from the locale-sensitive built-in money type.

### Core Workflow

```sql
CREATE EXTENSION pg_money;
CREATE TABLE invoices (id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                       total money_with_currency NOT NULL);
INSERT INTO invoices(total) VALUES ('USD 19.95'), (money_make(5.25, 'USD'));
SELECT sum(total), avg(total), money_format(sum(total)) FROM invoices;
SELECT money_split('USD 10.00', 3);
SELECT money_exchange('USD 100', 'EUR', 0.85);
```

### Objects and Semantics

`money_with_currency` stores a currency with an exact decimal amount; `money_minor` works with minor-unit amounts. `money_make`, `money_from_minor`, and `money_minor_make` construct values. `money_round` accepts an explicit rounding mode, and `money_split` allocates residual units while conserving the total.

Arithmetic and aggregation reject incompatible currencies instead of converting implicitly. Exchange rates are supplied explicitly or read from an application-owned table with `money_exchange_at`; the extension fetches no market rates. Check rounding and amount bounds at application boundaries.

### Operation

The 0.3.0 control file requires superuser installation and allows relocation. No preload is required. Treat exchange-rate tables and mutation privileges separately from read-only arithmetic; caller-supplied SQL data does not establish the correctness or freshness of a rate.
