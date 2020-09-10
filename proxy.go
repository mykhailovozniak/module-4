package module_4

import "fmt"

type Subject interface {
	Send() string
}

type Proxy struct {
	realSubject Subject
}

func (p *Proxy) Send() string {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	return "<strong>" + p.realSubject.Send() + "</strong>"
}

type RealSubject struct {
}

func (s *RealSubject) Send() string {
	return "Message Text"
}

func ProxyClientCode() {
	proxy := new(Proxy)


	fmt.Println(proxy.Send())
}