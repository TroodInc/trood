# Trood Metarepository

Welcome to the Trood Metarepository! Here, we gather and share tools, libraries, and resources to empower developers worldwide to build faster, better, and more efficiently.

Trood is an AI-powered technical director that manages the launch and further development of custom software applications. We automate critical processes like market research, product documentation, design, code generation, deployment, and integration—helping businesses and individuals bring their ideas to life with less risk and greater efficiency.

For the community, Trood provides a platform to collaborate, share, and showcase innovative projects while enabling developers to contribute to something greater. By participating, you can:

- Connect with like-minded peers and teams worldwide.
- Gain recognition for your work and expertise.
- Contribute to open-source projects that are changing how software is built.

Join us in building the future of software development. Whether you're sharing your own libraries, contributing to existing projects, or learning from others, there's a place for you here.

## [DOCS](docs/)

## Submodules

This repository contains the following submodules:

| Repository Name                   | Path                            | URL                                                                              | Description                                                  |
| --------------------------------- | ------------------------------- | -------------------------------------------------------------------------------- | ------------------------------------------------------------ |
| **blockchain/TRST0**              | `blockchain/TRST0`              | [blockchain/TRST0](https://github.com/TroodInc/TRST0)                            | Trood Revenue Share Token                                    |
| **business-objects/ORM**          | `business-objects/ORM`          | [business-objects/ORM](https://github.com/TroodInc/custodian)                    | Custorian, fast PostgreSQL ORM                               |
| **infrastructure/auth**           | `infrastructure/auth`           | [infrastructure/auth](https://github.com/TroodInc/trood-auth)                    | Trood Auth                                                   |
| **infrastructure/mail**           | `infrastructure/mail`           | [infrastructure/mail](https://github.com/TroodInc/trood-mail)                    | Mail Service                                                 |
| **infrastructure/fileservice**    | `infrastructure/fileservice`    | [infrastructure/fileservice](https://github.com/TroodInc/trood-fileservice)      | A lightweight file service                                   |
| **front-end/trood-core-template** | `front-end/trood-core-template` | [front-end/trood-core-template](https://github.com/TroodInc/trood-core-template) | React/Redux Front-End Boilerplate for Trood Business Objects |

## Get Involved

We believe in the power of collaboration. Here's how you can participate:

- **Contribute**: Submit pull requests, report issues, and improve the libraries.
- **Share**: Publish your own libraries in the Trood Metarepository to reach a global audience.
- **Collaborate**: Join discussions and work on projects with a vibrant community of developers.

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

### Useful Links

- [Trood Team Wiki](https://trood.com/teamspace): Access detailed documentation and resources to get started.
- [Join Trood Community](https://trood.com/launchpad): Connect with peers and explore opportunities to collaborate.

### Contact Us

If you have any questions or want to learn more, feel free to reach out or join our community channels. Together, we can redefine what's possible in software development.
