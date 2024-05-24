### Reading content from a file, line by line in GO

1. The os package has a method, Open() which opens a file to be read.

2. The built-in bufio package in GO has NewScanner method which returns a value that exposes
   methods that allows reading the content of an opended file line by line...

3. To this NewScanner() we pass a value that is of type Reader, since it's an
   interface that is implement by that file value returned from os.Open() method.

4. That exposed value, has a Scan() method which read the line and then move to next line, and this method returns a boolean indicating that if there are more lines to read in that file, so it can be used together with a for loop to continue reading untill there isn't any line left to read.

5. The read line can be accessed by calling Text() method on that scanner value, inside the for loop.

6. Always work with a pointer, to make sure you update that original value and not a copy.

### Selecting value emitted by one of two or more channels in GO

1. If we have one or more channels for a given go-routine and we know that only one of this channel will emit a value, we can use a special control structure built into GO.

2. select statement has similar idea to switch statement and therefor allows to select the case of the channel that emits the value eariler, so the channel that emites value first wins...

3. What select statement does is that it allows defining cases for waiting of emitting a value per channel and the case of the channel that emites the value first that will be executed and it will ignore other cases and moves on...

4. select statement allows simply waiting for one channel to emite a value and once that happened, it will move on so basically ignores/discards the other case(s)...

### Defering code execution in GO

1. If we a function call inside an outer function that is called in multiple places e.g file opened inside a custom read function and then we need to close that by calling file.Close() when an error occurs or the function is done successfully then we can use the defer keyword infront of that function call to defer it's execution untill the outer function finishes execustion either by returning an error or in case of success.

2. defer keyword allows avoding calling same function in multiple places instead you can simply put that call to function somewhere above in an appropate flow with defer keyword infront of it to allow Go to execute that function for you upon fouter function execution ends.
