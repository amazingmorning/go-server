package gzet

import (
	"fmt"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func (s *Server) Start() {
	fmt.Printf("[START] Server name: %s,listenner at IP: %s, Port %d is starting\n", s.Name, s.IP, s.Port)
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))

		if err != nil {
			fmt.Printf("ResolveTCPAddr error:%s", err)
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("ListenTCP error:%s", err)
			return
		}
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf("Accept error:%s", err)
				return
			}

			go func() {
				for {
					buf := make([]byte, 1024)
					count, err := conn.Read(buf)
					if err == nil {
						fmt.Println(count)
					}

					if _, err := conn.Write(buf[:count]); err != nil {
						fmt.Println("Write error!")
					}
				}

			}()
		}

	}()

}

func (s *Server) Serve() {
	s.Start()
	//阻塞,
	select {}
}

func (s *Server) Stop() {

}

func NewServer(name string) Server {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      9898,
	}
	return *s
}
