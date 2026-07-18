## Usage

Sources:

- [Upstream README](https://github.com/blm768/pg-libphonenumber/blob/753e2fa4be452620455a099aeda917648f2da70a/README.md)
- [Version 0.1.0 install SQL](https://github.com/blm768/pg-libphonenumber/blob/753e2fa4be452620455a099aeda917648f2da70a/sql/pg_libphonenumber--0.1.0.sql)
- [C++ SQL-callable implementation](https://github.com/blm768/pg-libphonenumber/blob/753e2fa4be452620455a099aeda917648f2da70a/src/pg_libphonenumber.cpp)
- [Packed representation and comparison](https://github.com/blm768/pg-libphonenumber/blob/753e2fa4be452620455a099aeda917648f2da70a/src/packed_phone_number.h)
- [PGXN release metadata](https://pgxn.org/dist/pg_libphonenumber/)

pg_libphonenumber 0.1.0 provides an eight-byte packed_phone_number type backed by Google's C++ libphonenumber, an explicit-region parser, international text output, a country-code accessor, and B-tree operators.

### Parse with an explicit region

```sql
CREATE EXTENSION pg_libphonenumber;
SELECT parse_packed_phone_number('03 7010 1234', 'AU')::text AS normalized;
SELECT phone_number_country_code(
  parse_packed_phone_number('+1 281 901 0011', 'US')
);
```

Use parse_packed_phone_number for national-format input. Direct packed_phone_number text input is hard-coded to the US region and is ambiguous for other countries.

### Caveats

- Upstream labels the extension alpha and partially implemented. PGXN has only the 0.1.0 release from 2017, and the repository was last changed in 2020.
- Parsing does not call libphonenumber's validity check. A syntactically parseable value is not evidence that a number is assigned, reachable, or appropriate for an account.
- Parsing and formatting functions are declared immutable even though results depend on the linked libphonenumber metadata version. Library upgrades can change normalized output and invalidate assumptions in stored expressions.
- The greater-than and greater-or-equal SQL operators incorrectly declare themselves as their own commutators. Planner rewrites can therefore change predicate meaning; do not rely on these operators or their B-tree class without repairing and retesting the SQL.
- Binary send and receive copy the native eight-byte object directly, without portable byte order or semantic validation. Avoid binary transfer across architectures or independently built versions.
- Ordering is an internal packed-value heuristic and is documented in source as potentially unintuitive. It is not linguistic, dialing, geographic, or numeric-string order.
