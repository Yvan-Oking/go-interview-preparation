# Go Scheduler

## Go's M-P-G Scheduler

Go's scheduler maps many goroutines (G) to a smaller number of OS threads (M) using processors (P) as the intermediary to distribute the workload efficiently.

### Components:

- **G (Goroutine)**: The actual goroutine
- **M (Machine/OS Thread)**: OS thread that executes goroutines
- **P (Processor)**: Logical processor that manages goroutines

### How It Works:

#### Local Queue Execution:
- Each processor (P) has its own local queue of ready-to-run goroutines
- When a goroutine (G) becomes ready (e.g., created by the `go` keyword), it is placed into the local queue of the current processor (P) that created it
- The M (OS thread) running on that P picks goroutines from this local queue and executes them one by one

#### Idle Processor Detection:
- If a processor (P) runs out of goroutines to execute (i.e., its local queue is empty), it becomes idle
- An idle processor will not just sit idle; it attempts to find work from other processors

#### Work Stealing:
- When a processor (P) is idle, it will try to steal goroutines from the local queues of other processors (P)
- The work-stealing process begins with the idle P selecting another P at random to "steal" from
- If the chosen P has goroutines, the idle P takes half of the goroutines from that queue and places them into its own local queue
- This ensures that the idle P now has work to do, and the load on the overloaded P is reduced

### Global Run Queue and Local Queues:

In addition to the local queues, Go's runtime has a global run queue:

- **Global Run Queue**: When new goroutines are created and the current processor's local queue is full, they are placed in the global queue
- Processors that become idle can also steal from the global queue
- If a processor's local queue is empty and it cannot steal any work from other processors, it may pull goroutines from the global run queue as a last resort

### Work Stealing Example:

Consider this simplified example with 4 processors (P0, P1, P2, P3):

- P0 has 10 goroutines
- P1 has 5 goroutines  
- P2 has no goroutines (idle)
- P3 has 8 goroutines

**Process:**
1. P2 is idle because its local queue is empty
2. P2 randomly selects P0 for work stealing
3. P2 takes half of P0's goroutines (5 goroutines)
4. Now P0 has 5 goroutines, and P2 has 5 goroutines
5. Both P0 and P2 can now continue executing their goroutines, balancing the workload

### Key Features:

- **Preemptive Scheduling**: Ensures that long-running goroutines do not monopolize CPU time
- **Load Balancing**: Work stealing distributes load evenly across processors
- **Efficiency**: Minimal overhead for managing thousands of goroutines
- **Scalability**: Adapts to available CPU cores

### Summary:

- **Go's M-P-G scheduler** maps many goroutines (G) to a smaller number of OS threads (M) using processors (P) as the intermediary to distribute the workload efficiently
- Each processor (P) has a local queue for goroutines ready to execute. When the queue is empty, it can steal work from other processors
- The global run queue holds goroutines waiting for a processor when local queues are full or unavailable
- Preemptive scheduling ensures that long-running goroutines do not monopolize CPU time
- This scheduling model is what makes Go particularly efficient for handling massive concurrency while keeping resource consumption low
