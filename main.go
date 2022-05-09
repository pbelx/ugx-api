package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	_ "github.com/mattn/go-sqlite3"
)

func getip(c *gin.Context) {
	dbs, err := sql.Open("sqlite3", "./logs.db")

	if err != nil {
		fmt.Println(err)
	}
	defer dbs.Close()
	uaj := c.GetHeader("User-Agent")
	urlx := c.FullPath()
	ipx := c.ClientIP()
	// timex := time.Now()
	// fmt.Println(uaj, urlx, ipx, timex)
	stmt := `INSERT INTO xlog (useragent,ip,url) VALUES (?,?,?)`
	dbs.Prepare(stmt)
	fmt.Println(ipx)

	dbs.Exec(stmt, uaj, ipx, urlx)

	rows, err := dbs.Query(`SELECT useragent,ip FROM xlog`)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {

		var useragent string
		var ip string
		rows.Scan(&useragent, &ip)
		// fmt.Println("w00t")
		// fmt.Printf(("%v,%v\n"), useragent, ip)

	}

}

func test(c *gin.Context) {
	getip(c)
	c.JSON(200, gin.H{
		"RadioStations": "=" +
			"sanyu,hot100,rxfm,hii,rcity,ugdjs,homeboyz,capitalke,power,",
		"Version": "Online Radio Server(V2)",
	})

}

func sanyu(c *gin.Context) {
	getip(c)
	resp, err := http.Get("http://s44.myradiostream.com:8138")
	if err != nil {
		c.JSON(200, gin.H{
			"message": "service offline",
		})
	} else {
		io.Copy(c.Writer, resp.Body)
	}

}
func foo(c *gin.Context) {
	resp, err := http.Get("http://127.0.0.1")
	if err != nil {
		c.JSON(200, gin.H{
			"message": "service offline",
		})

	} else {
		io.Copy(c.Writer, resp.Body)
	}

}
func hot100(c *gin.Context) {
	getip(c)
	resp, err := http.Get("https://www.hot100.ug/stream?")
	if err != nil {
		c.JSON(200, gin.H{
			"message": "service offline",
		})
	} else {
		io.Copy(c.Writer, resp.Body)
	}
}
func rxfm(c *gin.Context) {
	getip(c)
	resp, err := http.Get("https://c14.radioboss.fm/stream/223?1618402604653")
	if err != nil {
		c.JSON(200, gin.H{
			"message": "service offline",
		})
	} else {
		io.Copy(c.Writer, resp.Body)
	}

}
func hii(c *gin.Context) {
	resp, err := http.Get("http://eu9.fastcast4u.com:4556/stream?")
	if err != nil {
		c.JSON(200, gin.H{
			"message": "service offline",
		})
	} else {
		io.Copy(c.Writer, resp.Body)
	}

}
func rcity(c *gin.Context) {
	getip(c)
	resp, err := http.Get("http://162.244.80.106:10853/stream?")
	if err != nil {
		c.JSON(200, gin.H{
			"message": "service offline",
		})
	} else {
		io.Copy(c.Writer, resp.Body)
	}

}
func ugdjs(c *gin.Context) {
	getip(c)
	resp, err := http.Get("http://node-14.zeno.fm/vkunv86994zuv?")
	if err != nil {
		c.JSON(200, gin.H{
			"message": "service offline",
		})
	} else {
		io.Copy(c.Writer, resp.Body)
	}

}
func homeboyz(c *gin.Context) {
	getip(c)
	resp, err := http.Get("https://homeboyzradio-atunwadigital.streamguys1.com/homeboyzradio")
	if err != nil {
		c.JSON(200, gin.H{
			"message": "service offline",
		})
	} else {
		io.Copy(c.Writer, resp.Body)
	}

}
func capitalke(c *gin.Context) {
	getip(c)
	resp, err := http.Get("https://atunwadigital.streamguys1.com/capitalfm?")
	if err != nil {
		c.JSON(200, gin.H{
			"message": "service offline",
		})
	} else {
		io.Copy(c.Writer, resp.Body)
	}
}
func power(c *gin.Context) {
	getip(c)
	resp, err := http.Get("https://securestreams5.autopo.st:1941/stream?")
	if err != nil {
		c.JSON(200, gin.H{
			"message": "service offline",
		})
	} else {
		io.Copy(c.Writer, resp.Body)
	}
}
func flightF(c *gin.Context) {
	getip(c)

	type Flights struct {
		Airline    string `json:"airline,omitempty"`
		FlightCode string `json:"flightCode,omitempty"`
		ET         string `json:"expectedtime,omitempty"`
		ORIGIN     string `json:"origin,omitempty"`
		Status     string `json:"status,omitempty"`
		Category   string `json:"category,omitempty"`
	}
	type Flights2 struct {
		Airline     string `json:"airline,omitempty"`
		FlightCode  string `json:"flightCode,omitempty"`
		ET          string `json:"expectedtime,omitempty"`
		Destination string `json:"destination,omitempty"`
		Status      string `json:"status,omitempty"`
		Category    string `json:"category,omitempty"`
	}

	Flightlist := make([]Flights, 0)
	DFlightlist := make([]Flights2, 0)
	cc := colly.NewCollector()
	cc.OnHTML(".fids_table", func(e *colly.HTMLElement) {
		item := Flights{}
		itemx := Flights2{}
		// item.Category = e.ChildText("th:nth-child(4)")
		category := e.ChildText("th:nth-child(5)")
		// airline := e.ChildText("td:nth-child(1)")
		// fcode := e.ChildText("td:nth-child(2)")
		// status := e.ChildText("td:nth-child(6)")

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {

			if category == "ETA" {
				if el.ChildText("td:nth-child(1)") != "" {
					item.Category = el.ChildText("th:nth-child(4)")
					item.Airline = el.ChildText("td:nth-child(1)")
					// fmt.Println(item.Airline)
					item.Status = el.ChildText("td:nth-child(6)")
					item.ORIGIN = el.ChildText("td:nth-child(4)")
					item.FlightCode = el.ChildText("td:nth-child(3)")
					item.ET = el.ChildText("td:nth-child(5)")
					Flightlist = append(Flightlist, item)
					// item.Category = category
					// item.Airline = airline
					// item.Status = status
					// item.FlightCode = fcode
					// Flightlist = append(Flightlist, item)
				}
			} else if category == "ETD" {
				if el.ChildText("td:nth-child(1)") != "" {
					itemx.Category = el.ChildText("th:nth-child(4)")
					itemx.Airline = el.ChildText("td:nth-child(1)")

					itemx.Status = el.ChildText("td:nth-child(6)")
					itemx.Destination = el.ChildText("td:nth-child(4)")
					itemx.FlightCode = el.ChildText("td:nth-child(3)")
					itemx.ET = el.ChildText("td:nth-child(5)")
					DFlightlist = append(DFlightlist, itemx)
				}
			}
		})
	})
	cc.Visit("https://caa.go.ug/arrivals-and-departures/")

	c.JSON(200, gin.H{
		"data":  Flightlist,
		"data2": DFlightlist,
	})
}

