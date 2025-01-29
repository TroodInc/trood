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