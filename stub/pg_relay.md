## Usage

Sources:

- [Official upstream README](https://gitlab.com/pebble-it/pg_relay/-/blob/main/README.md)
- [Official extension control file](https://gitlab.com/pebble-it/pg_relay/-/blob/main/pg_relay.control)
- [Official project page](https://gitlab.com/pebble-it/pg_relay)

`pg_relay` — pg_relay lets your PostgreSQL database run a SQL action in response to an event — reliably, automatically, and with every outcome recorded. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_relay;

-- Simple: pass an order ID
SELECT pgrelay.notify('new_order', NEW.id::text);
-- Action: SELECT warehouse.reserve_stock($1::bigint)

-- Complex: pass multiple values as JSON
SELECT pgrelay.notify('order_shipped',
    json_build_object('order_id', NEW.id, 'tracking', NEW.tracking_code)::text);
-- Action: SELECT notify_customer.send_shipment_email($1)
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
