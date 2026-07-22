## Usage

Sources:

- [Official README for the 22.6 series](https://github.com/greenbone/pg-gvm/blob/v22.6.18/README.md)
- [Official host utility SQL](https://github.com/greenbone/pg-gvm/blob/v22.6.18/sql/hosts.in.sql)
- [Official iCalendar utility SQL](https://github.com/greenbone/pg-gvm/blob/v22.6.18/sql/ical.in.sql)
- [Official regular-expression utility SQL](https://github.com/greenbone/pg-gvm/blob/v22.6.18/sql/regexp.in.sql)

`pg-gvm` is Greenbone's PostgreSQL helper library for GVMd. Version/API series `22.6` provides C functions for Greenbone host-range expressions, iCalendar recurrence calculation, and regular-expression matching. It is an application support component rather than a general SQL utility collection.

### Core Workflow

Create the extension, then use its helpers with the same text formats produced by Greenbone components:

```sql
CREATE EXTENSION "pg-gvm";

SELECT hosts_contains(
  '192.168.123.1-192.168.123.20, 192.168.123.30',
  '192.168.123.10'
);

SELECT max_hosts(
  '192.168.123.1-192.168.123.20, 192.168.123.30-192.168.123.34',
  '192.168.123.10'
);

SELECT regexp('abc', '^[a-z]+$');
```

The first call tests membership in a Greenbone host specification. The second counts addresses in the range expression while excluding the second argument.

### API

- `hosts_contains(text, text) RETURNS boolean`: whether a host expression contains the requested host.
- `max_hosts(text, text) RETURNS integer`: count hosts represented by the first expression after exclusions in the second.
- `regexp(text, text) RETURNS boolean`: regular-expression match; invalid patterns return false in the upstream tests.
- `next_time_ical(text, bigint, text) RETURNS integer`: calculate the next recurrence from iCalendar text, a Unix timestamp, and timezone.
- `next_time_ical(text, bigint, text, integer) RETURNS integer`: same calculation with an occurrence offset.

### Caveats

The extension links to `libical`, GLib, and `libgvm-base`; package and library versions must match the Greenbone stack. The generated control file uses the project's API version and shared-library name, so install a coherent release rather than mixing artifacts from different tags.

Host syntax and iCalendar behavior follow Greenbone's libraries, not PostgreSQL native network/range types. Test boundary addresses, exclusions, timezones, daylight-saving transitions, malformed calendar input, and large range expressions before embedding these helpers in policy decisions.
