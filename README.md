
# Password Manager CLI

This is a simple password manager CLI built in Go, which allows you to manage your passwords from the command line. The CLI provides commands to add, update, delete, find, and list passwords. Password entries are stored using PostgreSQL and passwords are encrypted before storing.

## Prerequisites

- Golang 
- PostgreSQL 

## Features

- **Add Password**: Add a new password entry.
- **Update Password**: Update an existing password entry.
- **Delete Password**: Delete a password entry.
- **Find Password**: Find a password entry based on username or site.
- **List Passwords**: List all password entries.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/iamber12/password-manager-cli
    ```
2. Navigate to the project directory:
    ```sh
    cd password-manager-cli
    ```
3. Build the CLI:
    ```sh
    go build -o password-manager
    ```

## Commands

### Add Command

Adds a new password entry.

- **Flags**:
  - `-u, --username` (required): Username for the password
  - `-p, --password` (required): Password
  - `-s, --site` (required): Site

### Update Command

Updates an existing password entry.

- **Flags**:
  - `-u, --username` (required): Username for the password
  - `-p, --password` (required): Password
  - `-s, --site` (required): Site

### Delete Command

Deletes a password entry.

- **Flags**:
  - `-u, --username` (required): Username for the password
  - `-s, --site` (required): Site

### Find Command

Finds a password entry based on username or site.

- **Flags**:
  - `-u, --username`: Username for the password
  - `-s, --site`: Site

### List Command

Lists all password entries.

## Example

```sh
# Add a new password
./password-manager add -u user1 -p pass123 -s example.com

# Update the password
./password-manager update -u user1 -p newpass123 -s example.com

# Find the password
./password-manager find -u user1

# List all passwords
./password-manager list

# Delete the password
./password-manager delete -u user1 -s example.com
```

---

Happy managing your passwords securely!
