## Usage

Sources:

- [Official README](https://github.com/pramsey/pgsql-addressing-dictionary/blob/master/README.md)
- [Extension control file](https://github.com/pramsey/pgsql-addressing-dictionary/blob/master/addressing_dictionary.control)
- [Version 1.1 extension SQL](https://github.com/pramsey/pgsql-addressing-dictionary/blob/master/addressing_dictionary--1.1.sql)
- [PostgreSQL text-search dictionary documentation](https://www.postgresql.org/docs/current/textsearch-dictionaries.html)

addressing_dictionary installs address-oriented full-text-search configurations and dictionaries. It is useful for approximate local address lookup where street-type, direction, and ordinal variants should normalize to comparable lexemes; it is not a geocoder or a substitute for address parsing and validation.

### Core Workflow

Install the extension once in the database, normalize address text with the supplied configuration, and index the same expression used by queries.

```sql
CREATE EXTENSION addressing_dictionary;

SELECT to_tsvector('addressing_en', '1234 n main st');

CREATE INDEX address_search_idx
ON address_book
USING gin (to_tsvector('addressing_en', full_address));

SELECT *
FROM address_book
WHERE to_tsvector('addressing_en', full_address)
      @@ plainto_tsquery('addressing_en', '1234 north main street');
```

Version 1.1 keeps an abbreviated cardinal direction and also emits its expanded form, so an input such as a single-letter north token can produce both forms. This deliberately preserves ambiguity rather than assuming every occurrence has only one meaning.

### Installed Objects

- The English configuration combines a thesaurus, synonym dictionary, and stop-word dictionary for address tokens.
- The French configuration installs the corresponding French text-search objects.
- Both configurations are copied from PostgreSQL's simple configuration, so they normalize through the extension's data files instead of applying a general-language stemmer.

### Operational Boundaries

The extension is relocatable and needs no preload or restart. Its dictionary, synonym, stop-word, and thesaurus files must be installed on every server that may execute queries or restore the extension.

Full-text matching is heuristic. Abbreviations can be ambiguous, and nationwide data can produce many false or duplicate candidates. Use it to improve candidate search—especially for a bounded city or region—and retain separate structured-address or geospatial checks when correctness matters.
