# Open Stash API

This is the repo for the API for [Open Stash](https://github.com/emcassi)

## User Creation

### Name Rules

- Name must be at most 30 characters

### Password Rules

- Password must be at least 8 characters
- Password must be no more than 72 characters (This limitation is due to bcrypt)
- Password must contain at least one lowercase character
- Password must contain at least one uppercase character
- Password must contain at least one number
- Password must contain at least one special character (!@#$%^&\*-\_=+,.<>/?\\|`~()[]{})
