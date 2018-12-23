package main

import (
	"github.com/PacktPublishing/Hands-On-GUI-Application-Development-in-Go/client"
	"github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/layout"
	"github.com/fyne-io/fyne/widget"
	"time"
)

type composeUI struct {
	app    fyne.App
	server *client.EmailServer

	list *widget.Group

	message, subject, to *widget.Entry
}

func (c *composeUI) loadUI() fyne.Window {
	compose := c.app.NewWindow("GoMail Compose")

	c.subject = widget.NewEntry()
	c.subject.SetText("subject")
	toLabel := widget.NewLabel("To")
	c.to = widget.NewEntry()
	c.to.SetText("email")

	c.message = widget.NewMultilineEntry()
	c.message.SetText("content")

	send := widget.NewButton("Send", func() {
		email := client.NewMessage(c.subject.Text, c.message.Text,
			client.Email(c.to.Text), "", time.Now())
		c.server.Send(email)
		compose.Close()
	})
	send.Style = widget.PrimaryButton
	buttons := widget.NewHBox(
		layout.NewSpacer(),
		widget.NewButton("Cancel", func() {
			compose.Close()
		}),
		send)

	top := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(c.subject, nil, toLabel, nil),
		c.subject, toLabel, c.to)
	content := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(top, buttons, nil, nil),
		top, c.message, buttons)
	compose.SetContent(content)

	compose.Resize(fyne.NewSize(400, 320))
	return compose
}

func newCompose(mailApp fyne.App, server *client.EmailServer) *composeUI {
	ui := &composeUI{
		app:    mailApp,
		server: server,
	}

	return ui
}