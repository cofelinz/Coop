package network

import (
	"os"
	"fmt"
	"net"
	"testing"
	"time"
	"networ"
)

//默认的服务器地址
var (
	server = "127.0.0.1:9876"
)

func TestClient(t *testing.T)  {
	//拿到服务器地址信息
	hawkServer,err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		fmt.Printf("hawk server [%s] resolve error: [%s]",server,err.Error())
		os.Exit(1)
	}
	//连接服务器
	connection,err := net.DialTCP("tcp",nil,hawkServer)
	if err != nil {
		fmt.Printf("connect to hawk server error: [%s]",err.Error())
		os.Exit(1)
	}
	client := &TcpClient{
		connection:connection,
		hawkServer:hawkServer,
		stopChan:make(chan struct{}),
	}
	//启动接收
	go client.receivePackets()

	//发送心跳的goroutine
	go func() {
		heartBeatTick := time.Tick(2 * time.Second)
		for{
			select {
			case <-heartBeatTick:
				client.sendHeartPacket()
			case <-client.stopChan:
				return
			}
		}
	}()

	//测试用的，开300个goroutine每秒发送一个包
	for i:=0;i<300;i++ {
		go func() {
			sendTimer := time.After(1 * time.Second)
			for{
				select {
				case <-sendTimer:
					client.sendReportPacket()
					sendTimer = time.After(1 * time.Second)
				case <-client.stopChan:
					return
				}
			}
		}()
	}
	//等待退出
	<-client.stopChan
}