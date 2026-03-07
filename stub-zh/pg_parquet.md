

## 用法

`pg_parquet` 主要提供以下三项功能：

1. 将 PostgreSQL 表/查询结果导出为 Parquet 文件，
2. 将 Parquet 文件中的数据导入到 PostgreSQL 表中，
3. 查看 Parquet 文件的模式与元数据信息。

### 使用 COPY 在 Parquet 文件与 PostgreSQL 表之间导入导出



可以使用 PostgreSQL 的 `COPY` 命令来读写 Parquet 文件。以下示例演示了如何将包含复合类型的 PostgreSQL 表写入 Parquet 文件，然后再将该 Parquet 文件的内容读回同一张表。

```sql
-- 创建复合类型
CREATE TYPE product_item AS (id INT, name TEXT, price float4);
CREATE TYPE product AS (id INT, name TEXT, items product_item[]);

-- 创建包含复合类型的表
CREATE TABLE product_example (
    id int,
    product product,
    products product[],
    created_at TIMESTAMP,
    updated_at TIMESTAMPTZ
);

-- 向表中插入数据
insert into product_example values (
    1,
    ROW(1, 'product 1', ARRAY[ROW(1, 'item 1', 1.0), ROW(2, 'item 2', 2.0), NULL]::product_item[])::product,
    ARRAY[ROW(1, NULL, NULL)::product, NULL],
    now(),
    '2022-05-01 12:00:00-04'
);

-- 将表导出为 Parquet 文件
COPY product_example TO '/tmp/product_example.parquet' (format 'parquet', compression 'gzip');

-- 查看表内容
SELECT * FROM product_example;

-- 将 Parquet 文件导入到表中
COPY product_example FROM '/tmp/product_example.parquet';

-- 查看表内容
SELECT * FROM product_example;
```



此外，也可以使用 `COPY` 命令通过标准输入/输出读写 Parquet 流。以下是使用示例（必须指定 `format = parquet`）：

```bash
psql -d pg_parquet -p 28817 -h localhost -c "create table product_example_reconstructed (like product_example);"
 CREATE TABLE

psql -d pg_parquet -p 28817 -h localhost -c "copy product_example to stdout (format parquet);" | psql -d pg_parquet -p 28817 -h localhost -c "copy product_example_reconstructed from stdin (format parquet);"
 COPY 2
```



### 查看 Parquet 文件模式

调用 `SELECT * FROM parquet.schema(<uri>)` 可查看指定 URI 处 Parquet 文件的模式信息。

```sql
SELECT * FROM parquet.schema('/tmp/product_example.parquet') LIMIT 10;
             uri              |     name     | type_name  | type_length | repetition_type | num_children | converted_type | scale | precision | field_id | logical_type
------------------------------+--------------+------------+-------------+-----------------+--------------+----------------+-------+-----------+----------+--------------
 /tmp/product_example.parquet | arrow_schema |            |             |                 |            5 |                |       |           |          |
 /tmp/product_example.parquet | id           | INT32      |             | OPTIONAL        |              |                |       |           |        0 |
 /tmp/product_example.parquet | product      |            |             | OPTIONAL        |            3 |                |       |           |        1 |
 /tmp/product_example.parquet | id           | INT32      |             | OPTIONAL        |              |                |       |           |        2 |
 /tmp/product_example.parquet | name         | BYTE_ARRAY |             | OPTIONAL        |              | UTF8           |       |           |        3 | STRING
 /tmp/product_example.parquet | items        |            |             | OPTIONAL        |            1 | LIST           |       |           |        4 | LIST
 /tmp/product_example.parquet | list         |            |             | REPEATED        |            1 |                |       |           |          |
 /tmp/product_example.parquet | element        |            |             | OPTIONAL        |            3 |                |       |           |        5 |
 /tmp/product_example.parquet | id           | INT32      |             | OPTIONAL        |              |                |       |           |        6 |
 /tmp/product_example.parquet | name         | BYTE_ARRAY |             | OPTIONAL        |              | UTF8           |       |           |        7 | STRING
(10 rows)
```



