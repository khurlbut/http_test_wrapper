This small program demonstrates the usage of panic, defer and recover.

You can kinda think of these as throw, try and catch although they feel a bit more fluid/flexible.

If you remove the 'recover' method then the panic will get passed up the call stack until the program crashes. 
This is anologous to not catching a throw.
