**Sissejuhatus Go keelde**
Go, tuntud ka kui Golang, on Google'i arendatud avatud lähtekoodiga programmeerimiskeel, mis on loodud kiireks, lihtsaks ja usaldusväärseks tarkvara arendamiseks. Go keelt iseloomustab selle lihtsus, tõhusus ja suurepärane toetus samaaegseks töötlemiseks (concurrency).

**Sisseehitatud tüübid**
Go keeles on mitmeid sisseehitatud andmetüüpe, mis toetavad erinevaid andmetöötluse vajadusi:

**bool**: Tõeväärtuse tüüp, mis võib olla kas true või false.
**string**: Sõnetüüp, mis koosneb Unicode'i tähemärkidest.
**int, int8, int16, int32, int64**: Erimahulised täisarvud.
**uint, uint8, uint16, uint32, uint64, uintptr**: Erimahulised täisarvud, kuid ainult positiivsed väärtused.
**float32, float64**: Ujukomaarvud, mis võimaldavad esitada reaalarve.
**complex64, complex128**: Kompleksarvud, mis koosnevad kahest float tüübist (reaal- ja imaginaarosa).
**byte**: Samaväärne uint8-ga, kasutatakse sageli andmete töötlemisel.
**rune**: Samaväärne int32-ga, kasutatakse Unicode'i tähemärkide esitamiseks.

**Sisseehitatud funktsioonid**
Go keel sisaldab mitmeid funktsioone, mis toetavad sisseehitatud tüüpide ja andmestruktuuride töötlemist:

**append**: Lisab elemente viilu lõppu.
**cap**: Tagastab viilu, massiivi või kanali mahutavuse.
**close**: Sulgeb kanali.
**complex**: Loob kompleksarvu kahest ujukomaarvust.
**copy**: Kopeerib elemendid ühest viilust teise.
**delete**: Eemaldab võtme kaardilt.
**len**: Tagastab argumendi pikkuse.
**make**: Loob viilu, kaardi või kanali.
**new**: Reserveerib mälu tüübile vastava väärtuse jaoks.
**panic**: Katkestab normaalse funktsiooni käitumise.
**real, imag**: Tagastavad kompleksarvu reaalse või imaginaarosa.
**recover**: Aitab taastada normaalse programmi käitumise pärast paanikat.
See dokumentatsioon on mõeldud üldiseks juhendiks ja ei kata kõiki Go keele funktsioone ja võimalusi. Lisainformatsiooni saamiseks soovitame tutvuda Go ametliku dokumentatsiooniga.
https://go.dev/doc/
