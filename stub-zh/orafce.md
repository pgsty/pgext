

## 用法

> [orafce: 模拟 Oracle RDBMS 部分函数和包的函数与操作符](https://github.com/orafce/orafce)

### 日期函数

```sql
SELECT add_months(date '2005-05-31', 1);        -- 2005-06-30
SELECT last_day(date '2005-05-24');              -- 2005-05-31
SELECT next_day(date '2005-05-24', 'monday');    -- 2005-05-30
SELECT months_between(date '1995-02-02', date '1995-01-01'); -- 1.032...
SELECT trunc(date '2005-07-12', 'iw');           -- 2005-07-11
SELECT round(date '2005-07-12', 'yyyy');         -- 2006-01-01
```

### Oracle DATE 数据类型

```sql
SET search_path TO oracle, "$user", public, pg_catalog;
CREATE TABLE t (col1 date);
INSERT INTO t VALUES('2014-06-24 12:12:11'::date);  -- 包含时间部分
```

### 字符串函数（NVL、DECODE 等）

```sql
SELECT nvl('A', 'B');            -- A
SELECT nvl(NULL, 'B');           -- B
SELECT decode(1, 1, 'one', 2, 'two', 'other');  -- one
SELECT lnnvl(true);              -- false
SELECT nanvl(0.0/0.0, 999);     -- 999
```

### DUAL 表

```sql
SELECT * FROM dual;
```

### DBMS_OUTPUT 包

```sql
SELECT dbms_output.enable();
SELECT dbms_output.put_line('Hello');
SELECT dbms_output.get_line(line, status);  -- 获取输出
```

### DBMS_PIPE 包

```sql
SELECT dbms_pipe.create_pipe('my_pipe');
SELECT dbms_pipe.pack_message('message text');
SELECT dbms_pipe.send_message('my_pipe');
-- 在另一个会话中：
SELECT dbms_pipe.receive_message('my_pipe');
SELECT dbms_pipe.unpack_message_text();
```

### DBMS_ALERT 包

```sql
CALL dbms_alert.register('my_alert');
-- 在另一个会话中：
CALL dbms_alert.signal('my_alert', 'Alert message');
-- 回到第一个会话：
CALL dbms_alert.waitone('my_alert', name, message, status, 60);
```

### DBMS_UTILITY 包

```sql
SELECT dbms_utility.format_call_stack();
```

### UTL_FILE 包

```sql
CALL utl_file.fopen('/tmp', 'test.txt', 'w');
CALL utl_file.put_line(f, 'Hello World');
CALL utl_file.fclose(f);
```

### PLVstr / PLVchr 包

```sql
SELECT plvstr.left('Hello World', 5);     -- Hello
SELECT plvstr.right('Hello World', 5);    -- World
SELECT plvstr.rvrs('Hello');              -- olleH
SELECT plvchr.nth('Hello', 3);            -- l
SELECT plvchr.first('Hello');             -- H
SELECT plvchr.last('Hello');              -- o
```

### PLVsubst 包

```sql
SELECT plvsubst.string('My name is %s %s.', ARRAY['Pavel','Stehule']);
-- My name is Pavel Stehule.
```

### DBMS_ASSERT（SQL 注入防护）

```sql
SELECT dbms_assert.enquote_literal('some value');
SELECT dbms_assert.schema_name('public');
SELECT dbms_assert.object_name('my_table');
```

### VARCHAR2 和 NVARCHAR2 类型

该扩展提供 Oracle 兼容的 `varchar2` 和 `nvarchar2` 数据类型，分别以字节（varchar2）或字符（nvarchar2）为单位强制执行声明的长度。
