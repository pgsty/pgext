## Usage

Sources:

- [README.md](https://github.com/rdnkrkmz/max2min/blob/6de4805841dfd3797c3f967acfd2939fecedb6c9/README.md)
- [max2min--1.0.sql](https://github.com/rdnkrkmz/max2min/blob/6de4805841dfd3797c3f967acfd2939fecedb6c9/max2min--1.0.sql)
- [LICENSE](https://github.com/rdnkrkmz/max2min/blob/6de4805841dfd3797c3f967acfd2939fecedb6c9/LICENSE)
- [max2min.control](https://github.com/rdnkrkmz/max2min/blob/6de4805841dfd3797c3f967acfd2939fecedb6c9/max2min.control)

`max2min` 1.0 provides `max_to_min()` aggregates that return the largest and smallest values in one text result, with an optional delimiter.

### Core Workflow

```sql
CREATE EXTENSION max2min;
SELECT max_to_min(v) FROM (VALUES (5), (3), (10), (NULL::integer)) AS t(v);
SELECT max_to_min(v, ' .. ') FROM (VALUES (5), (3), (10)) AS t(v);
```

### Operational Boundaries

The versioned SQL implements numeric, integer, floating-point, character and date/time overloads. SQL NULL inputs are handled by the aggregate; ordering follows the input type rather than the formatted text. The result is presentation text, not a typed range. This is a SQL extension with no shared library or preload requirement. The control does not mark installation trusted, and upstream does not publish a PostgreSQL-major compatibility matrix.
