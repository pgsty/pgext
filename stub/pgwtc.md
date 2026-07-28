## Usage

Sources:

- [Official upstream README](https://github.com/gciolli/pgwtc/blob/4ecaabd7ebe48aae2e627fed470fd52e4b7281fd/README.md)
- [Official extension control file (pgwtc.control)](https://github.com/gciolli/pgwtc/blob/4ecaabd7ebe48aae2e627fed470fd52e4b7281fd/pgwtc.control)
- [Official extension SQL (pgwtc--0.1.sql)](https://github.com/gciolli/pgwtc/blob/4ecaabd7ebe48aae2e627fed470fd52e4b7281fd/pgwtc--0.1.sql)

`pgwtc` — This directory contains two extensions for PostgreSQL 17:. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pgwtc;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `lilypond_book_subjects()` is an extension function and returns `SETOF`.
- `lilypond_voice(src text , vox vox , start_id int DEFAULT NULL , max_count bigint DEFAULT NULL , add_key boolean DEFAULT false , add_time boolean DEFAULT false , add_clef boolean DEFAULT false , add_rest boolean DEFAULT false)` is an extension function and returns `text`.
- `answer` is an extension-defined type.
- `notes_pretty` is an extension-defined view.
- `subject_occurrences` is an extension-defined view.
- `subject_occurrences_pretty` is an extension-defined view.
- `subject_patterns` is an extension-defined view.
- `subjects` is an extension-defined view.
- `subjects_pretty` is an extension-defined view.
- `metadata` is a table installed or managed by the extension.
- `notes` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- Install the confirmed extension dependencies first: `ly2pg`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
