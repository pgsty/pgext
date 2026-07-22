## 用法

来源：

- [官方 README](https://github.com/shestero/pglas/blob/main/README.md)
- [官方 Rust 实现](https://github.com/shestero/pglas/blob/main/src/lib.rs)
- [官方 control 文件](https://github.com/shestero/pglas/blob/main/pglas.control)
- [官方 LAS 2.0 规范](https://www.cwls.org/wp-content/uploads/2017/02/Las2_Update_Feb2017.pdf)

`pglas` 从 PostgreSQL 服务器文件系统读取 Log ASCII Standard（LAS）2.0 测井文件。它可以报告文件空值标记、列出曲线名称，并返回一条曲线的深度/数值对。这里是测井 LAS 文件，不是同样使用 `.las` 后缀的 LiDAR 点云文件。

### 核心流程

由超级用户创建扩展，然后查询服务器本地 LAS 2.0 文件：

```sql
CREATE EXTENSION pglas;

SELECT las_na('/srv/well-logs/A10.las');

SELECT *
FROM las_curves('/srv/well-logs/A10.las');

SELECT *
FROM las_curve('/srv/well-logs/A10.las', 'Gamma')
ORDER BY "DEPT";
```

`las_curve` 把所选曲线与 `DEPT` 曲线配对，实现假定该深度曲线是文件第一条曲线。解析器识别 LAS 空值标记后，会将对应值返回为 SQL `NULL`。

### API

- `las_na(file text) RETURNS double precision`：文件声明的 N/A 标记。
- `las_curves(file text)`：返回 `(IDX bigint, CURVE text)`，枚举曲线头。
- `las_curve(file text, curve text)`：返回所选曲线的 `(DEPT double precision, VAL double precision)`。

### 注意事项

函数接受任意服务器端路径，并以 PostgreSQL 服务器进程权限打开文件。应只向可信角色授予 `EXECUTE`，并把可读文件放在专用目录；否则该 API 可能泄露服务器本地数据或文件系统错误细节。

只支持 LAS 2.0。实现假定第一条曲线是深度，并依赖 `lasrs` 解析器，因此把结果写入永久表前，应验证头部、空值标记、曲线长度、区分大小写的名称及畸形记录。打包的 pgrx 构建还必须匹配 PostgreSQL 大版本。
