## Usage

Sources:

- [Official project note](https://github.com/bjeanes/pg-dataclips/blob/e1c49098762c6b35d4354fdcb0d3adc0d942c2ea/docs/dataclip.md)
- [Extension control file](https://github.com/bjeanes/pg-dataclips/blob/e1c49098762c6b35d4354fdcb0d3adc0d942c2ea/dataclips.control)
- [Extension SQL](https://github.com/bjeanes/pg-dataclips/blob/e1c49098762c6b35d4354fdcb0d3adc0d942c2ea/sql/dataclips.sql)
- [C implementation](https://github.com/bjeanes/pg-dataclips/blob/e1c49098762c6b35d4354fdcb0d3adc0d942c2ea/src/dataclips.c)

`dataclips` is an archived, unfinished prototype intended to query Heroku Dataclips through PostgreSQL. Upstream explicitly says it is not functional; it is not a usable Dataclips client or FDW.

### Declared Interface

The distribution declares extension name `dataclips` and one polymorphic record-returning function:

```sql
CREATE EXTENSION dataclips;

SELECT *
FROM dataclip('gqhkylscctbsgrohaythgdmlfcnn')
  AS t(name text, data json, active boolean);
```

The query is the upstream design example, not a working workflow. Because `dataclip(text)` returns `SETOF record`, every call must supply a column definition list appropriate to the expected Dataclip output.

### Implementation State

The C function contains comments for the intended HTTP fetch, CSV/JSON parsing, tuple descriptor, and row production, but none of that logic is implemented. It hard-codes three calls and returns a zero `Datum` instead of constructed tuples. No foreign data wrapper objects exist.

The packaging is internally inconsistent as well: the Makefile builds a library named `dataclips`, while `dataclips.control` requests `$libdir/dataclip`; the SQL script's guard message also says `CREATE EXTENSION dataclip` even though the canonical extension name is `dataclips`. A stock build can therefore fail before the incomplete function is reached.

### Caveats

Do not deploy this extension or expose `dataclip(text)` to users. It has no implemented authentication, response validation, error model, or stable row mapping, and the repository is archived. The source is useful only as historical prototype material; a real integration needs a maintained client or FDW with explicit credential, network, and schema handling.
