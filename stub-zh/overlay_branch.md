## 用法

来源：

- [上游来源 1](https://github.com/postgamma/gammabranch/blob/3aeda5c228c9ab72373a1774b981f58cceb871b2/README.md)
- [上游来源 2](https://github.com/postgamma/gammabranch/blob/3aeda5c228c9ab72373a1774b981f58cceb871b2/overlay_branch.control)

`overlay_branch` 0.1 以 Gamma Branch 名义分发，使用 MVCC 增量覆盖存储实验性的草稿变更。

### 核心流程

加入预加载项后需重启，同时保留其他预加载库。`create_branch` 创建草稿，`use_branch` 为会话选择分支，`apply_branch` 在冲突检查后合并。示例返回 MAIN，但不应用草稿。

```ini
shared_preload_libraries = 'overlay_branch'
```

```sql
CREATE EXTENSION overlay_branch;
CREATE TABLE branch_products (id integer PRIMARY KEY, price integer);
INSERT INTO branch_products VALUES (1, 10);
SELECT create_branch('draft');
SELECT use_branch('draft');
UPDATE branch_products SET price = 20 WHERE id = 1;
SELECT use_branch(NULL);
SELECT * FROM branch_products;
```

### 运行边界

仅面向 PostgreSQL 17，且只应使用可丢弃环境。没有预加载时，写入重定向会静默失效，写入直接进入 MAIN。默认仅支持具有单列主键的普通堆表；分区表、外部表、外键及其他不支持的约束会被拒绝。MAIN 的并发变更可能使合并失败，快照模式有时间限制，并可能将会话强制切回 MAIN。
