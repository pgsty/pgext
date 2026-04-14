## Usage
> Sources: [README](https://github.com/frectonz/pg-when/blob/main/README.md) and [project repo](https://github.com/frectonz/pg-when).

`pg-when` is a PostgreSQL extension for creating time values from natural-language phrases.
It exposes the same parsed expression through multiple return formats:
`when_is`, `seconds_at`, `millis_at`, `micros_at`, and `nanos_at`.

The query syntax combines up to three parts:

```sql
SELECT when_is('<date> at <time> in <timezone>');
SELECT when_is('<date>');
SELECT when_is('<time> in <timezone>');
SELECT when_is('<date> at <time>');
```

If no timezone is provided, the extension defaults to UTC.

### Supported Components

`<date>` can be relative or exact.

Relative date examples from the README include:

- `today`
- `yesterday`
- `tomorrow`
- `next week`
- `last month`
- `this friday`
- `5 days ago`
- `in 2 years`

Exact dates can be written as:

- `YYYY-MM-DD` or `YYYY/MM/DD`
- `DD-MM-YYYY` or `DD/MM/YYYY`
- `Month D, YYYY`
- `D Month YYYY`

`<time>` can also be relative or exact.

Relative time examples include:

- `noon`
- `midnight`
- `morning`
- `evening`
- `next hour`
- `previous minute`
- `this hour`

Exact times can be written in 12-hour or 24-hour forms such as `8:30 pm` or `15:45`.

Timezones accept either IANA names or UTC offsets, for example `America/New_York` or `UTC-08:00`.

### Examples

```sql
SELECT when_is('5 days ago at this hour in Asia/Tokyo');
SELECT when_is('next friday at 8:00 pm in America/New_York');
SELECT when_is('in 2 months at midnight in UTC-8');
SELECT when_is('last monday at 22:30');
SELECT when_is('December 31, 2026 at evening');
```

### Deployment

The upstream README shows Docker images for PostgreSQL 13 through 18.
That matches the package metadata in this repository.
