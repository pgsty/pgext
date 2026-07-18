## Usage

Sources:

- [Upstream README](https://github.com/pramsey/pgsql-addressing-dictionary/blob/bf3590e038d820c8eb5c79135e5a78abf5fc9f2f/README.md)
- [Extension SQL](https://github.com/pramsey/pgsql-addressing-dictionary/blob/bf3590e038d820c8eb5c79135e5a78abf5fc9f2f/addressing_dictionary--1.1.sql)
- [Extension control file](https://github.com/pramsey/pgsql-addressing-dictionary/blob/bf3590e038d820c8eb5c79135e5a78abf5fc9f2f/addressing_dictionary.control)

`addressing_dictionary` installs full-text-search dictionaries and configurations for postal addresses. The documented `addressing_en` configuration normalizes common street types, directions, and number forms so variant spellings can produce matching lexemes.

Install the upstream dictionary files on the PostgreSQL server, then enable and use the extension:

```sql
CREATE EXTENSION addressing_dictionary;

SELECT to_tsvector('addressing_en', '1234 n main st');
SELECT to_tsquery('addressing_en', 'north & main & street');
```

### Operational notes

Version `1.1` is a relocatable, SQL-only extension, but its text-search dictionaries depend on files installed by the upstream build. The install SQL also defines `addressing_fr`; the README's worked examples and language description focus on `addressing_en`, so validate the French configuration against local address data before relying on it. Address normalization can yield false positives or multiple candidates and is not a substitute for geocoding or address validation.
