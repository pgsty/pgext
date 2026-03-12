

## Usage

> [periods: Periods and SYSTEM VERSIONING for PostgreSQL](https://github.com/xocolatl/periods)

This extension recreates the behavior defined in [SQL:2016](https://www.iso.org/standard/63556.html) (originally in SQL:2011) around periods and tables with `SYSTEM VERSIONING`. The idea is to figure out all the rules that PostgreSQL would like to adopt and to allow earlier versions of PostgreSQL to simulate the behavior once the feature is finally integrated.

### What is a period?

A period is a definition on a table which specifies a name and two columns. The period's name cannot be the same as any column name of the table.

```sql
-- Standard SQL
CREATE TABLE example (
    id bigint,
    start_date date,
    end_date date,
    PERIOD FOR validity (start_date, end_date)
);
```

Since extensions cannot modify PostgreSQL's grammar, we use functions, views, and triggers to get as close to the same thing as possible.

```sql
CREATE TABLE example (
    id bigint,
    start_date date,
    end_date date
);
SELECT periods.add_period('example', 'validity', 'start_date', 'end_date');
```

Defining a period constrains the two columns such that the start column's value must be strictly inferior to the end column's value, and that both columns be non-null. The period's value includes the start value but excludes the end value.

## Unique Constraints

Periods may be part of `PRIMARY KEY`s and `UNIQUE` constraints.

```sql
CREATE TABLE example (
    id bigint,
    start_date date,
    end_date date
);
SELECT periods.add_period('example', 'validity', 'start_date', 'end_date');
SELECT periods.add_unique_key('example', ARRAY['id'], 'validity');
```

The extension will create a unique constraint over all of the columns specified and the two columns of the period given. It will also create an exclusion constraint using gist to implement the `WITHOUT OVERLAPS` part of the constraint.

## Foreign Keys

If you can have unique keys with periods, you can also have foreign keys pointing at them.

```sql
SELECT periods.add_foreign_key('example2', 'ARRAY[ex_id]', 'validity', 'example_id_validity');
```

## Portions

The SQL standard allows syntax for updating or deleting just a portion of a period. Rows are inserted as needed for the portions not being updated or deleted.

```sql
-- Standard SQL
UPDATE example
FOR PORTION OF validity FROM '...' TO '...'
SET ...
WHERE ...;
```

This extension uses a view with an `INSTEAD OF` trigger to figure out what portion of the period you would like to modify:

```sql
UPDATE example__for_portion_of_validity
SET ...,
    start_date = ...,
    end_date = ...
WHERE ...;
```

In order to use this feature, the table must have a primary key.

## Predicates

The SQL standard provides for several predicates on periods, implemented as inlined functions:

```sql
-- "t" and "u" are tables with respective periods "p" and "q".
-- Both periods have underlying columns "s" and "e".

WHERE periods.contains(t.s, t.e, 42)            -- t.p CONTAINS 42
WHERE periods.contains(t.s, t.e, u.s, u.e)      -- t.p CONTAINS u.q
WHERE periods.equals(t.s, t.e, u.s, u.e)        -- t.p EQUALS u.q
WHERE periods.overlaps(t.s, t.e, u.s, u.e)      -- t.p OVERLAPS u.q
WHERE periods.precedes(t.s, t.e, u.s, u.e)      -- t.p PRECEDES u.q
WHERE periods.succeeds(t.s, t.e, u.s, u.e)      -- t.p SUCCEEDS u.q
WHERE periods.immediately_precedes(t.s, t.e, u.s, u.e)  -- t.p IMMEDIATELY PRECEDES u.q
WHERE periods.immediately_succeeds(t.s, t.e, u.s, u.e)  -- t.p IMMEDIATELY SUCCEEDS u.q
```


## System-Versioned Tables

### SYSTEM_TIME

If the period is named `SYSTEM_TIME`, then special rules apply. The type of the columns must be `date`, `timestamp without time zone`, or `timestamp with time zone`; and they are not modifiable by the user. This extension uses triggers to set the start column to `transaction_timestamp()` and the end column is always `'infinity'`.

***Note:*** It is generally unwise to use anything but `timestamp with time zone` because changes in the `TimeZone` configuration parameter or Daylight Savings Time changes can distort the history.

```sql
CREATE TABLE example (
    id bigint PRIMARY KEY,
    value text
);
SELECT periods.add_system_time_period('example', 'row_start', 'row_end');
```

The columns need not exist — they will be created by the extension.

### Excluding Columns

It might be desirable to prevent some columns from updating the `SYSTEM_TIME` values:

```sql
SELECT periods.add_system_time_period(
            'example',
            excluded_column_names => ARRAY['foo', 'bar']);
```

Excluded columns can be defined after the fact as well:

```sql
SELECT periods.set_system_time_period_excluded_columns(
            'example',
            ARRAY['foo', 'bar']);
```

### WITH SYSTEM VERSIONING

This special `SYSTEM_TIME` period can be used to keep track of changes in the table.

```sql
CREATE TABLE example (
    id bigint PRIMARY KEY,
    value text
);
SELECT periods.add_system_time_period('example', 'row_start', 'row_end');
SELECT periods.add_system_versioning('example');
```

This instructs the system to keep a record of all changes in the table. A separate history table is used. You can create the history table yourself and instruct the extension to use it if you want to do things like add partitioning.

### Temporal Querying

The SQL standard extends the `FROM` and `JOIN` clauses to allow specifying a point in time, or a range of time. This extension implements them through inlined functions:

```sql
SELECT * FROM t__as_of('...');                       -- FOR system_time AS OF '...'
SELECT * FROM t__from_to('...', '...');              -- FOR system_time FROM '...' TO '...'
SELECT * FROM t__between('...', '...');              -- FOR system_time BETWEEN '...' AND '...'
SELECT * FROM t__between_symmetric('...', '...');    -- FOR system_time BETWEEN SYMMETRIC '...' AND '...'
```

### Access Control

The history table as well as the helper functions all follow the ownership and access privileges of the base table. The history data is read-only. In order to trim old data, `SYSTEM VERSIONING` must be suspended:

```sql
BEGIN;
SELECT periods.drop_system_versioning('t');
GRANT DELETE ON TABLE t TO CURRENT_USER;
DELETE FROM t_history WHERE system_time_end < now() - interval '1 year';
SELECT periods.add_system_versioning('t');
COMMIT;
```

### Altering a Table with System Versioning

This extension prevents you from dropping objects while system versioning is active. The suggested way to make changes is:

```sql
BEGIN;
SELECT periods.drop_system_versioning('t');
ALTER TABLE t ...;
ALTER TABLE t_history ...;
SELECT periods.add_system_versioning('t');
COMMIT;
```

It is up to you to make sure you alter the history table in a way that is compatible with the main table. Re-activating system versioning will verify this.
