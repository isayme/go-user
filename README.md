## User
A simple user restful API service.

## APIs
### Signup
```
POST /v1/users/signup
// body
{
  "email": "your email",
  "username": "your username",
  "password": "your password"
}
```

### Login
```
POST /v1/users/login
// body
{
  "email": "your email",
  "password": "your password"
}
```

### Me
```
GET /v1/users/me
```
Provide access token in HTTP Header Authorization.

## Dev
> make dev

## Docker image
> make image

## MongoDB 
> db.users.createIndex({ email: 1 }, { unique: true })

## TODOs
- email verify
- reset password
