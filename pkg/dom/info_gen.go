package dom

var GeneratedInfo = &Info{
	counters: map[string]int{"DATA_PRZYJAZDU": 1, "DATA_ZDARZ": 1, "DATA_ZDARZENIA": 1, "DATA_ZGLOSZENIA": 1, "DROGA_PUBLICZNA": 1, "DZIELNICA": 1, "GMINA": 1, "GODZINA_ZDARZ": 1, "GPS_X_GUS": 1, "GPS_Y_GUS": 1, "ID": 1, "INFO_O_DRODZE": 1, "INNE_PRZYCZYNY": 1, "JEDNOSTKA_LIKWIDUJACA": 1, "JEDNOSTKA_MIEJSCA": 1, "JEDNOSTKA_OPERATORA": 1, "KIERUNEK": 1, "KM_HM": 1, "KOD_GUS": 1, "LICZBA_PASOW": 1, "MIEJSCE": 1, "MIEJSCOWOSC": 1, "NADZOR": 1, "NR_KW": 1, "NUMER_DOMU": 1, "ODLEGLOSC_SKRZYZ": 1, "POJAZD": 19, "POJAZDY": 1, "POWIAT": 1, "PREDKOSC_DOPUSZCZALNA": 1, "RONDO_WEZEL": 1, "STREFA_RUCHU": 1, "STREFA_ZAMIESZKALA": 1, "SZOS_KOD": 1, "SZOS_TABK_TYPE": 1, "SZRD_KOD": 1, "SZRD_TABK_TYPE": 1, "UCZESTNICY": 1, "ULICA_ADRES": 1, "ULICA_SKRZYZ": 1, "USZKODZENIA_POZA_POJAZDAMI": 1, "WARUNKI_ATMOSFERYCZNE": 1, "WOJ": 1, "WSP_GPS_X": 1, "WSP_GPS_Y": 1, "ZSSD_KOD": 1, "ZSSD_KOD2": 1, "_DataMod": 1, "_DataUtworzenia": 1, "_SzkicZdarzenia": 1, "__src": 1},
	children: map[string]*Info{
		"NADZOR":            {counters: map[string]int{"_": 1}},
		"JEDNOSTKA_MIEJSCA": {counters: map[string]int{"_": 1}},
		"NUMER_DOMU":        {counters: map[string]int{"_": 1}},
		"DROGA_PUBLICZNA":   {counters: map[string]int{"_": 1}},
		"UCZESTNICY": {
			counters: map[string]int{"OSOBA": 49},
			children: map[string]*Info{
				"OSOBA": {
					counters: map[string]int{"DATA_UR": 1, "ID": 1, "INFO_DODATKOWE": 1, "INWALIDA": 1, "LICZBA_LAT_KIEROWANIA": 1, "MIEJSCE_W_POJ": 1, "OBCOKRAJOWIEC": 1, "ODBLASKI": 1, "PLEC": 1, "POZIOM_ALKOHOLU": 1, "PRZYCZYNY_KIEROWCY": 1, "PRZYCZYNY_PIESZY": 1, "SOBY_KOD": 1, "SOBY_TABK_TYPE": 1, "SPAK_KOD": 1, "SPAK_TABK_TYPE": 1, "SRUZ_KOD": 1, "SRUZ_TABK_TYPE": 1, "SSRU_KOD": 1, "SSRU_TABK_TYPE": 1, "STUC_KOD": 1, "STUC_TABK_TYPE": 1, "SUSU_KOD": 1, "SUSU_TABK_TYPE": 1, "SUZZ_KOD": 1, "SUZZ_TABK_TYPE": 1, "ZBIEGL_Z_MIEJSCA": 1, "ZSPO_ID": 1, "ZSZD_ID": 1, "_DataMod": 1, "_DataUtworzenia": 1},
					children: map[string]*Info{
						"SPAK_KOD":        {counters: map[string]int{"_": 1}},
						"SSRU_TABK_TYPE":  {counters: map[string]int{"_": 1}},
						"DATA_UR":         {counters: map[string]int{"_": 1}},
						"POZIOM_ALKOHOLU": {counters: map[string]int{"_": 1}},
						"PRZYCZYNY_PIESZY": {
							counters: map[string]int{"SPPI_KOD": 1},
							children: map[string]*Info{
								"SPPI_KOD": {
									counters: map[string]int{"SPPI_KOD": 1, "SPPI_TABK_TYPE": 1, "ZSUC_ID": 1},
									children: map[string]*Info{
										"ZSUC_ID":        {counters: map[string]int{"_": 1}},
										"SPPI_KOD":       {counters: map[string]int{"_": 1}},
										"SPPI_TABK_TYPE": {counters: map[string]int{"_": 1}},
									},
								},
							},
						},
						"ODBLASKI":              {counters: map[string]int{"_": 1}},
						"SOBY_TABK_TYPE":        {counters: map[string]int{"_": 1}},
						"LICZBA_LAT_KIEROWANIA": {counters: map[string]int{"_": 1}},
						"MIEJSCE_W_POJ":         {counters: map[string]int{"_": 1}},
						"SUSU_TABK_TYPE":        {counters: map[string]int{"_": 1}},
						"STUC_TABK_TYPE":        {counters: map[string]int{"_": 1}},
						"SUZZ_TABK_TYPE":        {counters: map[string]int{"_": 1}},
						"INWALIDA":              {counters: map[string]int{"_": 1}},
						"ID":                    {counters: map[string]int{"_": 1}},
						"ZSZD_ID":               {counters: map[string]int{"_": 1}},
						"SRUZ_KOD":              {counters: map[string]int{"_": 1}},
						"SRUZ_TABK_TYPE":        {counters: map[string]int{"_": 1}},
						"SUSU_KOD":              {counters: map[string]int{"_": 1}},
						"OBCOKRAJOWIEC":         {counters: map[string]int{"_": 1}},
						"STUC_KOD":              {counters: map[string]int{"_": 1}},
						"ZBIEGL_Z_MIEJSCA":      {counters: map[string]int{"_": 1}},
						"ZSPO_ID":               {counters: map[string]int{"_": 1}},
						"SOBY_KOD":              {counters: map[string]int{"_": 1}},
						"PRZYCZYNY_KIEROWCY": {
							counters: map[string]int{"SPSZ_KOD": 1},
							children: map[string]*Info{
								"SPSZ_KOD": {
									counters: map[string]int{"SPSZ_KOD": 1, "SPSZ_TABK_TYPE": 1, "ZSUC_ID": 1},
									children: map[string]*Info{
										"SPSZ_KOD":       {counters: map[string]int{"_": 1}},
										"SPSZ_TABK_TYPE": {counters: map[string]int{"_": 1}},
										"ZSUC_ID":        {counters: map[string]int{"_": 1}},
									},
								},
							},
						},
						"INFO_DODATKOWE": {
							counters: map[string]int{"JAZDA_BEZ": 2, "POD_WPLYWEM": 2},
							children: map[string]*Info{
								"JAZDA_BEZ": {
									counters: map[string]int{"SUSB_KOD": 2, "SUSW_KOD": 2},
									children: map[string]*Info{
										"SUSW_KOD": {
											counters: map[string]int{"SUSW_KOD": 1, "SUSW_TABK_TYPE": 1, "ZSUC_ID": 1},
											children: map[string]*Info{
												"ZSUC_ID":        {counters: map[string]int{"_": 1}},
												"SUSW_KOD":       {counters: map[string]int{"_": 1}},
												"SUSW_TABK_TYPE": {counters: map[string]int{"_": 1}},
											},
										},
										"SUSB_KOD": {
											counters: map[string]int{"SUSB_KOD": 1, "SUSB_TABK_TYPE": 1, "ZSUC_ID": 1},
											children: map[string]*Info{
												"ZSUC_ID":        {counters: map[string]int{"_": 1}},
												"SUSB_KOD":       {counters: map[string]int{"_": 1}},
												"SUSB_TABK_TYPE": {counters: map[string]int{"_": 1}},
											},
										},
									},
								},
								"POD_WPLYWEM": {
									counters: map[string]int{"SUSB_KOD": 2, "SUSW_KOD": 2},
									children: map[string]*Info{
										"SUSB_KOD": {
											counters: map[string]int{"SUSB_KOD": 1, "SUSB_TABK_TYPE": 1, "ZSUC_ID": 1},
											children: map[string]*Info{
												"ZSUC_ID":        {counters: map[string]int{"_": 1}},
												"SUSB_KOD":       {counters: map[string]int{"_": 1}},
												"SUSB_TABK_TYPE": {counters: map[string]int{"_": 1}},
											},
										},
										"SUSW_KOD": {
											counters: map[string]int{"SUSW_KOD": 1, "SUSW_TABK_TYPE": 1, "ZSUC_ID": 1},
											children: map[string]*Info{
												"SUSW_TABK_TYPE": {counters: map[string]int{"_": 1}},
												"ZSUC_ID":        {counters: map[string]int{"_": 1}},
												"SUSW_KOD":       {counters: map[string]int{"_": 1}},
											},
										},
									},
								},
							},
						},
						"SUZZ_KOD":       {counters: map[string]int{"_": 1}},
						"SPAK_TABK_TYPE": {counters: map[string]int{"_": 1}},
						"SSRU_KOD":       {counters: map[string]int{"_": 1}},
						"PLEC":           {counters: map[string]int{"_": 1}},
					},
				},
			},
		},
		"INNE_PRZYCZYNY": {
			counters: map[string]int{"SPIP_KOD": 2},
			children: map[string]*Info{
				"SPIP_KOD": {
					counters: map[string]int{"SPIP_KOD": 1, "SPIP_TABK_TYPE": 1, "ZSZD_ID": 1},
					children: map[string]*Info{
						"ZSZD_ID":        {counters: map[string]int{"_": 1}},
						"SPIP_KOD":       {counters: map[string]int{"_": 1}},
						"SPIP_TABK_TYPE": {counters: map[string]int{"_": 1}},
					},
				},
			},
		},
		"POJAZD": {
			counters: map[string]int{"ID": 1, "INNE_CECHY_POJAZU": 1, "KRAJ_REJ": 1, "KRAJ_UBZ": 1, "MARKA": 1, "NR_POJAZDU": 1, "RODZAJ_POJAZDU": 1, "SPSU_KOD": 1, "SPSU_TABK_TYPE": 1, "STAN_POJAZDU": 1, "ZSZD_ID": 1},
			children: map[string]*Info{
				"KRAJ_UBZ":       {counters: map[string]int{"_": 1}},
				"ID":             {counters: map[string]int{"_": 1}},
				"RODZAJ_POJAZDU": {counters: map[string]int{"_": 1}},
				"SPSU_KOD":       {counters: map[string]int{"_": 1}},
				"SPSU_TABK_TYPE": {counters: map[string]int{"_": 1}},
				"INNE_CECHY_POJAZU": {
					counters: map[string]int{"SPIC_KOD": 1},
					children: map[string]*Info{
						"SPIC_KOD": {
							counters: map[string]int{"SPIC_KOD": 1, "SPIC_TABK_TYPE": 1, "ZSPO_ID": 1},
							children: map[string]*Info{
								"ZSPO_ID":        {counters: map[string]int{"_": 1}},
								"SPIC_KOD":       {counters: map[string]int{"_": 1}},
								"SPIC_TABK_TYPE": {counters: map[string]int{"_": 1}},
							},
						},
					},
				},
				"STAN_POJAZDU": {
					counters: map[string]int{"SPSP_KOD": 7},
					children: map[string]*Info{
						"SPSP_KOD": {
							counters: map[string]int{"SPSP_KOD": 1, "SPSP_TABK_TYPE": 1, "ZSPO_ID": 1},
							children: map[string]*Info{
								"SPSP_KOD":       {counters: map[string]int{"_": 1}},
								"SPSP_TABK_TYPE": {counters: map[string]int{"_": 1}},
								"ZSPO_ID":        {counters: map[string]int{"_": 1}},
							},
						},
					},
				},
				"KRAJ_REJ":   {counters: map[string]int{"_": 1}},
				"ZSZD_ID":    {counters: map[string]int{"_": 1}},
				"NR_POJAZDU": {counters: map[string]int{"_": 1}},
				"MARKA":      {counters: map[string]int{"_": 1}},
			},
		},
		"DATA_PRZYJAZDU":      {counters: map[string]int{"_": 1}},
		"JEDNOSTKA_OPERATORA": {counters: map[string]int{"_": 1}},
		"NR_KW":               {counters: map[string]int{"_": 1}},
		"GODZINA_ZDARZ":       {counters: map[string]int{"_": 1}},
		"INFO_O_DRODZE": {
			counters: map[string]int{"NAWIERZCHNIA": 1, "OZNAKOWANIE_POZIOME": 1, "RODZAJ_DROGI": 1, "STAN_NAWIERZCHNI": 6, "SYGNALIZACJA_SWIETLNA": 1},
			children: map[string]*Info{
				"RODZAJ_DROGI": {
					counters: map[string]int{"RODR_KOD": 1},
					children: map[string]*Info{
						"RODR_KOD": {
							counters: map[string]int{"RODR_KOD": 1, "RODR_TABK_TYPE": 1, "ZSZD_ID": 1},
							children: map[string]*Info{
								"RODR_KOD":       {counters: map[string]int{"_": 1}},
								"RODR_TABK_TYPE": {counters: map[string]int{"_": 1}},
								"ZSZD_ID":        {counters: map[string]int{"_": 1}},
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
								"SYSW_KOD":       {counters: map[string]int{"_": 1}},
								"SYSW_TABK_TYPE": {counters: map[string]int{"_": 1}},
								"ZSZD_ID":        {counters: map[string]int{"_": 1}},
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
								"ZSZD_ID":        {counters: map[string]int{"_": 1}},
								"OZPO_KOD":       {counters: map[string]int{"_": 1}},
								"OZPO_TABK_TYPE": {counters: map[string]int{"_": 1}},
							},
						},
					},
				},
				"STAN_NAWIERZCHNI": {
					counters: map[string]int{"STNA_KOD": 6},
					children: map[string]*Info{
						"STNA_KOD": {
							counters: map[string]int{"STNA_KOD": 1, "STNA_TABK_TYPE": 1, "ZSZD_ID": 1},
							children: map[string]*Info{
								"STNA_TABK_TYPE": {counters: map[string]int{"_": 1}},
								"ZSZD_ID":        {counters: map[string]int{"_": 1}},
								"STNA_KOD":       {counters: map[string]int{"_": 1}},
							},
						},
					},
				},
			},
		},
		"USZKODZENIA_POZA_POJAZDAMI": {
			counters: map[string]int{"SSUP_KOD": 5},
			children: map[string]*Info{
				"SSUP_KOD": {
					counters: map[string]int{"SSUP_KOD": 1, "SSUP_TABK_TYPE": 1, "ZSZD_ID": 1},
					children: map[string]*Info{
						"ZSZD_ID":        {counters: map[string]int{"_": 1}},
						"SSUP_KOD":       {counters: map[string]int{"_": 1}},
						"SSUP_TABK_TYPE": {counters: map[string]int{"_": 1}},
					},
				},
			},
		},
		"WSP_GPS_Y":             {counters: map[string]int{"_": 1}},
		"GPS_X_GUS":             {counters: map[string]int{"_": 1}},
		"GMINA":                 {counters: map[string]int{"_": 1}},
		"PREDKOSC_DOPUSZCZALNA": {counters: map[string]int{"_": 1}},
		"SZOS_KOD":              {counters: map[string]int{"_": 1}},
		"ULICA_SKRZYZ":          {counters: map[string]int{"_": 1}},
		"WARUNKI_ATMOSFERYCZNE": {
			counters: map[string]int{"SSWA_KOD": 5},
			children: map[string]*Info{
				"SSWA_KOD": {
					counters: map[string]int{"SSWA_KOD": 1, "SSWA_TABK_TYPE": 1, "ZSZD_ID": 1},
					children: map[string]*Info{
						"ZSZD_ID":        {counters: map[string]int{"_": 1}},
						"SSWA_KOD":       {counters: map[string]int{"_": 1}},
						"SSWA_TABK_TYPE": {counters: map[string]int{"_": 1}},
					},
				},
			},
		},
		"ZSSD_KOD":              {counters: map[string]int{"_": 1}},
		"LICZBA_PASOW":          {counters: map[string]int{"_": 1}},
		"JEDNOSTKA_LIKWIDUJACA": {counters: map[string]int{"_": 1}},
		"WOJ":                   {counters: map[string]int{"_": 1}},
		"POWIAT":                {counters: map[string]int{"_": 1}},
		"DATA_ZDARZENIA":        {counters: map[string]int{"_": 1}},
		"STREFA_ZAMIESZKALA":    {counters: map[string]int{"_": 1}},
		"GPS_Y_GUS":             {counters: map[string]int{"_": 1}},
		"SZOS_TABK_TYPE":        {counters: map[string]int{"_": 1}},
		"SZRD_KOD":              {counters: map[string]int{"_": 1}},
		"SZRD_TABK_TYPE":        {counters: map[string]int{"_": 1}},
		"KM_HM":                 {counters: map[string]int{"_": 1}},
		"KOD_GUS":               {counters: map[string]int{"_": 1}},
		"DATA_ZGLOSZENIA":       {counters: map[string]int{"_": 1}},
		"STREFA_RUCHU":          {counters: map[string]int{"_": 1}},
		"ID":                    {counters: map[string]int{"_": 1}},
		"ULICA_ADRES":           {counters: map[string]int{"_": 1}},
		"DATA_ZDARZ":            {counters: map[string]int{"_": 1}},
		"WSP_GPS_X":             {counters: map[string]int{"_": 1}},
		"RONDO_WEZEL":           {counters: map[string]int{"_": 1}},
		"POJAZDY": {
			counters: map[string]int{"POJAZD": 25},
			children: map[string]*Info{
				"POJAZD": {
					counters: map[string]int{"DATA_OST_BAD_TECH": 1, "ID": 1, "INNE_CECHY_POJAZU": 1, "KRAJ_REJ": 1, "KRAJ_UBZ": 1, "MARKA": 1, "NR_POJAZDU": 1, "RODZAJ_POJAZDU": 1, "ROK_PRODUKCJI": 1, "SPSU_KOD": 1, "SPSU_TABK_TYPE": 1, "STAN_POJAZDU": 1, "WYPOSAZENIE_DOD": 1, "ZSZD_ID": 1, "_DataMod": 1, "_DataUtworzenia": 1},
					children: map[string]*Info{
						"RODZAJ_POJAZDU": {counters: map[string]int{"_": 1}},
						"SPSU_KOD":       {counters: map[string]int{"_": 1}},
						"STAN_POJAZDU": {
							counters: map[string]int{"SPSP_KOD": 8},
							children: map[string]*Info{
								"SPSP_KOD": {
									counters: map[string]int{"SPSP_KOD": 1, "SPSP_TABK_TYPE": 1, "ZSPO_ID": 1},
									children: map[string]*Info{
										"ZSPO_ID":        {counters: map[string]int{"_": 1}},
										"SPSP_KOD":       {counters: map[string]int{"_": 1}},
										"SPSP_TABK_TYPE": {counters: map[string]int{"_": 1}},
									},
								},
							},
						},
						"ZSZD_ID":           {counters: map[string]int{"_": 1}},
						"DATA_OST_BAD_TECH": {counters: map[string]int{"_": 1}},
						"ID":                {counters: map[string]int{"_": 1}},
						"SPSU_TABK_TYPE":    {counters: map[string]int{"_": 1}},
						"INNE_CECHY_POJAZU": {
							counters: map[string]int{"SPIC_KOD": 2},
							children: map[string]*Info{
								"SPIC_KOD": {
									counters: map[string]int{"SPIC_KOD": 1, "SPIC_TABK_TYPE": 1, "ZSPO_ID": 1},
									children: map[string]*Info{
										"ZSPO_ID":        {counters: map[string]int{"_": 1}},
										"SPIC_KOD":       {counters: map[string]int{"_": 1}},
										"SPIC_TABK_TYPE": {counters: map[string]int{"_": 1}},
									},
								},
							},
						},
						"KRAJ_REJ":        {counters: map[string]int{"_": 1}},
						"KRAJ_UBZ":        {counters: map[string]int{"_": 1}},
						"ROK_PRODUKCJI":   {counters: map[string]int{"_": 1}},
						"WYPOSAZENIE_DOD": {counters: map[string]int{"_": 1}},
						"NR_POJAZDU":      {counters: map[string]int{"_": 1}},
						"MARKA":           {counters: map[string]int{"_": 1}},
					},
				},
			},
		},
		"ZSSD_KOD2":        {counters: map[string]int{"_": 1}},
		"ODLEGLOSC_SKRZYZ": {counters: map[string]int{"_": 1}},
		"KIERUNEK":         {counters: map[string]int{"_": 1}},
		"MIEJSCOWOSC":      {counters: map[string]int{"_": 1}},
		"MIEJSCE": {
			counters: map[string]int{"CHARAKT_MIEJSCA": 1, "GEOMETRIA_DROGI": 3, "OBSZAR_ZABUDOWANY": 1, "SKRZYZOWANIE": 1},
			children: map[string]*Info{
				"GEOMETRIA_DROGI": {
					counters: map[string]int{"GEOD_KOD": 2},
					children: map[string]*Info{
						"GEOD_KOD": {
							counters: map[string]int{"GEOD_KOD": 1, "GEOD_TABK_TYPE": 1, "ZSZD_ID": 1},
							children: map[string]*Info{
								"ZSZD_ID":        {counters: map[string]int{"_": 1}},
								"GEOD_KOD":       {counters: map[string]int{"_": 1}},
								"GEOD_TABK_TYPE": {counters: map[string]int{"_": 1}},
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
		"DZIELNICA": {counters: map[string]int{"_": 1}},
	},
}
