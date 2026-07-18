## Usage

Sources:

- [Extension SQL](https://github.com/r888888888/test_parser/blob/356c676d16a987f9226b3ea419b4e18612c61fef/test_parser--1.0.sql)
- [Extension control file](https://github.com/r888888888/test_parser/blob/356c676d16a987f9226b3ea419b4e18612c61fef/test_parser.control)
- [C implementation](https://github.com/r888888888/test_parser/blob/356c676d16a987f9226b3ea419b4e18612c61fef/test_parser.c)

`test_parser` version `1.0` installs a C text-search parser named `testparser`. Use PostgreSQL's parser diagnostic function to inspect the tokens it emits:

```sql
CREATE EXTENSION test_parser;

SELECT *
FROM ts_parse('testparser', 'The quick brown fox 123');
```

### Project scope

The fixed upstream source contains the parser callbacks and extension objects but no README, release notes, license file, or PostgreSQL compatibility matrix. Treat it as a minimal parser example: build and regression-test it against the exact server major version before use, and do not infer production support from the extension name.
