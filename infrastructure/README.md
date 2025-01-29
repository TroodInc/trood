## Capabilities of File Service

### Flexible integration with cloud storage

- Support for S3 (Amazon, MinIO), DigitalOcean Spaces, etc.
- Local disk (in development environments or if there is no cloud).
- Developers don’t need to set up integration themselves — File Service already knows how to store and serve files via URL.

### Document generation from templates

- Django-template (or similar) logic inside, support for PDF/PNG.
- Automatic data substitution from JSON.
- Can cover many use cases for generating acts, invoices, contracts "on the fly."

### Soft deletion mechanism

- When a file is deleted, a `deleted=true` flag is set.
- There is a periodic/manual mechanism for final cleanup (`cleanup`).

Useful for "safe" deletion and the possibility of "restoration" if needed.

## Capabilities of Auth Service (ABAC)

### Attribute-Based Access Control

- You can check access not only by roles but also by user attributes, object attributes, and environment (sbj, obj, ctx).
- Complex conditions with AND/OR, operators (in, not, lt/gt), etc.

For designing "granular" security: "The user sees objects where `object.owner == user.id` or `object.isPublic == true`."

### 2-Factor Authentication

- Ready-made logic for SMS/phone verification (can be extended).
- Management of factors and temporary tokens.

If additional security measures are needed, everything is built-in and available out of the box.

### Field masking mechanism

- You can hide certain fields in the response (e.g., "salary" for regular users).
- Convenient when the same entity is used by different user categories.

## Capabilities of Mail Service

### Connecting external mail servers

- SMTP/IMAP can be configured for any mail provider (Gmail, corporate Exchange, etc.).
- Developers don’t need to solve the problem of "how to communicate with the mail server," everything is implemented "out of the box."

### Grouping emails into threads

- Analysis of `reply-to` headers → merging emails into discussion threads.

This simplifies communication logic and searching for the entire history on a topic.

### Working with folders and contacts

- Creating folders, moving emails.
- Bulk operations, linking emails and contacts.

Mechanics very similar to a "standard email client," but available via API.

### Attachments

- Works in conjunction with File Service (an uploaded file can be immediately attached to an email).
- Management of large files/documents without manual workarounds.
