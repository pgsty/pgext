## Usage

Sources:

- [Upstream README](https://github.com/wttw/pgpsl/blob/8def780438e263506a8ae4beb9c754a31f9ee341/README.md)
- [Version 1.0.0 SQL definition](https://github.com/wttw/pgpsl/blob/8def780438e263506a8ae4beb9c754a31f9ee341/psl--1.0.0.sql)
- [Extension control file](https://github.com/wttw/pgpsl/blob/8def780438e263506a8ae4beb9c754a31f9ee341/psl.control)

The package `pgpsl` installs extension `psl` version `1.0.0`. It exposes one function that finds the registrable domain of a hostname using a Public Suffix List snapshot compiled into the library.

### Core Workflow

```sql
CREATE EXTENSION psl;

SELECT registered_domain('foo.bar.blighty.com');
SELECT registered_domain('www.blighty.co.uk');
SELECT registered_domain('co.uk');
```

`registered_domain(text)` folds the result to lowercase. It returns the enclosing registered domain for a hostname, the same value for an already-registered domain, and NULL for a public suffix or a name without dots. For an apparently valid but unknown top-level suffix, it falls back to the final two components.

### Data Freshness and Caveats

The Public Suffix List is fixed at build time and cannot be refreshed dynamically. Rebuild the extension with a reviewed list snapshot whenever suffix accuracy matters; changing the package alone does not update an already-loaded library.

Upstream explicitly warns that its list-preprocessing code is broken and that attempted fixes produced a corrupt embedded list and crashes. Test representative suffixes, internationalized names, invalid input, and unknown suffixes before use. Do not use this unmaintained function as the sole security boundary for cookies, tenant isolation, or authorization; a stale or corrupt list can misclassify domain ownership.
