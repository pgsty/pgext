## 用法

来源：

- [官方 README](https://github.com/polydbms/pg_sheet_fdw/blob/3850b6b84196eb9c1e93b7957f0e786a6f5b084b/README.md)
- [扩展 SQL](https://github.com/polydbms/pg_sheet_fdw/blob/3850b6b84196eb9c1e93b7957f0e786a6f5b084b/pg_sheet_fdw--0.1.2.sql)
- [FDW 实现](https://github.com/polydbms/pg_sheet_fdw/blob/3850b6b84196eb9c1e93b7957f0e786a6f5b084b/src/pg_sheet_fdw.c)

`pg_sheet_fdw` 0.1.2 是用于本地 Microsoft Excel `.xlsx` 工作表的只读外部数据包装器。它把一张工作表映射为外部表，并把单元格转换为用户声明的 PostgreSQL 列类型。电子表格仍是服务器端文件；它并不是 Google Sheets 或远程对象连接器。

### 核心流程

```sql
CREATE EXTENSION pg_sheet_fdw;
CREATE SERVER sheets FOREIGN DATA WRAPPER pg_sheet_fdw;

CREATE FOREIGN TABLE monthly_sales (
    sold_on date,
    sku text,
    quantity integer,
    amount numeric
)
SERVER sheets
OPTIONS (
    filepath '/srv/imports/sales.xlsx',
    sheetname 'January',
    skiprows '1',
    numberofthreads '4',
    batchsize '1000'
);

SELECT * FROM monthly_sales WHERE quantity > 0;
```

`filepath` 为必填项，且应使用绝对路径。`sheetname` 默认为第一张工作表，`skiprows` 默认为 0；省略 `numberofthreads` 时由解析器自动选择，显式值限制为 1–10。若省略 `batchsize`，实现会推导一个值，目标是把工作表拆成约 101 批，最小为 1,000 行。

### 数据与执行边界

PostgreSQL 服务器操作系统用户必须能够读取工作簿及路径中的每一级目录。应按工作表列顺序定义列，并使类型与真实单元格匹配。上游示例表明，某些类型不匹配可能被转换或饱和，而不是报错，因此应验证导入结果，并先转换到暂存表后再依赖其值。

FDW handler 只实现扫描回调，没有插入、更新或删除回调。规划器/执行器支持较基础，工作表解析器读取的数据可能远多于 SQL 过滤器最终返回的内容。大型工作簿会消耗后端 CPU 与内存，并可能在事务快照之外独立变化。README 报告了 PostgreSQL 13 测试和 PostgreSQL 16 CI；其他大版本上应同时验证构建兼容性与工作簿行为。
