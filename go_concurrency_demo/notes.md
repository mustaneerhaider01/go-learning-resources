### Channels in GO

1. Channels allow waiting for the completion of a single or multiple Go routines as well.

2. Channel is like a data transmission device in Go which is just a normal value that can be stored in variable.

3. We can define whihc kind of value can be emitted by the channel.

4. Channels can be created with the built-in make() function.

5. We can simple pass channel to a go routine and then inside that function we can explicitly tell Go that
   when that GO routine is completed by using "<-" operator which shows direction of data flow
   e.g doneChan <- true

6. Then in the place where we start that go routine, we can read from that channel and then GO will wait untill that value is emitted from that channe; before program exits e.g <-done

7. We can pass single channel to multiple go routines and inside each go routine manually specify when it ends, then in order to avoid race condition between go routines, we need to wait/read as many values from that channel as many go routines we have.

8. Another solution built-in to GO is, to use for loop to loop through values emitted by that one channel and keep on waiting for new values. The downside is that now GO doesn't when that channel is out of values or simply when we have no value left to emit and therefore GO throws error.

9. To tackle that error, we can explicity use close() function and pass it the channel once we know that we will be done with emittting the values through that channel, so basically when that long-taking task will be completed and then it's value will be emitted via that single channel.
