## Usage

Sources:

- [Upstream telephone type documentation](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/doc/telephone.md)
- [Upstream README at the reviewed commit](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/README.md)
- [Extension control file](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/telephone.control)

`telephone` adds a normalized telephone-number type. It normalizes punctuation and spacing, maps uppercase letters to keypad digits for identity comparisons, and supports pause, confirmation, extension, and partial-number directives. Unique constraints therefore compare a number's identity rather than its presentation.

```sql
CREATE EXTENSION telephone;

CREATE TABLE contact_phone (
  phone telephone PRIMARY KEY,
  label text NOT NULL
);

INSERT INTO contact_phone VALUES ('1 (800) 555-01AZ', 'support');
SELECT * FROM contact_phone
WHERE phone = '18005550129'::telephone;
```

Numbers beginning with `+` are checked against the extension's dialing-plan data and can expose domestic, E.123, service-type, and geographic-segment helpers. Digits mode remains available when a plan is unsupported.

### Caveats

Version 0.0.1 is labeled unstable, and international dialing-plan coverage is partial; do not treat it as a complete global validation library. Although the control file says relocatable, installation SQL hard-codes several helper functions in `public`. Review schema placement and test the exact calling plans your application accepts.
