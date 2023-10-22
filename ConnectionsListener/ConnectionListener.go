package ConnectionListener

import (
	"fmt"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	msgCh 	   chan []byte	
	clientPool []net.Conn
}

func NewServer(listenAddr string) *Server{
	return &Server{
		listenAddr: listenAddr,
		msgCh: make(chan []byte,10),
		clientPool: make([]net.Conn, 0),
	}
}

func (s *Server) Start() error{
	ln, err := net.Listen("tcp", "localhost" + s.listenAddr)

	if err != nil{
		return err
	}
	fmt.Println("Server Acceso")
	defer ln.Close()

	s.ln = ln

	defer close(s.msgCh)
	s.acceptLoop()
	return nil
}

func (s *Server) acceptLoop(){
	for {
		conn, err := s.ln.Accept()
		
		if err != nil{
			fmt.Println("accept error: ", err)
			continue
		}
		
		s.clientPool = append(s.clientPool, conn)

		go s.Login(conn)
	}

}

func (s *Server) Login(conn net.Conn){
	psw := "prova"
	defer conn.Close()
	buffer := make([]byte,2048)
	n, err := conn.Read(buffer)
		//line, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("reading error: ", err)
			conn.Close()
			return
		}


	
	
	//fmt.Println(line,psw)
	if string(buffer[:n]) == psw {
		conn.Write([]byte("Login Effettuato!!"))

		//continuo a ricevere dati dal destinatario
		s.readLoop(conn)
	}else{
		conn.Write([]byte("Password Sbagliata!!"))
		conn.Close()
	}

}

func (s *Server) readLoop(conn net.Conn){
	defer conn.Close()
	buffer := make([]byte,2048)
	//reader := bufio.NewReader(conn)
	for{
		n, err := conn.Read(buffer)
		//line, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("reading error: ", err)
			break
		}
		fmt.Print(string(buffer[:n]))
		//	s.msgCh <- []byte(line)
		
	}
}
