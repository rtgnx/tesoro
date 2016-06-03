package tesoro

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/unicode/norm"

	"net/url"

	"math"

	"crypto/hmac"
	"crypto/sha256"

	"github.com/conejoninja/tesoro/pb/messages"
	"github.com/conejoninja/tesoro/pb/types"
	"github.com/conejoninja/tesoro/transport"
	"github.com/golang/protobuf/proto"
	"github.com/zserge/hid"
)

const hardkey uint32 = 2147483648

type Client struct {
	t transport.TransportHID
}

type PasswordConfig struct {
	Version string  `json:version`
	Config  string  `json:config`
	Tags    []Tag   `json:tags`
	Entries []Entry `json:entries`
}

type Tag struct {
	Title  string `json:title`
	Icon   string `json:icon`
	Active string `json:active`
}

type Entry struct {
	Title    string `json:title`
	Username string `json:username`
	Nonce    string `json:nonce`
	Note     string `json:note`
	Password string `json:password`
	SafeNote string `json:safe_note`
	Tags     string `json:tags`
}

func (c *Client) SetTransport(device hid.Device) {
	c.t.SetDevice(device)
}

func (c *Client) CloseTransport() {
	c.t.Close()
}

func (c *Client) Header(msgType int, msg []byte) []byte {

	typebuf := make([]byte, 2)
	binary.BigEndian.PutUint16(typebuf, uint16(msgType))

	msgbuf := make([]byte, 4)
	binary.BigEndian.PutUint32(msgbuf, uint32(len(msg)))

	return append(typebuf, msgbuf...)
}

func (c *Client) Initialize() []byte {
	var m messages.Initialize
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_Initialize"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) Ping(str string, pinProtection, passphraseProtection, buttonProtection bool) []byte {
	var m messages.Ping
	m.Message = &str
	m.ButtonProtection = &buttonProtection
	m.PinProtection = &pinProtection
	m.PassphraseProtection = &passphraseProtection
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_Ping"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) ChangePin() []byte {
	var m messages.ChangePin
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_ChangePin"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) GetEntropy(size uint32) []byte {
	var m messages.GetEntropy
	m.Size = &size
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_GetEntropy"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) GetFeatures() []byte {
	var m messages.GetFeatures
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_GetFeatures"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) PinMatrixAck(str string) []byte {
	var m messages.PinMatrixAck
	m.Pin = &str
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_PinMatrixAck"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) GetAddress(addressN []uint32, showDisplay bool, coinName string) []byte {
	var m messages.GetAddress
	m.AddressN = addressN
	m.CoinName = &coinName
	m.ShowDisplay = &showDisplay
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_GetAddress"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) GetPublicKey(address []uint32) []byte {
	var m messages.GetPublicKey
	m.AddressN = address
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_GetPublicKey"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) SignMessage(message []byte) []byte {
	var m messages.SignMessage
	m.Message = norm.NFC.Bytes(message)
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_SignMessage"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) SignIdentity(uri string, challengeHidden []byte, challengeVisual string, index uint32) []byte {
	var m messages.SignIdentity
	identity := URIToIdentity(uri)
	identity.Index = &index
	m.Identity = &identity
	m.ChallengeHidden = challengeHidden
	m.ChallengeVisual = &challengeVisual
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_SignIdentity"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) SetLabel(label string) []byte {
	var m messages.ApplySettings
	m.Label = &label
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_ApplySettings"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) SetHomescreen(homescreen []byte) []byte {
	var m messages.ApplySettings
	m.Homescreen = homescreen
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_ApplySettings"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) VerifyMessage(address, signature string, message []byte) []byte {

	sign, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return []byte("Wrong signature")
	}

	var m messages.VerifyMessage
	m.Address = &address
	m.Signature = sign
	m.Message = norm.NFC.Bytes(message)
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_VerifyMessage"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) ButtonAck() []byte {
	var m messages.ButtonAck
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_ButtonAck"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) GetMasterKey() []byte {
	masterKey, _ := hex.DecodeString("2d650551248d792eabf628f451200d7f51cb63e46aadcbb1038aacb05e8c8aee2d650551248d792eabf628f451200d7f51cb63e46aadcbb1038aacb05e8c8aee")
	return c.CipherKeyValue(
		true,
		"Activate TREZOR Password Manager?",
		masterKey,
		StringToBIP32Path("m/10016'/0"),
		[]byte{},
		true,
		true,
	)
}

