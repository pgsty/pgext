## Usage

Sources:

- [Official README](https://github.com/alexeyklyukin/translate_isbn/blob/dd65576710245916ab14e7ff00d8f94058e63e23/README.md)
- [Extension SQL API](https://github.com/alexeyklyukin/translate_isbn/blob/dd65576710245916ab14e7ff00d8f94058e63e23/translate_isbn--1.0.sql)
- [Conversion implementation](https://github.com/alexeyklyukin/translate_isbn/blob/dd65576710245916ab14e7ff00d8f94058e63e23/translate_isbn.c)

`translate_isbn` 1.0 provides `translate_isbn1013(text)`, a search-expansion helper compatible with Evergreen library-system behavior. It recognizes whitespace-separated ISBN-like tokens, strips punctuation, checks/corrects check digits, and emits ISBN-10 and ISBN-13 alternatives where conversion is possible.

### Core Workflow

```sql
CREATE EXTENSION translate_isbn;

SELECT translate_isbn1013('9789997122575');
-- 9789997122575 9997122577

SELECT translate_isbn1013('ISBN 0-596-52068-9');
```

For each recognized 10- or 13-character candidate, the output includes the normalized original. If its check digit is wrong, a corrected form is appended. A valid conversion is then appended: ISBN-10 becomes an ISBN-13 with the `978` prefix, while only compatible ISBN-13 prefixes can become ISBN-10. Results are space-separated text intended to broaden a search term.

### Validation Boundary

The function validates digit shape and checksum but deliberately does not validate ISBN registration groups or publisher codes. Consequently, a checksum-valid but unassigned code may be converted. If no token is recognized, it returns the original input unchanged; mixed input can omit non-ISBN tokens once a candidate is found.

Do not use this function as a canonical ISBN data type, catalog authority, or publishing validator. The upstream README recommends PostgreSQL's `isn` contrib extension for fuller validation and typed storage. The C function is `STRICT` and `IMMUTABLE`, accepts text only, and can change formatting by removing punctuation and joining alternatives.
