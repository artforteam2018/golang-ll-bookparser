package requestsession

import "gopkg.in/h2non/gentleman.v2"

func BasicHeaders() map[string]interface{} {
	return map[string]interface{}{
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"accept-encoding":           "gzip, deflate, br, zstd",
		"accept-language":           "en-US,en;q=0.9,ru;q=0.8",
		"cache-control":             "max-age=0",
		"connection":                "keep-alive",
		"sec-ch-ua":                 "\"Google Chrome\";v=\"125\", \"Chromium\";v=\"125\", \"Not.A/Brand\";v=\"24\"",
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        "\"macOS\"",
		"sec-fetch-dest":            "document",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-site":            "none",
		"sec-fetch-user":            "?1",
		"upgrade-insecure-requests": "1",
		"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36",
		// "Cookie":                    "LiveLibId=bceh14h6t601773oec73pp8c7f; __ll_tum=2036088225; __ll_ab_mp=1; __utnt=g0_y0_a15721_u0_c0; __ll_unreg_session=bceh14h6t601773oec73pp8c7f; __ll_unreg_sessions_count=1; tmr_lvid=021298c5d1e983db7bc17cd0ac4cdcbd; tmr_lvidTS=1716715737558; _ym_uid=1716715738222215353; _ym_d=1716715738; _gid=GA1.2.630604727.1716715738; _ym_isad=2; domain_sid=ypvTwlZXQGkJPXKPHC0sY%3A1716715737978; __ll_fv=1716715740; __ll_dvs=5; ll_asid=1784656883; promoLLid=npcjb8nj26kts9ml5s307ml3g7; __llutmz=-180; __llutmf=0; __ll_popup_count_pviews=mailc1_; showed_vkid_onetap=1; __ll_unreg_r=60; __popupmail_showed=1000; __popupmail_showed_uc=1; __popupmail_showed_t=1000; iwatchyou=a182b75c63dae3147b052138a0a7a26a; __ll_popup_showed=mail_kv24unreg6_; __ll_popup_last_show=1716753907; __ll_popup_count_shows=mailc1_kv24unreg6c1_; __gr=g426c2_g1240c2_g1121c2_g433c1_g1217c1_g1136c1_g1133c1_g91c3_g1c3_; _ga_90RPM8SDHL=GS1.1.1716753869.3.1.1716754030.13.0.1147395951; _ga=GA1.2.955158677.1716715737; __ll_cp=8; __r=pc4869c2_pc4870c2_pc4862c1_pc4847c1_; tmr_detect=0%7C1716754032522; __ll_dv=1716754090",
	}
}

func setHeaders(req *gentleman.Request) {
	headers := BasicHeaders()

	for key, val := range headers {
		req.SetHeader(key, val.(string))
	}
}

func setHeadersForFullText(req *gentleman.Request, urlReferer string) {
	headers := map[string]string{
		"Accept":             "*/*",
		"Accept-Encoding":    "gzip, deflate, br, zstd",
		"Accept-Language":    "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7",
		"Connection":         "keep-alive",
		"Content-Length":     "62",
		"Content-Type":       "application/x-www-form-urlencoded; charset=UTF-8",
		"Host":               "www.livelib.ru",
		"Origin":             "https://www.livelib.ru",
		"Referer":            urlReferer,
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-origin",
		"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36",
		"X-Requested-With":   "XMLHttpRequest",
		"sec-ch-ua":          `"Google Chrome";v="125", "Chromium";v="125", "Not.A/Brand";v="24"`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": `"macOS"`,
	}

	for key, val := range headers {
		req.SetHeader(key, val)
	}
}
