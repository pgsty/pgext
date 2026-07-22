## Usage

Sources:

- [Official README](https://github.com/andy128k/pg-sphinx/blob/6a63ade3a0f7119474435d92b570d580638b4d90/README.org)
- [Extension SQL for 0.3](https://github.com/andy128k/pg-sphinx/blob/6a63ade3a0f7119474435d92b570d580638b4d90/sphinx--0.3.sql)
- [Extension control file](https://github.com/andy128k/pg-sphinx/blob/6a63ade3a0f7119474435d92b570d580638b4d90/sphinx.control)

`sphinx` 0.3 connects PostgreSQL functions to a Sphinx Search server. It supports search, document replacement and deletion, and snippet generation through Sphinx's SQL interface; the repository is archived, so use it only for maintaining a compatible legacy deployment.

### Core Workflow

Create the extension as a superuser, then set connection values in the extension-owned configuration table.

```sql
CREATE EXTENSION sphinx;

UPDATE sphinx_config SET value = '127.0.0.1' WHERE key = 'host';
UPDATE sphinx_config SET value = '9306' WHERE key = 'port';
UPDATE sphinx_config SET value = 'test_' WHERE key = 'prefix';

SELECT *
FROM sphinx_select('blog_posts', 'photo', '', '', 0, 20, '');
```

The optional prefix is prepended to index names, so the example addresses the Sphinx index named `test_blog_posts`.

### Main Objects

- `sphinx_select(...)` returns `(id, weight)` rows for a search.
- `sphinx_replace(index, id, data)` replaces a document; the data array alternates keys and values.
- `sphinx_delete(index, id)` removes a document.
- `sphinx_snippet(...)` and `sphinx_snippet_options(...)` generate highlighted excerpts.
- `sphinx_config` stores `host`, `port`, `username`, `password`, and `prefix` values and is included in extension dumps.

Version 0.3 grants all privileges on `sphinx_config` to public during installation. Revoke that grant and delegate only the minimum required access, especially if the table contains credentials. Calls depend synchronously on the external Sphinx service, and document updates are external side effects that do not participate in PostgreSQL transaction rollback.
