package adaptors

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ale8k/hexagonal-architecture-in-go-example/internal/ports"
)

type httpAdaptor struct {
	mailPort ports.MailPort
}

func NewHttpAdaptor(mailPort ports.MailPort) *httpAdaptor {
	return &httpAdaptor{
		mailPort: mailPort,
	}
}

func (h *httpAdaptor) Run() {
	http.HandleFunc("/mail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request to send mail come in")
		mail, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("failed to receive mail")
			return
		}

		defer r.Body.Close()
		h.mailPort.ReceiveMail(mail)
		w.Write([]byte("mail received"))
	})

	http.ListenAndServe(":8080", nil)
}
