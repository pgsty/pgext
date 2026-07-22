## Usage

Sources:

- [Official README](https://github.com/lacanoid/pgwiki/blob/dbcc433acd6fb2dacc78d0621103a8e2d5c187ec/README.md)
- [Official extension SQL](https://github.com/lacanoid/pgwiki/blob/dbcc433acd6fb2dacc78d0621103a8e2d5c187ec/wiki.sql)
- [Official extension control file](https://github.com/lacanoid/pgwiki/blob/dbcc433acd6fb2dacc78d0621103a8e2d5c187ec/wiki.control)

`wiki` implements a versioned wiki data model and SQL API inside PostgreSQL, with optional PL/Perl helpers for talking to an external MediaWiki instance. The reviewed `0.1` source is abandoned and targets old PostgreSQL releases; use it only after a security and compatibility audit.

### Prerequisites and Setup

The control file fixes the extension in the `wiki` schema, requires both `plperl` and untrusted `plperlu`, and marks installation as superuser-only. Its SQL dump also assigns objects to a role named `wiki`; inspect the script and create or map the expected owner before installation.

```sql
CREATE EXTENSION plperl;
CREATE EXTENSION plperlu;
CREATE EXTENSION wiki;
```

Because `plperlu` can access the host operating system, only a trusted database administrator should install or grant access to this extension.

### Core Workflow

Create or replace a page with `wiki.page_set`, then read either its current body or its structured metadata:

```sql
SELECT wiki.page_set('Welcome', 'Hello from PostgreSQL');

SELECT wiki.page_get('Welcome');
SELECT wiki.page_get_data('Welcome');
```

Each body change creates a revision. The SQL also supplies overloads for text, XML, and JSON bodies and functions for retrieving a page by its node or revision identifier.

### Important Objects

- `wiki.page_set(...)` writes a page body and returns the new revision identifier.
- `wiki.page_get(...)` returns the current or selected page body.
- `wiki.page_get_data(...)` returns page metadata as JSON.
- `wiki.node`, `wiki.body`, `wiki.revision`, and `wiki.links` hold pages, content revisions, and links.
- `wiki.foreign_page_get(...)` and `wiki.foreign_pages_search(...)` use `plperlu` and the Perl `MediaWiki::API` module to access a configured remote wiki.
- `wiki.html_tidy(...)` depends on the Perl `HTML::Tidy` module.

### Operational Caveats

Remote-wiki credentials are stored in the extension's `wiki.wiki` table, and the untrusted Perl functions can make network and host-level calls. Review table and function grants, revoke broad defaults, protect stored credentials, and restrict outbound network access. The upstream repository was last updated in 2019 and its documented compatibility stops at PostgreSQL `12`; no support for current PostgreSQL versions is established. Pin the exact source, test install/upgrade/restore paths, and prefer a maintained wiki or application-layer integration for new deployments.
