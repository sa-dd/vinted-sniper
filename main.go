package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Vinted API response structures
type VintedResponse struct {
	Items []VintedItem `json:"items"`
}

type VintedItem struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Status    string `json:"status"`
	SizeTitle string `json:"size_title"`
	User      struct {
		Login string `json:"login"`
	} `json:"user"`
	Photo struct {
		URL string `json:"url"`
	} `json:"photo"`
	Price struct {
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currency_code"`
	} `json:"price"`
}

// Configuration
var (
	discordToken   = "MTEzMjAyNTY1MzA5NzY2MDQ3OA.GSGaZJ.YB_dFAx6gul8SDngWZjeLTbyk8_eruttn4Zobg"
	discordChannel = "1377199803607420938"
	vintedURL      = "https://www.vinted.co.uk/api/v2/catalog/items?page=1&per_page=96&time=1748554179&search_text=&catalog_ids=&order=newest_first&catalog_from=0&size_ids=206,207,208,209,210,211,212,308,1624&brand_ids=53,14,535,8715,162,21099,345,245062,359177,484362,313669,378906,140618,1798422,597509,1065021,57144,345731,269830,99164,73458,719079,670432,299684,1985410,311812,291429,1037965,472855,511110,299838,8139,401801,3063,1412112,725037,164166,190014,46923,506331,13727,345562,335419,318349,276609,228018,45765,6526387,204980,4810534,296216,103684,205836,834670,559746,1195,24491,280255,672258,222762&status_ids=&color_ids=&material_ids="
	checkInterval  = 1 * time.Second
)

// State management
var (
	lastItemIDs  = make(map[int64]bool)
	stateMutex   sync.Mutex
	requestCount = 0 // Track number of requests made
)

func main() {
	if discordToken == "" || discordChannel == "" {
		log.Fatal("DISCORD_TOKEN and DISCORD_CHANNEL environment variables must be set")
	}

	// Create Discord session
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	// Start session
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}
	defer dg.Close()

	log.Println("Bot is now running. Press CTRL-C to exit.")

	// Start the Vinted polling
	go pollVintedAPI(dg)

	// Wait for termination signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func pollVintedAPI(s *discordgo.Session) {
	// Create HTTP client with keep-alive
	transport := &http.Transport{
		MaxIdleConns:        10,
		IdleConnTimeout:     90 * time.Second,
		DisableCompression:  false,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: false},
		MaxIdleConnsPerHost: 10,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}

	// Initial delay to let Discord connect
	time.Sleep(2 * time.Second)

	// Polling loop
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for range ticker.C {
		items, err := fetchVintedItems(client)
		if err != nil {
			log.Printf("Error fetching items: %v", err)
			continue
		}

		newItems := findNewItems(items)

		// Increment request counter
		requestCount++

		// Skip sending for first 2 requests
		if requestCount > 2 && len(newItems) > 0 {
			log.Printf("Found %d new items, sending to Discord", len(newItems))
			sendItemsToDiscord(s, newItems)
		} else if requestCount <= 2 {
			log.Printf("Skipping send for request #%d (initial setup)", requestCount)
		} else {
			log.Println("No new items found")
		}

		// Update state with current items
		updateItemState(items)
	}
}

