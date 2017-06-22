# gdebug

**Description:**

gdebug provides the following basic utilities to help in debugging go:

- Dump:
    Print an object 

- DumpComparison:
    Print a side by side comparison of two objects.    

- The dumps recursively expand the object
- The output is indented
- Note: gdebug uses json.marshal to stringify object. So only *public* members of a struct will be printed

**Installation:**

```
go get github.com/vigneshmr/gdebug
```

or if you use vendored packages, add the following line to your `glide.yaml` and `glide i`:

```yaml
- package: github.com/vigneshmr/gdebug
```

**Usage / Examples:**

- **Dump:**
    Print an object 

    Signature:
    ```go
    func Dump(obj interface{}, msg string)
    
    ```
    **Usage:**
    ```go
    Dump(testStruct{}, "msg")
    ```

- **DumpComparison:**
    Print a side by side comparison of two objects.    

    Signature:
    ```go
    func DumpComparison(expected interface{}, actual interface{})
    ```
    
    **Usage:**
    ```go
    DumpComparison(testStruct{}, testStruct{})
    ```

- **DumpComparisonIfDifferent:**
    Print a side by side comparison of two object if they differ in value.
    
    Signature:
    ```go
    func DumpComparisonIfDifferent(expected interface{}, actual interface{})    
    ```
    
   **Usage:**
    ```go
    DumpComparisonIfDifferent(testStruct{}, testStruct{})
    ```
