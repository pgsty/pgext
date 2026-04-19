## Usage

Sources: [official README](https://github.com/frectonz/pg-when/blob/main/README.md), [official repo](https://github.com/frectonz/pg-when)

`pg-when` parses a constrained natural-language time expression and returns either a PostgreSQL timestamp with time zone or Unix epoch values at different resolutions.

```sql
CREATE EXTENSION pg_when;

SELECT when_is('next friday at 8:00 pm in America/New_York');
SELECT seconds_at('next friday at 8:00 pm in America/New_York');
SELECT millis_at('next friday at 8:00 pm in America/New_York');
SELECT micros_at('next friday at 8:00 pm in America/New_York');
SELECT nanos_at('next friday at 8:00 pm in America/New_York');
```

### Supported Query Shape

The parser accepts up to three parts:

```sql
SELECT when_is('<date> at <time> in <timezone>');
SELECT when_is('<date>');
SELECT when_is('<time> in <timezone>');
SELECT when_is('<date> at <time>');
```

If no timezone is provided, upstream says the default is UTC.

### Common Inputs

- relative dates: `today`, `tomorrow`, `last month`, `this friday`, `5 days ago`, `in 2 years`
- exact dates: `YYYY-MM-DD`, `DD/MM/YYYY`, `January 10, 2004`, `10 Jan 2004`
- relative times: `noon`, `midnight`, `morning`, `evening`, `next hour`
- exact times: `8:30 pm`, `15:45`
- time zones: `America/New_York`, `Europe/London`, `UTC-08:00`, `UTC+05:30`

### Examples

```sql
SELECT when_is('5 days ago at this hour in Asia/Tokyo');
SELECT when_is('in 2 months at midnight in UTC-8');
SELECT when_is('December 31, 2026 at evening');
```

### Caveats

- The extension is aimed at the documented grammar above, not arbitrary English.
- Upstream distributes ready-made Docker images for PostgreSQL 13 through 18, but the stub should focus on SQL usage rather than container setup.
