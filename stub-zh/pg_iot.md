## 用法

来源：

- [README.md](https://github.com/pierreforstmann/pg_iot/blob/862938f0f7ad3824b3bd125154483b91bd720356/README.md)
- [pg_iot.control](https://github.com/pierreforstmann/pg_iot/blob/862938f0f7ad3824b3bd125154483b91bd720356/pg_iot.control)
- [pg_iot--0.0.1.sql](https://github.com/pierreforstmann/pg_iot/blob/862938f0f7ad3824b3bd125154483b91bd720356/pg_iot--0.0.1.sql)
- [pg_iot.c](https://github.com/pierreforstmann/pg_iot/blob/862938f0f7ad3824b3bd125154483b91bd720356/pg_iot.c)

`pg_iot` 将指定表标记为仅允许插入的表。0.0.1 版通过服务器钩子拒绝对已登记表执行 UPDATE、DELETE、MERGE、TRUNCATE 和 ALTER TABLE，同时保留 INSERT 与 SELECT。上游报告已在 PostgreSQL 14–19 上验证。

### 启用并登记表

将库加入现有预加载列表，然后重启 PostgreSQL：

```ini
shared_preload_libraries = 'pg_iot'
```

以超级用户在数据库中安装扩展，并按模式名和表名登记已有表：

```sql
CREATE EXTENSION pg_iot;
CREATE TABLE public.audit_log (event text);
SELECT iot.pg_iot_set('public', 'audit_log');
INSERT INTO public.audit_log VALUES ('opened');
SELECT * FROM public.audit_log;
SELECT * FROM iot.tables;
```

扩展不可迁移模式，会创建固定的 `iot` 模式。`iot.pg_iot_set(cstring, cstring)` 返回布尔值，PUBLIC 已被撤销执行权限。`iot.tables` 记录已登记表的模式、表名与 OID。配置预加载列表时应保留其他所需的库。

### 维护边界

文档所述接口没有将仅插入表恢复为普通表的操作。登记前应规划模式变更与数据保留策略。扩展依赖预加载的库及自身目录，不能防止数据库或操作系统管理员篡改归档。常规表权限仍决定谁能读取和插入数据。
