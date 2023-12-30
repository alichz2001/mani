# Storage and Gateway
#### _test project for Mani_


## Services

- Simple gateway and authorization server
- Storage service for store files

You can find project description in this [link](https://github.com/alichz2001/mani/blob/master/Take_home_challenge_for_Golang_at_Mani.pdf).


## Project requirements
- [x] Store service:
  - [x] Store files in filesystem.
  - [x] Store metadata(name, tags, type) in database.
  - [x] Fetch files by name.
  - [x] Fetch files by tags.
  - [ ] (optional) Encryption for files.
  - [x] (optional) Upload file size limit.
  - [x] (optional) Storage size limit.
- [x] Retrieval service:
  - [x] Authentication.
    - [x] Register user (sign up).
    - [x] Login user (sing in).
  - [x] Proxy requests to upstream service(Store service).
- [ ] (optional) Tests.
- [ ] (optional) Logging.
- [ ] (optional) Handle fault scenarios.
  - [x] Health-checking

## Used technologies
- **Fiber**: used as webserver and gateway. [link](https://docs.gofiber.io/)
- **ArangoDB**: used as primary database. [link](https://arangodb.com/)
- **Docker & Docker-compose**: for containerization and dev and test mode run.
- **Makefile**: for manage commands.


## Notes
- In this project, I've employed a layered architecture with three logical layers: Model, Service, and Controller. The simplicity of the current task doesn't necessitate additional abstractions or complex packaging. However, for future development, some refactoring may be beneficial. This structure is well-suited for testing purposes.
- Perhaps, for a production-level 'Store' service, ArangoDB alone might not be sufficient due to anticipated challenges related to Atomicity and Race-conditions. However, for this sample task, it eill be enough.
- Why 'Fiber'? So while both Gin and Echo are excellent frameworks with their respective advantages and drawbacks, my preference for a web server is Fiber. see this [link](https://medium.com/deno-the-complete-reference/go-gin-vs-fiber-hello-world-performance-6863e597b654)


## RUN

```shell
make run
```
-----------------
## API Documentation

**Note:** You can find the full API Postman collection in this [link](https://github.com/alichz2001/mani/blob/master/mani_task.postman_collection.json).

#### 1\. User Authentication

* POST http://<server>/v1/auth/signup
  > Create a new user and receive a JWT as a response.

* POST http://<server>/v1/auth/signin
  > Validate the username and password, and receive a JWT in the response.

#### 2\. File Storage

* POST http://<server>/v1/storage/file/upload
  > Upload a file and save its name and tags in the database.

* GET http://<server>/v1/storage/file/fetch
  > Get a list of files' metadata based on specified filters.
  **Response:**
  - Files with an exact name match OR any file that has one of the requested tags.
  - If there is no match, retrieve the oldest unexpired file.
  - If there is no unexpired file, the response will be an empty list.

* GET http://<server>/v1/storage/files/<file\_name>
  > Serve the file with an exact match to the given name.

## Personal TODO
- [x] rate-limiting and heath-checking for upstream services.
- [ ] make config more useful. Add reloading config options by signal, file monitor or even REST API.
- [ ] make better error handling. create const errors and return error messages in fiber ErrorHandler.
- [ ] use proper context for free unneeded resources.
- [ ] add metadata to proxies requests in gateway. information like user data etc.
- [ ] queue disk IO.
- [ ] add check for static file server to not serve expired files.
- [ ] add interfaces for more usable and maintainable codes.
- [ ] add graceful shutdown scenario.


