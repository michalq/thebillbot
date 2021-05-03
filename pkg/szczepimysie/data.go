package szczepimysie

import "time"

type Data struct {
	ObjectId            int       //"OBJECTID" : 3,
	Data                time.Time //"Data" : 1609227000000,
	SzczepieniaSuma     int       //"SZCZEPIENIA_SUMA" : 6423,
	SzczepieniaDziennie int       //"SZCZEPIENIA_DZIENNIE" : 4391,
	Dawka1Suma          int       //"DAWKA_1_SUMA" : 6423,
	Dawka2Suma          int       //"DAWKA_2_SUMA" : 0,
	Dawka2Dziennie      int       //"DAWKA_2_DZIENNIE" : 0,

	//"DATA_SHOW" : "29.12.2020",
	//"ODCZYNY_NIEPOZADANE" : 10,
	//"DAWKI_UTRACONE" : 6,
	//"SZCZEPIENIA_PLEC_NIEUSTALONO" : 0,
	//"SZCZEPIENIA_KOBIETY" : 329,
	//"SZCZEPIENIA_MEZCZYZNI" : 381,
	//"SZCZEPIENIA0_17" : 0,
	//"SZCZEPIENIA18_30" : 79,
	//"SZCZEPIENIA31_40" : 89,
	//"SZCZEPIENIA41_50" : 89,
	//"SZCZEPIENIA51_60" : 81,
	//"SZCZEPIENIA61_70" : 97,
	//"SZCZEPIENIA71_75" : 89,
	//"SZCZEPIENIA75_" : 89,
	//"SZCZEPIENIA_WIEK_NIEUSTALONO" : 6,
	//"STAN_MAGAZYN" : null,
	//"LICZBA_DAWEK_PUNKTY" : null,
	//"SUMA_DAWEK_POLSKA" : null,
	//"zamowienia_realizacja" : null
}
