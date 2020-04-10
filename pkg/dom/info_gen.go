package dom

var GeneratedInfo = &Info{
	counters: map[string]int{"DATA_ZDARZ": 1, "DATA_ZDARZENIA": 1, "DROGA_PUBLICZNA": 1, "GMINA": 1, "GODZINA_ZDARZ": 1, "ID": 1, "INFO_O_DRODZE": 1, "JEDNOSTKA_LIKWIDUJACA": 1, "JEDNOSTKA_MIEJSCA": 1, "JEDNOSTKA_OPERATORA": 1, "MIEJSCE": 1, "MIEJSCOWOSC": 1, "NR_KW": 1, "POJAZDY": 1, "POWIAT": 1, "PREDKOSC_DOPUSZCZALNA": 1, "SZOS_KOD": 1, "SZOS_TABK_TYPE": 1, "SZRD_KOD": 1, "SZRD_TABK_TYPE": 1, "UCZESTNICY": 1, "ULICA_ADRES": 1, "ULICA_SKRZYZ": 1, "WARUNKI_ATMOSFERYCZNE": 1, "WOJ": 1, "__src": 1},
	children: map[string]*Info{
		"ULICA_ADRES":           {counters: map[string]int{"_": 1}},
		"ULICA_SKRZYZ":          {counters: map[string]int{"_": 1}},
		"DATA_ZDARZ":            {counters: map[string]int{"_": 1}},
		"PREDKOSC_DOPUSZCZALNA": {counters: map[string]int{"_": 1}},
		"SZRD_TABK_TYPE":        {counters: map[string]int{"_": 1}},
		"DROGA_PUBLICZNA":       {counters: map[string]int{"_": 1}},
		"JEDNOSTKA_MIEJSCA":     {counters: map[string]int{"_": 1}},
		"JEDNOSTKA_OPERATORA":   {counters: map[string]int{"_": 1}},
		"INFO_O_DRODZE": {
			counters: map[string]int{"NAWIERZCHNIA": 1, "OZNAKOWANIE_POZIOME": 1, "RODZAJ_DROGI": 1, "STAN_NAWIERZCHNI": 1, "SYGNALIZACJA_SWIETLNA": 1},
			children: map[string]*Info{
				"STAN_NAWIERZCHNI": {
					counters: map[string]int{"STNA_KOD": 1},
					children: map[string]*Info{
						"STNA_KOD": {
							counters: map[string]int{"STNA_KOD": 1, "STNA_TABK_TYPE": 1, "ZSZD_ID": 1},
							children: map[string]*Info{
								"ZSZD_ID":        {counters: map[string]int{"_": 1}},
								"STNA_KOD":       {counters: map[string]int{"_": 1}},
								"STNA_TABK_TYPE": {counters: map[string]int{"_": 1}},
							},
						},
					},
				},
				"RODZAJ_DROGI": {
					counters: map[string]int{"RODR_KOD": 1},
					children: map[string]*Info{
						"RODR_KOD": {
							counters: map[string]int{"RODR_KOD": 1, "RODR_TABK_TYPE": 1, "ZSZD_ID": 1},
							children: map[string]*Info{
								"ZSZD_ID":        {counters: map[string]int{"_": 1}},
								"RODR_KOD":       {counters: map[string]int{"_": 1}},
								"RODR_TABK_TYPE": {counters: map[string]int{"_": 1}},
							},
						},
					},
				},
				"NAWIERZCHNIA": {
					counters: map[string]int{"NADR_KOD": 1},
					children: map[string]*Info{
						"NADR_KOD": {
							counters: map[string]int{"NADR_KOD": 1, "NADR_TABK_TYPE": 1, "ZSZD_ID": 1},
							children: map[string]*Info{
								"ZSZD_ID":        {counters: map[string]int{"_": 1}},
								"NADR_KOD":       {counters: map[string]int{"_": 1}},
								"NADR_TABK_TYPE": {counters: map[string]int{"_": 1}},
							},
						},
					},
				},
				"SYGNALIZACJA_SWIETLNA": {
					counters: map[string]int{"SYSW_KOD": 1},
					children: map[string]*Info{
						"SYSW_KOD": {
							counters: map[string]int{"SYSW_KOD": 1, "SYSW_TABK_TYPE": 1, "ZSZD_ID": 1},
							children: map[string]*Info{
								"ZSZD_ID":        {counters: map[string]int{"_": 1}},
								"SYSW_KOD":       {counters: map[string]int{"_": 1}},
								"SYSW_TABK_TYPE": {counters: map[string]int{"_": 1}},
							},
						},
					},
				},
				"OZNAKOWANIE_POZIOME": {
					counters: map[string]int{"OZPO_KOD": 1},
					children: map[string]*Info{
						"OZPO_KOD": {
							counters: map[string]int{"OZPO_KOD": 1, "OZPO_TABK_TYPE": 1, "ZSZD_ID": 1},
							children: map[string]*Info{
								"OZPO_KOD":       {counters: map[string]int{"_": 1}},
								"OZPO_TABK_TYPE": {counters: map[string]int{"_": 1}},
								"ZSZD_ID":        {counters: map[string]int{"_": 1}},
							},
						},
					},
				},
			},
		},
		"POJAZDY": {
			counters: map[string]int{"POJAZD": 2},
			children: map[string]*Info{
				"POJAZD": {
					counters: map[string]int{"ID": 1, "MARKA": 1, "NR_POJAZDU": 1, "RODZAJ_POJAZDU": 1, "SPSU_KOD": 1, "SPSU_TABK_TYPE": 1, "ZSZD_ID": 1},
					children: map[string]*Info{
						"NR_POJAZDU":     {counters: map[string]int{"_": 1}},
						"RODZAJ_POJAZDU": {counters: map[string]int{"_": 1}},
						"MARKA":          {counters: map[string]int{"_": 1}},
						"SPSU_KOD":       {counters: map[string]int{"_": 1}},
						"SPSU_TABK_TYPE": {counters: map[string]int{"_": 1}},
						"ID":             {counters: map[string]int{"_": 1}},
						"ZSZD_ID":        {counters: map[string]int{"_": 1}},
					},
				},
			},
		},
		"MIEJSCOWOSC":    {counters: map[string]int{"_": 1}},
		"SZOS_TABK_TYPE": {counters: map[string]int{"_": 1}},
		"WARUNKI_ATMOSFERYCZNE": {
			counters: map[string]int{"SSWA_KOD": 1},
			children: map[string]*Info{
				"SSWA_KOD": {
					counters: map[string]int{"SSWA_KOD": 1, "SSWA_TABK_TYPE": 1, "ZSZD_ID": 1},
					children: map[string]*Info{
						"SSWA_KOD":       {counters: map[string]int{"_": 1}},
						"SSWA_TABK_TYPE": {counters: map[string]int{"_": 1}},
						"ZSZD_ID":        {counters: map[string]int{"_": 1}},
					},
				},
			},
		},
		"NR_KW":          {counters: map[string]int{"_": 1}},
		"POWIAT":         {counters: map[string]int{"_": 1}},
		"DATA_ZDARZENIA": {counters: map[string]int{"_": 1}},
		"GODZINA_ZDARZ":  {counters: map[string]int{"_": 1}},
		"SZOS_KOD":       {counters: map[string]int{"_": 1}},
		"SZRD_KOD":       {counters: map[string]int{"_": 1}},
		"MIEJSCE": {
			counters: map[string]int{"CHARAKT_MIEJSCA": 1, "OBSZAR_ZABUDOWANY": 1, "SKRZYZOWANIE": 1},
			children: map[string]*Info{
				"OBSZAR_ZABUDOWANY": {
					counters: map[string]int{"ZABU_KOD": 1},
					children: map[string]*Info{
						"ZABU_KOD": {
							counters: map[string]int{"ZABU_KOD": 1, "ZABU_TABK_TYPE": 1, "ZSZD_ID": 1},
							children: map[string]*Info{
								"ZSZD_ID":        {counters: map[string]int{"_": 1}},
								"ZABU_KOD":       {counters: map[string]int{"_": 1}},
								"ZABU_TABK_TYPE": {counters: map[string]int{"_": 1}},
							},
						},
					},
				},
				"SKRZYZOWANIE": {
					counters: map[string]int{"SKRZ_KOD": 1},
					children: map[string]*Info{
						"SKRZ_KOD": {
							counters: map[string]int{"SKRZ_KOD": 1, "SKRZ_TABK_TYPE": 1, "ZSZD_ID": 1},
							children: map[string]*Info{
								"ZSZD_ID":        {counters: map[string]int{"_": 1}},
								"SKRZ_KOD":       {counters: map[string]int{"_": 1}},
								"SKRZ_TABK_TYPE": {counters: map[string]int{"_": 1}},
							},
						},
					},
				},
				"CHARAKT_MIEJSCA": {
					counters: map[string]int{"CHMZ_KOD": 1},
					children: map[string]*Info{
						"CHMZ_KOD": {
							counters: map[string]int{"CHMZ_KOD": 1, "CHMZ_TABK_TYPE": 1, "ZSZD_ID": 1},
							children: map[string]*Info{
								"ZSZD_ID":        {counters: map[string]int{"_": 1}},
								"CHMZ_KOD":       {counters: map[string]int{"_": 1}},
								"CHMZ_TABK_TYPE": {counters: map[string]int{"_": 1}},
							},
						},
					},
				},
			},
		},
		"UCZESTNICY": {
			counters: map[string]int{"OSOBA": 2},
			children: map[string]*Info{
				"OSOBA": {
					counters: map[string]int{"DATA_UR": 1, "ID": 1, "LICZBA_LAT_KIEROWANIA": 1, "PLEC": 1, "PRZYCZYNY_KIEROWCY": 1, "SOBY_KOD": 1, "SOBY_TABK_TYPE": 1, "SRUZ_KOD": 1, "SRUZ_TABK_TYPE": 1, "SSRU_KOD": 1, "SSRU_TABK_TYPE": 1, "SUSU_KOD": 1, "SUSU_TABK_TYPE": 1, "ZSPO_ID": 1, "ZSZD_ID": 1},
					children: map[string]*Info{
						"SRUZ_KOD":              {counters: map[string]int{"_": 1}},
						"SRUZ_TABK_TYPE":        {counters: map[string]int{"_": 1}},
						"LICZBA_LAT_KIEROWANIA": {counters: map[string]int{"_": 1}},
						"ID":                    {counters: map[string]int{"_": 1}},
						"ZSZD_ID":               {counters: map[string]int{"_": 1}},
						"SSRU_KOD":              {counters: map[string]int{"_": 1}},
						"SOBY_KOD":              {counters: map[string]int{"_": 1}},
						"PLEC":                  {counters: map[string]int{"_": 1}},
						"PRZYCZYNY_KIEROWCY": {
							counters: map[string]int{"SPSZ_KOD": 1},
							children: map[string]*Info{
								"SPSZ_KOD": {
									counters: map[string]int{"SPSZ_KOD": 1, "SPSZ_TABK_TYPE": 1, "ZSUC_ID": 1},
									children: map[string]*Info{
										"ZSUC_ID":        {counters: map[string]int{"_": 1}},
										"SPSZ_KOD":       {counters: map[string]int{"_": 1}},
										"SPSZ_TABK_TYPE": {counters: map[string]int{"_": 1}},
									},
								},
							},
						},
						"ZSPO_ID":        {counters: map[string]int{"_": 1}},
						"SUSU_KOD":       {counters: map[string]int{"_": 1}},
						"SUSU_TABK_TYPE": {counters: map[string]int{"_": 1}},
						"SSRU_TABK_TYPE": {counters: map[string]int{"_": 1}},
						"DATA_UR":        {counters: map[string]int{"_": 1}},
						"SOBY_TABK_TYPE": {counters: map[string]int{"_": 1}},
					},
				},
			},
		},
		"JEDNOSTKA_LIKWIDUJACA": {counters: map[string]int{"_": 1}},
		"WOJ":                   {counters: map[string]int{"_": 1}},
		"ID":                    {counters: map[string]int{"_": 1}},
		"GMINA":                 {counters: map[string]int{"_": 1}},
	},
}
