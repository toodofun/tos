package system

import (
	"encoding/json"
	"github.com/MR5356/tos/config"
	"github.com/MR5356/tos/constant"
	"github.com/avast/retry-go"
	"github.com/go-resty/resty/v2"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

const (
	holidayAPI = "https://www.shuyz.com/githubfiles/china-holiday-calender/master/holidayAPI.json"
)

var (
	service *Service
	once    sync.Once
)

type Service struct {
	httpClient *resty.Client
}

func GetService() *Service {
	client := resty.New()
	once.Do(func() {
		service = &Service{
			httpClient: client,
		}
	})
	return service
}

// GetTimestamp 获取当前时间戳[毫秒级]
func (s *Service) GetTimestamp() int64 {
	return time.Now().UnixMilli()
}

func (s *Service) GetSystemInfo() *Info {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "UNKNOWN"
	}

	hostInfo, err := host.Info()
	if err != nil {
		hostInfo = new(host.InfoStat)
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		cpuInfo = make([]cpu.InfoStat, 0)
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		memInfo = new(mem.VirtualMemoryStat)
	}

	return &Info{
		Basic: Basic{
			Version:   constant.Version,
			Commit:    constant.Commit,
			BuildTime: constant.BuildTime,
			Hostname:  hostname,
		},
		Hardware: Hardware{
			Cpu:  cpuInfo,
			Mem:  *memInfo,
			Host: *hostInfo,
		},
	}
}

func (s *Service) GetNetworkInfo() []net.InterfaceStat {
	netInfo, err := net.Interfaces()
	if err != nil {
		netInfo = make([]net.InterfaceStat, 0)
	}
	return netInfo
}

func (s *Service) GetHolidayAPI() (res *HolidayAPI) {
	err := retry.Do(func() error {
		resp, err := s.httpClient.SetTimeout(time.Second * 3).R().
			Get(holidayAPI)
		if err != nil {
			return err
		}
		return json.Unmarshal(resp.Body(), &res)
	},
		retry.Attempts(config.Current().Robust.Retries),
		retry.Delay(0),
		retry.LastErrorOnly(true),
		retry.DelayType(retry.DefaultDelayType),
		retry.OnRetry(func(n uint, err error) {
			logrus.Warnf("[%d/%d]retry to get holiday api error: %v", n, config.Current().Robust.Retries, err)
		}),
	)

	if err != nil {
		_ = json.Unmarshal([]byte(defaultHoliday), &res)
	}
	return
}

func (s *Service) Initialize() error {
	return nil
}
