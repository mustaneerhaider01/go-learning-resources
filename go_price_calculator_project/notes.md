### Reading content from a line by line in GO

1. The os package has a method, Open() which opens a file to be read.

2. The built-in bufio package in GO has NewScanner method which returns a value that exposes
   methods that allows reading the content of an opended file line by line...

3. To this NewScanner() we pass a value that is of type Reader, since it's an
   interface that is implement by that file value returned from os.Open() method.

4. That exposed value, has a Scan() method which read the line and then move to next line, and this method returns a boolean indicating that if there are more lines to read in that file, so it can be used together with a for loop to continue reading untill there isn't any line left to read.

5. The read line can be accessed by calling Text() method on that scanner value, inside the for loop.

6. Always work with a pointer, to make sure you update that original value and not a copy.
