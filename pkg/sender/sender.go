package sender

var (
	defaultClientSender Sender = new(httpclient)

	userAgentGooglebot    = `Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; UserAgentGooglebot/2.1; +http://www.google.com/bot.html) Chrome/W.X.Y.Z Safari/537.36`
	userAgentGoogleChrome = `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36`
)

// Sender .
type Sender interface {
	Send(p *Processor) error
}

// httpclient .
type httpclient struct{}

// Send .
func (c *httpclient) Send(p *Processor) error {
	// TODO implement me
	panic("implement me")
}
