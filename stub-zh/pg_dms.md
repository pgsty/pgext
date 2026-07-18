## 用法

来源：

- [锁定提交的上游 README](https://github.com/abris-platform/pg_dms/blob/ec456d7a0f0ded981d7e4c516ccdeb20fd8a1197/README.md)
- [0.0.1 版安装 SQL](https://github.com/abris-platform/pg_dms/blob/ec456d7a0f0ded981d7e4c516ccdeb20fd8a1197/pg_dms--0.0.1.sql)
- [锁定提交的 C 实现](https://github.com/abris-platform/pg_dms/blob/ec456d7a0f0ded981d7e4c516ccdeb20fd8a1197/pg_dms.c)
- [锁定提交的回归示例](https://github.com/abris-platform/pg_dms/blob/ec456d7a0f0ded981d7e4c516ccdeb20fd8a1197/sql/pg_dms_test.sql)
- [锁定提交的扩展控制文件](https://github.com/abris-platform/pg_dms/blob/ec456d7a0f0ded981d7e4c516ccdeb20fd8a1197/pg_dms.control)

pg_dms 0.0.1 是一个特定应用的文档管理数据模型。它实现 pg_dms_id、pg_dms_family 和 pg_dms_ref 类型、跨类型 B-tree 比较、UUID 转换、内嵌操作历史、记录 JSON 与哈希辅助函数，以及注册表导入/导出表和触发器。它不是通用文档存储。

### 标识符与操作示例

```sql
CREATE EXTENSION pg_dms CASCADE;

CREATE TABLE public.directory (
  key pg_dms_id NOT NULL DEFAULT uuid_generate_v4(),
  num integer NOT NULL
);

INSERT INTO public.directory (num) VALUES (1);

UPDATE public.directory
SET key = pg_dms_setaction(
  key,
  100,
  'public.directory'::regclass,
  uuid_generate_v4()
);

SELECT key, pg_dms_getstatus(key), pg_dms_getlevel(key)
FROM public.directory;
```

pg_dms_id 文本值由两个使用下划线分隔的 UUID 组成。新 ID 会记录创建角色和事务开始时间戳。每次 setaction 调用都会向可变长值追加一个操作结构，因此索引键和行大小会随历史增长。

### SQL 安全、模式与兼容性限制

控制文件宣称扩展可重定位，但它的 SQL 将表和函数硬编码到 public，并将操作符硬编码到 pg_catalog。它还会创建名称宽泛的 action_list 和 registry 对象。只应在检查名称和操作符冲突后，安装到隔离的应用数据库；不应移动扩展模式。

注册表辅助函数和触发器会通过直接拼接 JSON 提供的模式名、表名、列名和值来构造动态 INSERT，没有对标识符或字面量做引用。不可信输入可以改变生成的 SQL。应从 PUBLIC 撤销所有函数和注册表权限；在使用参数化且对模式做白名单的代码替换这条导入路径之前，不应使用它。

多个 C 函数被声明为 IMMUTABLE，但它们会读取当前角色、事务时间或当前时钟。规划器可能错误缓存或重排这些调用。记录辅助函数依赖 PostgreSQL 内部元组 API，SQL 会创建 PostgreSQL 10 时代的过时表选项，且没有提供当代主版本兼容矩阵。

仓库自 2018 年后没有更改，也没有许可证文件。应将它视为无维护的历史应用代码：针对精确目标服务器构建，检查每个创建对象和授权，验证二进制排序和哈希，并只在具备独立备份和应用专用回归测试时使用。
