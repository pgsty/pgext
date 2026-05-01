## 用法

来源：[PGXN metadata](https://api.pgxn.org/dist/pg_bikram_sambat.json)、[PGXN source tree](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/)、[type SQL](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/sql/types.sql)、[function SQL](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/sql/functions.sql)、[operator SQL](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/sql/operators.sql)、[cast SQL](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/sql/casts.sql)、[regression examples](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/tests/pg_regress/sql/003_functions.sql)、[TODO](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/todos.txt)

来源说明：CSV 中的 GitHub URL 不可用，因此本文依据官方 PGXN 元数据与源码包 SQL。

`pg_bikram_sambat` 添加 `bs_date` 类型，用于 Bikram Sambat 日期，并提供转换、格式化、比较和 btree 索引支持。按普通 PostgreSQL 扩展安装：

```sql
CREATE EXTENSION pg_bikram_sambat;
```

### 日期类型

`bs_date` 存储 Bikram Sambat 日期，并以 `YYYY-MM-DD` 显示。文本输入接受用 `/`、`-` 或 `.` 分隔的年/月/日值；当年份出现在最后一个字段时，输入解析器也接受日优先字符串。

```sql
SELECT '2057/10/19'::bs_date;
SELECT CAST('2057-10-19' AS bs_date);
SELECT '19.10.2057'::bs_date;
```

像其他 PostgreSQL 类型一样在表中使用：

```sql
CREATE TABLE events (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  ad_date date,
  bs bs_date NOT NULL
);

INSERT INTO events (ad_date, bs)
VALUES
  ('2001-02-01', '2057/10/19'),
  ('1972-02-17', '2028/11/05');
```

### 转换函数

`ad_to_bs(date)` 将 Gregorian `date` 转换为 `bs_date`：

```sql
SELECT ad_to_bs('2001-02-01'::date);  -- 2057-10-19
SELECT ad_to_bs('1972-02-17'::date);  -- 2028-11-05
```

`current_bs_date()` 返回转换为 `bs_date` 的当前事务时间戳，因此同一事务内重复调用是稳定的：

```sql
SELECT current_bs_date();
SELECT pg_typeof(current_bs_date());  -- bs_date
```

版本 `0.1.0` 不暴露 SQL `bs_to_ad()` 函数，也没有直接的 `bs_date` 到 `date` cast；上游 TODO 文件将这些列为后续工作。

### 格式化

该扩展为 `bs_date` 重载 PostgreSQL `to_char`：

```sql
SELECT to_char('2057/10/19'::bs_date, 'YYYY-MM-DD');
SELECT to_char('2057/10/19'::bs_date, 'DD/MM/YYYY');
SELECT to_char('2057/10/19'::bs_date, 'Month DD, YYYY');
SELECT to_char('2057/10/19'::bs_date, 'Day, DD Month YYYY');
```

支持的日期格式 token 为 `YYYY`、`YY`、`Month`、`Mon`、`MM`、`Day`、`Dy` 和 `DD`。月份与星期名称会跟随格式 token 的大小写，因此 `MONTH`、`Month` 和 `month` 分别生成大写、标题式大小写和小写英文名称。

第三个参数传入 `dev` 时，使用 Devanagari 数字、月份名称和星期名称：

```sql
SELECT to_char('2057/10/19'::bs_date, 'YYYY-MM-DD', 'dev');
SELECT to_char('2057/10/19'::bs_date, 'Day, DD Month YYYY', 'dev');
```

### 操作符与索引

`bs_date` 支持比较操作符 `=`、`<>`、`>`、`>=`、`<` 和 `<=`。默认 btree 操作符类 `bs_date_ops` 支持普通 btree 索引、范围谓词和排序：

```sql
CREATE INDEX events_bs_idx ON events (bs);

SELECT * FROM events WHERE bs >= '2057/01/01' ORDER BY bs;
SELECT * FROM events WHERE bs BETWEEN '2056/01/01' AND '2058/12/12';
```

### 注意事项

打包的转换数据集覆盖 BS 年份 `2000` 到 `2100`，并以 `1943-04-14` AD 作为 `2000-01-01` BS 的参考日期。参考日期之前或映射 BS 范围之外的日期会抛出 PostgreSQL 错误。该扩展定义了从 `text` 到 `bs_date` 的隐式 cast，但没有定义从任意数值类型到 `bs_date` 的 cast。
