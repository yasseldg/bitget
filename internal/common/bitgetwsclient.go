package common

import (
	"fmt"
	"sync"
	"time"

	"github.com/yasseldg/bitget/config"
	"github.com/yasseldg/bitget/constants"
	"github.com/yasseldg/bitget/internal"
	"github.com/yasseldg/bitget/internal/model"
	slog "github.com/yasseldg/bitget/xy/logger"

	"github.com/gorilla/websocket"
	"github.com/robfig/cron"
)

type BitgetBaseWsClient struct {
	NeedLogin        bool
	Connection       bool
	LoginStatus      bool
	Listener         OnReceive
	ErrorListener    OnReceive
	Ticker           *time.Ticker
	SendMutex        *sync.Mutex
	WebSocketClient  *websocket.Conn
	LastReceivedTime time.Time
	AllSubscribe     *model.Set
	Signer           *Signer
	ListenerMap      map[model.SubscribeReq]OnReceive
	BooksMap         map[model.SubscribeReq]model.BookInfo
}

func (p *BitgetBaseWsClient) Init() *BitgetBaseWsClient {
	creds := config.GetDefaultCredentials()

	p.Connection = false
	p.AllSubscribe = model.NewSet()
	p.Signer = new(Signer).Init(creds.SecretKey)
	p.ListenerMap = make(map[model.SubscribeReq]OnReceive)
	p.BooksMap = make(map[model.SubscribeReq]model.BookInfo)
	p.SendMutex = &sync.Mutex{}
	p.LastReceivedTime = time.Now()

	if constants.TimerIntervalSecond > 0 {
		p.Ticker = time.NewTicker(constants.TimerIntervalSecond * time.Second)
	}

	return p
}

func (p *BitgetBaseWsClient) SetListener(msgListener OnReceive, errorListener OnReceive) {
	p.Listener = msgListener
	p.ErrorListener = errorListener
}

func (p *BitgetBaseWsClient) Connect() {
	p.ExecuterPing()
	go p.tickerLoop()
}

func (p *BitgetBaseWsClient) ConnectWebSocket() {
	var err error
	slog.Info("WebSocket connecting...")

	p.WebSocketClient, _, err = websocket.DefaultDialer.Dial(config.WsUrl, nil)
	if err != nil {
		slog.Error("WebSocket connected error: %s\n", err)
		return
	}

	slog.Info("WebSocket connected")
	p.Connection = true
}

func (p *BitgetBaseWsClient) Login() {
	creds := config.GetDefaultCredentials()

	timesStamp := internal.TimesStampSec()
	sign := p.Signer.Sign(constants.WsAuthMethod, constants.WsAuthPath, "", timesStamp)

	loginReq := model.WsLoginReq{
		ApiKey:     creds.ApiKey,
		Passphrase: creds.PASSPHRASE,
		Timestamp:  timesStamp,
		Sign:       sign,
	}
	var args []interface{}
	args = append(args, loginReq)

	baseReq := model.WsBaseReq{
		Op:   constants.WsOpLogin,
		Args: args,
	}
	p.SendByType(baseReq)
}

func (p *BitgetBaseWsClient) StartReadLoop() {
	go p.readLoop()
}

func (p *BitgetBaseWsClient) StartTickerLoop() {
	if constants.TimerIntervalSecond > 0 {
		go p.tickerLoop()
	}
}

func (p *BitgetBaseWsClient) ExecuterPing() {
	c := cron.New()
	_ = c.AddFunc(fmt.Sprintf("*/%d * * * * *", constants.PingIntervalSecond), p.ping)
	c.Start()
}

func (p *BitgetBaseWsClient) ping() {
	p.Send("ping")
}

func (p *BitgetBaseWsClient) SendByType(req model.WsBaseReq) {
	json, _ := internal.ToJson(req)
	p.Send(json)
}

func (p *BitgetBaseWsClient) Send(data string) {
	if p.WebSocketClient == nil {
		slog.Error("WebSocket sent error: no connection available")
		return
	}

	if data == constants.PingMessage {
		slog.Debug("WebSocket sendMessage: %s", data)
	} else {
		slog.Info("WebSocket sendMessage: %s", data)
	}

	p.SendMutex.Lock()
	err := p.WebSocketClient.WriteMessage(websocket.TextMessage, []byte(data))
	p.SendMutex.Unlock()
	if err != nil {
		slog.Error("WebSocket sent error: data=%s, error=%s", data, err)
	}
}

func (p *BitgetBaseWsClient) SubscribeAll() {

	var args []interface{}
	for _, reg := range p.AllSubscribe.List() {
		args = append(args, reg)
	}

	wsBaseReq := model.WsBaseReq{
		Op:   constants.WsOpSubscribe,
		Args: args,
	}
	p.SendByType(wsBaseReq)
}

