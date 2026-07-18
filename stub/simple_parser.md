## Usage

Sources:

- [Official extension control file](https://github.com/fairingrey/pg_simple_parser/blob/cdeaa0f1ed3b89d1961312c01e529f1c7c068084/simple_parser.control)

`simple_parser` — C full-text-search parser that splits input only at ASCII space characters, preserving punctuation such as underscores inside tokens.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "simple_parser";
SELECT extversion
FROM pg_extension
WHERE extname = 'simple_parser';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
