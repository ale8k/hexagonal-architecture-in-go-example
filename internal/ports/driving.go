package ports

type MailPort interface {
	ReceiveMail(mail any)
}