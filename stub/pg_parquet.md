
## Usage

There are mainly 3 things that you can do with `pg_parquet`:

1. You can export Postgres tables/queries to Parquet files,
2. You can ingest data from Parquet files to Postgres tables,
3. You can inspect the schema and metadata of Parquet files.

### COPY to/from Parquet files from/to Postgres tables



You can use PostgreSQL's `COPY` command to read and write from/to Parquet files. Below is an example of how to write a PostgreSQL table, with complex types, into a Parquet file and then to read the Parquet file content back into the same table.

```sql
-- create composite types
CREATE TYPE product_item AS (id INT, name TEXT, price float4);
CREATE TYPE product AS (id INT, name TEXT, items product_item[]);

-- create a table with complex types
CREATE TABLE product_example (
    id int,
    product product,
    products product[],
    created_at TIMESTAMP,
    updated_at TIMESTAMPTZ
);

-- insert some rows into the table
insert into product_example values (
    1,
    ROW(1, 'product 1', ARRAY[ROW(1, 'item 1', 1.0), ROW(2, 'item 2', 2.0), NULL]::product_item[])::product,
    ARRAY[ROW(1, NULL, NULL)::product, NULL],
    now(),
    '2022-05-01 12:00:00-04'
);

-- copy the table to a parquet file
COPY product_example TO '/tmp/product_example.parquet' (format 'parquet', compression 'gzip');

-- show table
SELECT * FROM product_example;

-- copy the parquet file to the table
COPY product_example FROM '/tmp/product_example.parquet';

-- show table
SELECT * FROM product_example;
```



You can also use `COPY` command to read and write Parquet stream from/to standard input and output. Below is an example usage (you have to specify `format = parquet`):

```bash
psql -d pg_parquet -p 28817 -h localhost -c "create table product_example_reconstructed (like product_example);"
 CREATE TABLE

psql -d pg_parquet -p 28817 -h localhost -c "copy product_example to stdout (format parquet);" | psql -d pg_parquet -p 28817 -h localhost -c "copy product_example_reconstructed from stdin (format parquet);"
 COPY 2
```



### Inspect Parquet schema

You can call `SELECT * FROM parquet.schema(<uri>)` to discover the schema of the Parquet file at given uri.

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



### Inspect Parquet metadata

You can call `SELECT * FROM parquet.metadata(<uri>)` to discover the detailed metadata of the Parquet file, such as column statistics, at given uri.

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



You can call `SELECT * FROM parquet.file_metadata(<uri>)` to discover file level metadata of the Parquet file, such as format version, at given uri.

```
SELECT * FROM parquet.file_metadata('/tmp/product_example.parquet')
             uri              | created_by | num_rows | num_row_groups | format_version 
------------------------------+------------+----------+----------------+----------------
 /tmp/product_example.parquet | pg_parquet |        1 |              1 | 1
(1 row)
```



You can call `SELECT * FROM parquet.kv_metadata(<uri>)` to query custom key-value metadata of the Parquet file at given uri.

```
SELECT uri, encode(key, 'escape') as key, encode(value, 'escape') as value FROM parquet.kv_metadata('/tmp/product_example.parquet');
             uri              |     key      |    value
------------------------------+--------------+---------------------
 /tmp/product_example.parquet | ARROW:schema | /////5gIAAAQAAAA ...
(1 row)
```



### Inspect Parquet column statistics

You can call `SELECT * FROM parquet.column_stats(<uri>)` to discover the column statistics of the Parquet file, such as min and max value for the column, at given uri.

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