func fetchVintedItems(client *http.Client) ([]VintedItem, error) {
	req, err := http.NewRequest("GET", vintedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	// Set required headers
	req.Header = http.Header{
		"authority":          {"www.vinted.co.uk"},
		"accept":             {"application/json, text/plain, */*"},
		"accept-language":    {"en-uk-fr"},
		"cache-control":      {"max-age=0"},
		"cookie":             {"v_udt=NzdDNytmMnY2NVFLbjVKR0R2d25BcFdMOGhJbC0ta3pOYk9jMEdIRjFnUWE0MC0tNEN0TUZ5d2tyQW9vTkMrSWx4dXdYdz09; anon_id=a8c8ace7-ebf8-4144-9e8f-70a50132ec5a; anonymous-locale=en-uk-fr; _gcl_au=1.1.569003728.1748377126; _ga=GA1.1.1510117370.1748377127; _fbp=fb.2.1748377127218.761401381429566609; domain_selected=true; OptanonAlertBoxClosed=2025-05-27T20:19:00.203Z; eupubconsent-v2=CQSD0XgQSD0XgAcABBENBsFsAP_gAEPgAChQLNtT_G__bWlr-T73aftkeYxP99h77sQxBgbJE-4FzLvW_JwXx2E5NAzatqIKmRIAu3TBIQNlHJDURVCgaogVryDMaEyUoTNKJ6BkiFMRI2NYCFxvm4tjeQCY5vr991c1mB-t7dr83dzyy4hHn3a5_2S1WJCdAYetDfv8bBOT-9IOd_x8v4v4_F7pE2-eS1n_pWvp7D9-Yls_9X299_bbff7Pn_ful_-_X_vf_n37v943n77v_4LMgAmGhUQRlkQAhEoGEECABQVhARQIAgAASBogIATBgU5AwAXWEiAEAKAAYIAQAAgwABAAAJAAhEAFABQIAAIBAoAAwAIBgIAGBgADABYCAQAAgOgYpgQQCBYAJGZFBpgSgAJBAS2VCCQBAgrhCEWeAQQIiYKAAAEAAoAAEB4LAYkkBKxIIAuIJoAACAAAKIECBFJ2YAgoDNlqLwZPoytMAwfMEzSmAZAEQRkZJsQm_aYeOQogAAAA.f_wACHwAAAAA; OTAdditionalConsentString=1~43.46.55.61.70.83.89.93.108.117.122.124.135.143.144.147.149.159.192.196.211.228.230.239.259.266.286.291.311.318.320.322.323.327.367.371.385.394.407.415.424.430.436.445.486.491.494.495.522.523.540.550.559.560.568.574.576.584.587.591.737.802.803.820.821.839.864.899.904.922.931.938.959.979.981.985.1003.1027.1031.1040.1046.1051.1053.1067.1092.1095.1097.1099.1107.1109.1135.1143.1149.1152.1162.1166.1186.1188.1205.1215.1226.1227.1230.1252.1268.1270.1276.1284.1290.1301.1307.1312.1345.1356.1375.1403.1415.1416.1421.1423.1440.1449.1455.1495.1512.1516.1525.1540.1548.1555.1558.1570.1577.1579.1583.1584.1603.1616.1638.1651.1653.1659.1667.1677.1678.1682.1697.1699.1703.1712.1716.1721.1725.1732.1745.1750.1765.1782.1786.1800.1810.1825.1827.1832.1838.1840.1842.1843.1845.1859.1866.1870.1878.1880.1889.1899.1917.1929.1942.1944.1962.1963.1964.1967.1968.1969.1978.1985.1987.2003.2008.2027.2035.2039.2047.2052.2056.2064.2068.2072.2074.2088.2090.2103.2107.2109.2115.2124.2130.2133.2135.2137.2140.2147.2156.2166.2177.2186.2205.2213.2216.2219.2220.2222.2225.2234.2253.2275.2279.2282.2292.2309.2312.2316.2322.2325.2328.2331.2335.2336.2343.2354.2358.2359.2370.2376.2377.2387.2400.2403.2405.2407.2411.2414.2416.2418.2425.2440.2447.2461.2465.2468.2472.2477.2481.2484.2486.2488.2493.2498.2501.2510.2517.2526.2527.2532.2535.2542.2552.2563.2564.2567.2568.2569.2571.2572.2575.2577.2583.2584.2596.2604.2605.2608.2609.2610.2612.2614.2621.2628.2629.2633.2636.2642.2643.2645.2646.2650.2651.2652.2656.2657.2658.2660.2661.2669.2670.2677.2681.2684.2687.2690.2695.2698.2713.2714.2729.2739.2767.2768.2770.2772.2784.2787.2791.2792.2798.2801.2805.2812.2813.2816.2817.2821.2822.2827.2830.2831.2833.2834.2838.2839.2844.2846.2849.2850.2852.2854.2860.2862.2863.2865.2867.2869.2873.2874.2875.2876.2878.2880.2881.2882.2883.2884.2886.2887.2888.2889.2891.2893.2894.2895.2897.2898.2900.2901.2908.2909.2916.2917.2918.2919.2920.2922.2923.2927.2929.2930.2931.2940.2941.2947.2949.2950.2956.2958.2961.2963.2964.2965.2966.2968.2973.2975.2979.2980.2981.2983.2985.2986.2987.2994.2995.2997.2999.3000.3002.3003.3005.3008.3009.3010.3012.3016.3017.3018.3019.3028.3034.3038.3043.3052.3053.3055.3058.3059.3063.3066.3068.3070.3073.3074.3075.3076.3077.3089.3090.3093.3094.3095.3097.3099.3100.3106.3109.3112.3117.3119.3126.3127.3128.3130.3135.3136.3145.3150.3151.3154.3155.3163.3167.3172.3173.3182.3183.3184.3185.3187.3188.3189.3190.3194.3196.3209.3210.3211.3214.3215.3217.3219.3222.3223.3225.3226.3227.3228.3230.3231.3234.3235.3236.3237.3238.3240.3244.3245.3250.3251.3253.3257.3260.3270.3272.3281.3288.3290.3292.3293.3296.3299.3300.3306.3307.3309.3314.3315.3316.3318.3324.3328.3330.3331.3531.3731.3831.4131.4531.4631.4731.4831.5231.6931.7235.7831.7931.8931.9731.10231.10631.10831.11031.11531.12831.13632.13731.14034.14133.14237.14332.15731.16831.16931.21233.23031.25131.25731.25931.26031.26831.27731.27831.28031.28731.28831.29631.32531.33631.34231.34631.36831.39131.39531.40632.41531.43631.43731.43831; _cc_id=139f05f7879e2e62552c07fd0379937d; panoramaId_expiry=1748981941229; panoramaId=ee703d5830adb2438f77e4f647b316d53938a921ade9f78b20411bf22ac1e5f1; panoramaIdType=panoIndiv; v_uid=272304855; v_sid=e9248bee-1748425864; anon_id=a8c8ace7-ebf8-4144-9e8f-70a50132ec5a; RoktRecogniser=99142fcc-5146-478b-a58f-9f446929cb29; cto_bidid=Cjk6Yl9WdTNBclpnTlZDM1ZCQzhNZHllODJzVDElMkZtcm5pZGhteDNZVjJYazlTeWhscG5RWkVSaSUyRmUzR2J4OUdJeEVoNTBRTCUyRlBTT2J4TXJBWWlXTUZieGxjUlhoTGhCSWhoVmdqUzN5VHExZXdTbyUzRA; cto_dna_bundle=VwEd2V84RG1haGs0T1pUdHZadkRnQmhGdnRsU1NhbUpjS05MSElFWTJuZ3ZRc2xUcXN2dGkyRjQwcG0yUUpYS3glMkZsT2pQck1Ha3lOeHpiQmFXQnVSTldTcnhBJTNEJTNE; __gads=ID=7137fc25046da2f5:T=1748377140:RT=1748548339:S=ALNI_MbuJqCY6btlS49psCjL6F2AKIEonw; __gpi=UID=000010e132d6dad9:T=1748377140:RT=1748548339:S=ALNI_MYyEJntAKDvDEkrlNG-FYJ4mNungA; __eoi=ID=fd6d529e9e32c422:T=1748377140:RT=1748548339:S=AA-AfjadAHZXehL7sYt49FZqkLgU; viewport_size=725; datadome=lZme1DigCoNAscXS1nY0eRNH0KgwSZxbbnBnpw17OtM1bVyy7XAefh5OZ2ro4XuXX_jwE_P7nzSnvrB2FHZm8Y3p4VGfhtrE7_CgwTn~qvpeofkVdDDunY2_sItHdZpT; cto_bundle=GyUzVl84RG1haGs0T1pUdHZadkRnQmhGdnRoODRsbUZxWm85a0t3WEUzdHZrTng1SkpjQ3lZZ3BuQVE5RFBaYnNiRkN6dFZlMktDOFp5NTU1VXJkOFQlMkZxNXdqV3REY3olMkJSNUR2UFVmdk1VeDduUDJBV24zZlhsRk9abUZBUm1JTG5JTW1IeWRQazYwNFklMkJ4RjB4dGVwNm1XSFElM0QlM0Q; _ga_C1K5578GF1=GS2.1.s1748600102$o13$g1$t1748600208$j60$l0$h0; _ga_ZJHK1N3D75=GS2.1.s1748600102$o13$g1$t1748600208$j60$l0$h0; __cf_bm=AnuNRj6faE_N4sMkOyMyXtN0nN3LtAu7GjpWEazLIOc-1748620138-1.0.1.1-SAAWqPTtg96VdAb5k5HAd5BSMAm5vGfx7O29PTht2fEOW4J65Kl6Nm_X9lQyP5XVsHv4BhJz5G2M_NqelXDOmaOkZgk4PxhCF4nETiBXzk9EmFfo9EDwxTFEtj7PnTV8; cf_clearance=LfXe0bbOLSwzjmM5p8vD0339Ib03DBWUhCdjO_Drl1Q-1748620139-1.2.1.1-NYUEvsx7qr4jUJk1XsG83X6nUcWSp.TWrACtvVUKnfLlnKoTugGUbOwqSTwRGM5XfOuftlp5LKSsh7VSQiGMLG34SMJeQZSPl4yygCz1qk9JoHM8iVoJaPzIsbSXzdiyN8DJN5CFqUvZFsCsmYdDo1SCN9uZsDHR8uHumohY0mHsP9rt59YF9jwEHh2fuwGKLZVadL7nvQou.b9mtD14yj7oCOl3RDh9KhZVBSerZUILEAbsvrKIcfEUl_hY0xTvHLkk99aA39byANMwWWF1qHHtWvU3gHPNpWOQ6jjbihz_7WxChvtY99p7wgQpd_vI.wBI6ixKCNQTAKdoHDjtSnqGOHOHSw9qNBZHmmdAlZo; access_token_web=eyJraWQiOiJFNTdZZHJ1SHBsQWp1MmNObzFEb3JIM2oyN0J1NS1zX09QNVB3UGlobjVNIiwiYWxnIjoiUFMyNTYifQ.eyJhcHBfaWQiOjQsImNsaWVudF9pZCI6IndlYiIsImF1ZCI6ImZyLmNvcmUuYXBpIiwiaXNzIjoidmludGVkLWlhbS1zZXJ2aWNlIiwic3ViIjoiMjcyMzA0ODU1IiwiaWF0IjoxNzQ4NjIwMTQwLCJzaWQiOiJlOTI0OGJlZS0xNzQ4NDI1ODY0Iiwic2NvcGUiOiJwdWJsaWMgdXNlciIsImV4cCI6MTc0ODYyNzM0MCwicHVycG9zZSI6ImFjY2VzcyIsImxvZ2luX3R5cGUiOjMsImFjdCI6eyJzdWIiOiIyNzIzMDQ4NTUifSwiYWNjb3VudF9pZCI6MjEzMjgyMTYxfQ.GzveY1AmtaElrnF94J6fSDeZ0ecBpjv4nLPR2YKtiMy4hcYE1k-QNV3AugyDxXWST4rrVgPsoixcHGpiea0uf7aHr-buj_eFzN75-MNuHljOzEiWFDfzVUN_kjDMvNswER547sqN9oMWsGzvIoAPNE3Z_QgyFChFqRekUhtDGE33PySMLHF9ZoxqzEapXzIZIzhYDJBWaiHqeG_RAd5j1hiz1qTuPCi342J90RcuwBUrE6QbwqUWlN4HSwSOyMf_m38WB17jRf99qJ2TzqauNsocWxQw7HH83lEmrDe-duWCkg9Aww_YpaHYmxywTe9EZISoqRTcL_V5d7GkEBe5tQ; refresh_token_web=eyJraWQiOiJFNTdZZHJ1SHBsQWp1MmNObzFEb3JIM2oyN0J1NS1zX09QNVB3UGlobjVNIiwiYWxnIjoiUFMyNTYifQ.eyJhcHBfaWQiOjQsImNsaWVudF9pZCI6IndlYiIsImF1ZCI6ImZyLmNvcmUuYXBpIiwiaXNzIjoidmludGVkLWlhbS1zZXJ2aWNlIiwic3ViIjoiMjcyMzA0ODU1IiwiaWF0IjoxNzQ4NjIwMTQwLCJzaWQiOiJlOTI0OGJlZS0xNzQ4NDI1ODY0Iiwic2NvcGUiOiJwdWJsaWMgdXNlciIsImV4cCI6MTc0OTIyNDk0MCwicHVycG9zZSI6InJlZnJlc2giLCJsb2dpbl90eXBlIjozLCJhY3QiOnsic3ViIjoiMjcyMzA0ODU1In0sImFjY291bnRfaWQiOjIxMzI4MjE2MX0.YktzTXQOuAQPCfAVZz-mLLnBo-kL_698rJQONEaeywoGOGQ7StxHDMMB-8N2qCIww7KOQb7ovo8dMmlLOR-ZKgiZGHTqOhqqioOvT3E6t9lblW3lbr0EQBh4vMQjdjSQ6ORX1PdJ6g-MFjS-qM53bu6Da_BZbUo9MXtqUwrMHzqej5H4wiPKXcTQvgWgBrzDPdOfLt3BJ2kFfqKKp2qJoRJXp4FUPF8yFsdAUKCQxcRCCZPHhtuYpCCa75mFfCOn0UgPxDLU-IJ2J4fzniXCygZUK9_ekh-pkZtYt35QfNWtjwPqJnMicItjpUfev3fML9wm0NVYeJI43fxnrSJSUg; _dd_s=rum=0&expire=1748620913918; OptanonConsent=isGpcEnabled=0&datestamp=Fri+May+30+2025+20%3A46%3A57+GMT%2B0500+(Pakistan+Standard+Time)&version=202312.1.0&browserGpcFlag=0&isIABGlobal=false&hosts=&consentId=272304855&interactionCount=3&landingPath=NotLandingPage&groups=C0001%3A1%2CC0002%3A1%2CC0003%3A1%2CC0004%3A1%2CC0005%3A1%2CV2STACK42%3A1%2CC0015%3A1%2CC0035%3A1%2CC0038%3A1&genVendors=V2%3A1%2CV1%3A1%2C&geolocation=PK%3BPB&AwaitingReconsent=false; _vinted_fr_session=UnJ3Q2Vsc1pKaDVBVEE5N2E0U1ZSU0Q4cXY3Vld5Lys2b1NJcSt2NVc5c3hlQzhFMXdseTkwdDZoQUtCOUtMNkVMVndpUVFwZEg5SXRucDdGRlNPd0k2Tll2VlFDeWE0UTVvN0c5RGVHWVlwcjFsZHJIdndxM0RjSmpqMFVBbjBybkk0RXN0ZVFZT01JZ1NmMUxiOTlqNThwNUROQTZjdGVZYXFKdkNMakVtTlhEQlMzY0pnMVFFd3pycGlXbjJlbGVJNHlPdnU3VVQySHRmd2RJZGRmWjNOVWFXUlgzL2lmcHRzSXNLWXc0a3gyVSt3WVJvMWRGK3Q1RU1vczB0SU5xUUMrR2ZMUTJsY0RLay9HRllnZERWM0RnZGZFd1ROTGlKY1VSSzFUNzV2TUdDT0ZWTjZDbEVZLy9MYzhHVDdPZmJ4aEowbjQ2ZGIrbmd2QkZCTTVscU5EUSt1THRsUU9jaGNYL3ZpWnRzZ1IzUFprUHZzcnUzcitTUWZSUG90NEJOckFzM1p0MGtSQ09kZlZ4ZmErM3hnSGF0VVlQMVcwbXZ5ZmRuZWQvUUpXNDdsS1VHSldWcm5yUUJyYlJyTkIxWE1IN3hVK0dBU0dsS0FuVU81eThqSHVuMzFxT2VCdEt3elowdkl6cGdWaFVZWEhJTE14dCtFRWlMWlJCZ2o5MXdtdWpWUUR1UC9LQ3hGeE1PSWkzdS9lZ3owTU1EVE5JNmtvekpOb25UdnRyUWdwcG01WGpNOTgzU254eExJMjA1YlZyZHR0UU0yV1Yvang0VGxBOFNrTWUzNDhxU0VxckJmSVFRY0JwVGVudnAyZVFSRzZ1SU5QZnRaTitqRy0tOUhiVlJLQWpCWk9QTGQyVzg2aWlQQT09--5fc6397dfd9cc8c5a6e98f2acf3cd40bb36d777b; banners_ui_state=PENDING"}, // Get cookie from env
		"priority":           {"u=1, i"},
		"referer":            {"https://www.vinted.co.uk/catalog"},
		"sec-ch-ua":          {"\"Chromium\";v=\"136\", \"Microsoft Edge\";v=\"136\", \"Not.A/Brand\";v=\"99\""},
		"sec-ch-ua-mobile":   {"?0"},
		"sec-ch-ua-platform": {"\"Windows\""},
		"sec-fetch-dest":     {"empty"},
		"sec-fetch-mode":     {"cors"},
		"sec-fetch-site":     {"same-origin"},
		"user-agent":         {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36 Edg/136.0.0.0"},
		"x-anon-id":          {"a8c8ace7-ebf8-4144-9e8f-70a50132ec5a"},
		"x-csrf-token":       {"75f6c9fa-dc8e-4e52-a000-e09dd4084b3e"},
		"x-money-object":     {"true"},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response failed: %w", err)
	}

	var response VintedResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	return response.Items, nil
}

func findNewItems(items []VintedItem) []VintedItem {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	var newItems []VintedItem
	for _, item := range items {
		if !lastItemIDs[item.ID] {
			newItems = append(newItems, item)
		} else {
			// Items are ordered newest first, so we can stop at first known item
			break
		}
	}
	return newItems
}

func updateItemState(items []VintedItem) {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	// Reset state
	lastItemIDs = make(map[int64]bool)

	// Add current items to state
	for _, item := range items {
		lastItemIDs[item.ID] = true
	}
}

func sendItemsToDiscord(s *discordgo.Session, items []VintedItem) {
	for _, item := range items {
		embed := createItemEmbed(item)
		_, err := s.ChannelMessageSendEmbed(discordChannel, embed)
		if err != nil {
			log.Printf("Failed to send message to Discord: %v", err)
		}
		// Avoid rate limits
		time.Sleep(500 * time.Millisecond)
	}
}

func createItemEmbed(item VintedItem) *discordgo.MessageEmbed {
	price := fmt.Sprintf("%s %s", item.Price.Amount, item.Price.CurrencyCode)

	return &discordgo.MessageEmbed{
		Title:       item.Title,
		URL:         item.URL,
		Description: "New item listed on Vinted!",
		Color:       0x1FC598, // Vinted green
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: item.Photo.URL,
		},
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Price", Value: price, Inline: true},
			{Name: "Condition", Value: item.Status, Inline: true},
			{Name: "Size", Value: item.SizeTitle, Inline: true},
			{Name: "Seller", Value: item.User.Login, Inline: true},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Vinted Monitor",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
