# 🐲 Errors

Errors is an alternate version of the standard errors package in Golang. It provides a simple and easy-to-use interface to create domain-specific, HTTP, and plain text errors. It also has all the functionality of the standard errors package

## Features

- **Domain-specific errors**: With Errors, you can create custom error that are specific to your application domain. This means you can provide meaningful error messages that make sense to your users, without having to decipher obscure error codes.

- **HTTP errors**: Errors provides a set of HTTP-specific error types that make it easy to represent common HTTP error codes in a consistent and easy-to-use way. From BadRequestError to InternalServerError, we've got you covered!

- **Plain text errors**: Sometimes you just need a simple error message that doesn't involve HTTP codes or domain-specific error codes. With Errors, you can create plain text errors that get the job done.

- **Lightweight and easy to use**: Errors is a lightweight and easy-to-use package that doesn't require any special setup or configuration. You can start using it right away, without any learning curve.

- **Stack Tracing**: With errors, you can create error chains causing, and It can be traced easily.

## Usage

Using Errors is as easy as 1-2-3:

1. Import package

   ```golang
   import "github.com/ahmetcanozcan/errors"
   ```

2. Use the New function to create errors:

   ```golang
    // Domain-specific error
    ErrSomethingWrong := errors.New("Something went wrong", 1003)

    // Domain-specific with HTTP status
    ErrResourceNotFound := errors.New("Resource not found", 1004, http.StatusNotFound)

    // Plain text error
    err := errors.New("Something went wrong")
   ```

3. Just use them:

   ```golang
    func myFunction() error {
      if err := getSomething(); err != nil {
        return ErrSomethingWrong.WithCause(err)
        // return ErrSomethingWrong
        // return errors.New("something wrong")
        // return errors.New("something wrong").WithCause(err)
      }
    }

    func myHandler(c *fiber.Ctx) error {
      if err := myFunction(); err != nil {
        return err
      }

      return c.SendString("Ok!")
    }

    // see more detail: https://docs.gofiber.io/guide/error-handling/#custom-error-handler
    func errHandler(c *fiber.Ctx, err error) error {
      code := errors.Code(err)
      status := errors.Status(err)
      payload := map[string]any{
        "code": code,
        "message": err.Error()
      }
      return c.Status(status).JSON(payload)
    }

   ```

4. stack tracing - optional

   ```golang
     func errHandler(c *fiber.Ctx, err error) error {
       code := errors.Code(err)
       status := errors.Status(err)
       payload := map[string]any{
         "code": code,
         "message": err.Error()
       }

       stack := errors.StackTrace(err)
       for _, e := range stack {
         fmt.Println("stack err", e)
       }
       return c.Status(status).JSON(payload)
     }
   ```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
