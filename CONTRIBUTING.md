# Contribution Guide for Organization Repositories

This guide outlines the best practices and guidelines for contributing to our organization's repositories, currently private but planned for open-source release.

## Terminology Definitions

To ensure clarity, the following terms are used consistently throughout this document:

- **MUST**: A requirement that is mandatory; failure to follow will result in rejection of the contribution.
- **MUST NOT**: A strict prohibition; violating this rule may lead to rejection or security concerns.
- **SHOULD**: A recommended practice; not strictly enforced but highly encouraged for consistency and best practices.
- **SHOULD NOT**: A discouraged practice; while not outright prohibited, contributors are urged to avoid it unless justified.

## 1. Repository Structure & Branching

- **Main Branch (`main`)**: This branch **should** be protected; all changes **must** go through Pull Requests (PRs).
- **Development Branch (`develop`)**: Active development **should** happen here before merging into `main`.
- **Feature Branches**: Contributors **should** create a branch per feature/fix, following the format:
  - `feat/auth-improvements`
  - `fix/api-timeout`
  - `chore/dependency-upgrade`
- **Release Branches (`release/x.y.z`)**: These branches **should** be used for staging and testing before merging into `main`.

## 2. Commit Message Guidelines (Conventional Commits)

Commit messages **must** follow the [Conventional Commit](https://www.conventionalcommits.org/) format:

```
type(scope): message
```

### **Allowed Commit Types:**

- **feat**: A new feature
- **fix**: A bug fix
- **docs**: Documentation updates
- **style**: Code style changes (formatting, missing semi-colons, etc.)
- **refactor**: Code changes that neither fix a bug nor add a feature
- **perf**: Performance improvements
- **test**: Adding or updating tests
- **chore**: Maintenance tasks (e.g., dependency updates, build scripts)
- **ci**: CI/CD-related changes

### **Examples:**

✅ `feat(auth): add JWT token support`

✅ `fix(api): resolve timeout issue in user service`

✅ `chore(deps): update dependency axios to v0.24.0`

❌ `fixed bug`

❌ `updated code`

## 3. Contribution Process

1. **Clone & Setup**

- You can use HTTP or SSH for cloning:
  - HTTP is recommended if you do not have SSH keys set up or are using a temporary environment.
  - SSH is recommended for frequent contributors as it provides secure, password-less authentication.
- Using HTTP:

     ```sh
     git clone https://github.com/genefriendway/<repo-name>.git
     ```

- Using SSH([Set up SSH keys](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/about-ssh))

     ```sh
     git clone git@github.com:genefriendway/<repo-name>.git
     ```

- Checkout a new feature branch:

     ```sh
     cd <repo-name>
     git checkout -b feat/your-feature
     ```

2. **Make Changes & Commit**

   ```sh
   git add .
   git commit -m "feat(module): describe change"
   ```

3. **Push & Create PR**

   ```sh
   git push origin feat/your-feature
   ```

   - Open a PR against `develop` (or `main` if applicable).
   - Fill out the PR template.
   - Request review from relevant maintainers.

## 4. Code Review Guidelines

- PRs **should** be small, focused, and well-documented.
- Contributors **must** follow the repository’s coding standards (e.g., linting, formatting, type safety).
- Necessary unit/integration tests **should** be written.
- All reviewer comments **must** be addressed before merging.
- Draft PRs **should** be used for work-in-progress submissions.

## 5. Code Review Process

1. **Automated Checks**: Code **must** pass all CI/CD tests (linting, formatting, unit tests).
2. **Reviewer Assignment**: A maintainer or senior developer **should** be assigned to review the PR.
3. **Feedback & Revisions**:
   - Reviewers **should** provide feedback through inline comments.
   - Contributors **must** address all comments and push updates before requesting another review.
4. **Approval & Merging**:
   - Once approved, the PR **must** be merged by a maintainer.
   - **All PRs must be squash merged.**
   - The final commit message **must** follow Conventional Commit guidelines.
   - Feature branches **should** be deleted after merging.

## 6. Security & Compliance

- Secrets (API keys, passwords, credentials) **must not** be committed.
- Environment variables **should** be stored in `.env` files and **must** be added to `.gitignore`.
- Security best practices **must** be followed when handling data.

## 7. Issues & Discussions

- GitHub Issues **should** be used to report bugs, request features, and track progress.
- Contributors **should** engage in discussions and document architectural decisions when necessary.
- If unsure about an approach, contributors **must** discuss before implementing.

## 8. Preparing for Open Source Release

- Documentation (README, setup guides, API references) **must** be clear and up-to-date.
- Licenses (MIT, Apache, etc.) and contributor agreements **must** be in place.
- Private/internal references **must** be cleaned up before making the repository public.
- Test coverage **should** be improved and CI/CD workflows **should** be added.
