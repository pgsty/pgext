## Usage

Sources:

- [README](https://github.com/giuliosmall/pg_typesafe/blob/v0.0.1/README.md)
- [Control file](https://github.com/giuliosmall/pg_typesafe/blob/v0.0.1/typesafe.control)
- [typesafe--0.0.1.sql](https://github.com/giuliosmall/pg_typesafe/blob/v0.0.1/typesafe--0.0.1.sql)

`typesafe` calls TypeSafe AI for classification, probability-style answers, and scoring from SQL. This pre-alpha client transmits the supplied text to an external API and requires API credentials and libcurl 7.61 or newer.

### Core Workflow

Install the library and create the extension. Supply the key through the server environment variable `TYPESAFE_API_KEY` before starting PostgreSQL, through `TYPESAFE_API_KEY_FILE`, or through a server-readable file selected by the superuser-only `typesafe.api_key_file` setting.

```sql
CREATE EXTENSION typesafe;
SELECT typesafe_noul('My payouts have been failing for three days.',
                     'Does this convey urgency?');
SELECT * FROM typesafe_detect_many(
  ARRAY['Service is unavailable', 'Thank you for your help'], 'Is this urgent?');
```

### API and Privileges

`typesafe_noul` returns a probability-style result. The detection, classification, scoring, and question helpers have scalar and batch forms; batch calls reduce repeated HTTP work. `EXECUTE` is revoked from `PUBLIC`; grant only the exact function signatures the application needs.

The v0.0.1 release now documents PostgreSQL 15–18 support. The extension is relocatable and does not require preload. The session credential setting `typesafe.api_key` can appear in SQL logs; prefer the environment or protected file. Network latency, service limits, and API failures become query concerns, so avoid treating remote model output as a database integrity guarantee.
