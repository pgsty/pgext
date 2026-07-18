## Usage

Sources:

- [Upstream README](https://github.com/robe2/postgis_letters/blob/a0dd76d45570ca875efe695c45be4dc6ff903bd2/README.md)
- [Version 0.0.5 control file](https://github.com/robe2/postgis_letters/blob/a0dd76d45570ca875efe695c45be4dc6ff903bd2/postgis_letters.control.in)
- [Geometry function definition](https://github.com/robe2/postgis_letters/blob/a0dd76d45570ca875efe695c45be4dc6ff903bd2/sql_bits/postgis_letters.sql.in)
- [Bundled font data](https://github.com/robe2/postgis_letters/blob/a0dd76d45570ca875efe695c45be4dc6ff903bd2/sql_bits/font_set.sql.in)

postgis_letters 0.0.5 is a pure-SQL PostGIS add-on that turns text into multipolygon glyph geometry. Installation creates a font_set table and loads the bundled Kankin vector font.

### Render text as geometry

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_letters;
SELECT ST_LettersAsGeometry('Circle', 'kankin');
```

The optional SRID, text height, and start position arguments place and scale the result. Supply sizes and positions in the coordinate units of the target geometry.

### Caveats

- Assigning an SRID labels the output coordinates; it does not project the glyphs. Transform a suitable planar result when geographic or projected accuracy matters.
- Characters absent from the selected font produce no glyph. Validate the supported character set before storing results.
- Complex text creates large multipolygon values and can be expensive to store, render, index, or intersect.
- The SQL function refers to font_set without schema qualification. Install in a controlled schema and use a controlled search path to avoid resolving an unintended table.
- Extension dumps preserve only font_set rows whose packaged_font flag is false; bundled rows are replaced by install or upgrade scripts. Back up custom fonts and set the flag deliberately.
- The README targets PostgreSQL 9.1 and PostGIS 2.0 or later, but upstream publishes no current compatibility matrix or automated tests.
