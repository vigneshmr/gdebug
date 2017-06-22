gdebug
======

Description:
============

gdebug provides the following basic utilities to help in debugging go:

- Dump:
    Print an object 

- DumpComparison:
    Print a side by side comparison of two objects.    

- The dumps recursively expand the object
- The output is indented
- Note: gdebug uses json.marshal to stringify object. So only *public* members of a struct will be printed

Installation:
============

```
go get github.com/vigneshmr/gdebug
```

**or** 

if you use vendored packages, add the following line to your `glide.yaml` and `glide i`:

```yaml
- package: github.com/vigneshmr/gdebug
```

**or** if you have a large project and need this import just for adhoc debugging, you can add it directly to your `glide.lock`:

```yaml
- name:     github.com/vigneshmr/gdebug
  vcs:      git
```

Usage / Examples:
=================

Directly calling into package for adhoc debugging:
--------------------------------------------------

### Dump:

Print an object 

**Signature:**
```go
func Dump(obj interface{}, msg string)

```
**Usage:**
```go
Dump(testStruct{}, "msg")
```

### DumpComparison:

Print a side by side comparison of two objects.    

**Signature:**
```go
func DumpComparison(expected interface{}, actual interface{})
```

**Usage:**
```go
DumpComparison(testStruct{}, testStruct{})
```

**Output:**

<img src="https://user-images.githubusercontent.com/1412892/27418738-eb42765e-56d1-11e7-9dad-c753daed6540.png" width=60%>
    
### DumpComparisonIfDifferent:

Print a side by side comparison of two object if they differ in value.

**Signature:**
```go
func DumpComparisonIfDifferent(expected interface{}, actual interface{})    
```

**Usage:**
```go
DumpComparisonIfDifferent(testStruct{}, testStruct{})
```
