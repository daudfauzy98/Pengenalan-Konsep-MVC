package model

import (
	"fmt"
	"log"
	"strings"

	"firebase.google.com/go/db"
)

// Antrian dideskripsikan disini
type Antrian struct {
	ID     string `json:"id"`
	Status bool   `json:"status"`
}

// AddAntrian function is to add queue to firebase
func AddAntrian() (bool, error) {
	_, dataAntrian, _ := GetAntrian()
	var ID string
	var antrianRef *db.Ref

	ref := client.NewRef("antrian")

	if dataAntrian == nil {
		ID = fmt.Sprintf("B-0")
		antrianRef = ref.Child("0")
	} else {
		ID = fmt.Sprintf("B-%d", len(dataAntrian))
		antrianRef = ref.Child(fmt.Sprintf("%d", len(dataAntrian)))
	}

	antrian := Antrian{
		ID:     ID,
		Status: false,
	}

	if err := antrianRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}

// GetAntrian function to get all from firebase
func GetAntrian() (bool, []map[string]interface{}, error) {
	var data []map[string]interface{}
	ref := client.NewRef("antrian")

	err := ref.Get(ctx, &data)
	if err != nil {
		log.Fatalln("Error reading from database:", err)
		return false, nil, err
	}

	return true, data, nil
}

// UpdateAntrian function is to update queue in firebase
func UpdateAntrian(idAntrian string) (bool, error) {
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-") // B-0 -> [B, 0]
	childRef := ref.Child(id[1])
	antrian := Antrian{
		ID:     idAntrian,
		Status: true,
	}

	if err := childRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}

// DeleteAntrian function is to delete queue from firebase
func DeleteAntrian(idAntrian string) (bool, error) {
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])

	if err := childRef.Delete(ctx); err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}
