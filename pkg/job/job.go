package job

//type Job struct {
//	group    errgroup.Group //errgroup，参考我的文章，专门讲这个原理
//	app      *fx.App        //fx 实例
//	provides []interface{}
//	invokes  []interface{}
//	supplys  []interface{}
//	CronSeTu *qw_pixiv_setu.CronTab
//}
//
//func NewJob() *Job {
//	return &Job{}
//}
//func (job *Job) Run() {
//	pterm.Info.Println("Run")
//	job.app = fx.New(
//		fx.Provide(job.provides...),
//		fx.Invoke(job.invokes...),
//		fx.Supply(job.supplys...),
//		fx.Provide(qw_pixiv_setu.NewCronTab),
//		fx.Supply(job),
//		fx.Populate(&job.CronSeTu), //给CronSeTu 实例赋值
//		fx.NopLogger,               //禁用fx 默认logger
//	)
//	pterm.Info.Println(job.CronSeTu)
//	job.group.Go(job.CronSeTu.Run)
//	err := job.group.Wait() //等待子协程退出
//	if err != nil {
//		panic(err)
//	}
//}
//func (job *Job) Provide(ctr ...interface{}) {
//	pterm.Info.Println("Provide")
//	job.provides = append(job.provides, ctr...)
//}
//func (job *Job) Invoke(invokes ...interface{}) {
//	pterm.Info.Println("Invoke")
//	job.invokes = append(job.invokes, invokes...)
//}
//func (job *Job) Supply(objs ...interface{}) {
//	pterm.Info.Println("Supply")
//	job.supplys = append(job.supplys, objs...)
//}
