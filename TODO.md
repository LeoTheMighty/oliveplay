# TODO

Below is a high-level checklist of what remains to be done in this project. As you complete tasks, check them off!

## General

- [x] Finalize Docker Compose file to wire up all services properly
- [x] Set up environment variables and `.env` files for local dev vs. production
- [ ] Provide documentation on environment configurations

## App (Frontend)

- [ ] Implement basic UI skeleton for user login, cat profile, feeding schedule
- [ ] Integrate with API endpoints for real-time data (diet, analytics, etc.)
- [ ] Write tests for basic user flows and ensure lint rules are followed

## API (Backend)

- [ ] Define endpoints for cat profile and feeding schedule management
- [ ] Implement authentication and authorization layer (JWT or session-based)
- [ ] Integrate with Migrator service to ensure database schema is up to date
- [ ] Write tests to cover core business logic; ensure lint rules pass

## Worker (Background Jobs)

- [ ] Plan out tasks such as notifications, feed scheduling checks, data processing
- [ ] Implement queues or scheduling logic for real-time or delayed tasks
- [ ] Write tests for long-running processes and concurrency checks

## Migrator (Database Service)

- [ ] Configure all migration files to set up the initial schema
- [ ] Add migrations for cat profiles, feeding logs, user accounts
- [ ] Validate the approach for rolling back failed migrations
- [ ] Write tests to ensure migrations run correctly in different environments

## Controller (On-Board Controller in Embedded Rust)

- [ ] Set up build pipeline for embedded Rust platform
- [ ] Implement hardware interfaces to control motors, sensors, etc.
- [ ] Expose a robust set of tests to ensure reliability (simulate hardware)
- [ ] Document how to flash firmware onto the hardware device

## Infrastructure & DevOps

- [ ] Build continuous integration (CI) pipeline for automated tests and lint checks
- [ ] Integrate Nx caching strategies for faster builds/tests
- [ ] Establish a continuous delivery (CD) workflow to push images to a container registry
- [ ] Set up monitoring and logging solutions for production

## Documentation & Community

- [ ] Create user guide for cat owners (possibly integrated within the App UI)
- [ ] Write contributor guidelines (CONTRIBUTING.md)
- [ ] Add code of conduct
- [ ] Plan for an open-source license

---

**Keep an eye on**: This list is a living document. As we refine the project and gather feedback, new tasks may be added or existing tasks reprioritized. Always remember: a well-fed cat is a happy cat! 