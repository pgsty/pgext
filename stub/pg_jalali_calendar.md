## Usage

Sources:

- [Official Rust implementation](https://github.com/alisol911/pg_jalali_calendar/blob/9682677a9d23b357b5a25db605c7ead5eb1abf50/src/lib.rs)
- [Official Cargo manifest](https://github.com/alisol911/pg_jalali_calendar/blob/9682677a9d23b357b5a25db605c7ead5eb1abf50/Cargo.toml)
- [Extension control file](https://github.com/alisol911/pg_jalali_calendar/blob/9682677a9d23b357b5a25db605c7ead5eb1abf50/pg_jalali_calendar.control)

`pg_jalali_calendar` provides text-based conversions and calculations for the Persian (Jalali) calendar, plus two Persian-language number helpers. Its date API consumes and returns strings rather than PostgreSQL `date` or `timestamp` values.

### Core Workflow

Create the extension, convert between Jalali `YYYY/MM/DD` strings and Gregorian `YYYY-MM-DD` strings, and validate untrusted input before calling functions that raise errors on invalid dates.

```sql
CREATE EXTENSION pg_jalali_calendar;

SELECT jalali_date_to_gregorian('1404/01/01');
SELECT gregorian_date_to_jalali('2025-03-21');
SELECT jalali_date_is_valid('1403/12/30');
```

Date arithmetic returns Jalali text. Day addition accepts positive or negative values, but month addition in version `0.1.0` accepts only positive values.

```sql
SELECT jalali_date_add_days('1404/01/01', -10);
SELECT jalali_date_add_months('1404/01/31', 7);
SELECT jalali_date_diff('1404/01/01', '1404/02/01');
```

### Important Objects

- `jalali_date_to_gregorian(text)`, `gregorian_date_to_jalali(text)`: calendar conversions.
- `jalali_date_add_days(text, integer)`, `jalali_date_add_months(text, integer)`: date arithmetic.
- `jalali_date_diff(text, text)`, `jalali_date_diff_with_addition(text, text, integer)`: signed day differences.
- `jalali_date_now()`: current Jalali date calculated from the UTC date.
- `jalali_date_is_valid(text)`, `jalali_date_is_leap_year(text)`: validation and leap-year checks.
- `jalali_date_period_state(text, integer)`: returns `Start`, `End`, `Middle`, or `Unknown` for the source-defined monthly period boundary.
- `number_to_farsi_word(bigint)`, `farsi_word_add_suffix(text)`: Persian number words and ordinal suffixes.

### Semantics and Compatibility

Month addition clamps a day that does not exist in the target month to that month's last valid day. It rejects zero and negative month counts in this release. Most functions panic into a PostgreSQL error for malformed or out-of-range text; `jalali_date_is_valid(text)` is the safe predicate for screening input. `jalali_date_now()` uses UTC, not the session time zone.

The pinned `0.1.0` Cargo manifest enables only the pgrx PostgreSQL 18 feature by default, and the control file is non-relocatable and marked both superuser-required and trusted. Confirm the actual packaged build and privileges on the target platform. If values must participate in PostgreSQL indexing, range queries, time zones, or interval logic, store a native Gregorian `date` alongside the Jalali presentation string.
