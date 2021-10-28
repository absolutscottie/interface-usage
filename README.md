### Using Interfaces for Easy Mocks and Maintainable Code
This example demonstrates the effectiveness of using an interface to define behavior to accomplish the following:
* easy to mock interface behavior via the use of code generation (i.e. [mockery](https://github.com/vektra/mockery))
* separation of data access interfaces and underlying implementation 

## The Interface
Check out the code in `/internal/dao/dao.go` which defines an interface. Using `mockery` we can easily generate the mock found in `/internal/dao/mock_DAO.go` with the following command.

```
brew install mockery
mockery --all --inpackage
```

## The implemenation
The other advantage of using an interface is being able to separate the interface from the implementation. The code in `/internal/postgres` is a placeholder implementation of a postgres database which implicitly implements the DAO interface. 

This allows the implemenation to be swapped out in the event that we find a better option with minimal pain since the interface remains unchanged.