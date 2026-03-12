

## 用法

> [bzip: Bzip 压缩与解压缩](https://github.com/steve-chavez/pg_bzip)

### 函数

- `bzcat(data bytea) returns bytea` -- 解压 bzip2 数据，类似于 `bzcat` 命令。
- `bzip2(data bytea, compression_level int default 9) returns bytea` -- 使用 bzip2 压缩数据。

### 示例

解压 bzip2 压缩的文件：

```sql
SELECT convert_from(bzcat(pg_read_binary_file('/path/to/data.csv.bz2')), 'utf8') AS contents;
```

使用 bzip2 压缩数据：

```sql
SELECT bzip2(repeat('my text to be compressed', 1000)::bytea) AS compressed;
```

使用自定义压缩级别（1-9）进行压缩：

```sql
SELECT bzip2('some data'::bytea, 5);
```
