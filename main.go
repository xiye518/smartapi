package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
	"smartapi/internal/common"
	"smartapi/internal/log"
	"smartapi/internal/service/apis"
	"smartapi/internal/service/models"
	"strconv"
	"time"
	"util/cmdutil"
)

var (
	configFile  string
	v           bool
	serviceName string
)

var (
	// VERSION 版本信息
	VERSION string
	// BUILD 构建时间
	BUILD string
	// COMMITSHA1 git commit ID
	COMMITSHA1 string
)

func usage() {
	fmt.Printf(`Info:
	version: %s
	release time: %s
	commit sha1: %s
`, VERSION, BUILD, COMMITSHA1)
	fmt.Println("Usage:")
	flag.PrintDefaults()
}

func parseArgs() {
	flag.StringVar(&configFile, "config", "./etc/config.json", "config file")
	flag.BoolVar(&v, "v", false, "show version")
	flag.StringVar(&serviceName, "serviceName", "smartapi", "running service name")

	flag.Parse()
}

//定时任务
func TimerTask() {
	t := time.NewTicker(15 * time.Second)
	for {
		select {
		case <-t.C:
			checkCpuPercent()
		}
	}
}

//检查cpu百分比，达到50%即写pprof文件
func checkCpuPercent() error {
	cmd := "top -bn1 | grep " + serviceName + " | awk '{print $9}'"
	cpuStr, err := cmdutil.GetShellOutput(cmd)
	if err != nil {
		log.Errorf("GetShellOutput: %s", err)
		return err
	}
	log.Debugf("common: %s", cmd)
	log.Debugf("cpu percent: %s", cpuStr)

	reg := regexp.MustCompile(`(?s).+?([0-9]+).+?`)
	res := reg.FindAllStringSubmatch(cpuStr, -1)
	log.Debug("res:", res)
	if len(res) == 0 {
		return nil
	}

	log.Debugf("res[0][1]: %v", res[0][1])
	cpu, err := strconv.Atoi(res[0][1])
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("cpu percent: %d", cpu)
	if cpu > 50 {
		t := time.Now().Format("20060102150405")
		fp, _ := os.Create(fmt.Sprintf("/tmp/%s_%s.pprof", serviceName, t))
		if err != nil {
			log.Error(err)
			return err
		}
		defer fp.Close()
		if err = pprof.StartCPUProfile(fp); err != nil {
			log.Error(err)
			return err
		}
		//cpu五秒采样
		time.Sleep(5 * time.Second)
		pprof.StopCPUProfile()
	}

	return nil
}

func reload() {
	bs, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Sprintf("read config file from: %s with err: %s", configFile, err.Error()))
	}
	cfg, err := common.LoadConfig(bs)
	if err != nil {
		panic(err)
		return
	}
	// 配置新日志，输出到日志文件同时打印到控制台，
	if err := log.SetLogger(cfg.LogConfig.RollType, cfg.LogConfig.Dir, cfg.LogConfig.File, cfg.LogConfig.Count, cfg.LogConfig.Size, cfg.LogConfig.Uint, cfg.LogConfig.Level, cfg.LogConfig.Compress); err != nil {
		panic(fmt.Sprintf("init log with err: %s", err.Error()))
	}
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	go log.HandleSignalChangeLogLevel()
	go http.ListenAndServe(cfg.PprofAddr, nil) //
	TimerTask()
	// 初始化db
	err = models.InitDB(cfg.MysqlConfig)
	if err != nil {
		panic(err)
	}

	// 开启http服务
	defer models.DB.Close()
	router := apis.RunGinService()
	err = router.Run(fmt.Sprintf(":%d", cfg.ServicePort))
	if err != nil {
		panic(err)
	}
}

func main() {
	parseArgs()
	if v {
		usage()
		return
	}
	reload()
	return
}
