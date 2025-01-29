
## Get Started
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
