// STUDENT PROBLEM SET — Lesson 055.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Replace a memory repository with PostgreSQL while keeping the service interface unchanged.
// 2. Extend it with one explicit boundary case: Write a test for creating and retrieving a task through HTTP against a real PostgreSQL instance.
// 3. Rebuild the same behavior with one changed constraint: Write a second case with one changed constraint and verify it. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 055: Go PostgreSQL Repositories and Tests
// Your task: implement missing-row and canceled-context behavior. validate title and assign unique IDs safely under concurrent calls. return Task{},errors.New("implement the required behavior")
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main
import ("context"; "errors")
type Task struct { ID int64; Title string; Done bool }
type Repository interface { Get(context.Context,int64)(Task,error); Create(context.Context,string)(Task,error) }
type MemoryRepository struct { tasks map[int64]Task }
func(r *MemoryRepository)Get(ctx context.Context,id int64)(Task,error){
    // TODO: implement the assigned behavior.
    return Task{}, nil
}
func(r *MemoryRepository)Create(ctx context.Context,title string)(Task,error){
    // TODO: implement the assigned behavior.
    return Task{}, nil
}
func main(){}
