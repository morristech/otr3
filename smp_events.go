package otr3

// SMPEvent define the events used to indicate status of SMP to the UI
type SMPEvent int

const (
	// SMPEventError means abort the current auth and update the auth progress dialog with progress_percent. This event is only sent when we receive a message for another message state than we are in
	SMPEventError SMPEvent = iota
	// SMPEventAbort means update the auth progress dialog with progress_percent
	SMPEventAbort
	// SMPEventCheated means abort the current auth and update the auth progress dialog with progress_percent
	SMPEventCheated
	// SMPEventAskForAnswer means ask the user to answer the secret question
	SMPEventAskForAnswer
	// SMPEventAskForSecret means prompt the user to enter a shared secret
	SMPEventAskForSecret
	// SMPEventInProgress means update the auth progress dialog with progress_percent
	SMPEventInProgress
	// SMPEventSuccess means update the auth progress dialog with progress_percent
	SMPEventSuccess
	// SMPEventFailure means update the auth progress dialog with progress_percent
	SMPEventFailure
)

func (c *Conversation) smpEvent(e SMPEvent, percent int) {
	c.getEventHandler().HandleSMPEvent(e, percent, "")
}

func (c *Conversation) smpEventWithQuestion(e SMPEvent, percent int, question string) {
	c.getEventHandler().HandleSMPEvent(e, percent, question)
}