func (c *Client) ClearSession() []byte {
	var m messages.ClearSession
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_ClearSession"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) CipherKeyValue(encrypt bool, key string, value []byte, address []uint32, iv []byte, askOnEncrypt, askOnDecrypt bool) []byte {
	var m messages.CipherKeyValue
	m.Key = &key
	if encrypt {
		paddedValue := make([]byte, 16*int(math.Ceil(float64(len(value))/16)))
		copy(paddedValue, value)
		m.Value = paddedValue
	} else {
		var err error
		m.Value, err = hex.DecodeString(string(value))
		if err != nil {
			fmt.Println("ERROR Decoding string")
		}
	}
	m.AddressN = address
	if len(iv) > 0 {
		m.Iv = iv
	}
	m.Encrypt = &encrypt
	m.AskOnEncrypt = &askOnEncrypt
	m.AskOnDecrypt = &askOnDecrypt
	marshalled, err := proto.Marshal(&m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(int(messages.MessageType_value["MessageType_CipherKeyValue"]), marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) Call(msg []byte) (string, uint16) {
	c.t.Write(msg)
	return c.ReadUntil()
}

func (c *Client) ReadUntil() (string, uint16) {
	var str string
	var msgType uint16
	for {
		str, msgType = c.Read()
		if msgType != 999 { //timeout
			break
		}
	}

	return str, msgType
}

func (c *Client) Read() (string, uint16) {
	marshalled, msgType, msgLength, err := c.t.Read()
	if err != nil {
		return "Error reading", 999
	}
	if msgLength <= 0 {
		fmt.Println("Empty message", msgType)
		return "", msgType
	}

	str := "Uncaught message type " + strconv.Itoa(int(msgType))
	if msgType == 2 {
		var msg messages.Success
		err = proto.Unmarshal(marshalled, &msg)
		if err != nil {
			str = "Error unmarshalling (2)"
		} else {
			str = msg.GetMessage()
		}
	} else if msgType == 3 {
		var msg messages.Failure
		err = proto.Unmarshal(marshalled, &msg)
		if err != nil {
			str = "Error unmarshalling (3)"
		} else {
			str = msg.GetMessage()
		}
	} else if msgType == 10 {
		var msg messages.Entropy
		err = proto.Unmarshal(marshalled, &msg)
		if err != nil {
			str = "Error unmarshalling (10)"
		} else {
			str = hex.EncodeToString(msg.GetEntropy())
		}
	} else if msgType == 12 {
		var msg messages.PublicKey
		err = proto.Unmarshal(marshalled, &msg)
		if err != nil {
			str = "Error unmarshalling (12)"
		} else {
			str = msg.GetXpub()
		}
	} else if msgType == 17 {
		var msg messages.Features
		err = proto.Unmarshal(marshalled, &msg)
		if err != nil {
			str = "Error unmarshalling (17)"
		} else {
			ftsJSON, _ := json.Marshal(&msg)
			str = string(ftsJSON)
		}
	} else if msgType == 18 {
		var msg messages.PinMatrixRequest
		err = proto.Unmarshal(marshalled, &msg)
		if err != nil {
			str = "Error unmarshalling (18)"
		} else {
			msgSubType := msg.GetType()
			if msgSubType == 1 {
				str = "Please enter current PIN:"
			} else if msgSubType == 2 {
				str = "Please enter new PIN:"
			} else {
				str = "Please re-enter new PIN:"
			}
		}
	} else if msgType == 26 {
		var msg messages.ButtonRequest
		err = proto.Unmarshal(marshalled, &msg)
		if err != nil {
			str = "Error unmarshalling (26)"
		} else {
			str = "Confirm action on TREZOR device"
		}
	} else if msgType == 30 {
		var msg messages.Address
		err = proto.Unmarshal(marshalled, &msg)
		if err != nil {
			str = "Error unmarshalling (30)"
		} else {
			str = msg.GetAddress()
		}
	} else if msgType == 40 {
		var msg messages.MessageSignature
		err = proto.Unmarshal(marshalled, &msg)
		if err != nil {
			str = "Error unmarshalling (40)"
		} else {
			smJSON, _ := json.Marshal(&msg)
			str = string(smJSON)
		}
	} else if msgType == 48 {
		var msg messages.CipheredKeyValue
		err = proto.Unmarshal(marshalled, &msg)
		if err != nil {
			str = "Error unmarshalling (48)"
		} else {
			str = string(msg.GetValue())
		}
	} else if msgType == 54 {
		var msg messages.SignedIdentity
		err = proto.Unmarshal(marshalled, &msg)
		if err != nil {
			str = "Error unmarshalling (54)"
		} else {
			smJSON, _ := json.Marshal(&msg)
			str = string(smJSON)
		}
	}
	return str, msgType
}

func BIP32Path(keys []uint32) string {
	path := "m"
	for _, key := range keys {
		path += "/"
		if key < hardkey {
			path += string(key)
		} else {

			path += string(key-hardkey) + "'"
		}
	}
	return path
}

func StringToBIP32Path(str string) []uint32 {

	if !ValidBIP32(str) {
		return []uint32{}
	}

	re := regexp.MustCompile("([/]+)")
	str = re.ReplaceAllString(str, "/")

	keys := strings.Split(str, "/")
	path := make([]uint32, len(keys)-1)
	for k := 1; k < len(keys); k++ {
		i, _ := strconv.Atoi(strings.Replace(keys[k], "'", "", -1))
		if strings.Contains(keys[k], "'") {
			path[k-1] = hardened(uint32(i))
		} else {
			path[k-1] = uint32(i)
		}
	}
	return path
}

func ValidBIP32(path string) bool {
	re := regexp.MustCompile("([/]+)")
	path = re.ReplaceAllString(path, "/")

	re = regexp.MustCompile("^m/")
	path = re.ReplaceAllString(path, "")

	re = regexp.MustCompile("'/")
	path = re.ReplaceAllString(path+"/", "")

	re = regexp.MustCompile("[0-9/]+")
	path = re.ReplaceAllString(path, "")

	return path == ""
}

func PNGToString(filename string) ([]byte, error) {
	img := make([]byte, 1024)
	infile, err := os.Open(filename)
	if err != nil {
		return img, err
	}
	defer infile.Close()

	src, _, err := image.Decode(infile)
	if err != nil {
		return img, err
	}

	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	if w != 128 || h != 64 {
		err = errors.New("Wrong homescreen size")
		return img, err
	}

	imgBin := ""
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			color := src.At(i, j)
			r, g, b, _ := color.RGBA()
			if (r + g + b) > 0 {
				imgBin += "1"
			} else {
				imgBin += "0"
			}
		}
	}
	k := 0
	for i := 0; i < len(imgBin); i += 8 {
		if s, err := strconv.ParseUint(imgBin[i:i+8], 2, 32); err == nil {
			img[k] = byte(s)
			k++
		}
	}
	return img, nil
}

