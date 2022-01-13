package notifications

import "fmt"

type FakeMailer struct {
	Sent []*Message
}

func NewFakeMailer() *FakeMailer {
	return &FakeMailer{
		Sent: make([]*Message, 0),
	}
}

func (fm *FakeMailer) Send(messages ...*Message) (int, error) {
	sentEmailsCount := 0
	for _, msg := range messages {
		fm.Sent = append(fm.Sent, msg)
		sentEmailsCount++
	}
	return sentEmailsCount, nil
}

type FakeDisconnectedMailer struct{}

func NewFakeDisconnectedMailer() *FakeDisconnectedMailer {
	return &FakeDisconnectedMailer{}
}

func (fdm *FakeDisconnectedMailer) Send(messages ...*Message) (int, error) {
	return 0, fmt.Errorf("connect: connection refused")
}

func (ns *NotificationService) MailQueuePop() *Message {
	return <-ns.mailQueue
}
