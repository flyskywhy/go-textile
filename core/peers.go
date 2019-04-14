package core

import (
	"fmt"

	peer "github.com/libp2p/go-libp2p-peer"
	"github.com/textileio/go-textile/ipfs"
	"github.com/textileio/go-textile/pb"
	"github.com/textileio/go-textile/util"
)

// PeerUser returns a user object with the most recently updated contact for the given id
// Note: If no underlying contact is found, this will return an blank object w/ a
// generic user name for display-only purposes.
func (t *Textile) PeerUser(id string) *pb.User {
	p := t.datastore.Peers().GetBest(id)
	if p == nil {
		return &pb.User{
			Name: ipfs.ShortenID(id),
		}
	}
	return &pb.User{
		Address: p.Address,
		Name:    toName(p),
		Avatar:  p.Avatar,
	}
}

// addPeer adds or updates a peer
func (t *Textile) addPeer(peer *pb.Peer) error {
	x := t.datastore.Peers().Get(peer.Id)
	if x != nil && (peer.Updated == nil || util.ProtoTsIsNewer(x.Updated, peer.Updated)) {
		return nil
	}

	// peer is new / newer, update
	if err := t.datastore.Peers().AddOrUpdate(peer); err != nil {
		return err
	}

	// ensure new update is actually different before announcing to account
	if x != nil {
		if peersEqual(x, peer) {
			return nil
		}
	}

	thrd := t.AccountThread()
	if thrd == nil {
		return fmt.Errorf("account thread not found")
	}

	if _, err := thrd.annouce(&pb.ThreadAnnounce{Peer: peer}); err != nil {
		return err
	}
	return nil
}

// publishPeer publishes this peer's info to the cafe network
func (t *Textile) publishPeer() error {
	self := t.datastore.Peers().Get(t.node.Identity.Pretty())
	if self == nil {
		return nil
	}

	sessions := t.datastore.CafeSessions().List().Items
	if len(sessions) == 0 {
		return nil
	}
	for _, session := range sessions {
		pid, err := peer.IDB58Decode(session.Id)
		if err != nil {
			return err
		}
		if err := t.cafe.PublishPeer(self, pid); err != nil {
			return err
		}
	}
	return nil
}

// updatePeerInboxes sets own peer inboxes from the current cafe sessions
func (t *Textile) updatePeerInboxes() error {
	var inboxes []*pb.Cafe
	for _, session := range t.datastore.CafeSessions().List().Items {
		inboxes = append(inboxes, session.Cafe)
	}
	return t.datastore.Peers().UpdateInboxes(t.node.Identity.Pretty(), inboxes)
}

// toName returns a peer's name or trimmed address
func toName(peer *pb.Peer) string {
	if peer == nil || peer.Address == "" {
		return ""
	}
	if peer.Name != "" {
		return peer.Name
	}
	if len(peer.Address) >= 7 {
		return peer.Address[:7]
	}
	return ""
}

// peersEqual returns whether or not the two peers are identical
// Note: this does not consider Peer.Created or Peer.Updated
func peersEqual(a *pb.Peer, b *pb.Peer) bool {
	if a.Id != b.Id {
		return false
	}
	if a.Address != b.Address {
		return false
	}
	if a.Name != b.Name {
		return false
	}
	if a.Avatar != b.Avatar {
		return false
	}
	if len(a.Inboxes) != len(b.Inboxes) {
		return false
	}
	ac := make(map[string]*pb.Cafe)
	for _, c := range a.Inboxes {
		ac[c.Peer] = c
	}
	for _, j := range b.Inboxes {
		i, ok := ac[j.Peer]
		if !ok {
			return false
		}
		if !cafesEqual(i, j) {
			return false
		}
	}
	return true
}
