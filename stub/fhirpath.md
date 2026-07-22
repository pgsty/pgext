## Usage

Sources:

- [Official README](https://github.com/niquola/fhirpath-pg/blob/e43be1ab1394c7ae4d47f6e90985064a9bca1407/README.md)
- [Extension SQL definition](https://github.com/niquola/fhirpath-pg/blob/e43be1ab1394c7ae4d47f6e90985064a9bca1407/fhirpath--1.0.sql)
- [Extension control file](https://github.com/niquola/fhirpath-pg/blob/e43be1ab1394c7ae4d47f6e90985064a9bca1407/fhirpath.control)

`fhirpath` is a native C implementation of FHIRPath/FluentPath evaluation over PostgreSQL `jsonb`. Use it to compile a path into the `fhirpath` type and extract or normalize values from FHIR-shaped JSON documents. The upstream repository provides a small, low-level SQL surface rather than a complete application framework.

### Core Workflow

Install the extension, cast a path expression to `fhirpath`, and pass it with a JSON document to an extractor:

```sql
CREATE EXTENSION fhirpath;

SELECT fhirpath_extract(
    '{"resourceType":"Patient","active":true}'::jsonb,
    'resourceType'::fhirpath
);
```

The specialized accessors accept additional text arguments used by the implementation to select or normalize values. Verify their behavior against representative FHIR documents before embedding them in generated columns or indexes.

### User-Facing Objects

- `fhirpath`: variable-length type whose input parser compiles a path expression.
- `fhirpath_extract(jsonb, fhirpath)`: returns the selected value as `jsonb`.
- `fhirpath_as_string(jsonb, fhirpath, text)`: returns a selected value as `text`.
- `fhirpath_as_token(jsonb, fhirpath, text)` and `fhirpath_as_reference(jsonb, fhirpath, text)`: return `text[]` token or reference representations.
- `fhirpath_as_number(jsonb, fhirpath, text, text)`: returns `numeric`.
- `fhirpath_as_date(jsonb, fhirpath, text, text)` and `fhirpath_date_bound(text, text)`: return `timestamptz` values.
- `fhirpath_exists(jsonb, fhirpath, text)`: tests whether a selected value exists.

### Operational Notes

The control file marks `fhirpath` relocatable and does not require preloading. Its SQL functions are declared `STRICT IMMUTABLE`, which makes them usable in expression indexes when their inputs are immutable. Upstream publishes no formal release and the README is primarily development material; it does not claim complete conformance with a particular FHIRPath specification release. Treat version `1.0` as an experimental API, test path grammar and edge cases directly, and pin the reviewed source revision.
