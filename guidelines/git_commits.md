## Commit Message Format

```
<type>(<scope>): <short summary>

<optional body>

<optional footer>
```

- **Type**: Describes the purpose of the commit. Choose one from the following:

  - `feat`: A new feature for the user or a significant change.
  - `fix`: A bug fix.
  - `docs`: Documentation changes.
  - `style`: Code style changes (e.g., formatting).
  - `refactor`: Code refactor without changing its behavior.
  - `test`: Adding or modifying tests.
  - `chore`: Routine tasks, maintenance, or tooling changes.

- **Scope (optional)**: Indicates the module, file, or aspect of the project the commit affects.

- **Short Summary**: A concise description of the change, preferably in the imperative mood (e.g., "Fix bug" instead of "Fixed bug").

- **Extended Description (optional)**: A more detailed explanation of the changes. Include any relevant context, reasoning, or technical details.

- **Footer (optional)**: Additional information, such as breaking changes, issue references, or other related information.

Here's an example:

```
feat(user-auth): add email verification process

- Introduces a new email verification flow for user registration.
- Implements email confirmation endpoint.
- Updates user model to include an 'is_verified' field.
- Adds unit tests for the new functionality.
- Fixes #123

```