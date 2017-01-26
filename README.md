# gdebug

gdebug provides the following basic utilities to help in debugging go:

- Dump:
    Print an object 

- DumpComparison:
    Print a side by side comparison of two objects.    

- The dumps recursively expand the object
- The output is indented
- Note: gdebug uses json.marshal to stringify object. So only *public* members of a struct will be printed
