package main
import ("context"; "errors")
type Task struct { ID int64; Title string; Done bool }
type Repository interface { Get(context.Context,int64)(Task,error); Create(context.Context,string)(Task,error) }
type MemoryRepository struct { tasks map[int64]Task }
func(r *MemoryRepository)Get(ctx context.Context,id int64)(Task,error){
 // TODO: implement missing-row and canceled-context behavior.
 task,ok:=r.tasks[id];if !ok{return Task{},errors.New("not found")};return task,nil
}
func(r *MemoryRepository)Create(ctx context.Context,title string)(Task,error){
 // TODO: validate title and assign unique IDs safely under concurrent calls.
 return Task{},errors.New("not implemented")
}
func main(){}
