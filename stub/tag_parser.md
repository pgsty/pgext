## Usage

Sources:

- [Official upstream README](https://github.com/serxka/tag_parser/blob/18280fc48cc212866a864b7f1741d7b8c75d12d8/Readme.md)
- [Official extension control file (tag_parser.control)](https://github.com/serxka/tag_parser/blob/18280fc48cc212866a864b7f1741d7b8c75d12d8/tag_parser.control)
- [Official extension SQL (tag_parser--0.1.0.sql)](https://github.com/serxka/tag_parser/blob/18280fc48cc212866a864b7f1741d7b8c75d12d8/tag_parser--0.1.0.sql)

`tag_parser` — A simple full text search parser for PostgreSQL, written in C. Is much simpler than the default tsearch parser, it will only split into lexeme's at a comma boundary. This can easily be changed to any other character in the source code by changing the definition of BREAK_CHAR. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION tag_parser;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `tagpsr_end(internal)` is an extension function and returns `void`.
- `tagpsr_gettoken(internal, internal, internal)` is an extension function and returns `internal`.
- `tagpsr_lextype(internal)` is an extension function and returns `internal`.
- `tagpsr_start(internal, int4)` is an extension function and returns `internal`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
