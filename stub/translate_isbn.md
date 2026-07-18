## Usage

Sources:

- [Official extension control file](https://github.com/alexeyklyukin/translate_isbn/blob/dd65576710245916ab14e7ff00d8f94058e63e23/translate_isbn.control)
- [Official upstream documentation](https://github.com/alexeyklyukin/translate_isbn/blob/dd65576710245916ab14e7ff00d8f94058e63e23/README.md)

`translate_isbn` — C function that converts text ISBN codes between ISBN-10 and ISBN-13 forms for Evergreen-compatible search expansion.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "translate_isbn";
SELECT extversion
FROM pg_extension
WHERE extname = 'translate_isbn';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
