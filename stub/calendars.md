## Usage

Sources:

- [Official extension control file](https://github.com/macdice/calendars/blob/22d7900c371ad142505667c14c7d9605dccf077d/calendars.control)
- [Official upstream documentation](https://github.com/macdice/calendars/blob/22d7900c371ad142505667c14c7d9605dccf077d/README)

`calendars` 1.0 is an unfinished experiment intended to add named-calendar overloads for date formatting and parsing. The checked-in Ethiopian conversion routines do not implement real calendar conversion, so the extension is unsuitable for application data.

### Core Workflow

```sql
CREATE EXTENSION calendars;

SELECT to_char(current_date, 'YYYY-MM-DD', 'Ethiopian');
SELECT to_date('2000-01-01', 'YYYY-MM-DD', 'Ethiopian');
```

The three-argument `to_char(date, text, text)` and `to_date(text, text, text)` overloads select a calendar by name. In the reviewed source, the first call returns a literal unfinished marker and the second returns a fixed date instead of converting its input.

### Caveats

Do not use these results for storage, validation, reporting, or date arithmetic. The generic overload names can also affect function resolution in databases that define similar signatures. Keep this extension confined to source inspection or isolated experiments until an upstream implementation and test vectors exist.
