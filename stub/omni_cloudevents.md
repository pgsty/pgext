


## Usage

> [omni_cloudevents: CloudEvents support](https://docs.omnigres.org/omni_cloudevents/cloud_events/)

The `omni_cloudevents` extension enables creation, validation, and publication of standardized CloudEvents from SQL. It is a templated extension.

### Setup

```sql
SELECT omni_cloudevents.instantiate(schema => 'omni_cloudevents');
SET search_path TO omni_cloudevents, public;
```

### Creating Events

```sql
SELECT cloudevent(
    id     => gen_random_uuid(),
    source => 'https://service.com/endpoint',
    type   => 'user.login'
);
```

Optional parameters: `datacontenttype`, `dataschema`, `subject`, `ts` (defaults to now), `data` (any PG type), `specversion` (default `1.0`).

### Publishing Events

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

Published events are persisted in the `cloudevents_egress` table.

### NOTICE Publisher

```sql
SELECT omni_cloudevents.create_notice_publisher();
-- To observe uncommitted events:
SELECT omni_cloudevents.create_notice_publisher(publish_uncommitted => true);
-- Delete a publisher:
SELECT omni_cloudevents.delete_publisher(name);
```

Events generate JSON-formatted NOTICE messages with full event details.
