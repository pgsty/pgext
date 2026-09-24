## Usage

Sources:

- [README.md](https://github.com/pt-immer/kamu-money-pg/blob/a83ceeaeb6d5c97cb6071c23c47bec172306842a/README.md)
- [DESIGN.md](https://github.com/pt-immer/kamu-money-pg/blob/a83ceeaeb6d5c97cb6071c23c47bec172306842a/DESIGN.md)
- [kamu-money-pg/kmoney.control](https://github.com/pt-immer/kamu-money-pg/blob/a83ceeaeb6d5c97cb6071c23c47bec172306842a/kamu-money-pg/kmoney.control)
- [kamu-money-pg/Cargo.toml](https://github.com/pt-immer/kamu-money-pg/blob/a83ceeaeb6d5c97cb6071c23c47bec172306842a/kamu-money-pg/Cargo.toml)

`kmoney` supplies fixed-width, scale-18 monetary types whose SQL type encodes an ISO 4217 currency, plus a mixed-currency carrier for interchange.

### Enablement

Release 0.2.1 exposes PostgreSQL-major features for 15–18. Install the matching pgrx build, then create the non-relocatable, untrusted extension as a superuser:

```sql
CREATE EXTENSION kmoney;
```

No preload or restart is required. Upstream builds through a pgrx fork with feature-gated YugabyteDB adaptations; validate the exact server/kernel artifact before deployment.

### Currency-Specific Types

Each currency has a distinct 16-byte type such as `kmoney_usd` or `kmoney_idr`. The parser accepts either an untagged amount or the matching currency tag; cross-currency operators do not exist.

```sql
SELECT '10.50'::kmoney_usd;
SELECT ('1.25'::kmoney_usd + '2.75'::kmoney_usd)::text;
SELECT '1.00'::kmoney_usd + '1.00'::kmoney_idr;
```

`kmoney_mixed` stores the currency code with the amount for heterogeneous values. It intentionally has no `sum` aggregate; aggregate a currency-specific column instead.

### Division, Allocation, and Indexing

Division returns quotient and residue so rounding loss is explicit. Allocation conserves the original amount across integer weights.

```sql
SELECT quotient, residue
FROM kmoney_usd_div('10.00'::kmoney_usd, 3, 'half_even');

SELECT unnest(kmoney_usd_allocate('10.00'::kmoney_usd, ARRAY[1,1,1]));
```

Version 0.2.1 deliberately ships no B-tree or hash operator classes, so these custom types cannot back ordinary B-tree/hash indexes. Values are bounded to the magnitude of `numeric(36,18)` and reject excess precision instead of silently rounding.

The extension now lives in the dedicated kamu-money-pg repository. The 0.2.1 workspace uses pgrx 0.19.2 through the documented YugabyteDB fork; the PostgreSQL 15–18 type and arithmetic contract remains unchanged.
