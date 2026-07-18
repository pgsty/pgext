## Usage

Sources:

- [Project README](https://github.com/ArturFormella/tcn_json/blob/b4d5bcd733905a28c7d3401242861e1333202fea/README.md)
- [Extension control file](https://github.com/ArturFormella/tcn_json/blob/b4d5bcd733905a28c7d3401242861e1333202fea/tcn_json.control)
- [Version 1.1 SQL API](https://github.com/ArturFormella/tcn_json/blob/b4d5bcd733905a28c7d3401242861e1333202fea/tcn_json--1.1.sql)
- [Trigger implementation](https://github.com/ArturFormella/tcn_json/blob/b4d5bcd733905a28c7d3401242861e1333202fea/tcn_json.c)

`tcn_json` 1.1 is a modified form of PostgreSQL's triggered-change-notification extension. An AFTER row trigger sends a transactional notification containing a label, an operation code, and a JSON-like object of the row's primary-key values.

### Publish row-change notifications

The target table must have a valid primary key. The first trigger argument selects the channel and the optional second argument replaces the table-name label:

```sql
CREATE EXTENSION tcn_json;

CREATE TRIGGER orders_tcn_json
AFTER INSERT OR UPDATE OR DELETE ON orders
FOR EACH ROW
EXECUTE FUNCTION triggered_change_notification_json('orders_changes', 'orders');

LISTEN orders_changes;
```

With two trigger arguments, a typical payload has an envelope such as `"orders",I,{"id":"42"}`. With no arguments, the channel is `tcn_json` and the table name is used as the label. The operation is `I`, `U`, or `D`.

### Notification boundary

Delivery follows PostgreSQL LISTEN/NOTIFY semantics: notifications become visible only after commit, duplicate payloads in one transaction may be folded, disconnected consumers do not receive history, and payload size is limited. Use a durable outbox or logical decoding when consumers require replay, acknowledgement, or guaranteed delivery.

The trigger sends only primary-key values, not changed columns or a complete row. It raises an error on a table without a primary key, which aborts the modifying statement. Treat the payload as this extension's text envelope rather than as a standalone JSON document, validate it before parsing, and remember that primary-key updates and unusual string values need explicit application tests. The reviewed source targets PostgreSQL 9.5 and 9.6-era build instructions and publishes no current compatibility matrix.
