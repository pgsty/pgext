## Usage

Sources:

- [Official README](https://github.com/ttfkam/pg_geekspeak/blob/8d176e395f7f8a98a3c4dd2d48a1ee1e359b4e47/README.md)
- [Official extension SQL](https://github.com/ttfkam/pg_geekspeak/blob/8d176e395f7f8a98a3c4dd2d48a1ee1e359b4e47/geekspeak--1.0.0.sql)
- [Official extension control file](https://github.com/ttfkam/pg_geekspeak/blob/8d176e395f7f8a98a3c4dd2d48a1ee1e359b4e47/geekspeak.control)

`geekspeak` installs the complete application schema used by the GeekSpeak podcast website: episodes, people, sessions, headlines, program segments, authentication helpers, search, and filesystem-backed media listings. It is application-specific data definition, not a reusable podcast toolkit, and installation creates many globally named objects in the selected schema.

### Prerequisites and Installation

Version `1.0.0` requires `btree_gist`, `isn`, `multicorn`, `pgcrypto`, and `plpgsql`. Its SQL grants to two pre-existing roles and defines Multicorn filesystem tables, so create or map those roles deliberately before installation:

```sql
CREATE ROLE geekspeak_org NOLOGIN;
CREATE ROLE geekspeak_audit NOLOGIN;
CREATE EXTENSION geekspeak CASCADE;
```

The schema hard-codes `/var/www/geekspeak.org/media` as the root for `episode_media_fdt` and `episode_misc_fdt`. The PostgreSQL service account and Multicorn Python environment must be able to access that directory before refreshing the materialized media view.

### Main Object Groups

- `episodes`, `participants`, `people`, `locations`, `headlines`, `bits`, and `bit_templates` hold show and editorial data.
- `passwords`, `acls`, `sessions`, and `active_sessions` implement the site's password, role, and persistent-session model.
- `episode_media_fdt`, `episode_misc_fdt`, and `episode_media` expose and cache files through `multicorn.fsfdw.FilesystemFdw`.
- `login()`, `logout()`, `register()`, `recover()`, `confirm()`, and `update_password()` implement application authentication flows.
- `episode_as_json()`, `headlines_as_json()`, `bits_as_json()`, and related helpers shape API output.
- `geekspeak.session_duration` controls the session lifetime and defaults to one month.

### Example Read Path

Once application data and media have been provisioned, read through the packaged helper functions and views:

```sql
SELECT episode_num(current_date) AS current_episode;
SELECT * FROM active_sessions;
REFRESH MATERIALIZED VIEW episode_media;
SELECT * FROM episode_media ORDER BY num, filename;
```

### Operational Boundaries

Review the entire SQL before installing into a nonempty schema: it creates tables, triggers, a server, foreign tables, a materialized view, functions with fixed ownership, grants, absolute paths, and historical application assumptions. Password hashing, session expiry, IP tracking, ACLs, filesystem exposure, and API functions form a security boundary; do not reuse them without a contemporary security review. Test the old PostgreSQL 10-era SQL and Multicorn wrapper against the exact target versions, and plan independent migrations, backups, media permissions, and uninstall behavior for this application schema.
