package network

import "testing"

func TestServer(t *testing.T) {
	//类似于初始化套接字，绑定端口
	hawkServer, err := net.ResolveTCPAddr("tcp", server)
	checkErr(err)
	//侦听
	listen, err := net.ListenTCP("tcp", hawkServer)
	checkErr(err)
	//记得关闭
	defer listen.Close()
	tcpServer := &TcpServer{
		listener:listen,
		hawkServer:hawkServer,
	}
	fmt.Println("start server successful......")
	//开始接收请求
	for {
		conn, err := tcpServer.listener.Accept()
		fmt.Println("accept tcp client %s",conn.RemoteAddr().String())
		checkErr(err)
		// 每次建立一个连接就放到单独的协程内做处理
		go Handle(conn)
	}
}