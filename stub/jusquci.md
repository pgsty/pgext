## Usage

Sources:

- [Upstream README](https://github.com/thjbdvlt/jusquci/blob/9238c91a290c8b0e1d036a53a148caf56ac1c2b6/README.md)
- [Extension control file](https://github.com/thjbdvlt/jusquci/blob/9238c91a290c8b0e1d036a53a148caf56ac1c2b6/postgresql/jusquci.control)
- [Text-search objects](https://github.com/thjbdvlt/jusquci/blob/9238c91a290c8b0e1d036a53a148caf56ac1c2b6/postgresql/jusquci--1.0.sql)
- [PostgreSQL build and stopword installation](https://github.com/thjbdvlt/jusquci/blob/9238c91a290c8b0e1d036a53a148caf56ac1c2b6/postgresql/makefile)

`jusquci` supplies a PostgreSQL text-search parser for contemporary French. Its tokenizer recognizes elisions, compound and inclusive forms, punctuation runs, URLs, emoticons, and related modern text patterns. Installing the extension creates the `jusquci` parser, the `jusquci_stop` dictionary, and the `jusquci` text-search configuration.

### Parse French text

```sql
CREATE EXTENSION jusquci;

SELECT to_tsvector(
  'jusquci',
  'le quotidien,s''invente-t-il par mille manière de braconner???'
);

SELECT *
FROM ts_parse('jusquci', 'peut-être--là lecteur-rice-x-s');
```

The separate `french_jusquci.stop` file must be copied into PostgreSQL's text-search data directory. The upstream build exposes this as `make install install_stop`; without that file, creation of `jusquci_stop` can fail.

Version `1.0` is implemented as a C parser and is marked superuser-only for extension creation. Upstream says it has only been tested on Debian Linux and PostgreSQL 16, so test token types, stopword behavior, search ranking, encoding, and target-version compatibility before replacing an existing French configuration.
