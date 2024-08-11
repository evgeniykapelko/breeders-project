package main

import (
	"breeders/models"
	"breeders/pets"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
	"net/http"
	"net/url"
	"time"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) DogOfMonth(w http.ResponseWriter, r *http.Request) {
	breed, _ := app.App.Models.DogBreed.GetBreedByName("German Shepherd Dog")

	dom, _ := app.App.Models.Dog.GetDogOfMonthByID(1)

	layout := "2006-01-02"
	dob, _ := time.Parse(layout, "2023-11-01")
	dog := models.DogOfMonth{
		Dog: &models.Dog{
			ID:          1,
			DogName:     "Sam",
			BreederID:   breed.ID,
			Color:       "Black",
			DateOfBirth: dob,
			Description: "Sam is a very good boy",
			Weight:      20,
			Breed:       *breed,
		},
		Video: dom.Video,
		Image: dom.Image,
	}

	data := make(map[string]interface{})
	data["dog"] = dog

	app.render(w, "dog-of-month.page.gohtml", &templateData{Data: data})
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))
}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dogBreeds, err := app.App.Models.DogBreed.All()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dogBreeds)
}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	p, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetBreed("mixed breed").
		SetWeight(15).
		SetDescription("Dog Builder").
		SetColor("Black and White").
		SetAge(3).
		SetAgeEstimate(true).
		Build()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}

	_ = t.WriteJSON(w, http.StatusOK, p)

}

func (app *application) CreateCatWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	p, err := pets.NewPetBuilder().
		SetSpecies("cat").
		SetBreed("mixed breed").
		SetWeight(15).
		SetDescription("Cat Builder").
		SetGeographicOrigin("Canada").
		SetColor("Black and White").
		SetAge(1).
		SetAgeEstimate(true).
		Build()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}

	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	catBreeds, err := app.App.CatService.GetAllBreeds()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}

	_ = t.WriteJSON(w, http.StatusOK, catBreeds)
}

func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	species := chi.URLParam(r, "species")
	b := chi.URLParam(r, "breed")

	breed, _ := url.QueryUnescape(b)

	pet, err := pets.NewPetWithBreedFromAbstractFactory(species, breed)
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, pet)
}
