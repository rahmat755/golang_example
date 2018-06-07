package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadFromUrl(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]+".zip"
	fmt.Println("Downloading", url, "to", fileName)

	if _, err := os.Stat(fileName); err != nil {
		output, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error while creating", fileName, "-", err)
			return
		}
		defer output.Close()

		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error while downloading", url, "-", err)
			return
		}
		defer response.Body.Close()
		n, err := io.Copy(output, response.Body)
		if err != nil {
			fmt.Println("Error while downloading", url, "-", err)
			return
		}

		fmt.Println(n, "bytes downloaded.")
	} else {
		fmt.Println("File already exists")
	}
}
func DownLoad(){
	var regions = map [string]string{
		"tat" : "131",
		"altay": "118",
		"amur": "167",
		"archangelsk":"171",
		"astrahanskaya_obl":"91",
		"belgorodskaya_obl":"170",
		"bryanskaya_obl":"104",
		"vladimirskaya_obl":"103",
		"volgogradskaya_obl":"90",
		"vologodskaya_obl":"169",
		"voronegskaya_obl":"102",
		"moscow":"101",
		"sank_peterburg":"125",
		"sevastopol":"168",
		"evreyskaya_avt_obl": "166",
		"zabaykalskiy_kray":"117",
		"ivanovskaya_obl":"163",
		"irkutskaya_obl":"162",
		"kabardino_balkanskaya_resp": "161",
		"kaliningradskaya_obl": "160",
		"kalugskaya_obl":"159",
		"kamchatskiy_kray": "165",
		"karachayevo_cherkesskaya_resp": "121",
		"kemerovskaya_obl":"158",
		"kirovskaya_obl": "157",
		"kostromskaya_obl": "156",
		"krasnodarskiy_kray": "89",
		"krasnoyarskiy_kray": "155",
		"kurganskaya_obl": "154",
		"kurskaya_obl": "153",
		"leningradskaya_obl": "124",
		"lipedskaya_obl": "152",
		"magadanskaya_obl":"151",
		"moskowskaya_obl":"150",
		"murmanskaya_obl": "149",
		"neneckiy_avt_okr":"148",
		"nignegorodskaya_obl": "147",
		"novgorodskaya_obl": "146",
		"novosibirskaya_obl": "145",
		"omskaya_obl": "144",
		"orenburgskaya_obl": "142",
		"orlovskaya_obl":"100",
		"penzenskaya_obl": "143",
		"permskiy_kray": "141",
		"primorskiy_kray":"87",
		"pskovskaya_obl": "123",
		"resp_adigea":  "137",
		"resp_altay":  "136",
		"resp_bashkortostan":  "140",
		"resp_buryatiya":  "116",
		"resp_dagestan":  "120",
		"resp_ingushetiya": "135",
		"resp_kalmikiya":"88" ,
		"resp_kareliya": "134",
		"resp_komi":  "122",
		"resp_krim": "133",
		"resp_mariy_el":  "139",
		"resp_mordoviya":  "138",
		"resp_saha(yakutiya)":  "164",
		"resp_severnaya_osetiya_alaniya":  "132",
		"resp_tiva": "115",
		"resp_hakasiya": "130",
		"rostovskaya_obl":"129",
		"ryazanskaya_obl": "99",
		"samarskaya_obl": "128",
		"saratovskaya_obl": "127",
		"sahalinskaya_obl": "126",
		"sverdlovskaya_obl": "110",
		"smolenskaya_obl":  "98",
		"stavropolskiy_kray": "119",
		"tambovskaya_obl":"114",
		"tverskaya_obl": "97",
		"tomskaya_obl":"113",
		"tulskaya_obl":"96",
		"tumenskaya_obl":"109",
		"udmurtskaya_resp": "112",
		"ulyanovskaya_obl":"111",
		"habarovskiy_kray":"105",
		"hanti_mansiyskiy_avt_okr": "108",
		"chelyabinskaya_obl":"107",
		"chechenskaya_resp":"94",
		"chuvashskaya_resp": "93",
		"chukotskiy_avt_okr":"92",
		"yamalo_neneckiy_avt_okr":"106",
		"yaroslavskaya_obl": "95",}
	var uprorgs = map [string]string{
		"tat": "46",
		"altay" : "33",
		"amur": "82",
		"archangelsk" : "86",
		"astrahanskaya_obl": "6",
		"belgorodskaya_obl": "85",
		"bryanskaya_obl": "19",
		"vladimirskaya_obl": "18",
		"volgogradskaya_obl": "5",
		"vologodskaya_obl": "84",
		"voronegskaya_obl": "17",
		"moscow": "16",
		"sank_peterburg": "40",
		"sevastopol": "83",
		"evreyskaya_avt_obl": "81",
		"zabaykalskiy_kray": "32",
		"ivanovskaya_obl": "78",
		"irkutskaya_obl": "77",
		"kabardino_balkanskaya_resp": "76",
		"kaliningradskaya_obl": "75",
		"kalugskaya_obl": "74",
		"kamchatskiy_kray": "80",
		"karachayevo_cherkesskaya_resp": "36",
		"kemerovskaya_obl": "73",
		"kirovskaya_obl": "72",
		"kostromskaya_obl": "71",
		"krasnodarskiy_kray": "4",
		"krasnoyarskiy_kray": "70",
		"kurganskaya_obl": "69",
		"kurskaya_obl": "68",
		"leningradskaya_obl": "39",
		"lipedskaya_obl": "67",
		"magadanskaya_obl": "66",
		"moskowskaya_obl": "65",
		"murmanskaya_obl": "64",
		"neneckiy_avt_okr": "63",
		"nignegorodskaya_obl": "62",
		"novgorodskaya_obl": "61",
		"novosibirskaya_obl": "60",
		"omskaya_obl": "59",
		"orenburgskaya_obl": "57",
		"orlovskaya_obl": "15",
		"penzenskaya_obl": "58",
		"permskiy_kray": "56",
		"primorskiy_kray": "2",
		"pskovskaya_obl": "38",
		"resp_adigea": "52",
		"resp_altay": "51",
		"resp_bashkortostan": "55",
		"resp_buryatiya": "31",
		"resp_dagestan": "35",
		"resp_ingushetiya": "50",
		"resp_kalmikiya": "3",
		"resp_kareliya": "49",
		"resp_komi": "37",
		"resp_krim": "48",
		"resp_mariy_el": "54",
		"resp_mordoviya": "53",
		"resp_saha(yakutiya)": "79",
		"resp_severnaya_osetiya_alaniya": "47",
		"resp_tiva": "30",
		"resp_hakasiya": "45",
		"rostovskaya_obl": "44",
		"ryazanskaya_obl": "14",
		"samarskaya_obl": "43",
		"saratovskaya_obl": "42",
		"sahalinskaya_obl": "41",
		"sverdlovskaya_obl": "25",
		"smolenskaya_obl": "13",
		"stavropolskiy_kray": "34",
		"tambovskaya_obl": "29",
		"tverskaya_obl": "12",
		"tomskaya_obl": "28",
		"tulskaya_obl": "11",
		"tumenskaya_obl": "24",
		"udmurtskaya_resp": "27",
		"ulyanovskaya_obl": "26",
		"habarovskiy_kray": "20",
		"hanti_mansiyskiy_avt_okr": "23",
		"chelyabinskaya_obl": "22",
		"chechenskaya_resp": "9",
		"chuvashskaya_resp": "8",
		"chukotskiy_avt_okr": "7",
		"yamalo_neneckiy_avt_okr": "21",
		"yaroslavskaya_obl": "10",
	}

	for _,value := range regions {
		url := "https://www.reformagkh.ru/opendata/export/" + (value)
		DownloadFromUrl(url)
	}
	fmt.Println("Ended loading house data: \r\n")
	fmt.Println("----")
	fmt.Println("----")
	fmt.Println("----")
	fmt.Println("starting loading uprorgs data\r\n")
	for _,val:=range uprorgs{
		url:= "https://www.reformagkh.ru/opendata/export/" + (val)
		DownloadFromUrl(url)
	}
	println("Done!")
}
func main() {
	DownLoad()
}