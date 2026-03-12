

## Usage

> [asn1oid: ASN.1 OID data type for PostgreSQL](https://github.com/df7cb/pgsql-asn1oid)

The `asn1oid` extension provides a data type for storing and comparing ASN.1 Object Identifiers (OIDs).

```sql
CREATE EXTENSION asn1oid;

SELECT '1.3.6.1.4.1'::asn1oid;
   asn1oid
─────────────
 1.3.6.1.4.1
```

### Data Type

The `asn1oid` type stores ASN.1 OID values in dotted-decimal notation (e.g., `1.3.6.1.4.1.311`). These are hierarchical identifiers used in SNMP, LDAP, X.509 certificates, and other standards.

### Operators

Standard comparison operators are supported for ordering and equality: `=`, `<>`, `<`, `>`, `<=`, `>=`.

### Casts

The type supports casting to and from `text`.
