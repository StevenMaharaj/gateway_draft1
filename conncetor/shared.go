package conncetor

type Connector interface {
  Connect() error
  Disconnect() error
  IsConnected() bool
}