func (p *BitgetBaseWsClient) tickerLoop() {
	slog.Info("WebSocket TickerLoop started")
	for {
		<-p.Ticker.C
		slog.Debug("WebSocket ticker")
		elapsedSecond := time.Since(p.LastReceivedTime).Seconds()
		if elapsedSecond > constants.PingIntervalSecond {
			slog.Info("WebSocket reconnect...")
			p.DisconnectWebSocket()
			p.ConnectWebSocket()
			if p.Connection {
				p.LastReceivedTime = time.Now()
				// Subscribe All again
				p.SubscribeAll()
			}
		}
	}
}

func (p *BitgetBaseWsClient) DisconnectWebSocket() {
	if p.WebSocketClient == nil {
		return
	}

	slog.Info("WebSocket disconnecting...")

	err := p.WebSocketClient.Close()
	if err != nil {
		slog.Error("WebSocket disconnect error: %s\n", err)
		return
	}

	slog.Info("WebSocket disconnected")
}

func (p *BitgetBaseWsClient) readLoop() {
	for {
		if p.WebSocketClient == nil {
			slog.Warn("WebSocket Read error: no connection available")
			time.Sleep(time.Second)
			continue
		}

		_, buf, err := p.WebSocketClient.ReadMessage()
		if err != nil {
			slog.Warn("WebSocket Read error: %s", err)
			time.Sleep(time.Second)
			continue
		}

		p.LastReceivedTime = time.Now()

		message := string(buf)
		if message == constants.PongMessage {
			slog.Debug("WebSocket Keep connected: %s", message)
			continue
		}

		jsonMap := internal.JSONToMap(message)

		v, e := jsonMap["code"]
		if e && int(v.(float64)) != 0 {
			p.ErrorListener(message)
			continue
		}

		v, e = jsonMap["event"]
		if e && v == constants.WsOpLogin {
			slog.Info("WebSocket login msg: %s", message)
			p.LoginStatus = true
			continue
		}

		if !p.CheckSum(jsonMap) {
			continue
		}

		v, e = jsonMap["data"]
		if e {
			listener := p.GetListener(jsonMap["arg"])

			listener(message)
			continue
		}
		p.handleMessage(message)
	}
}

func (p *BitgetBaseWsClient) CheckSum(jsonMap map[string]interface{}) bool {
	argV, argE := jsonMap["arg"]
	actionV, actionE := jsonMap["action"]

	if !argE || !actionE {
		return true
	}
	channelV, _ := argV.(map[string]interface{})["channel"]
	if channelV != "books" {
		return true
	}
	dataV, _ := jsonMap["data"]
	if dataV == nil {
		return true
	}
	data := dataV.([]interface{})[0]

	asks := data.(map[string]interface{})["asks"].([]interface{})
	bids := data.(map[string]interface{})["bids"].([]interface{})
	checksum := data.(map[string]interface{})["checksum"].(float64)
	bookInfo := model.BookInfo{
		Asks:     asks,
		Bids:     bids,
		Checksum: fmt.Sprintf("%f", checksum),
	}

	mapData := argV.(map[string]interface{})
	subscribeReq := model.SubscribeReq{
		InstType: fmt.Sprintf("%v", mapData["instType"]),
		Channel:  fmt.Sprintf("%v", mapData["channel"]),
		InstId:   fmt.Sprintf("%v", mapData["instId"]),
	}

	if actionV == "snapshot" {
		p.BooksMap[subscribeReq] = bookInfo
		return true
	}
	if actionV == "update" {
		allBookInfo, subE := p.BooksMap[subscribeReq]

		if !subE {
			return true
		}
		newBoooks := allBookInfo.Merge(bookInfo)
		p.BooksMap[subscribeReq] = newBoooks
		return newBoooks.CheckSum(uint32(checksum))
	}
	return true
}

func (p *BitgetBaseWsClient) GetListener(argJson interface{}) OnReceive {

	mapData := argJson.(map[string]interface{})

	subscribeReq := model.SubscribeReq{
		InstType: fmt.Sprintf("%v", mapData["instType"]),
		Channel:  fmt.Sprintf("%v", mapData["channel"]),
		InstId:   fmt.Sprintf("%v", mapData["instId"]),
	}

	v, e := p.ListenerMap[subscribeReq]

	if !e {
		return p.Listener
	}
	return v
}

type OnReceive func(message string)

func (p *BitgetBaseWsClient) handleMessage(msg string) {
	slog.Info("WebSocket default: %s ", msg)
}
