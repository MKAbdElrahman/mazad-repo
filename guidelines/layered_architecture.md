 Four-layer architecture  backend API

1. **Handler Layer (`handler`):**
   - Receives incoming API requests.
   - Handles request parsing and validation.
   - Invokes appropriate services to process the business logic.

2. **Service Layer (`service`):**
   - Contains the core business logic and application-specific rules.
   - Implements the actual functionality of the API.
   - May interact with the data access layer to retrieve or update data.

3. **Data Access Layer (`store`):**
   - Manages the storage and retrieval of data from a database or another data storage mechanism.
   - Performs CRUD operations on the database.

4. **Infrastructure Layer (`infra`):**
   - Provides infrastructure services that support the operation of the API.
   - Includes components for logging, security, configuration, and communication with external services or systems.

In this architecture, the handler layer handles the incoming requests, delegates the processing to the Service layer, which contains the core business logic. The Service layer may interact with the Data Access layer to access the database. The Infrastructure layer provides necessary support services.