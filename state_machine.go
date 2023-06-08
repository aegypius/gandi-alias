package main

import (
	"context"
	"fmt"

	"github.com/aegypius/gandi-alias/aliases"
	"github.com/charmbracelet/log"
	"github.com/qmuntal/stateless"
)

type trigger string
type state int64

const (
	triggerEmailStored  trigger = "emailStored"
	triggerAliasQueried trigger = "aliasQueried"
	triggerAliasAdded   trigger = "aliasAdded"
	triggerRestart      trigger = "restart"
)

const (
	stateNew            state = iota
	stateEmailStored    state = iota
	stateAliasesQueried state = iota
	stateAliasAdded     state = iota
)

func initStateMachine() *stateless.StateMachine {

	// Create State machine
	stateMachine := stateless.NewStateMachine(stateNew)

	stateMachine.Configure(stateNew).
		Permit(triggerEmailStored, stateEmailStored)

	stateMachine.Configure(stateEmailStored).
		Permit(triggerAliasQueried, stateAliasesQueried).
		OnEntry(func(_ context.Context, args ...interface{}) error {
			return OnEmailStored(args[0].(model))
		})

	stateMachine.Configure(stateAliasesQueried).
		Permit(triggerAliasAdded, stateAliasAdded)

	stateMachine.Configure(stateAliasAdded).
		PermitReentry(triggerAliasAdded)

	stateMachine.Configure(stateAliasAdded).
		Permit(triggerRestart, stateNew)

	return stateMachine
}

type item struct {
	title string
}

func (i item) Title() string       { return i.title }
func (i item) FilterValue() string { return i.title }

func OnEmailStored(m model) (err error) {
	var items []string

	items, err = aliases.ListAliases(aliases.EmailAddress(m.email.Value()))
	if err != nil {
		log.Error(err)
		return
	}

	m.aliases.Title = fmt.Sprintf("Aliases for %s", m.email.Value())
	for _, alias := range items {
		log.Info(alias)
		m.aliases.InsertItem(-1, item{alias})
	}

	m.state.Fire(triggerAliasQueried)

	return
}
