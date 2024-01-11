package server

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
	"openlettings.com/template"
	"openlettings.com/types"
)

func (s *Server) GetLogin(w http.ResponseWriter, r *http.Request) {
	Login := template.Login()
	template.MainLayout(Login).Render(r.Context(), w)
}

func (s *Server) GetHomePage(w http.ResponseWriter, r *http.Request) {
	Home := template.Home()
	template.MainLayout(Home).Render(r.Context(), w)
}

func (s *Server) GetRegistration(w http.ResponseWriter, r *http.Request) {
	Register := template.RegisterPage()
	template.MainLayout(Register).Render(r.Context(), w)
}

func (s *Server) GetAdminListings(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user-id")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if userId == nil {
		http.Error(w, "Unauthorized", http.StatusForbidden)
	}

	empty1 := types.NewRange(0, 0)
	empty2 := types.NewRange(0, 0)
	propertyFilter := types.NewPropertyFilter("", "", *empty1, *empty2, userId.(int))
	properties, err := s.db.GetProperties(propertyFilter, page, pageSize)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	Listings := template.MyListingsPage(properties)
	template.AdminLayout(Listings).Render(r.Context(), w)
}

func (s *Server) GetListings(w http.ResponseWriter, r *http.Request) {

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if pageSize == 0 {
		pageSize = 10
	}
	propertyType := r.URL.Query().Get("type")
	address := r.URL.Query().Get("address")
	upperBedCount, _ := strconv.Atoi(r.URL.Query().Get("upperBedCount"))
	lowerBedcount, _ := strconv.Atoi(r.URL.Query().Get("lowerBedCount"))
	lowerPriceRange, _ := strconv.Atoi(r.URL.Query().Get("lowerPriceRange"))
	upperPriceRange, _ := strconv.Atoi(r.URL.Query().Get("upperPriceRange"))

	priceRange := types.NewRange(upperPriceRange, lowerPriceRange)
	bedRange := types.NewRange(upperBedCount, lowerBedcount)

	propertyFilter := types.NewPropertyFilter(propertyType, address, *priceRange, *bedRange, 0)

	property_count, err := s.db.GetPropertyCount()

	if err != nil {
		fmt.Println(err.Error())
	}

	hasNextPage := false

	if property_count-(page*pageSize) > 0 {
		hasNextPage = true
	}
	fmt.Println("property count ", property_count)
	fmt.Println(hasNextPage)

	properties, err := s.db.GetProperties(propertyFilter, page, pageSize)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	rawQuery := r.URL.Path
	re := regexp.MustCompile("page=\\d+")
	var nextPage string
	if re.MatchString(rawQuery) {
		nextPage = string(re.ReplaceAll([]byte(rawQuery), []byte(fmt.Sprintf("page=%d", page+1))))
	} else {
		nextPage = rawQuery + "?page=2"
	}

	if page == 1 {
		Listings := template.ListingsPage(*propertyFilter, properties, hasNextPage, nextPage)
		template.MainLayout(Listings).Render(r.Context(), w)
		return
	}
	fmt.Println("this should happen")
	template.Listings(properties, hasNextPage, nextPage).Render(r.Context(), w)

}

func (s *Server) GetListingDetails(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user-id")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	property, err := s.db.GetPropertyDetails(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if userId == nil {
		ListingDetails := template.ListingDetails(*property)
		template.MainLayout(ListingDetails).Render(r.Context(), w)
		return
	}

	template.AdminListingDetails().Render(r.Context(), w)

}

func (s *Server) GetNewPropertyForm(w http.ResponseWriter, r *http.Request) {
	CreateNewPropertyForm := template.NewPropertyForm()
	template.AdminLayout(CreateNewPropertyForm).Render(r.Context(), w)
}
