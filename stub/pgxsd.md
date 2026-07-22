## Usage

Sources:

- [Official version 1.0 SQL](https://github.com/johto/pgxsd/blob/ff89bb220cf47ab287b900b47f9eb645088b4422/pgxsd--1.0.sql)
- [Official validator implementation](https://github.com/johto/pgxsd/blob/ff89bb220cf47ab287b900b47f9eb645088b4422/pgxsd.c)
- [Official regression runner](https://github.com/johto/pgxsd/blob/ff89bb220cf47ab287b900b47f9eb645088b4422/regress/run.pl)
- [Official control file](https://github.com/johto/pgxsd/blob/ff89bb220cf47ab287b900b47f9eb645088b4422/pgxsd.control)

`pgxsd` validates PostgreSQL `xml` values against XML Schema documents stored in the database. It keeps XSD documents in the fixed `pgxsd` schema and resolves imported schema locations from that table rather than the network.

### Core Workflow

Store each schema under the exact location used as the validator's second argument or by XSD imports, then validate an XML document:

```sql
CREATE EXTENSION pgxsd;

INSERT INTO pgxsd.schemata (schema_location, document)
VALUES (
  'invoice.xsd',
  '<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
     <xs:element name="invoice" type="xs:integer"/>
   </xs:schema>'::xml
);

SELECT pgxsd.schema_validate(
  '<invoice>42</invoice>'::xml,
  'invoice.xsd'
);
```

Success returns `void`. An invalid document or schema raises an `invalid_xml_document` error; the upstream regression runner expects SQLSTATE `2200M` for validation failures.

### Objects

- `pgxsd.schemata(schema_location text PRIMARY KEY, document xml)`: XSD registry and import resolver.
- `pgxsd.schema_validate(doc xml, schemaname text) RETURNS void`: parse the named schema and validate `doc` with libxml2.

### Caveats

The extension requires PostgreSQL built with libxml/libxml2 schema support. Schema locations are application keys, not fetched URLs; every imported XSD location must exist exactly once in `pgxsd.schemata`.

Version `1.0` declares `schema_validate` as `IMMUTABLE` even though it reads the mutable `schemata` table through SPI. That volatility declaration is unsound: do not use it in stored generated expressions or indexes, and force re-evaluation after schema changes. XML/XSD parsing and entity expansion can also consume significant CPU/memory; restrict who can insert schemas or submit documents, set workload limits, and test maliciously nested input.