### 查看 Parquet 元数据

调用 `SELECT * FROM parquet.metadata(<uri>)` 可查看指定 URI 处 Parquet 文件的详细元数据，例如列统计信息等。

```
SELECT uri, row_group_id, row_group_num_rows, row_group_num_columns, row_group_bytes, column_id, file_offset, num_values, path_in_schema, type_name FROM parquet.metadata('/tmp/product_example.parquet') LIMIT 1;
             uri              | row_group_id | row_group_num_rows | row_group_num_columns | row_group_bytes | column_id | file_offset | num_values | path_in_schema | type_name
------------------------------+--------------+--------------------+-----------------------+-----------------+-----------+-------------+------------+----------------+-----------
 /tmp/product_example.parquet |            0 |                  1 |                    13 |             842 |         0 |           0 |          1 | id             | INT32
(1 row)
```



```
SELECT stats_null_count, stats_distinct_count, stats_min, stats_max, compression, encodings, index_page_offset, dictionary_page_offset, data_page_offset, total_compressed_size, total_uncompressed_size FROM parquet.metadata('/tmp/product_example.parquet') LIMIT 1;
 stats_null_count | stats_distinct_count | stats_min | stats_max |    compression     |        encodings         | index_page_offset | dictionary_page_offset | data_page_offset | total_compressed_size | total_uncompressed_size
------------------+----------------------+-----------+-----------+--------------------+--------------------------+-------------------+------------------------+------------------+-----------------------+-------------------------
                0 |                      | 1         | 1         | GZIP(GzipLevel(6)) | PLAIN,RLE,RLE_DICTIONARY |                   |                      4 |               42 |                   101 |                      61
(1 row)
```



调用 `SELECT * FROM parquet.file_metadata(<uri>)` 可查看指定 URI 处 Parquet 文件的文件级元数据，例如格式版本等。

```
SELECT * FROM parquet.file_metadata('/tmp/product_example.parquet')
             uri              | created_by | num_rows | num_row_groups | format_version
------------------------------+------------+----------+----------------+----------------
 /tmp/product_example.parquet | pg_parquet |        1 |              1 | 1
(1 row)
```



调用 `SELECT * FROM parquet.kv_metadata(<uri>)` 可查询指定 URI 处 Parquet 文件的自定义键值元数据。

```
SELECT uri, encode(key, 'escape') as key, encode(value, 'escape') as value FROM parquet.kv_metadata('/tmp/product_example.parquet');
             uri              |     key      |    value
------------------------------+--------------+---------------------
 /tmp/product_example.parquet | ARROW:schema | /////5gIAAAQAAAA ...
(1 row)
```



### 查看 Parquet 列统计信息

调用 `SELECT * FROM parquet.column_stats(<uri>)` 可查看指定 URI 处 Parquet 文件的列统计信息，例如列的最小值和最大值等。

```
SELECT * FROM parquet.column_stats('/tmp/product_example.parquet')
 column_id | field_id |         stats_min          |         stats_max          | stats_null_count | stats_distinct_count
-----------+----------+----------------------------+----------------------------+------------------+----------------------
         4 |        7 | item 1                     | item 2                     |                1 |
         6 |       11 | 1                          | 1                          |                1 |
         7 |       12 |                            |                            |                2 |
        10 |       17 |                            |                            |                2 |
         0 |        0 | 1                          | 1                          |                0 |
        11 |       18 | 2025-03-11 14:01:22.045739 | 2025-03-11 14:01:22.045739 |                0 |
         3 |        6 | 1                          | 2                          |                1 |
        12 |       19 | 2022-05-01 19:00:00+03     | 2022-05-01 19:00:00+03     |                0 |
         8 |       15 |                            |                            |                2 |
         5 |        8 | 1                          | 2                          |                1 |
         9 |       16 |                            |                            |                2 |
         1 |        2 | 1                          | 1                          |                0 |
         2 |        3 | product 1                  | product 1                  |                0 |
(13 rows)
```