## Capabilities of Custodian

### Broad data search and selection capabilities
- RQL queries (filters, sorting, pagination).
- Built-in `only/exclude` parameters for selective field output.
- Depth management for nested relationships.

These mechanisms allow efficient collection of complex data in a "tree" structure in one request, saving effort on writing manual joins.

### Relationship management (including onDelete and generic relationships)
- Multiple deletion modes (cascade, restrict, setNull).
- Support for "generic" references, where one record can reference multiple potential object types.

This provides great flexibility in modeling: you can describe polymorphic relationships between objects without additional "manual" code.

### Calling external services on data changes (actions)
- Ability to "subscribe" to create/update/delete events and trigger arbitrary HTTP endpoints.
- Inclusion of required fields (`includeValues`) when sending, including support for generic logic (various types of related objects).

This is convenient when you need to automatically integrate with other systems (notify CRM, trigger calculations, etc.) without writing your own "middleware."

### Fine support for bulk operations
- POST, PATCH, DELETE multiple records in one request.
- Full transactionality (either all records are changed, or a rollback occurs).

For large volumes of data, this reduces network overhead and simplifies code.

### Optimistic locking (CAS)
- Prevents conflicts when editing the same record in parallel.
- If CAS is not needed, it can be disabled.

### Convenient migrations
- Schema management through a full-fledged migration mechanism (history, branching, rollbacks).
- Generation and application of migrations without manual SQL writing.
- Ability to combine migrations with metadata (when you need to quickly change the structure on the fly).

## Additional capabilities

### Metadata cache management
- Custodian has a caching mechanism, which avoids recreating object structures on every request.
- The cache is only reset when the schema changes (when you change metadata).

This provides a balance between flexibility (dynamic schemas) and performance.

### Event model and notifications
- When performing CRUD operations, Custodian can call "actions," Mail Service can process new email triggers, and File Service can update readiness status.
- You can set up chains like "after creating a record → send a notification to an external service" without writing manual "observer" code.

## Example usage scenario

### Describing a model in Custodian:
```json
{
  "name": "order",
  "key": "id",
  "fields": [
    {
      "name": "id",
      "type": "number",
      "optional": true,
      "default": { "func": "nextval" }
    },
    { "name": "clientName", "type": "string", "optional": false },
    { "name": "status", "type": "string" }
  ]
}
```
- Sent this to `/custodian/meta/order` → the table and endpoints `/custodian/data/order` are created automatically.

### Files:
- Uploaded an invoice template to File Service → `/api/v1.0/templates/`.
- When generating an invoice: `/api/v1.0/files/from_template` → the response contains a ready-made PDF with order data.

### Mail:
- Created a Mailbox `/mailboxes` (SMTP), sent an email to the client along with the PDF from File Service.
- Received incoming mail via `/mailboxes/<id>/fetch`, processed it in an external service.

### Rights via Auth Service (ABAC):
- Set up a rule: "A user with the manager role can only retrieve and update their own orders."
- When accessing `/custodian/data/order`, filtering/checking is automatically applied.

**Result**: You don’t manually write CRUD, SQL, file uploads, SMTP scripts, etc. — everything is ready to use and integrate.
