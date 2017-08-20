# goqueue
This package provides an implementations of queues in Golang and analyze they're runtime.

Go does not provide a queue like container, if you google around you may find
some suggestions and stackoverflow and so on.

I needed a queue and thought about different ways to implement them and wondered
which one would be the fastes. So I tried different implementations and
want to share them and the results with you.

Since Go does not support generics you can't simply use the implementation
in this package but have to use it as a "recipe" for your type. Sucks I know,
but that's how it is.

I've only implemented FIFO queues here, but implementing a stack should be
easy given the recipe.

## Implementations
I've tried three different implementations:

 1. A simple one that uses a slice, appends elements on the end of the slice on Push and updates the slice on Pop (`s = s[1:]`)
 2. One that uses a linked list to retrieve the elements
 3. An extendable slice that I found on github:  [LIFO Stack and FIFO Queue in golang](https://gist.github.com/moraes/2141121) by [moraes](https://gist.github.com/moraes) so thanks for that. Hope it's okay that I adjusted it a bit and put it here.

I ran to different benchmarks: One that first pushes 100,000 elements, then pops 50,000 elements and finally inserts 75,000 elements again.
The second one alternates between pushes and pops, it does the following 250,000 times: Pop one element, push two elements (I think that's something you usually do with a queue...). I've also tested it with smaller and bigger benchmarks.

Here are the results on my computer, the table shows the time / op (average for one run) as go benchmark outputs them:

| Benchmark                            | Slice            | Linked            | Extendable       |
|--------------------------------------|------------------|-------------------|------------------|
| Sequential (100000, 50000, 75000)    | 1.989.646 ns / op  | 12.391.950 ns / op  | 4.511.841 ns / op  |
| Alternate (250000)                   | 5.482.742 ns / op  | 36.237.653 ns / op  | 11.865.900 ns / op |
| Sequential (10000, 5000, 7500)       | 144.381 ns / op   | 692.702 ns / op    | 442.928 ns / op   |
| Alternate (100000)                   | 2.395.953 ns / op  | 13.396.773 ns / op  | 4.617.221 ns / op  |
| Sequential (1000000, 250000, 250000) | 12.120.915 ns / op | 91.740.777 ns / op  | 41.156.517 ns / op |
| Alternate (1000000)                  | 20.503.286 ns / op | 138.518.414 ns / op | 58.408.381 ns / op |

As you can see using a slice is always the fastest option, much faster than using the extendable version. The linked version is always the slowest one (by far). So I think the easiest thing is to use just slices, easy to understand and fast.

I'd also suggest to set `s[0] = nil` in the Pop method (better for garbage collection).
