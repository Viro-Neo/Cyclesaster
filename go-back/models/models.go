package models

type DataFilters struct {
	Id          int
	Day         string
	Month       string
	Year        string
	Birth_year  string
	Departement string
	Gender      string
	Road        string
}

type Accident struct {
	accident_id int
	jour        string
	mois        string
	an          string
	hrmn        string
	lum         string
	dep         string
	com         string
	agg         string
	intt        string
	atm         string
	col         string
	adr         string
	lat         string
	long        string
}

type User struct {
	num_acc     int
	id_vehicule int
	num_veh     int
	place       int
	catu        int
	grav        int
	sexe        int
	an_nais     int
	trajet      int
	secu        int
	secu1       int
	secu2       int
	secu3       int
	locp        int
	actp        int
	etatp       int
}

type Vehicle struct {
	num_acc     int
	id_vehicule int
	num_veh     int
	senc        int
	catv        int
	obs         int
	obsm        int
	choc        int
	manv        int
	motor       int
	occutc      int
}

type Place struct {
	num_acc int
	catr    int
	voie    int
	v1      int
	v2      int
	circ    int
	nbv     int
	prof    int
	pr      int
	pr1     int
	plan    int
	lartpc  int
	larrout int
	surf    int
	infra   int
	situ    int
	vma     int
}
