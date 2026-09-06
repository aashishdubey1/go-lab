# Concurrency

hello

#### What's the big idea here , why we need this ?

maybe because the computers waits , it wait for network calls , disk i/o , user inputs
if code runs sequentially , the cpu sits idle also the whole process takes longer time

we need something that enables running multiple tasks at same time

what's **Parallelism** then ? i've heard this here and there

okay so parallelism is about doing tasks at same time while concurrency based on idea of doing tasks by switching

in concurrency it might takes multiple core and execute task if available then we can call this
parallelism

so this is the whole idea of concurrency and parallelism , how go does that ~~maybe that is part of another discussion~~
or not ?

let's write some code and see it by myself ( the `fetchServer()` one )

#### Now let's see how go achive concurrency

`what the heck is **GMP** model ?`

some gyan from here there , so traditional lang like c++ and java uses os threads.

- An os thread has fixed memory stack (1-2 mb)
  - what's this memory stack thing and why bother mentioning it here ?
- maybe it's memory where process is gonna run ? (**LEARN IT**)
- another one is , switching between os threads requires cpu kernal call, which are expensive
  - what's expensive means here ? why expensive ?
- if you create 100,000 os thread that's gonna fuck your machine due to memory exhaustion

if that's the case then why those languages uses os threads ? are they dumb

so go introduce something like `Goroutines` which are light weight **user space** threads
managed by **go runtime**, not the os threads

- go routine stack starts tiny (around 2kb) and grows/shrink dynamically
- we can easily launch 100,000+ goroutines concurrently on a single laptop .
  - what if all goroutiens takes more memory space , then again fucked up ?
  - who control that assignment of memory space (thinks need Operating System course ?)

#### How ?

### Goroutines(go)

independent concurrent thread of execution
to turn sync function call into async backround tasks
place keyword go before any function call

so if we write go infront of any function that will start executing in async way ? ( let's try that)
okay so i put that go in front of all function and nothing happens
here's what happen i thought
the all three functions goes to run async and nothing is holding main function to wait for those functions
but where those functions are running ? and where main is running ?
does main is another goroutine ?

here's llm answers

- program start with launching the `main` goroutine.
- main hits `go fetchData()` it tells go scheduler -> run this function asynchronously when you get chance
- same goes for remaining functions
- main reaches the end of the main() functions and exits
- when main program exits, the entire go program shuts down immediately , killing all child goroutines

and to answer where are they all running

- they are all running inside same process, managed by Go's internal Schedular.
- go runtimes maps all these goroutines into os threads.
- They execute in parallel across your CPU cores or switch off concurrently on a single core.

okay so now what ?
we can add spwan goroutines but how to stop them ?

#### sync.WaitGroup

go provides thread-safe counter. it has 3 methods

> what's the thread-safe ?

- Add(n): Increment counter by n (tell it how many tasks to wait for).
- Done(): Decrement counter by 1 (call inside goroutine when finished).
- Wait(): Block until counter reaches 0.

let's implement that in our previos code
another task ?

```
Write a program that processes 5 items in parallel. Each item takes 200ms to process.

Use a loop to launch 5 goroutines.

Use sync.WaitGroup to ensure main() waits for all 5 to complete.

Ensure loop variables are passed safely so item IDs print correctly (1, 2, 3, 4, 5 in any order).
```

okay that's it ?
maybe i need some more practice and have to solve some questions before moving forward
