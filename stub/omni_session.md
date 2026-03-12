


## Usage

> [omni_session: Session management](https://docs.omnigres.org/omni_session/session_management/)

The `omni_session` extension provides standardized session management, creating an unlogged `omni_session.sessions` table for active sessions. Works primarily with the HTTP stack.

### Session Handler for HTTP

**Request handler** -- Extracts session UUID from a `_session` cookie:

```sql
SELECT omni_session.session_handler(request);
-- Optional: session_handler(request, cookie_name => 'my_session')
```

Creates a new session if the cookie is missing or invalid, sets the `omni_session.session` transaction variable.

**Response handler** -- Sets the session cookie on the response:

```sql
SELECT omni_session.session_handler(response);
```

Cookie parameters:

| Parameter     | Default     | Description                    |
|:--------------|:------------|:-------------------------------|
| `cookie_name` | `_session`  | Cookie identifier              |
| `http_only`   | `true`      | Prevents JavaScript access     |
| `secure`      | `true`      | HTTPS-only                     |
| `same_site`   | `Lax`       | Cross-site behavior            |
| `domain`      | `null`      | Domain scope                   |
| `max_age`     | `null`      | Lifetime in seconds            |
| `expires`     | `null`      | Expiration timestamp           |
| `path`        | `null`      | URL path scope                 |

### Direct Session ID Handling

```sql
-- Create or validate a session:
SELECT omni_session.session_handler(null::omni_session.session_id);
-- Returns new session ID and sets omni_session.session transaction variable
```
