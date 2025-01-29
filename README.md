## What is it and why is it needed?

Trood Core is a set of microservices that provide various functional blocks for typical business applications. The project is focused on a quick start when creating corporate solutions: CRM systems, internal portals, document management systems, etc. Instead of writing typical blocks from scratch, you get ready-made services:

### Custodian (Data structure management and interaction with the database):
- Generation of REST endpoints and database tables (PostgreSQL) from JSON descriptions.
- Migration mechanism (with history and branching).
- Bulk operations (mass addition, updating, deletion).
- Integration with external systems via Webhooks.

### File Service (File management and processing):
- Support for metadata, extensions, types, and templates.
- Document templates and on-the-fly file generation.
- Integration with ABAC for file access management.

### Auth Service (Authorization and access management):
- High flexibility in access management thanks to ABAC.
- Easy integration with other systems via REST API.
- Detailed control over users and roles.
- Built-in security mechanisms (multi-factor authentication, tokens).

### Mail Service (Email handling):
- Mailbox management (SMTP/IMAP).
- Sending and receiving emails, attachments.
- Working with folders, email threads, contacts.

Together, these services cover a wide range of application needs: data storage, file handling, email management, and access rights configuration — all without reinventing the wheel or manually writing repetitive code.

## How can this help developers?

### Quick start and reduced routine
- **CRUD out of the box**: Describe objects in JSON — Custodian creates tables and REST endpoints.
- **Bulk operations and search (RQL)**: No need to manually write SQL queries or tweak filter logic.
- **Minimal code duplication**: Quickly define fields and relationships through metadata. Custodian automatically synchronizes the database schema.

### Ready-made mail and file modules
- **Mail Service** simplifies sending/receiving emails, including attachments, folders, and threads. It speeds up typical use cases (mailings, notifications, capturing incoming emails).
- **File Service** solves file storage and document generation from templates. Developers don’t need to write logic for uploading/storing files, etc.

### Flexible authorization (ABAC)
- Fine-tuned access control without manual "if user in role" checks: Auth Service allows setting rules based on user attributes (subject), object attributes (object), and context (context).
- Example: "A manager can only see their orders, while an administrator can see all."

### Scalable microservice architecture
- Each service (Custodian, File, Mail, Auth) has its own Docker container/repository.
- You can scale narrowly, for example, if you need to handle high email load without affecting other blocks.

### Convenient integrations
- REST API in all services — easily integrates with any Frontend framework (React, Vue, Angular) or mobile applications.
- External calls on create/update/delete events (hooks, actions).

## What’s next?
If you want to test Trood Core in practice:
1. Launch the services (Docker or locally), connect Postgres.
2. Describe one or two entities in Custodian (via JSON metadata) and try adding a few records.
3. Set up File Service, upload a file, and insert its link into a Custodian object.
4. Connect Auth Service (2FA, ABAC), test access restrictions.
5. Mail Service: set up a mailbox, send yourself an email, process incoming mail.

You’ll see that many things that usually require several libraries and significant effort "just work" almost "out of the box."

If you need to convince clients that this is faster and cheaper than custom development or limited no-code solutions, just demonstrate:
- **Speed**: How quickly you can set up basic CRUD and file/mail logic.
- **Flexibility**: If needed, you can write additional code, create services around it, without being limited by a "black box."

Use Trood Core as a foundation, and AI tools (Cursor, ChatGPT, etc.) — to speed up writing auxiliary logic, tests, frontend, etc. This combination gives developers the maximum advantage in time and quality.
