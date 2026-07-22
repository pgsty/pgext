## Usage

Sources:

- [Extension control file](https://github.com/niquola/fhir_jsonb/blob/b4d8339cdebc98a80df5b155ba97373e548bba85/fhir_jsonb.control)
- [Extension SQL](https://github.com/niquola/fhir_jsonb/blob/b4d8339cdebc98a80df5b155ba97373e548bba85/fhir_jsonb--1.0.sql)
- [Official regression query](https://github.com/niquola/fhir_jsonb/blob/b4d8339cdebc98a80df5b155ba97373e548bba85/sql/fhir_jsonb.sql)

`fhir_jsonb` is a destructive synthetic-data fixture for experimenting with PostgreSQL JSONB queries over FHIR-shaped documents. It is not a FHIR server, validator, schema implementation, or production data model; install it only in a disposable database or isolated schema.

### Core Workflow

Installation immediately drops and recreates four unqualified tables and populates them with generated data. Use a dedicated schema with no valuable objects of the same names:

```sql
CREATE SCHEMA fhir_fixture;
CREATE EXTENSION fhir_jsonb WITH SCHEMA fhir_fixture;

SELECT count(*) FROM fhir_fixture.patients;
SELECT doc #>> '{name,0,text}' AS patient_name
FROM fhir_fixture.patients
LIMIT 5;
```

The installation script creates `patients` with 100 rows, `encounters` with 1,000 rows, and `observations` and `conditions` with 10,000 rows each. Every table has an integer `id` and a `jsonb` column named `doc`.

### Objects and Querying

Besides the four fixture tables, the extension installs helper functions such as `patient_random_is_active()`, `encounter_random_status(integer)`, `observation_random_name(integer)`, and `random_array_element(varchar[])`. These are data generators, not stable healthcare APIs.

The official regression query installs `jsquery` separately and uses its `@@` operator. Neither the control file nor the installation script declares `jsquery` as a dependency, so install it independently only if you intend to run those historical examples. Ordinary JSONB operators such as `#>>` do not require it.

### Safety and Compatibility

The installation SQL executes `DROP TABLE IF EXISTS` for `patients`, `encounters`, `observations`, and `conditions` before creating the fixture. Existing tables with those names in the extension target schema can be destroyed. Never install it in an application schema or production database.

Generated documents resemble an old subset of FHIR resources but include fake personal data and provide no FHIR-version conformance, validation, terminology binding, or reference-integrity guarantees. The reviewed source dates from 2014 and has no maintained PostgreSQL compatibility matrix; test it on the exact server version and treat the resulting objects as disposable benchmark data.
