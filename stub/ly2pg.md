## Usage

Sources:

- [Official upstream README](https://github.com/gciolli/pgwtc/blob/4ecaabd7ebe48aae2e627fed470fd52e4b7281fd/README.md)
- [Official extension control file (ly2pg.control)](https://github.com/gciolli/pgwtc/blob/4ecaabd7ebe48aae2e627fed470fd52e4b7281fd/ly2pg.control)
- [Official extension SQL (ly2pg--0.1.sql)](https://github.com/gciolli/pgwtc/blob/4ecaabd7ebe48aae2e627fed470fd52e4b7281fd/ly2pg--0.1.sql)

`ly2pg` — This directory contains two extensions for PostgreSQL 17:. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION ly2pg;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `clavis2lilypond(IN clavis , OUT key text , OUT is_maior boolean)` is an extension function and returns `record`.
- `clavis2nota(IN clavis , nota OUT nota , is_maior OUT boolean)` is an extension function and returns `record`.
- `duration2ticks(text)` is an extension function and returns `int`.
- `int_error(text)` is an extension function and returns `int`.
- `lilypond(nota)` is an extension function and returns `text`.
- `lilypond(nota, text)` is an extension function and returns `text`.
- `lilypond(nota, text[])` is an extension function and returns `text`.
- `lilypond_book_subjects()` is an extension function and returns `SETOF text`.
- `lilypond_voice(src text , vox vox , start_id int DEFAULT NULL , max_count bigint DEFAULT NULL , add_key boolean DEFAULT false , add_time boolean DEFAULT false , add_clef boolean DEFAULT false , add_rest boolean DEFAULT false)` is an extension function and returns `text`.
- `nota(jsonb)` is an extension function and returns `nota`.
- `nota_add(nota, int)` is an extension function and returns `nota`.
- `nota_at_clavis(nota, clavis)` is an extension function and returns `int`.
- `tempo(text)` is an extension function and returns `tempo`.
- `tempo2text(tempo)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
