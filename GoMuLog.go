package GoMuLog

import (
	"GoMuLog/Helper"
	"fmt"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"
)


const (
	GDLogMod_FLAG_PRINT      int = 1  // 打印到屏幕
	GDLogMod_FLAG_WRITEFILE  int = 2  // 写入到文件
	GDLogMod_FLAG_THREAD     int = 4  // 启动多线程/异步
	GDLogMod_FLAG_SHORTFUILE int = 8  // 显示文件位置
	GDLogMod_FLAG_DATE       int = 16 // 显示时间
)

type GDLog struct {
	// 
	Mode uint; // 标识符号
	isStart bool // 是否开始  开始之后不允许修改
	mode_lock sync.Mutex 
	
	// 多线程模式下 的通道  同步模式下无效数据
	print_chan chan log_msg
	write_chan chan log_msg
	
	// 保存的文件位置
	dirname string
	dirname_lock sync.Mutex
	
	// 记录是否有协程在运行
	is_run_print_coroutine bool
	is_run_print_coroutine_lock sync.Mutex
	
	is_run_wtite_coroutine bool
	is_run_write_coroutine_lock sync.Mutex
}

type log_msg struct{
	msg string
	date_str string
	info_type int // 0 debug  1 info(default) 2 warring 3 error
}

// 构造一个
func NewGDLog()*NewGDLog{
	return &NewGDLog{
		Mode : 1 | 16,
		print_chan: make(chan log_msg),
		write_chan : make(chan log_msg),
		isStart : false,
		dirname : "./out",
		is_run_print_coroutine: false,
		is_run_wtite_coroutine: false,
	}
}

// 开始运行，开始之后无法修改mode
func (this *GDLog)Start(){
	this.isStart = true
}

// 异步显示日志
func (this *GDLog)async_print_log(log log_msg){
	this.print_chan <- log
	if !this.get_is_run_print_coroutine() {
		go this.async_print()
	}
}

// 同步显示日志
func (this *GDLog)print_log(log log_msg){
	fmt.Print(log.msg)
}

// 异步写入日志
func (this *GDLog)async_wite_log(log log_msg){
	this.write_chan <- log
	if !this.get_is_run_write_coroutine() {
		go this.async_write()
	}
}

// 同步写入日志
func (this *GDLog)wite_log(log log_msg){
	this.filename_lock.Lock()	
	defer this.filename_lock.Unlock()
	filename = this.dirname + "/" + log.date_str + ".log"
	Helper.AppendToFile(filename, log.msg)
}

///----------------------------------------------------
func (this *GDLog)async_print(){
	if !this.try_is_run_print_coroutine(false, true) {
		defer this.try_is_run_print_coroutine(true, false)
		for{
			data := <-this.print_chan  // 堵塞的方式
			// data, ok := <-this.print_chan // 不阻塞的方式
			if len(data) > 0{
				fmt.Println("123123")
			}else{
				break;
			}
		}
	}
}

func (this *GDLog)async_write(){
	if !this.try_is_run_write_coroutine(false, true) {
		defer this.try_is_run_write_coroutine(true, false)
	}
}

func (this *GDLog) PrintDebug(param string) {
	if !this.isStart {
		fmt.Print("GOMuLog is not init!")
	}
	showlog := ""
	date_str := timeObj.Format("2006-01-02 03:04:05")
	// 是否打印时间
	if this.Mode&16 == 16 {
		timeObj := time.Now()
		showlog += date_str
		// showlog += fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", timeObj.Year(), timeObj.Month(),
		//     timeObj.Day(), timeObj.Hour(), timeObj.Minute(), timeObj.Second())
	}
	// 是否打印位置信息
	if check_mode(this.Mode, GDLogMod_FLAG_SHORTFUILE) {
		pc, file, line, ok := runtime.Caller(1)
		if ok {
			filename := path.Base(file)
			funcname := runtime.FuncForPC(pc).Name() 
			showlog = showlog + fmt.Sprintf(" %s:%d %s  ", filename, line, funcname)                                                                                                                                                                                                                                                                                                                                                                    python
		}
		else {
			showlog += " runtime.Caller Failed "
		}
	}
	// 是否异步显示
	// if check_mode(this.Mode, GDLogMod_FLAG_THREAD) {

	// }
	// else{

	// }
	fmt.Println(showlog)
}

// 获取状态
func (this *GDLog)get_is_run_print_coroutine()bool{
	this.is_run_print_coroutine_lock.Lock()
	defer this.is_run_print_coroutine_lock.Unlock()
	ret := this.is_run_print_coroutine
	return ret
}

// 尝试获取is_run_print_coroutine 如果is_run_print_coroutine为param1 则将is_run_print_coroutine设置为param2
func (this *GDLog)try_is_run_print_coroutine(param1, param2 bool)bool{
	this.is_run_print_coroutine_lock.Lock()
	defer this.is_run_print_coroutine_lock.Unlock()
	ret := false
	if this.is_run_print_coroutine == param1{
		ret = true
		this.is_run_print_coroutine = param2
	}
	return ret
}

// 获取状态
func (this *GDLog)get_is_run_write_coroutine()bool{
	this.is_run_write_coroutine_lock.Lock()
	defer this.is_run_write_coroutine_lock.Unlock()
	ret := this.is_run_wtite_coroutine
	return ret
}

// 尝试获取is_run_write_coroutine 如果is_run_write_coroutine为param1 则将is_run_write_coroutine设置为param2
func (this *GDLog)try_is_run_write_coroutine(param1, param2 bool)bool{
	this.is_run_write_coroutine_lock.Lock()
	defer this.is_run_write_coroutine_lock.Unlock()
	ret := false
	if this.is_run_wtite_coroutine == param1{
		ret = true
		this.is_run_wtite_coroutine = param2
	}
	return ret
}

// 增加标记值
func (this *GDLog)AddModeFlag(flag uint){
	if !this.isStart {
		this.mode_lock.Lock()()
		defer this.mode_lock.Unlock()
		this.Mode = this.Mode | flag
	}
}

// 移除标记值
func (this *GDLog)RemoveModeFlag(flag uint){
	if !this.isStart {
		this.mode_lock.Lock()()
		defer this.mode_lock.Unlock()
		this.Mode = this.Mode & ^flag
	}
}

// 设置模式
func (this *GDLog)SetMode(param uint){
	if !this.isStart {
		this.mode_lock.Lock()()
		defer this.mode_lock.Unlock()
		this.Mode = flag
	}
}


// 判断对应值是否为1
func (this *GDLog)check_mode(flag uint)bool{
	this.mode_lock.Lock()()
	defer this.mode_lock.Unlock()
	ret := this.Mode & flag == flag
	return ret
}