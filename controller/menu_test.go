package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/JevaPrahaysuma/managemnet.git/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	// "gorm.io/driver/postgres"
	//"gorm.io/gorm"
	"github.com/jinzhu/gorm"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetMenuPostgres(t *testing.T) {
	r := SetUpRouter()
	// Handlers for testing
	r.GET("/user/", GetMenuPostgres)
	req, _ := http.NewRequest("GET", "/user/?q=Hotel C&vanues&events=&startPrice=0&endPrice=110", nil)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	//mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))

	var (
		menu_item_id         = 976059
		name_dish            = "Mug of Famous U. S. Grant Hotel Coffee"
		price                = 0.05
		sponsor              = "U.S. Grant Hotel"
		physical_description = "22.5x14cm"
		call_number          = "1921-0045_wotm"
		date                 = "1921-05-17"
		location             = "U.S. Grant Hotel"
		currency             = "Dollars"
		currency_symbol      = "$"
		status               = "complete"
		created_at           = "2012-06-20T21:08:00Z"
		updated_at           = "2012-06-20T21:08:00Z"
	)
	rows := sqlmock.NewRows([]string{"menu_item_id", "name_dish", "price", "sponsor",
		"physical_description", "call_number", "date", "location",
		"currency", "currency_symbol", "status", "created_at", "updated_at"}).AddRow(menu_item_id, name_dish, price, sponsor, physical_description, call_number, date, location, currency, currency_symbol, status, created_at, updated_at)

	config.DB, err = gorm.Open("postgres", db)
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT menu_items.id as menu_item_id, dishes.name as name_dish, 
		dishes.description as description_dish, menu_items.price, menus.name as name_menu, 
		menus.sponsor, menus.event, menus.vanue, menus.place, menus.physical_description, menus.occasion, menus.notes, menus.call_number, menus.keywords, menus.language, menus.date, menus.location, menus.location_type, menus.currency, menus.currency_symbol, menus.status, menu_items.created_at, menu_items.updated_at FROM "menu_items" LEFT JOIN menu_pages on menu_items.menu_page_id = menu_pages.id LEFT JOIN dishes on menu_items.dish_id = dishes.id LEFT JOIN menus on menu_pages.menu_id = menus.id WHERE ((menus.name LIKE '%Hotel C%' OR dishes.name LIKE '%Hotel C%') AND (menu_items.price >= 0 AND menu_items.price <= 110)) LIMIT 10`),
	).WillReturnRows(rows)
	//require.NoError(s.T(), err)

	w := httptest.NewRecorder()
	fmt.Println(w.Body.String())
	r.ServeHTTP(w, req)
	want := `{"status":200,"message":"success","data":{"data":[{"menu_item_id":976059,"name_dish":"Mug of Famous U. S. Grant Hotel Coffee","price":0.05,"sponsor":"U.S. Grant Hotel","physical_description":"22.5x14cm","call_number":"1921-0045_wotm","date":"1921-05-17","location":"U.S. Grant Hotel","currency":"Dollars","currency_symbol":"$","status":"complete","created_at":"2012-06-20T21:08:00Z","updated_at":"2012-06-20T21:08:00Z"}]}}`
	//var menu []models.ListMenu
	//json.Unmarshal(w.Body.Bytes(), &menu)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Equal(t, want, w.Body.String())
}
