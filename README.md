# User Management Api

## How to run the application

1. Clone the application with `git@github.com/MukulManni/UserApi.git`
2. Once the application is cloned, create a local postgresql database and change the url strings with local strings
3. Unzip the archive and change directory to UserApi
4. Create go module with command:
```shell
    go mod init UserApi
```
5. Then build application with:
```shell
    go build
```
6. Run:
```shell
    sudo ./UserApi
```

## Endpoints Description

### Create User
```shell
    curl -X POST -d '{"name":"TestUser","dob":"13-07-2002","address":"Utopia","description":"Hello from User"}' https://userapi-1.herokuapp.com/create
```

### Get User Details
```shell
    curl https://userapi-1.herokuapp.com/get/2
```

### Update User
```shell
    curl -X PUT -d '{"name":"TestedUser"}' https://userapi-1.herokuapp.com/update/2
```

### Delete User
```shell
    curl -X DELETE https://userapi-1.herokuapp.com/delete/2
```

## Test Driven Development Description

1. Test to check the working of get user details request
2. Test to check the working of create user request
3. Test to check the working of update user details request
4. Test to check the working of delete request

To run all the unit test cases, please do the following -
`go test -v`

## Hope You like this test driven Api written in Go with Gin