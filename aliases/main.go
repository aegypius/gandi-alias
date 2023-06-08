package aliases

import (
	"os"
	"strings"

	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/email"
)

var client = email.New(config.Config{
	APIKey: os.Getenv("GANDI_API_KEY"),
})

type EmailAddress string

func (m EmailAddress) GetDomain() (domain string, err error) {
	return strings.Split(string(m), "@")[1], nil
}

func selectMailBox(emailAddress EmailAddress) (m email.ListMailboxResponse, err error) {
	domain, err := emailAddress.GetDomain()
	if err != nil {
		return m, err
	}

	mailboxes, err := client.ListMailboxes(domain)
	if err != nil {
		return m, err
	}

	for _, m := range mailboxes {
		if m.Address == string(emailAddress) {
			return m, nil
		}
	}
	return
}

func listAliases(m email.ListMailboxResponse) (aliases []string, err error) {
	mbox, mboxErr := client.GetMailbox(m.Domain, m.ID)

	if mboxErr != nil {
		return nil, mboxErr
	}

	return mbox.Aliases, nil
}

func ListAliases(emailAddress EmailAddress) (aliases []string, err error) {
	m, err := selectMailBox(emailAddress)
	if err != nil {
		return
	}

	return listAliases(m)
}

func AddAlias(mailbox EmailAddress, alias string) (err error) {
	m, err := selectMailBox(mailbox)
	if err != nil {
		return
	}

	aliases, err := listAliases(m)
	if err != nil {
		return
	}

	client.UpdateEmail(m.Domain, m.ID, email.UpdateEmailRequest{
		Aliases: append(aliases, alias),
	})

	return
}
