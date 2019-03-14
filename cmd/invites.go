package cmd

import (
	"errors"
)

func init() {
	register(&invitesCmd{})
}

var errMissingInviteId = errors.New("missing invite id")

type invitesCmd struct {
	Create createInvitesCmd `command:"create" description:"Create account-to-account or external invites to a thread"`
	List   lsInvitesCmd     `command:"ls" description:"List thread invites"`
	Accept acceptInvitesCmd `command:"accept" description:"Accept an invite to a thread"`
	Ignore ignoreInvitesCmd `command:"ignore" description:"Ignore direct invite to a thread"`
}

func (x *invitesCmd) Name() string {
	return "invites"
}

func (x *invitesCmd) Short() string {
	return "Manage thread invites"
}

func (x *invitesCmd) Long() string {
	return `
Invites allow other users to join threads. There are two types of
invites, direct account-to-account and external:

- Account-to-account invites are encrypted with the invitee's account address (public key).
- External invites are encrypted with a single-use key and are useful for onboarding new users.
`
}

type createInvitesCmd struct {
	Client  ClientOptions `group:"Client Options"`
	Thread  string        `short:"t" long:"thread" description:"Thread ID. Omit for default."`
	Address string        `short:"a" long:"address" description:"Account address. Omit to create an external invite."`
}

func (x *createInvitesCmd) Usage() string {
	return `

Creates a direct account-to-account or external invite to a thread.
Omit the --address option to create an external invite.
Omit the --thread option to use the default thread (if selected).
`
}

func (x *createInvitesCmd) Execute(args []string) error {
	setApi(x.Client)
	if x.Thread == "" {
		x.Thread = "default"
	}

	if x.Address != "" {

	}

	res, err := executeJsonCmd(POST, "invites", params{
		opts: map[string]string{
			"thread":  x.Thread,
			"address": x.Address,
		},
	}, nil)
	if err != nil {
		return err
	}
	if res != "" {
		output(res)
	} else {
		output("ok")
	}
	return nil
}

type lsInvitesCmd struct {
	Client ClientOptions `group:"Client Options"`
}

func (x *lsInvitesCmd) Usage() string {
	return `

Lists all pending thread invites.`
}

func (x *lsInvitesCmd) Execute(_ []string) error {
	setApi(x.Client)

	res, err := executeJsonCmd(GET, "invites", params{}, nil)
	if err != nil {
		return err
	}
	output(res)
	return nil
}

type acceptInvitesCmd struct {
	Client ClientOptions `group:"Client Options"`
	Key    string        `short:"k" long:"key" description:"Key for an external invite."`
}

func (x *acceptInvitesCmd) Usage() string {
	return `

Accepts a direct account-to-account or external invite to a thread.
Use the --key option with an external invite.
`
}

func (x *acceptInvitesCmd) Execute(args []string) error {
	setApi(x.Client)
	if len(args) == 0 {
		return errMissingInviteId
	}

	res, err := executeJsonCmd(POST, "invites/"+args[0]+"/accept", params{
		args: args,
		opts: map[string]string{
			"key": x.Key,
		},
	}, nil)
	if err != nil {
		return err
	}
	output(res)
	return nil
}

type ignoreInvitesCmd struct {
	Client ClientOptions `group:"Client Options"`
}

func (x *ignoreInvitesCmd) Usage() string {
	return `

Ignores a direct account-to-account invite to a thread.
`
}

func (x *ignoreInvitesCmd) Execute(args []string) error {
	setApi(x.Client)
	if len(args) == 0 {
		return errMissingInviteId
	}

	res, err := executeStringCmd(POST, "invites/"+args[0]+"/ignore", params{
		args: args,
	})
	if err != nil {
		return err
	}
	output(res)
	return nil
}
