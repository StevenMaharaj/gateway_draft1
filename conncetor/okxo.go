package conncetor

import pb "github.com/stevenmaharaj/gateway_draft1/core"

type OKXOConnector struct {
  cmdCh chan string
  dtCh chan string
  isConn bool

}


func NewOKXOConnector() *OKXOConnector {
  return &OKXOConnector{
    cmdCh: make(chan string),
    dtCh: make(chan string),
  }
}


func (c *OKXOConnector) Connect() error {
  if !c.isConn {
    // TODD
    stream()
  } else {
    return errror.New("OKXOConnector is already connected")
}

func (c *OKXOConnector) Disconnect() error {
  return nil
}

func (c *OKXOConnector) IsConnected() bool {
  return true
}


func stream() {
  
  // 1. connect to websocket
  // 2. send subscribe message
  // 3. receive data
  // 4. send data to channel
  // 5. if error, send error to channel
  for {
    order := pb.Event{}

  }
}

