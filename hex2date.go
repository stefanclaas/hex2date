package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// A function that prints all time zones for a given region
func printTimeZones(region string, seconds int64) {
	// A map of regions to time zones
	timeZones := map[string][]string{
	"all":		 {"Africa/Abidjan", "Africa/Accra", "Africa/Addis_Ababa", "Africa/Algiers",
                          "Africa/Asmara", "Africa/Asmera", "Africa/Bamako", "Africa/Bangui",
			  "Africa/Banjul", "Africa/Bissau", "Africa/Blantyre", "Africa/Brazzaville",
			  "Africa/Bujumbura", "Africa/Cairo", "Africa/Casablanca", "Africa/Ceuta",
			  "Africa/Conakry", "Africa/Dakar", "Africa/Dar_es_Salaam", "Africa/Djibouti",
			  "Africa/Douala", "Africa/El_Aaiun", "Africa/Freetown", "Africa/Gaborone",
			  "Africa/Harare", "Africa/Johannesburg", "Africa/Juba", "Africa/Kampala",
			  "Africa/Khartoum", "Africa/Kigali", "Africa/Kinshasa", "Africa/Lagos",
			  "Africa/Libreville", "Africa/Lome", "Africa/Luanda", "Africa/Lubumbashi",
			  "Africa/Lusaka", "Africa/Malabo", "Africa/Maputo", "Africa/Maseru", "Africa/Mbabane",
			  "Africa/Mogadishu", "Africa/Monrovia", "Africa/Nairobi", "Africa/Ndjamena",
			  "Africa/Niamey", "Africa/Nouakchott", "Africa/Ouagadougou", "Africa/Porto-Novo",
			  "Africa/Sao_Tome", "Africa/Timbuktu", "Africa/Tripoli", "Africa/Tunis",
			  "Africa/Windhoek",
			  "America/Adak", "America/Anchorage", "America/Anguilla", "America/Antigua",
			  "America/Araguaina", "America/Argentina/Buenos_Aires", "America/Argentina/Catamarca",
			  "America/Argentina/ComodRivadavia", "America/Argentina/Cordoba", "America/Argentina/Jujuy",
			  "America/Argentina/La_Rioja", "America/Argentina/Mendoza", "America/Argentina/Rio_Gallegos",
			  "America/Argentina/Salta", "America/Argentina/San_Juan", "America/Argentina/San_Luis",
			  "America/Argentina/Tucuman", "America/Argentina/Ushuaia", "America/Aruba", "America/Asuncion",
			  "America/Atikokan", "America/Atka", "America/Bahia", "America/Bahia_Banderas", "America/Barbados",
			  "America/Belem", "America/Belize", "America/Blanc-Sablon", "America/Boa_Vista", "America/Bogota",
			  "America/Boise", "America/Buenos_Aires", "America/Cambridge_Bay", "America/Campo_Grande",
			  "America/Cancun", "America/Caracas", "America/Catamarca", "America/Cayenne", "America/Cayman",
			  "America/Chicago", "America/Chihuahua", "America/Ciudad_Juarez", "America/Coral_Harbour",
			  "America/Cordoba", "America/Costa_Rica", "America/Creston", "America/Cuiaba", "America/Curacao",
			  "America/Danmarkshavn", "America/Dawson", "America/Dawson_Creek", "America/Denver", "America/Detroit",
			  "America/Dominica", "America/Edmonton", "America/Eirunepe", "America/El_Salvador", "America/Ensenada",
			  "America/Fort_Nelson", "America/Fort_Wayne", "America/Fortaleza", "America/Glace_Bay", "America/Godthab",
			  "America/Goose_Bay", "America/Grand_Turk", "America/Grenada", "America/Guadeloupe", "America/Guatemala",
			  "America/Guayaquil", "America/Guyana", "America/Halifax", "America/Havana", "America/Hermosillo",
			  "America/Indiana/Indianapolis", "America/Indiana/Knox", "America/Indiana/Marengo", "America/Indiana/Petersburg",
			  "America/Indiana/Tell_City", "America/Indiana/Vevay", "America/Indiana/Vincennes", "America/Indiana/Winamac",
			  "America/Indianapolis", "America/Inuvik", "America/Iqaluit", "America/Jamaica", "America/Jujuy",
			  "America/Juneau", "America/Kentucky/Louisville", "America/Kentucky/Monticello", "America/Knox_IN",
			  "America/Kralendijk", "America/La_Paz", "America/Lima", "America/Los_Angeles", "America/Louisville",
			  "America/Lower_Princes", "America/Maceio", "America/Managua", "America/Manaus", "America/Marigot",
			  "America/Martinique", "America/Matamoros", "America/Mazatlan", "America/Mendoza", "America/Menominee",
			  "America/Merida", "America/Metlakatla", "America/Mexico_City", "America/Miquelon", "America/Moncton",
			  "America/Monterrey", "America/Montevideo", "America/Montreal", "America/Montserrat", "America/Nassau",
			  "America/New_York", "America/Nipigon", "America/Nome", "America/Noronha", "America/North_Dakota/Beulah",
			  "America/North_Dakota/Center", "America/North_Dakota/New_Salem", "America/Nuuk", "America/Ojinaga",
			  "America/Panama", "America/Pangnirtung", "America/Paramaribo", "America/Phoenix", "America/Port-au-Prince",
			  "America/Port_of_Spain", "America/Porto_Acre", "America/Porto_Velho", "America/Puerto_Rico",
			  "America/Punta_Arenas", "America/Rainy_River", "America/Rankin_Inlet", "America/Recife", "America/Regina",
			  "America/Resolute", "America/Rio_Branco", "America/Rosario", "America/Santa_Isabel", "America/Santarem",
			  "America/Santiago", "America/Santo_Domingo", "America/Sao_Paulo", "America/Scoresbysund", "America/Shiprock",
			  "America/Sitka", "America/St_Barthelemy", "America/St_Johns", "America/St_Kitts", "America/St_Lucia",
			  "America/St_Thomas", "America/St_Vincent", "America/Swift_Current", "America/Tegucigalpa", "America/Thule",
			  "America/Thunder_Bay", "America/Tijuana", "America/Toronto", "America/Tortola", "America/Vancouver",
			  "America/Virgin", "America/Whitehorse", "America/Winnipeg", "America/Yakutat", "America/Yellowknife",
			  "Antarctica/Casey", "Antarctica/Davis", "Antarctica/DumontDUrville", "Antarctica/Macquarie", "Antarctica/Mawson",
			  "Antarctica/McMurdo", "Antarctica/Palmer", "Antarctica/Rothera", "Antarctica/South_Pole", "Antarctica/Syowa",
		          "Antarctica/Troll", "Antarctica/Vostok",
			  "Arctic/Longyearbyen",
			  "Asia/Aden", "Asia/Almaty", "Asia/Amman", "Asia/Anadyr", "Asia/Aqtau", "Asia/Aqtobe", "Asia/Ashgabat",
			  "Asia/Ashkhabad", "Asia/Atyrau", "Asia/Baghdad", "Asia/Bahrain", "Asia/Baku", "Asia/Bangkok", "Asia/Barnaul",
			  "Asia/Beirut", "Asia/Bishkek", "Asia/Brunei", "Asia/Calcutta", "Asia/Chita", "Asia/Choibalsan", "Asia/Chongqing",
			  "Asia/Chungking", "Asia/Colombo", "Asia/Dacca", "Asia/Damascus", "Asia/Dhaka", "Asia/Dili", "Asia/Dubai",
			  "Asia/Dushanbe", "Asia/Famagusta", "Asia/Gaza", "Asia/Harbin", "Asia/Hebron", "Asia/Ho_Chi_Minh", "Asia/Hong_Kong",
			  "Asia/Hovd", "Asia/Irkutsk", "Asia/Istanbul", "Asia/Jakarta", "Asia/Jayapura", "Asia/Jerusalem", "Asia/Kabul",
			  "Asia/Kamchatka", "Asia/Karachi", "Asia/Kashgar", "Asia/Kathmandu", "Asia/Katmandu", "Asia/Khandyga", "Asia/Kolkata",
			  "Asia/Krasnoyarsk", "Asia/Kuala_Lumpur", "Asia/Kuching", "Asia/Kuwait", "Asia/Macao", "Asia/Macau", "Asia/Magadan",
			  "Asia/Makassar", "Asia/Manila", "Asia/Muscat", "Asia/Nicosia", "Asia/Novokuznetsk", "Asia/Novosibirsk", "Asia/Omsk",
			  "Asia/Oral", "Asia/Phnom_Penh", "Asia/Pontianak", "Asia/Pyongyang", "Asia/Qatar", "Asia/Qostanay", "Asia/Qyzylorda",
			  "Asia/Rangoon", "Asia/Riyadh", "Asia/Saigon", "Asia/Sakhalin", "Asia/Samarkand", "Asia/Seoul", "Asia/Shanghai",
			  "Asia/Singapore", "Asia/Srednekolymsk", "Asia/Taipei", "Asia/Tashkent", "Asia/Tbilisi", "Asia/Tehran", "Asia/Tel_Aviv",
			  "Asia/Thimbu", "Asia/Thimphu", "Asia/Tokyo", "Asia/Tomsk", "Asia/Ujung_Pandang", "Asia/Ulaanbaatar", "Asia/Ulan_Bator",
			  "Asia/Urumqi", "Asia/Ust-Nera", "Asia/Vientiane", "Asia/Vladivostok", "Asia/Yakutsk", "Asia/Yangon", "Asia/Yekaterinburg",
			  "Asia/Yerevan",
			  "Atlantic/Azores", "Atlantic/Bermuda", "Atlantic/Canary", "Atlantic/Cape_Verde", "Atlantic/Faeroe", "Atlantic/Faroe",
			  "Atlantic/Jan_Mayen", "Atlantic/Madeira", "Atlantic/Reykjavik", "Atlantic/South_Georgia", "Atlantic/St_Helena",
			  "Atlantic/Stanley",
			  "Australia/ACT", "Australia/Adelaide", "Australia/Brisbane", "Australia/Broken_Hill", "Australia/Canberra", "Australia/Currie",
			  "Australia/Darwin", "Australia/Eucla", "Australia/Hobart", "Australia/LHI", "Australia/Lindeman", "Australia/Lord_Howe",
			  "Australia/Melbourne", "Australia/NSW", "Australia/North", "Australia/Perth", "Australia/Queensland", "Australia/South",
			  "Australia/Sydney", "Australia/Tasmania", "Australia/Victoria", "Australia/West", "Australia/Yancowinna",
			  "Brazil/Acre", "Brazil/DeNoronha", "Brazil/East", "Brazil/West",
			  "CET",
			  "CST6CDT",
			  "Canada/Atlantic", "Canada/Central", "Canada/Eastern", "Canada/Mountain", "Canada/Newfoundland", "Canada/Pacific",
			  "Canada/Saskatchewan", "Canada/Yukon",
			  "Chile/Continental", "Chile/EasterIsland",
		  	  "Cuba",
			  "EET",
			  "EST",
			  "EST5EDT",
			  "Egypt",
			  "Eire",
			  "Etc/GMT", "Etc/GMT+0", "Etc/GMT+1", "Etc/GMT+2", "Etc/GMT+3", "Etc/GMT+4", "Etc/GMT+5", "Etc/GMT+6", "Etc/GMT+7",
			  "Etc/GMT+8", "Etc/GMT+9", "Etc/GMT+10", "Etc/GMT+11", "Etc/GMT+12", "Etc/GMT-0", "Etc/GMT-1", "Etc/GMT-2", "Etc/GMT-3",
			  "Etc/GMT-4", "Etc/GMT-5", "Etc/GMT-6", "Etc/GMT-7", "Etc/GMT-8", "Etc/GMT-9", "Etc/GMT-10", "Etc/GMT-11", "Etc/GMT-12",
			  "Etc/GMT-13", "Etc/GMT-14", "Etc/GMT0", "Etc/Greenwich", "Etc/UCT", "Etc/UTC", "Etc/Universal", "Etc/Zulu",
			  "Europe/Amsterdam", "Europe/Andorra", "Europe/Astrakhan", "Europe/Athens", "Europe/Belfast", "Europe/Belgrade",
			  "Europe/Berlin", "Europe/Bratislava", "Europe/Brussels", "Europe/Bucharest", "Europe/Budapest", "Europe/Busingen",
			  "Europe/Chisinau", "Europe/Copenhagen", "Europe/Dublin", "Europe/Gibraltar", "Europe/Guernsey", "Europe/Helsinki",
			  "Europe/Isle_of_Man", "Europe/Istanbul", "Europe/Jersey", "Europe/Kaliningrad", "Europe/Kiev", "Europe/Kirov",
			  "Europe/Kyiv", "Europe/Lisbon", "Europe/Ljubljana", "Europe/London", "Europe/Luxembourg", "Europe/Madrid", "Europe/Malta",
			  "Europe/Mariehamn", "Europe/Minsk", "Europe/Monaco", "Europe/Moscow", "Europe/Nicosia", "Europe/Oslo", "Europe/Paris",
			  "Europe/Podgorica", "Europe/Prague", "Europe/Riga", "Europe/Rome", "Europe/Samara", "Europe/San_Marino", "Europe/Sarajevo",
			  "Europe/Saratov", "Europe/Simferopol", "Europe/Skopje", "Europe/Sofia", "Europe/Stockholm", "Europe/Tallinn", "Europe/Tirane",
			  "Europe/Tiraspol", "Europe/Ulyanovsk", "Europe/Uzhgorod", "Europe/Vaduz", "Europe/Vatican", "Europe/Vienna", "Europe/Vilnius",
			  "Europe/Volgograd", "Europe/Warsaw", "Europe/Zagreb", "Europe/Zaporozhye", "Europe/Zurich",
			  "Factory",
			  "GB",
			  "GB-Eire",
			  "GMT",
			  "GMT+0",
			  "GMT-0",
			  "GMT0",
			  "Greenwich",
			  "HST",
			  "Hongkong",
			  "Iceland",
			  "Indian/Antananarivo", "Indian/Chagos", "Indian/Christmas", "Indian/Cocos", "Indian/Comoro", "Indian/Kerguelen", "Indian/Mahe",
			  "Indian/Maldives", "Indian/Mauritius", "Indian/Mayotte", "Indian/Reunion",
			  "Iran",
			  "Israel",
			  "Jamaica",
			  "Japan",
			  "Kwajalein",
			  "Libya",
			  "MET",
			  "MST",
			  "MST7MDT",
			  "Mexico/BajaNorte", "Mexico/BajaSur", "Mexico/General",
			  "NZ",
			  "NZ-CHAT",
			  "Navajo",
			  "PRC",
			  "PST8PDT",
			  "Pacific/Apia", "Pacific/Auckland", "Pacific/Bougainville", "Pacific/Chatham", "Pacific/Chuuk", "Pacific/Easter",
			  "Pacific/Efate", "Pacific/Enderbury", "Pacific/Fakaofo", "Pacific/Fiji", "Pacific/Funafuti", "Pacific/Galapagos",
			  "Pacific/Gambier", "Pacific/Guadalcanal", "Pacific/Guam", "Pacific/Honolulu", "Pacific/Johnston", "Pacific/Kanton",
			  "Pacific/Kiritimati", "Pacific/Kosrae", "Pacific/Kwajalein", "Pacific/Majuro", "Pacific/Marquesas", "Pacific/Midway",
			  "Pacific/Nauru", "Pacific/Niue", "Pacific/Norfolk", "Pacific/Noumea", "Pacific/Pago_Pago", "Pacific/Palau",
			  "Pacific/Pitcairn", "Pacific/Pohnpei", "Pacific/Ponape", "Pacific/Port_Moresby", "Pacific/Rarotonga", "Pacific/Saipan",
			  "Pacific/Samoa", "Pacific/Tahiti", "Pacific/Tarawa", "Pacific/Tongatapu", "Pacific/Truk", "Pacific/Wake", "Pacific/Wallis",
			  "Pacific/Yap",
			  "Poland",
			  "Portugal",
			  "ROC",
			  "ROK",
			  "Singapore",
			  "Turkey",
			  "UCT",
			  "US/Alaska", "US/Aleutian", "US/Arizona", "US/Central", "US/East-Indiana", "US/Eastern", "US/Hawaii", "US/Indiana-Starke",
			  "US/Michigan", "US/Mountain", "US/Pacific", "US/Samoa",
			  "UTC",
			  "Universal",
			  "W-SU",
			  "WET",
			  "Zulu"},
	"Africa": 	 {"Africa/Abidjan", "Africa/Accra", "Africa/Addis_Ababa", "Africa/Algiers",
                          "Africa/Asmara", "Africa/Asmera", "Africa/Bamako", "Africa/Bangui", "Africa/Bangui",
			  "Africa/Banjul", "Africa/Bissau", "Africa/Blantyre", "Africa/Brazzaville",
			  "Africa/Bujumbura", "Africa/Cairo", "Africa/Casablanca", "Africa/Ceuta",
			  "Africa/Conakry", "Africa/Dakar", "Africa/Dar_es_Salaam", "Africa/Djibouti",
			  "Africa/Douala", "Africa/El_Aaiun", "Africa/Freetown", "Africa/Gaborone",
			  "Africa/Harare", "Africa/Johannesburg", "Africa/Juba", "Africa/Kampala",
			  "Africa/Khartoum", "Africa/Kigali", "Africa/Kinshasa", "Africa/Lagos",
			  "Africa/Libreville", "Africa/Lome", "Africa/Luanda", "Africa/Lubumbashi",
			  "Africa/Lusaka", "Africa/Malabo", "Africa/Maputo", "Africa/Maseru", "Africa/Mbabane",
			  "Africa/Mogadishu", "Africa/Monrovia", "Africa/Nairobi", "Africa/Ndjamena",
			  "Africa/Niamey", "Africa/Nouakchott", "Africa/Ouagadougou", "Africa/Porto-Novo",
			  "Africa/Sao_Tome", "Africa/Timbuktu", "Africa/Tripoli", "Africa/Tunis",
			  "Africa/Windhoek"},
	"America":	 {"America/Adak", "America/Anchorage", "America/Anguilla", "America/Antigua",
			  "America/Araguaina", "America/Argentina/Buenos_Aires", "America/Argentina/Catamarca",
			  "America/Argentina/ComodRivadavia", "America/Argentina/Cordoba", "America/Argentina/Jujuy",
			  "America/Argentina/La_Rioja", "America/Argentina/Mendoza", "America/Argentina/Rio_Gallegos",
			  "America/Argentina/Salta", "America/Argentina/San_Juan", "America/Argentina/San_Luis",
			  "America/Argentina/Tucuman", "America/Argentina/Ushuaia", "America/Aruba", "America/Asuncion",
			  "America/Atikokan", "America/Atka", "America/Bahia", "America/Bahia_Banderas", "America/Barbados",
			  "America/Belem", "America/Belize", "America/Blanc-Sablon", "America/Boa_Vista", "America/Bogota",
			  "America/Boise", "America/Buenos_Aires", "America/Cambridge_Bay", "America/Campo_Grande",
			  "America/Cancun", "America/Caracas", "America/Catamarca", "America/Cayenne", "America/Cayman",
			  "America/Chicago", "America/Chihuahua", "America/Ciudad_Juarez", "America/Coral_Harbour",
			  "America/Cordoba", "America/Costa_Rica", "America/Creston", "America/Cuiaba", "America/Curacao",
			  "America/Danmarkshavn", "America/Dawson", "America/Dawson_Creek", "America/Denver", "America/Detroit",
			  "America/Dominica", "America/Edmonton", "America/Eirunepe", "America/El_Salvador", "America/Ensenada",
			  "America/Fort_Nelson", "America/Fort_Wayne", "America/Fortaleza", "America/Glace_Bay", "America/Godthab",
			  "America/Goose_Bay", "America/Grand_Turk", "America/Grenada", "America/Guadeloupe", "America/Guatemala",
			  "America/Guayaquil", "America/Guyana", "America/Halifax", "America/Havana", "America/Hermosillo",
			  "America/Indiana/Indianapolis", "America/Indiana/Knox", "America/Indiana/Marengo", "America/Indiana/Petersburg",
			  "America/Indiana/Tell_City", "America/Indiana/Vevay", "America/Indiana/Vincennes", "America/Indiana/Winamac",
			  "America/Indianapolis", "America/Inuvik", "America/Iqaluit", "America/Jamaica", "America/Jujuy",
			  "America/Juneau", "America/Kentucky/Louisville", "America/Kentucky/Monticello", "America/Knox_IN",
			  "America/Kralendijk", "America/La_Paz", "America/Lima", "America/Los_Angeles", "America/Louisville",
			  "America/Lower_Princes", "America/Maceio", "America/Managua", "America/Manaus", "America/Marigot",
			  "America/Martinique", "America/Matamoros", "America/Mazatlan", "America/Mendoza", "America/Menominee",
			  "America/Merida", "America/Metlakatla", "America/Mexico_City", "America/Miquelon", "America/Moncton",
			  "America/Monterrey", "America/Montevideo", "America/Montreal", "America/Montserrat", "America/Nassau",
			  "America/New_York", "America/Nipigon", "America/Nome", "America/Noronha", "America/North_Dakota/Beulah",
			  "America/North_Dakota/Center", "America/North_Dakota/New_Salem", "America/Nuuk", "America/Ojinaga",
			  "America/Panama", "America/Pangnirtung", "America/Paramaribo", "America/Phoenix", "America/Port-au-Prince",
			  "America/Port_of_Spain", "America/Porto_Acre", "America/Porto_Velho", "America/Puerto_Rico",
			  "America/Punta_Arenas", "America/Rainy_River", "America/Rankin_Inlet", "America/Recife", "America/Regina",
			  "America/Resolute", "America/Rio_Branco", "America/Rosario", "America/Santa_Isabel", "America/Santarem",
			  "America/Santiago", "America/Santo_Domingo", "America/Sao_Paulo", "America/Scoresbysund", "America/Shiprock",
			  "America/Sitka", "America/St_Barthelemy", "America/St_Johns", "America/St_Kitts", "America/St_Lucia",
			  "America/St_Thomas", "America/St_Vincent", "America/Swift_Current", "America/Tegucigalpa", "America/Thule",
			  "America/Thunder_Bay", "America/Tijuana", "America/Toronto", "America/Tortola", "America/Vancouver",
			  "America/Virgin", "America/Whitehorse", "America/Winnipeg", "America/Yakutat", "America/Yellowknife"},
	"Antarctica":	 {"Antarctica/Casey", "Antarctica/Davis", "Antarctica/DumontDUrville", "Antarctica/Macquarie", "Antarctica/Mawson",
			  "Antarctica/McMurdo", "Antarctica/Palmer", "Antarctica/Rothera", "Antarctica/South_Pole", "Antarctica/Syowa",
			  "Antarctica/Troll", "Antarctica/Vostok"},
	"Arctic":	 {"Arctic/Longyearbyen"},
	"Asia":		 {"Asia/Aden", "Asia/Almaty", "Asia/Amman", "Asia/Anadyr", "Asia/Aqtau", "Asia/Aqtobe", "Asia/Ashgabat",
			  "Asia/Ashkhabad", "Asia/Atyrau", "Asia/Baghdad", "Asia/Bahrain", "Asia/Baku", "Asia/Bangkok", "Asia/Barnaul",
			  "Asia/Beirut", "Asia/Bishkek", "Asia/Brunei", "Asia/Calcutta", "Asia/Chita", "Asia/Choibalsan", "Asia/Chongqing",
			  "Asia/Chungking", "Asia/Colombo", "Asia/Dacca", "Asia/Damascus", "Asia/Dhaka", "Asia/Dili", "Asia/Dubai",
			  "Asia/Dushanbe", "Asia/Famagusta", "Asia/Gaza", "Asia/Harbin", "Asia/Hebron", "Asia/Ho_Chi_Minh", "Asia/Hong_Kong",
			  "Asia/Hovd", "Asia/Irkutsk", "Asia/Istanbul", "Asia/Jakarta", "Asia/Jayapura", "Asia/Jerusalem", "Asia/Kabul",
			  "Asia/Kamchatka", "Asia/Karachi", "Asia/Kashgar", "Asia/Kathmandu", "Asia/Katmandu", "Asia/Khandyga", "Asia/Kolkata",
			  "Asia/Krasnoyarsk", "Asia/Kuala_Lumpur", "Asia/Kuching", "Asia/Kuwait", "Asia/Macao", "Asia/Macau", "Asia/Magadan",
			  "Asia/Makassar", "Asia/Manila", "Asia/Muscat", "Asia/Nicosia", "Asia/Novokuznetsk", "Asia/Novosibirsk", "Asia/Omsk",
			  "Asia/Oral", "Asia/Phnom_Penh", "Asia/Pontianak", "Asia/Pyongyang", "Asia/Qatar", "Asia/Qostanay", "Asia/Qyzylorda",
			  "Asia/Rangoon", "Asia/Riyadh", "Asia/Saigon", "Asia/Sakhalin", "Asia/Samarkand", "Asia/Seoul", "Asia/Shanghai",
			  "Asia/Singapore", "Asia/Srednekolymsk", "Asia/Taipei", "Asia/Tashkent", "Asia/Tbilisi", "Asia/Tehran", "Asia/Tel_Aviv",
			  "Asia/Thimbu", "Asia/Thimphu", "Asia/Tokyo", "Asia/Tomsk", "Asia/Ujung_Pandang", "Asia/Ulaanbaatar", "Asia/Ulan_Bator",
			  "Asia/Urumqi", "Asia/Ust-Nera", "Asia/Vientiane", "Asia/Vladivostok", "Asia/Yakutsk", "Asia/Yangon", "Asia/Yekaterinburg",
			  "Asia/Yerevan"},
	"Atlantic":	 {"Atlantic/Azores", "Atlantic/Bermuda", "Atlantic/Canary", "Atlantic/Cape_Verde", "Atlantic/Faeroe", "Atlantic/Faroe",
			  "Atlantic/Jan_Mayen", "Atlantic/Madeira", "Atlantic/Reykjavik", "Atlantic/South_Georgia", "Atlantic/St_Helena",
			  "Atlantic/Stanley"},
	"Australia":	 {"Australia/ACT", "Australia/Adelaide", "Australia/Brisbane", "Australia/Broken_Hill", "Australia/Canberra", "Australia/Currie",
			  "Australia/Darwin", "Australia/Eucla", "Australia/Hobart", "Australia/LHI", "Australia/Lindeman", "Australia/Lord_Howe",
			  "Australia/Melbourne", "Australia/NSW", "Australia/North", "Australia/Perth", "Australia/Queensland", "Australia/South",
			  "Australia/Sydney", "Australia/Tasmania", "Australia/Victoria", "Australia/West", "Australia/Yancowinna"},
 	"Brazil":	 {"Brazil/Acre", "Brazil/DeNoronha", "Brazil/East", "Brazil/West"},
	"CET":		 {"CET"},
	"CST6CDT":	 {"CST6CDT"},
	"Canada":	 {"Canada/Atlantic", "Canada/Central", "Canada/Eastern", "Canada/Mountain", "Canada/Newfoundland", "Canada/Pacific",
			  "Canada/Saskatchewan", "Canada/Yukon"},
	"Chile":	 {"Chile/Continental", "Chile/EasterIsland"},
	"Cuba":		 {"Cuba"},
	"EET":		 {"EET"},
	"EST":		 {"EST"},
	"EST5EDT":	 {"EST5EDT"},
	"Egypt":	 {"Egypt"},
	"Eire":	 	 {"Eire"},
	"Etc":		 {"Etc/GMT", "Etc/GMT+0", "Etc/GMT+1", "Etc/GMT+2", "Etc/GMT+3", "Etc/GMT+4", "Etc/GMT+5", "Etc/GMT+6", "Etc/GMT+7",
			  "Etc/GMT+8", "Etc/GMT+9", "Etc/GMT+10", "Etc/GMT+11", "Etc/GMT+12", "Etc/GMT-0", "Etc/GMT-1", "Etc/GMT-2", "Etc/GMT-3",
			  "Etc/GMT-4", "Etc/GMT-5", "Etc/GMT-6", "Etc/GMT-7", "Etc/GMT-8", "Etc/GMT-9", "Etc/GMT-10", "Etc/GMT-11", "Etc/GMT-12",
			  "Etc/GMT-13", "Etc/GMT-14", "Etc/GMT0", "Etc/Greenwich", "Etc/UCT", "Etc/UTC", "Etc/Universal", "Etc/Zulu"},
	"Europe":	 {"Europe/Amsterdam", "Europe/Andorra", "Europe/Astrakhan", "Europe/Athens", "Europe/Belfast", "Europe/Belgrade",
			  "Europe/Berlin", "Europe/Bratislava", "Europe/Brussels", "Europe/Bucharest", "Europe/Budapest", "Europe/Busingen",
			  "Europe/Chisinau", "Europe/Copenhagen", "Europe/Dublin", "Europe/Gibraltar", "Europe/Guernsey", "Europe/Helsinki",
			  "Europe/Isle_of_Man", "Europe/Istanbul", "Europe/Jersey", "Europe/Kaliningrad", "Europe/Kiev", "Europe/Kirov",
			  "Europe/Kyiv", "Europe/Lisbon", "Europe/Ljubljana", "Europe/London", "Europe/Luxembourg", "Europe/Madrid", "Europe/Malta",
			  "Europe/Mariehamn", "Europe/Minsk", "Europe/Monaco", "Europe/Moscow", "Europe/Nicosia", "Europe/Oslo", "Europe/Paris",
			  "Europe/Podgorica", "Europe/Prague", "Europe/Riga", "Europe/Rome", "Europe/Samara", "Europe/San_Marino", "Europe/Sarajevo",
			  "Europe/Saratov", "Europe/Simferopol", "Europe/Skopje", "Europe/Sofia", "Europe/Stockholm", "Europe/Tallinn", "Europe/Tirane",
			  "Europe/Tiraspol", "Europe/Ulyanovsk", "Europe/Uzhgorod", "Europe/Vaduz", "Europe/Vatican", "Europe/Vienna", "Europe/Vilnius",
			  "Europe/Volgograd", "Europe/Warsaw", "Europe/Zagreb", "Europe/Zaporozhye", "Europe/Zurich"},
	"Factory":	 {"Factory"},
	"GB":		 {"GB"},
	"GB-Eire":	 {"GB-Eire"},
	"GMT":		 {"GMT"}, 
	"GMT+0":	 {"GMT+0"},
	"GMT-0":	 {"GMT-0"},
	"GMT0":		 {"GMT0"},
	"Greenwich":	 {"Greenwich"},
	"HST":		 {"HST"},
	"Hongkong":	 {"Hongkong"},
	"Iceland":	 {"Iceland"},
	"Indian":	 {"Indian/Antananarivo", "Indian/Chagos", "Indian/Christmas", "Indian/Cocos", "Indian/Comoro", "Indian/Kerguelen", "Indian/Mahe",
			  "Indian/Maldives", "Indian/Mauritius", "Indian/Mayotte", "Indian/Reunion"},
	"Iran":		 {"Iran"},
	"Israel":	 {"Israel"},
	"Jamaica":	 {"Jamaica"},
	"Japan":	 {"Japan"},
	"Kwajalein":	 {"Kwajalein"},
	"Libya":	 {"Libya"},
	"MET":	 	 {"MET"},
	"MST":	 	 {"MST"},
	"MST7MDT":	 {"MST7MDT"},
	"Mexico":	 {"Mexico/BajaNorte", "Mexico/BajaSur", "Mexico/General"},
	"NZ":		 {"NZ"},
	"NZ-CHAT":	 {"NZ-CHAT"},
	"Navajo":	 {"Navajo"},
	"PRC":		 {"PRC"},
	"PST8PDT":	 {"PST8PDT"},
	"Pacific":	 {"Pacific/Apia", "Pacific/Auckland", "Pacific/Bougainville", "Pacific/Chatham", "Pacific/Chuuk", "Pacific/Easter",
			  "Pacific/Efate", "Pacific/Enderbury", "Pacific/Fakaofo", "Pacific/Fiji", "Pacific/Funafuti", "Pacific/Galapagos",
			  "Pacific/Gambier", "Pacific/Guadalcanal", "Pacific/Guam", "Pacific/Honolulu", "Pacific/Johnston", "Pacific/Kanton",
			  "Pacific/Kiritimati", "Pacific/Kosrae", "Pacific/Kwajalein", "Pacific/Majuro", "Pacific/Marquesas", "Pacific/Midway",
			  "Pacific/Nauru", "Pacific/Niue", "Pacific/Norfolk", "Pacific/Noumea", "Pacific/Pago_Pago", "Pacific/Palau",
			  "Pacific/Pitcairn", "Pacific/Pohnpei", "Pacific/Ponape", "Pacific/Port_Moresby", "Pacific/Rarotonga", "Pacific/Saipan",
			  "Pacific/Samoa", "Pacific/Tahiti", "Pacific/Tarawa", "Pacific/Tongatapu", "Pacific/Truk", "Pacific/Wake", "Pacific/Wallis",
			  "Pacific/Yap"},
	"Poland":	 {"Poland"},
	"Portugal":	 {"Portugal"},
	"ROC":		 {"ROC"},
	"ROK":		 {"ROK"},
	"Singapore":	 {"Singapore"},
	"Turkey":	 {"Turkey"},
	"UCT":		 {"UCT"},
	
	"US":		 {"US/Alaska", "US/Aleutian", "US/Arizona", "US/Central", "US/East-Indiana", "US/Eastern", "US/Hawaii", "US/Indiana-Starke",
			  "US/Michigan", "US/Mountain", "US/Pacific", "US/Samoa"},
	"UTC":		 {"UTC"},
	"Universal":	 {"Universal"},
	"W-SU":		 {"W-SU"},
	"WET":		 {"WET"},
	"Zulu":		 {"Zulu"}, 
	}

	// Check if the specified region is in the map
	zones, ok := timeZones[region]
	if !ok {
		fmt.Println("Region not found.")
		return
	}

	// Output all time zones for the region with full date information in tabulated columns
	fmt.Println("\nTime Zone\t\t\t\tLocal Time\n---------\t\t\t\t----------\n")
	for _, zone := range zones {
		loc, err := time.LoadLocation(zone)
		if err != nil {
			fmt.Printf("Error loading timezone %s: %v\n", zone, err)
			continue
		}
		// Convert the provided seconds to the local time in the specified timezone
		localTime := time.Unix(seconds, 0).In(loc)
		fmt.Printf("%-32s\t%s\n", zone, localTime.Format("Mon 02 Jan 2006 15:04:05 MST"))
	}
}

