Go_201
====
### 1. Pass by value
Think about the code below:
```
func (p person) updateName(newFirstName string) {
    p.firstName = newFirstName
}

// Todo: update the fname, and the old fname is Jim
jim.updateName("Jimmy")

jim.print() // use the print() function, but I will still get "Jim" as my fname, because this is pass by value
```
Go is a pass by value language, whenever pass a value into a function, go copies that value, so when we update the value it will never change the primary one.<br>

Check the code below: <br>
```
// Todo: update the fname, use pass by reference
jimPointer := &jim
jimPointer.updateName("Jimmy") 
jim.print() // use the print() function


func (pointerToPerson *person) updateName(newFirstName string) {
    (*pointerToPerson).firstName = newFirstName
}

// Structs with receiver functions
func (p person) print() {
    fmt.Printf("%+v", p)
}
```

&variable -> turn the value to address <br>
*pointer  -> turn the address to value <br>
Notice: 1. func (pointerToPerson *person) updateName(newFirstName string) -> the *person means a type of pointer that points at a person


### 2. Pointer shortcut
jim.updateName() -> Go will gloss it
eg.
```
jim.updateName("Jimmy") 
jim.print() // use the print() function


func (pointerToPerson *person) updateName(newFirstName string) {
    (*pointerToPerson).firstName = newFirstName
}

// Structs with receiver functions
func (p person) print() {
    fmt.Printf("%+v", p)
}
```

### 3.  Value type will need to use pointers to change these things in a function, but reference type will not.
Reference type will have a location stores its data structure in the ram, and a location to store the data. <br>

For example, slice will have a location has its structure: length, ptr to head, capcity<br>
            and ptr point to the data it contain, which will be another location.
            When we call this reference type, it will copy the structure part, and the copied structure part has 
            ptr to head and point to the same location of the data the original contains.<br>
   
Value types: int, float, string, bool, Structs<br>
Reference typs: slices, maps, channels, pointers, functions<br>


### 4. Map 
-- Simple example of map
    eg.``` colors := map[string]string {}```<br>
    The first string means the keys in the map will be string type.<br>
    the second string means the value in the map will be string type.<br>

-- Two ways to initiate a map
    ```colors := make(map[string]string)``` and 
    ```var colors map[string]string```

-- iterating over map in Go
```
    func printMap(arg map[string]string) {   -> arg: argument name
        for key, value := range arg {
            ...
        }
    }
```

### 5. Map vs. Structs
                Map                                                    Structs
    All keys and values must be the same type             Values can be different type
    Keys are indexed, we can iterate over them            Keys don't support indexing 
    Reference type                                        Value Type
                                                          use to represent a "thing" with a lot of different propertis

### 6. make([]byte, 99999)
make(type of a slice, the number of elements that we want to slice to be initialized with)<br>
  
Questions to think about:<br>
  1. Every value has a type
  2. Every function will need to specify the type of the return value
  3. Question: What if I have need to print an image, a photo, a text, a http request body<br>
              so does this means that every function I ever write has to be rewritten to accommodate different types even the logic is the same?

### 7. Rules of interface
interface is about helping you to reuse the code.
```
type bot interface {
    getGreeting(string, int) (string, error)
    // name     argument type   return type
}
```
Concrete type: you can create value direcly -> map, int, struct, string, englishbot  <br>
interface type: you can't create value -> bot

### 8. Reader interface: to resolve the question I had above
Use reader interface to take some source of data and kind of import that data into our application with a common point of contact that many other pieces of code can easily work with.
```
type Reader interface {
    Read(p []byte) (n int, err error)
}
```
Read function: <br>
  To take radically different sources of input and translate them into some common medium that all the other functions can easily work with.
  Cannot resize, so when we use make() put a big number in.

### 9. Code: func (c *Client) Get(url string) (resp *Response, err error)
Notice:Get will return a pointer of response and a error type

Code: 
```
type Response struct {
            Status     string // e.g. "200 OK"
            StatusCode int    // e.g. 200
            Proto      string // e.g. "HTTP/1.0"
            ProtoMajor int    // e.g. 1
            ProtoMinor int    // e.g. 0
            Header Header
            Body io.ReadCloser
            ...
            }
```
Notice: Response is a struct, and it has a body type in it

    |  Response struct     |
     ----------------------
    | status  -> string    |
    | status code -> int   |
    | body -> io.ReadCloser|   -->      |  Response struct     |
                                         ----------------------
                                        | Reader               |  ->  |  io.Reader                   |
                                        | Closer               |       ------------------------------
                                                                      | Read (p []byte) (n int, err) |


                                                                   -> |  io.Closer                  |
                                                                       ------------------------------
                                                                      | Close() (error)             |
                                        
### 10. Writer Interface:
Writer interface describes something that can tale info and send it outside of our program
Code:
```
type Writer interface {
  Write(p []byte) (n int, err error)
}
```
Write function receives a bite slice and then returns an integer and error.
     
### 11. Code: io.Copy(os.Stdout, resp.Body)
io.Copy(something that implements the writer interface,  
                 something that implements the reader interface)<br>
Notice: writer interface example in the code, -> os.Stdout -> value of type of File -> File has a function called "write" 
                 
         

### 12. 
```
  io.Copy(dst Writer, src Reader) (written int64, err error){
        return copyBuffer(dst, scr, nil)
    }

    // copyBuffer is the actual implementation of Copy and CopyBuffer.
    // if buf is nil, one is allocated.
    func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
        // If the reader has a WriteTo method, use it to do the copy.
        // Avoids an allocation and a copy.
        if wt, ok := src.(WriterTo); ok {
            return wt.WriteTo(dst)
        }
        // Similarly, if the writer has a ReadFrom method, use it to do the copy.
        if rt, ok := dst.(ReaderFrom); ok {
            return rt.ReadFrom(src)
        }
        if buf == nil {
            size := 32 * 1024
            if l, ok := src.(*LimitedReader); ok && int64(size) > l.N {
                if l.N < 1 {
                    size = 1
                } else {
                    size = int(l.N)
                }
            }
            buf = make([]byte, size)
        }
        for {
            nr, er := src.Read(buf)
            if nr > 0 {
                nw, ew := dst.Write(buf[0:nr])
                if nw < 0 || nr < nw {
                    nw = 0
                    if ew == nil {
                        ew = errInvalidWrite
                    }
                }
                written += int64(nw)
                if ew != nil {
                    err = ew
                    break
                }
                if nr != nw {
                    err = ErrShortWrite
                    break
                }
            }
            if er != nil {
                if er != EOF {
                    err = er
                }
                break
            }
        }
        return written, err
    }
```



                
