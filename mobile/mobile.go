package mobile

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/op/go-logging"

	tcore "github.com/textileio/textile-go/core"
	"github.com/textileio/textile-go/net"

	"gx/ipfs/QmZoWKhxUmZ2seW4BzX6fJkNR8hh9PsGModr7q171yq2SS/go-libp2p-peer"
	libp2p "gx/ipfs/QmaPbCnUMBohSGo3KnxEa2bHqyJVVeEEcwtqJAYxerieBo/go-libp2p-crypto"

	"github.com/textileio/textile-go/central/models"
)

var log = logging.MustGetLogger("mobile")

type Wrapper struct {
	RepoPath string
}

func NewNode(repoPath string, centralApiURL string) (*Wrapper, error) {
	var m Mobile
	return m.NewNode(repoPath, centralApiURL)
}

type Mobile struct{}

// Create a gomobile compatible wrapper around TextileNode
func (m *Mobile) NewNode(repoPath string, centralApiURL string) (*Wrapper, error) {
	node, err := tcore.NewNode(repoPath, centralApiURL, true, logging.DEBUG)
	if err != nil {
		return nil, err
	}
	tcore.Node = node

	return &Wrapper{RepoPath: repoPath}, nil
}

func (w *Wrapper) Start() error {
	if err := tcore.Node.Start(); err != nil {
		if err == tcore.ErrNodeRunning {
			return nil
		}
		return err
	}

	// join existing rooms
	albums := tcore.Node.Datastore.Albums().GetAlbums("")
	datacs := make([]chan string, len(albums))
	for i, a := range albums {
		go tcore.Node.JoinRoom(a.Id, datacs[i])
	}

	return nil
}

func (w *Wrapper) Stop() error {
	if err := tcore.Node.Stop(); err != nil && err != tcore.ErrNodeNotRunning {
		return err
	}
	return nil
}

func (w *Wrapper) SignUpWithEmail(username string, password string, email string, referral string) error {
	// build registration
	reg := &models.Registration{
		Username: username,
		Password: password,
		Identity: &models.Identity{
			Type:  models.EmailAddress,
			Value: email,
		},
		Referral: referral,
	}

	// signup
	return tcore.Node.SignUp(reg)
}

func (w *Wrapper) SignIn(username string, password string) error {
	// build creds
	creds := &models.Credentials{
		Username: username,
		Password: password,
	}

	// signin
	return tcore.Node.SignIn(creds)
}

func (w *Wrapper) SignOut() error {
	return tcore.Node.SignOut()
}

func (w *Wrapper) IsSignedIn() bool {
	_, err := tcore.Node.Datastore.Config().GetUsername()
	return err == nil
}

func (w *Wrapper) GetUsername() (string, error) {
	un, err := tcore.Node.Datastore.Config().GetUsername()
	if err != nil {
		return "", err
	}
	return un, nil
}

func (w *Wrapper) GetAccessToken() (string, error) {
	at, _, err := tcore.Node.Datastore.Config().GetTokens()
	if err != nil {
		return "", err
	}
	return at, nil
}

func (w *Wrapper) GetGatewayPassword() string {
	return tcore.Node.GatewayPassword
}

func (w *Wrapper) AddPhoto(path string, thumb string, thread string) (*net.MultipartRequest, error) {
	return tcore.Node.AddPhoto(path, thumb, thread)
}

func (w *Wrapper) SharePhoto(hash string, thread string) (*net.MultipartRequest, error) {
	return tcore.Node.SharePhoto(hash, thread)
}

func (w *Wrapper) GetPhotos(offsetId string, limit int, thread string) (string, error) {
	list := tcore.Node.GetPhotos(offsetId, limit, thread)
	if list == nil {
		list = &tcore.PhotoList{
			Items: make([]tcore.PhotoListItem, 0),
		}
	}

	// gomobile does not allow slices. so, convert to json
	jsonb, err := json.Marshal(list)
	if err != nil {
		log.Errorf("error marshaling json: %s", err)
		return "", err
	}

	return string(jsonb), nil
}

func (w *Wrapper) GetFileBase64(path string) (string, error) {
	b, err := tcore.Node.GetFile(path, nil)
	if err != nil {
		return "error", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func (w *Wrapper) GetPeerID() (string, error) {
	if !tcore.Node.Online() {
		return "", tcore.ErrNodeNotRunning
	}
	return tcore.Node.IpfsNode.Identity.Pretty(), nil
}

func (w *Wrapper) PairDesktop(pkb64 string) (string, error) {
	if !tcore.Node.Online() {
		return "", tcore.ErrNodeNotRunning
	}
	log.Info("pairing with desktop...")

	pkb, err := base64.StdEncoding.DecodeString(pkb64)
	if err != nil {
		log.Errorf("error decoding string: %s: %s", pkb64, err)
		return "", err
	}

	pk, err := libp2p.UnmarshalPublicKey(pkb)
	if err != nil {
		log.Errorf("error unmarshaling pub key: %s", err)
		return "", err
	}

	// the phrase will be used by the desktop client to create
	// the private key needed to decrypt photos
	// we invite the desktop to _read and write_ to our default album
	da := tcore.Node.Datastore.Albums().GetAlbumByName("default")
	if da == nil {
		err = errors.New("default album not found")
		log.Error(err.Error())
		return "", err
	}
	// encypt with the desktop's pub key
	cph, err := net.Encrypt(pk, []byte(da.Mnemonic))
	if err != nil {
		log.Errorf("encrypt failed: %s", err)
		return "", err
	}

	// get the topic to pair with from the pub key
	peerID, err := peer.IDFromPublicKey(pk)
	if err != nil {
		log.Errorf("id from public key failed: %s", err)
		return "", err
	}
	topic := peerID.Pretty()

	// finally, publish the encrypted phrase
	err = tcore.Node.IpfsNode.Floodsub.Publish(topic, cph)
	if err != nil {
		log.Errorf("publish %s failed: %s", topic, err)
		return "", err
	}
	log.Infof("published key phrase to desktop: %s", topic)

	// try a ping
	err = tcore.Node.PingPeer(topic, 1, make(chan string))
	if err != nil {
		log.Errorf("ping %s failed: %s", topic, err)
	}

	return topic, nil
}
