# vide-oh
Platforma za gledanje i deljenje video sadržaja

# Funkcionalnosti
Sistem poznaje tri tipa korisnika:
- neregistrovan korisnik
- registrovan korisnik
- administrator

U daljem tekstu su pobrojane funkcionalnosti kojim dati tip korisnika ima pristup.

Neregistrovan korisnik:
- registracija na sistem
- prijava na sistem
- pretraga videa u sistemu (sa prikazom thumbnail-a)
- gledanje videa (video će se streamovati (a ne skidati) i gledati u okviru ove aplikacije)

Registrovan korisnik:
- pretraga videa u sistemu
- gledanje videa
- upload videa na sistem (zajedno sa podacima poput naslova, opisa i sl.)
- pregled, izmena i brisanje sopstvenih videa
- pregled prosečne ocene i svih komentara na videu
- CRUD sopstvenih komentara i ocena na svakom videu
- izmena profila

Administrator:
- funkcionalnosti registrovanog korisnika osim CRUD-a sopstvenih videa
- blokiranje korisnika (o čemu se korisnik obaveštava putem imejla)
- brisanje videa ili izmena podataka o videu koje smatra neprikladnim
- brisanje neprikladnih komentara

# Arhitektura
Sistem se oslanja na mikroservisnu arhitekturu.
Komponente sistema su sledeće:
- Mikroservis za korisnike (autentifikacija, autorizacija, dodavanje i izmena korisnika, blokiranje korisnika...). Tehnologije: Go i PostgreSQL.
- Mikroservis za videe (CRUD videa, video streaming, pretraga videa...). Tehnologije: Python (Flask, Django ili FastAPI) i PostgreSQL (za podatke o videima).
- Mikroservis za komentare i ocene. Tehnologije: Go ili Rust i PostgreSQL.
- Centralna klijentska, frontend aplikacija koja podržava funkcionalnosti svih mikroservisa. Tehnologije: Vue.js.

# Potencijalna proširenja za diplomski rad
- HTTPS
- Docker + deployment
- Još jedan mikroservis koji podrazumeva dodavanje novog tipa korisnika (tehnička podrška) sa kojim će korisnici moći da četuju u slučaju da im je potrebna podrška (implementacija upotrebom WebSocket-a). Takođe, potencijalno može da se implementira chatbot za ove svrhe
- Provera prisustva neprikladnog sadržaja u videima pre nego što postanu dostupni za pregled uz pomoć algoritama mašinskog učenja
- Upravljanje brzinom reprodukcije videa 
- Neki jednostavniji video editing prilikom uploadovanja
- Download videa
- Kompresija/optimizacija streaming procesa
- ...
