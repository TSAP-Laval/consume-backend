package consume

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"

	"github.com/jinzhu/gorm"
	"github.com/tsap-laval/tsap-common/common"
	// piss
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// SeedData est une fonction permettant de populer la base de données
func SeedData() error {
	var err error

	db, err := gorm.Open("sqlite3", "database.db")
	db.LogMode(true)
	if err != nil {
		return err
	}
	defer db.Close()

	db.CreateTable(&common.Entraineur{}, &common.Joueur{},
		&common.Saison{}, &common.Lieu{}, &common.TypeAction{},
		&common.Sport{}, &common.Zone{}, &common.Niveau{},
		&common.Equipe{}, &common.Partie{}, &common.Action{},
		&common.Administrateur{}, &common.Position{},
		&common.JoueurPositionPartie{}, &common.Metrique{},
		&common.Video{})

	var joueursData []common.Joueur
	err = jsonLoad("data/joueurs.json", &joueursData)
	if err != nil {
		return err
	}
	for _, joueur := range joueursData {
		db.Create(&joueur)
	}

	var entraineursData []common.Entraineur
	err = jsonLoad("data/entraineurs.json", &entraineursData)
	if err != nil {
		return err
	}
	for _, entraineur := range entraineursData {
		db.Create(&entraineur)
	}

	var lieuxData []common.Lieu
	err = jsonLoad("data/lieux.json", &lieuxData)
	if err != nil {
		return err
	}
	for _, lieu := range lieuxData {
		db.Create(&lieu)
	}

	var saisonData []common.Saison
	err = jsonLoad("data/saisons.json", &saisonData)
	if err != nil {
		return err
	}
	for _, lieu := range saisonData {
		db.Create(&lieu)
	}

	var typeActionData []common.TypeAction
	err = jsonLoad("data/typesactions.json", &typeActionData)
	if err != nil {
		return err
	}
	for _, lieu := range typeActionData {
		db.Create(&lieu)
	}

	var sportData []common.Sport
	err = jsonLoad("data/sports.json", &sportData)
	if err != nil {
		return err
	}
	for _, sport := range sportData {
		db.Create(&sport)
	}

	var niveauData []common.Niveau
	err = jsonLoad("data/niveaux.json", &niveauData)
	if err != nil {
		return err
	}
	for _, niveau := range niveauData {
		db.Create(&niveau)
	}

	var equipeData []common.Equipe
	err = jsonLoad("data/equipes.json", &equipeData)
	if err != nil {
		return err
	}
	for i, equipe := range equipeData {
		x := &common.Sport{}
		if i%2 == 0 {
			db.First(x)
			equipe.Sport = *x
		} else {
			db.Last(x)
			equipe.Sport = *x
		}
		y := &common.Niveau{}
		db.First(y, rand.Intn(6)+1)
		equipe.Niveau = *y
		db.Create(&equipe)
	}

	admin := common.Administrateur{Email: "admin@admin.com", PassHash: "admin"}
	db.Create(admin)

	var positionData []common.Position
	err = jsonLoad("data/positions.json", &positionData)
	if err != nil {
		return err
	}
	for _, position := range positionData {
		db.Create(&position)
	}

	var metriqueData []common.Metrique
	err = jsonLoad("data/metriques.json", &metriqueData)
	if err != nil {
		return err
	}
	for _, metrique := range metriqueData {
		db.Create(&metrique)
	}

	video := common.Video{Path: "aucun video", AnalyseTermine: false}
	db.Create(video)

	var partieData []common.Partie
	err = jsonLoad("data/parties.json", &partieData)
	if err != nil {
		return err
	}
	for _, partie := range partieData {
		nb1 := 1
		nb2 := 2
		for nb1 != nb2 {
			nb1 = rand.Intn(6) + 1
			nb2 = rand.Intn(6) + 1
		}
		equipe1 := &common.Equipe{}
		db.First(equipe1, nb1)
		partie.EquipeMaison = *equipe1
		equipe2 := &common.Equipe{}
		db.First(equipe2, nb2)
		partie.EquipeAdverse = *equipe2
		saison := &common.Saison{}
		db.First(saison, rand.Intn(3)+1)
		partie.Saison = *saison
		lieu := &common.Lieu{}
		db.First(lieu, rand.Intn(100)+1)
		partie.Lieu = *lieu
		video := &common.Video{}
		db.First(video)
		partie.Video = *video

		db.Create(&partie)
	}

	zoneInsert := common.Zone{Nom: "offensive"}
	db.Create(zoneInsert)
	zoneInsert = common.Zone{Nom: "defensive"}
	db.Create(zoneInsert)

	var joueurpositionpartieData []common.JoueurPositionPartie
	err = jsonLoad("data/jpp.json", &joueurpositionpartieData)
	if err != nil {
		return err
	}
	for _, jpp := range joueurpositionpartieData {
		db.Create(&jpp)
	}

	var actionData []common.Action
	err = jsonLoad("data/actions.json", &actionData)
	if err != nil {
		return err
	}
	for _, action := range actionData {
		db.Create(&action)
	}

	return err
}

func jsonLoad(path string, out interface{}) error {
	raw, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, out)

	if err != nil {
		return err
	}

	return nil
}