func main() {
	// Flag for decimal input
	decimal := flag.Bool("d", false, "Set this flag to process decimal values")
	// Flag for time zones
	timeZone := flag.String("tz", "", "Specify a region to output all its time zones")

	flag.Parse()

	// Usage instructions
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [-d] [-t Region] [hex|dec value]\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Check if the TimeZone flag was set
	if *timeZone != "" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()

			var seconds int64
			var err error

			// Process the input value based on the set flag
			if *decimal {
				seconds, err = strconv.ParseInt(input, 10, 64)
			} else {
				seconds, err = strconv.ParseInt(input, 16, 64)
			}

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing the value: %v\n", err)
				continue
			}

			printTimeZones(*timeZone, seconds)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		}
		return // Exit the program after outputting the time zones
	}

	// If the TimeZone flag is not set, process input as Unix timestamp
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		var seconds int64
		var err error

		// Process the input value based on the set flag
		if *decimal {
			seconds, err = strconv.ParseInt(input, 10, 64)
		} else {
			seconds, err = strconv.ParseInt(input, 16, 64)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing the value: %v\n", err)
			continue
		}

		// Convert Unix timestamp to a date in UTC
		date := time.Unix(seconds, 0).UTC()
		// Output the date in UTC
		fmt.Println(date.Format("Monday, 02. January 2006 15:04:05 +0000 (UTC)"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
	}
}
