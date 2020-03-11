package webconsole

import (
	"net/http"
	"strconv"
	"strings"
)

type Console struct {
	hubs    map[string]*ConsoleHub
	Timeout int
	Port    int16
}

type ConsoleHub struct {
	text string
}

func i2s(i int) string {
	return strconv.Itoa(i)
}
func (c *Console) CreateHub(name string) *ConsoleHub {
	if c.hubs == nil {
		c.hubs = make(map[string]*ConsoleHub)
	}
	hub := &ConsoleHub{}
	c.hubs[name] = hub
	return hub
}

const LINEFEED = "----------------------------------------"

func (c *Console) Handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	w.Header().Add("Refresh", i2s(c.Timeout)+"; url="+url)
	w.Header().Set("Content-Type", "text/plain")
	url = strings.TrimPrefix(url, "/")
	hub, ok := c.hubs[url]
	if ok {
		w.Write([]byte(hub.text))
	} else {
		s := url + "\n\n"
		for key, value := range c.hubs {
			ds := "----- [" + key + "] "
			c := 40 - len(ds)
			if c < 0 {
				c = 0
			}
			s += ds + strings.Repeat("-", 40-len(ds)) + "\n" + value.text + "\n" + LINEFEED + "\n"
		}
		w.Write([]byte(s))
	}
}

func (c *ConsoleHub) Log(s string) {
	c.text += s
}

func (c *ConsoleHub) Clear() {
	c.text = ""
}

func (c *Console) Start() {
	go c.start()
}

func (c *Console) start() {
	http.HandleFunc("/", c.Handler)
	p := i2s(int(c.Port))
	hub := c.CreateHub(".")
	hub.Log("Starting HTTP server on :" + p)
	hub.Log(http.ListenAndServe(":"+p, nil).Error())
}