func URIToIdentity(uri string) types.IdentityType {
	var identity types.IdentityType
	u, err := url.Parse(uri)
	if err != nil {
		return identity
	}

	defaultPort := ""
	identity.Proto = &u.Scheme
	user := ""
	if u.User != nil {
		user = u.User.String()
	}
	identity.User = &user
	tmp := strings.Split(u.Host, ":")
	identity.Host = &tmp[0]
	if len(tmp) > 1 {
		identity.Port = &tmp[1]
	} else {
		identity.Port = &defaultPort
	}
	identity.Path = &u.Path
	return identity
}

func hardened(key uint32) uint32 {
	return hardkey + key
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func AES256GCMMEncrypt(plainText, key []byte) (string, string) {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	cipheredText := aesgcm.Seal(nil, nonce, plainText, nil)
	return hex.EncodeToString(cipheredText), hex.EncodeToString(nonce)
}

func AES256GCMDecrypt(cipheredText, key, nonce, tag []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("TEXT", string(cipheredText), len(cipheredText))
	fmt.Println("NONCE", string(nonce), len(nonce))
	fmt.Println("TAG", string(tag), len(tag))
	plainText, err := aesgcm.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plainText)
}

func GetFileEncKey(masterKey string) (string, string, string) {
	fileKey := masterKey[:len(masterKey)/2]
	encKey := masterKey[len(masterKey)/2:]
	filename_mess := []byte("5f91add3fa1c3c76e90c90a3bd0999e2bd7833d06a483fe884ee60397aca277a")
	mac := hmac.New(sha256.New, []byte(fileKey))
	mac.Write(filename_mess)
	tmpMac := mac.Sum(nil)
	digest := hex.EncodeToString(tmpMac)
	filename := digest + ".pswd"
	return filename, fileKey, encKey
}

func DecryptStorage(content, key string) string {
	cipherKey, _ := hex.DecodeString(key)
	return AES256GCMDecrypt([]byte(content[28:]+content[12:28]), cipherKey, []byte(content[:12]), []byte(content[12:28]))
}