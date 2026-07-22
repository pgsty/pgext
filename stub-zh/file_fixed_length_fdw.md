## 用法

来源：

- [官方扩展控制文件](https://github.com/adunstan/file_fixed_length_record_fdw/blob/53f7875d71cd457ae62daf142156af927630c999/file_fixed_length_fdw.control)
- [官方扩展实现](https://github.com/adunstan/file_fixed_length_record_fdw/blob/53f7875d71cd457ae62daf142156af927630c999/file_fixed_length_fdw.c)

`file_fixed_length_fdw` 将没有分隔符的定长记录公开为只读外部表。它适合每条记录都有相同字节布局，并且可以接受每条记录返回一个 `text[]` 值的场景。

### 核心流程

创建包装器、服务器，以及仅含一个 `text[]` 列的外部表。长度列表同时决定数组元素数量和每个元素消耗的字节数。

```sql
CREATE EXTENSION file_fixed_length_fdw;

CREATE SERVER fixed_files
  FOREIGN DATA WRAPPER file_fixed_length_fdw;

CREATE FOREIGN TABLE fixed_records (fields text[])
  SERVER fixed_files
  OPTIONS (
    filename '/srv/import/customer.dat',
    field_lengths '8,30,10',
    trim 'true',
    record_separator 'lf',
    encoding 'UTF8'
  );

SELECT fields FROM fixed_records;
```

### 选项与对象

- `filename` 是必填项，由 PostgreSQL 服务端进程读取。
- `field_lengths` 是逗号分隔的字节长度列表；外部表本身仍必须只有一个列。
- `trim` 会移除每个字段首尾的 ASCII 空格。
- `record_separator` 接受 `none`、`lf`、`cr` 或 `crlf`，分别给每条物理记录增加零、一或两个字节。
- `encoding` 校验字段字节并转换到服务端编码；未设置时使用客户端编码。

扩展会安装 `file_fixed_length_fdw` 外部数据包装器及其处理器和校验器。它只实现扫描与重新扫描，没有写入路径。

### 运维说明

由于 `filename` 会授予服务端文件访问能力，只有超级用户可以设置或修改外部表选项。实现要求每次读取都得到一条完整的定长记录；末尾不完整的短记录会报错。长度描述的是字节而非字符，因此应谨慎测试多字节编码。上游 control 文件的注释误提 SNMP，但实现与仓库都明确这是定长文件包装器。
