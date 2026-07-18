## Usage

Sources:

- [Upstream WIP documentation](https://github.com/postgrespro/pg_pageprep/blob/981b279f1af697adb954bcc6de87297371c0bcaf/README.md)
- [Extension control file](https://github.com/postgrespro/pg_pageprep/blob/981b279f1af697adb954bcc6de87297371c0bcaf/pg_pageprep.control)

`pg_pageprep` is a work-in-progress migration aid that prepares ordinary heap pages for PostgresPro Enterprise's 64-bit transaction-ID page format. The reviewed code targets migrations from PostgreSQL or PostgresPro Standard 10/11 to PostgresPro Enterprise 11.

```conf
shared_preload_libraries = 'pg_pageprep'
pg_pageprep.database = 'postgres'
pg_pageprep.role = 'postgres'
```

After restart, create the extension in every database and manage workers with the supplied script or functions. Processing may change relation fillfactor to 90; after upgrade, stop workers and restore original fillfactors. This is invasive, version-specific migration software: rehearse on a restored copy, inventory unsupported relations, monitor bloat and I/O, take a verified backup, and use only with the matching PostgresPro migration procedure.
