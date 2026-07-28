## Usage

Sources:

- [Official upstream README](https://github.com/siose-innova/pg_wui/blob/8b5e9016999bfcfddbbefd8dd2f17dfe5a69e6c3/README.md)
- [Official extension control file (pg_wui.control)](https://github.com/siose-innova/pg_wui/blob/8b5e9016999bfcfddbbefd8dd2f17dfe5a69e6c3/src/pg_wui.control)

`pg_wui` — PostgreSQL extension for wildland urban interface data analysis using SIOSE databases. Use it for the corresponding spatial data or geospatial workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pg_wui;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.2`.
- Install the confirmed extension dependencies first: `postgis`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
