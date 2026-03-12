


## 用法

> [omni_cloudevents: CloudEvents 支持](https://docs.omnigres.org/omni_cloudevents/cloud_events/)

`omni_cloudevents` 扩展支持在 SQL 中创建、验证和发布标准化的 CloudEvents。它是一个模板化扩展。

### 初始化

```sql
SELECT omni_cloudevents.instantiate(schema => 'omni_cloudevents');
SET search_path TO omni_cloudevents, public;
```

### 创建事件

```sql
SELECT cloudevent(
    id     => gen_random_uuid(),
    source => 'https://service.com/endpoint',
    type   => 'user.login'
);
```

可选参数：`datacontenttype`、`dataschema`、`subject`、`ts`（默认为当前时间）、`data`（任意 PG 类型）、`specversion`（默认 `1.0`）。

### 发布事件

```sql
SELECT publish(
    cloudevent(
        id     => gen_random_uuid(),
        source => 'https://api.yourservice.com/sys',
        type   => 'file.uploaded',
        data   => 'data-lake-bucket-123'::text
    )
);
```

已发布的事件持久化在 `cloudevents_egress` 表中。

### NOTICE 发布器

```sql
SELECT omni_cloudevents.create_notice_publisher();
-- 观察未提交的事件：
SELECT omni_cloudevents.create_notice_publisher(publish_uncommitted => true);
-- 删除发布器：
SELECT omni_cloudevents.delete_publisher(name);
```

事件生成包含完整事件详情的 JSON 格式 NOTICE 消息。
