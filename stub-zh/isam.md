## 用法

来源：

- [isam.control](https://github.com/jpmoll/bd2/blob/afb0845e259dcaa3af8715e9351ffd67ce5d2d97/isam.control)
- [README.md](https://github.com/jpmoll/bd2/blob/afb0845e259dcaa3af8715e9351ffd67ce5d2d97/README.md)
- [sql/isam--1.0.sql](https://github.com/jpmoll/bd2/blob/afb0845e259dcaa3af8715e9351ffd67ce5d2d97/sql/isam--1.0.sql)
- [src/isam_am.c](https://github.com/jpmoll/bd2/blob/afb0845e259dcaa3af8715e9351ffd67ce5d2d97/src/isam_am.c)

`isam` 是面向 PostgreSQL 18 的教学索引访问方法，支持单个 INTEGER 键的等值／范围查询以及带溢出页的插入；忽略 NULL 键，允许重复键。

### 核心用法

```sql
CREATE EXTENSION isam;
CREATE TABLE isam_sample (id integer);
INSERT INTO isam_sample VALUES (1), (2), (3);
CREATE INDEX isam_sample_idx ON isam_sample USING isam (id);
SELECT * FROM isam_sample WHERE id BETWEEN 1 AND 2;
SELECT isam_info('isam_sample_idx');
LOAD 'isam';
SHOW isam.fill_factor;
```

### 运行边界

需要超级用户安装。填充率默认 70，范围为 10–100。索引使用 PostgreSQL 数据目录下的独立文件，绕过共享缓冲区和 WAL，不提供崩溃恢复。读写采用索引级锁，不保证有序返回，不支持 UNLOGGED 表。重建可减少溢出链，但删除或重建索引会留下需人工清理的旧文件。仅适合可恢复数据上的实验；未找到明确许可证。
