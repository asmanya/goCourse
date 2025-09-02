package main

import (
	"fmt"
	"sync"
	"time"
)

// ===== CONSTRUCTION EXAMPLE =====

type Worker struct {
	ID   int
	Task string
}

// PerformTask simulates a worker performing a task
func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WorkerID %d started %s.\n", w.ID, w.Task)
	time.Sleep(time.Second)

	fmt.Printf("workerID %d finished %s.\n", w.ID, w.Task)
}

func main() {
	var wg sync.WaitGroup
	
	tasks := []string{"digging", "laying bricks", "painting"}

	for i, task := range tasks {
		worker := Worker{
			ID: i+1,
			Task: task,
		}
		wg.Add(1)
		go worker.PerformTask(&wg)
	} 

	// wait for all workers to finish
	wg.Wait()

	// construction is finished
	fmt.Println("construction is finished")
}

/*
// ===== EXAMPLES WITH CHANNELS =============

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting.\n", id)
	time.Sleep(time.Second) // simulate some work

	for task := range tasks{
		results <- task * 2
	}

	fmt.Printf("Worker %d finished.\n", id)
}

func main() {
	var wg sync.WaitGroup

	numWorkers := 3
	numJobs := 5
	results := make(chan int, numJobs)
	tasks := make(chan int, numJobs)

	wg.Add(numWorkers)

	for i := range numWorkers {
		go worker(i+1, tasks, results, &wg)
	}

	for i := range numJobs {
		tasks <- i
	}
	close(tasks)


	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}

}

======= BASIC EXAMPLE WITHOUT USING CHANNELS ==========
func main() {
	var wg sync.WaitGroup
	numWorkers := 3

	wg.Add(numWorkers) // tells how many goroutines to wait for - makes a count

	// launch workers
	for i := range numWorkers {
		go worker(i, &wg)
	}

	wg.Wait() // waits till all goroutines are finished
	fmt.Println("All workers finished.")
}

func worker(id int, wg *sync.WaitGroup) {
	// blocking mechanism
	defer wg.Done() // marke go routine as done, reduce count by 1

	fmt.Printf("Worker %d starting.\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished.\n", id)
}

*/
