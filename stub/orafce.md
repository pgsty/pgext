

## Usage

> [orafce: Functions and operators that emulate a subset of functions and packages from the Oracle RDBMS](https://github.com/orafce/orafce)

### Date Functions

```sql
SELECT add_months(date '2005-05-31', 1);        -- 2005-06-30
SELECT last_day(date '2005-05-24');              -- 2005-05-31
SELECT next_day(date '2005-05-24', 'monday');    -- 2005-05-30
SELECT months_between(date '1995-02-02', date '1995-01-01'); -- 1.032...
SELECT trunc(date '2005-07-12', 'iw');           -- 2005-07-11
SELECT round(date '2005-07-12', 'yyyy');         -- 2006-01-01
```

### Oracle DATE Data Type

```sql
SET search_path TO oracle, "$user", public, pg_catalog;
CREATE TABLE t (col1 date);
INSERT INTO t VALUES('2014-06-24 12:12:11'::date);  -- includes time component
```

### String Functions (NVL, DECODE, etc.)

```sql
SELECT nvl('A', 'B');            -- A
SELECT nvl(NULL, 'B');           -- B
SELECT decode(1, 1, 'one', 2, 'two', 'other');  -- one
SELECT lnnvl(true);              -- false
SELECT nanvl(0.0/0.0, 999);     -- 999
```

### DUAL Table

```sql
SELECT * FROM dual;
```

### Package DBMS_OUTPUT

```sql
SELECT dbms_output.enable();
SELECT dbms_output.put_line('Hello');
SELECT dbms_output.get_line(line, status);  -- retrieves output
```

### Package DBMS_PIPE

```sql
SELECT dbms_pipe.create_pipe('my_pipe');
SELECT dbms_pipe.pack_message('message text');
SELECT dbms_pipe.send_message('my_pipe');
-- In another session:
SELECT dbms_pipe.receive_message('my_pipe');
SELECT dbms_pipe.unpack_message_text();
```

### Package DBMS_ALERT

```sql
CALL dbms_alert.register('my_alert');
-- In another session:
CALL dbms_alert.signal('my_alert', 'Alert message');
-- Back in first session:
CALL dbms_alert.waitone('my_alert', name, message, status, 60);
```

### Package DBMS_UTILITY

```sql
SELECT dbms_utility.format_call_stack();
```

### Package UTL_FILE

```sql
CALL utl_file.fopen('/tmp', 'test.txt', 'w');
CALL utl_file.put_line(f, 'Hello World');
CALL utl_file.fclose(f);
```

### Package PLVstr / PLVchr

```sql
SELECT plvstr.left('Hello World', 5);     -- Hello
SELECT plvstr.right('Hello World', 5);    -- World
SELECT plvstr.rvrs('Hello');              -- olleH
SELECT plvchr.nth('Hello', 3);            -- l
SELECT plvchr.first('Hello');             -- H
SELECT plvchr.last('Hello');              -- o
```

### Package PLVsubst

```sql
SELECT plvsubst.string('My name is %s %s.', ARRAY['Pavel','Stehule']);
-- My name is Pavel Stehule.
```

### DBMS_ASSERT (SQL Injection Protection)

```sql
SELECT dbms_assert.enquote_literal('some value');
SELECT dbms_assert.schema_name('public');
SELECT dbms_assert.object_name('my_table');
```

### VARCHAR2 and NVARCHAR2 Types

The extension provides Oracle-compatible `varchar2` and `nvarchar2` data types that enforce the declared length in bytes (varchar2) or characters (nvarchar2).