func Forexrates(c *gin.Context) {
	getip(c)
	type rates struct {
		EUR float64 `json:"EUR"`
		GBP float64 `json:"GBP"`
		KES float64 `json:"KES"`
		UGX float64 `json:"UGX"`
	}

	type jsondata struct {
		Base       string `json:"base"`
		Disclaimer string `json:"disclaimer"`
		License    string `json:"license"`
		Rates      rates  `json:"rates"`
		Timestamp  string `json:"timestamp"`
	}

	fmt.Println("Fetching Rates now...")
	resp, err := http.Get("https://openexchangerates.org/api/latest.json?app_id=868aea88493f4ecfb54940830249a5c0&symbols=UGX,EUR,GBP,KES")
	if err != nil {
		c.JSON(200, gin.H{
			"message": "service offline",
		})
	} else {
		reply, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "service offline 404",
			})
		}

		// fmt.Println(string(reply))

		var xrates jsondata
		json.Unmarshal([]byte(reply), &xrates)
		//get ugx rate
		realugx := math.Floor(float64(xrates.Rates.UGX))
		//get other currencies compared to ugx
		// fmt.Println(realugx)

		Euros := math.Floor(realugx / xrates.Rates.EUR)
		Pounds := math.Floor(realugx / xrates.Rates.GBP)
		Ksh := math.Floor(realugx / xrates.Rates.KES)

		// fmt.Printf("Euros %v ,Pounds %v,Ksh %v", Euros, Pounds, Ksh)
		c.JSON(200, gin.H{
			"Euros":  Euros,
			"Pounds": Pounds,
			"KSH":    Ksh,
			"USD":    realugx,
		})
	}
}

func quotes(c *gin.Context) {
	getip(c)

	type xquotes []struct {
		Q string `json:"q"`
		H string `json:"h"`
		A string `json:"a"`
	}

	resp, err := http.Get("https://zenquotes.io/api/random")
	if err != nil {
		c.JSON(200, gin.H{
			"message": "service offline",
		})
	} else {
		bodydata, err := ioutil.ReadAll(resp.Body)
		// fmt.Printf(("%T"), bodydata)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "no data response",
			})
		}
		var qq xquotes

		json.Unmarshal([]byte(bodydata), &qq)
		author := qq[0].A
		quote := qq[0].Q

		c.JSON(200, gin.H{
			"Author": author,
			"Quote":  quote,
		})
	}

}

func main() {
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("gin go")
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.GET("/", test)
	//radios start
	r.GET("/sanyu", sanyu)
	r.GET("/hot100", hot100)
	r.GET("/rxfm", rxfm)
	r.GET("/hii", hii)
	r.GET("/rcity", rcity)
	r.GET("/ugdjs", ugdjs)
	r.GET("/homeboyz", homeboyz)
	r.GET("/capitalke", capitalke)
	r.GET("/power", power)
	//radios end

	//forexrates
	r.GET("/forex", Forexrates)
	// quotes
	r.GET("/quotes", quotes)
	//flight info
	r.GET("/flight", flightF)
	r.GET("/foo", foo)

	liststations := "V2 Stations URL=>,sanyu,hot100,rxfm,hii" +
		"rcity,ugdjs,homeboyz" +
		",capitalke,power"
	fmt.Println("Server up port 8080")
	fmt.Print(liststations)
	r.Run()
}
