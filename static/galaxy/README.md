# pgext.cloud

Static full-window visualization report for the PostgreSQL extension galaxy.

## Source

- Database: `data`
- Table: `pgext.universe`
- Default unit: one node per extension, using `name`
- Package mode: one node per package, using `lead = true` rows deduplicated by `pkg`
- Star field: `star_cnt`
- Watch/fork fields: `watch_cnt`, `fork_cnt`

The report keeps both a reusable JSON file and a JS wrapper. The HTML loads the JS wrapper so it works from `file://` without a local server.

## Dimensions

The pgext.cloud report supports 32 color dimensions:

- Category
- Language
- License
- Extension Type
- Cloud Vendor
- Kernel Fork
- Repository
- Catalog Status
- Package Availability
- Binary Repository
- PG Major Coverage
- Star Tier
- Watch Tier
- Fork Tier
- Last Updated Time
- Last Release Time
- Contrib
- Lead
- Has Binary
- Has Library
- Need DDL
- Need Load
- Trusted
- Relocatable
- Binary Matrix
- Package Coverage
- Dependency Load
- Dependent Gravity
- Schema Model
- PGRX Version
- PGXN Presence
- PGEXT Build Support

## Files

- `index.html`: report entry point
- `assets/app.css`: report styling
- `assets/app.js`: canvas rendering and interactions
- `assets/galaxy.webp`: optimized WebP background from `~/galaxy.png`
- `data/pgext-universe-data.json`: reusable static data
- `data/pgext-universe-data.js`: direct-open data wrapper
- `data/pgext-universe.csv`: complete CSV snapshot from `pgext.universe`
- `assets/pgsql.svg`: copied PostgreSQL logo source; the page also embeds this SVG inline in `index.html`

## Regenerate Data

```bash
psql data -At -f tmp/data/sql/export_pgext_universe.sql > html/pgext-universe/data/pgext-universe-data.json
{ printf 'window.PGEXT_UNIVERSE_DATA = '; cat html/pgext-universe/data/pgext-universe-data.json; printf ';\n'; } > html/pgext-universe/data/pgext-universe-data.js
psql data -c "\\copy (select * from pgext.universe order by id, name) to '/Users/vonng/pgsty/e4e/html/pgext-universe/data/pgext-universe.csv' csv header"
```
