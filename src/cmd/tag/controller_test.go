package tag

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/myrachanto/microservice/tag/src/middle"
	"github.com/myrachanto/microservice/tag/src/pasetos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	authorisationHeaderKey = "Authorization"
	authorisationType      = "Bearer"
)

var data = &pasetos.Data{
	Username: "myrachanto",
	Email:    "myrachanto@gmail.com",
	// Shops: [],
}

// end to end testing
var (
	cont = NewtagController(NewtagService(NewtagRepo()))
)

func formdata() *bytes.Buffer {
	// buf := new(bytes.Buffer)
	// w := multipart.NewWriter(buf)
	// label, _ := w.CreateFormField("name")
	// label.Write([]byte("Anto"))
	// title, _ := w.CreateFormField("title")
	// title.Write([]byte("title here"))
	// description, _ := w.CreateFormField("description")
	// description.Write([]byte("description here"))

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	bodyWriter.WriteField("name", "Anto")
	bodyWriter.WriteField("title", "title here")
	bodyWriter.WriteField("description", "description here")
	// buf := new(bytes.Buffer)
	// Create file field
	// fw, err := w.CreateFormFile("upload", tarball)
	// if err != nil {
	// 	return err
	// }
	// fd, err := os.Open(tarball)
	// if err != nil {
	// 	return err
	// }
	// defer fd.Close()
	// Write file field from file to upload
	// _, err = io.Copy(fw, fd)
	// if err != nil {
	// 	return err
	// }
	// Important if you do not close the multipart writer you will not have a
	// terminating boundry
	// w.Close()

	// return buf
	bodyWriter.Close()
	return bodyBuf
}

func addAuthorization(t *testing.T, request *http.Request, Authorizationtype string, data *pasetos.Data, duration time.Duration) {
	PasetoMaker, _ := pasetos.NewPasetoMaker()
	token, err := PasetoMaker.CreateToken(data, duration)
	require.EqualValues(t, nil, err)
	authorizationHeader := fmt.Sprintf("%s %s", Authorizationtype, token)
	request.Header.Set(authorisationHeaderKey, authorizationHeader)

}
func CreateTestProduct(t *testing.T, json []byte) *httptest.ResponseRecorder {
	e := echo.New()
	endpoint := "/api/tags1"

	request, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(json))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	addAuthorization(t, request, authorisationType, data, time.Minute)
	e.POST(endpoint, cont.Create1, middle.PasetoAuthMiddleware)

	response := httptest.NewRecorder()
	e.ServeHTTP(response, request)
	return response
}
func TestCreatetag(t *testing.T) {
	testCases := []struct {
		name          string
		jsond         []byte
		ResponseCode  int
		ResponceName  string
		ResponceTitle string
		errs          string
	}{
		{name: "ok",
			jsond:         []byte(`{"name":"name","title":"this title","description":"description"}`),
			ResponseCode:  http.StatusCreated,
			ResponceName:  "name",
			ResponceTitle: "this title",
			errs:          "",
		},
		{name: "empty name",
			jsond:         []byte(`{"name":"","title":"this title","description":"description"}`),
			ResponseCode:  http.StatusBadRequest,
			ResponceName:  "",
			ResponceTitle: "",
			errs:          "bad request",
		},
		{name: "empty title",
			jsond:         []byte(`{"name":"name","title":"","description":"description"}`),
			ResponseCode:  http.StatusBadRequest,
			ResponceName:  "",
			ResponceTitle: "",
			errs:          "bad request",
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			response := CreateTestProduct(t, tc.jsond)
			tag := &Tag{}
			json.NewDecoder(response.Body).Decode(&tag)
			assert.EqualValues(t, tc.ResponseCode, response.Code)
			assert.EqualValues(t, tc.ResponceName, tag.Name)
			assert.EqualValues(t, tc.ResponceTitle, tag.Title)
			//cleanup
			service := NewtagService(NewtagRepo())
			service.Delete(tag.Code)
		})
	}

}

// func TestCreatetag1(t *testing.T) {
// 	e := echo.New()
// 	// endpoint := "http://localhost:2200/api/tags"
// 	endpoint := "/api/tags"

// 	// request, _ := http.NewRequest("POST", endpoint, formdata())
// 	request, _ := http.NewRequest("POST", endpoint, formdata())
// 	request.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
// 	addAuthorization(t, request, authorisationType, data, time.Minute)
// 	e.POST(endpoint, cont.Create, middle.PasetoAuthMiddleware)

// 	fmt.Println("step1>>>>>>>>>>>>>-----------", request)
// 	response := httptest.NewRecorder()
// 	fmt.Println("step2>>>>>>>>>>>>>-----------")
// 	// c := e.NewContext(request, response)

//		fmt.Println("step3>>>>>>>>>>>>>-----------")
//		e.ServeHTTP(response, request)
//		// client := &http.Client{}
//		// res, _ := client.Do(request)
//		// fmt.Println("step4>>>>>>>>>>>>>-----------", res.StatusCode)
//		// Assertions
//		// if assert.NoError(t, cont.Create(c)) {
//		// 	assert.Equal(t, http.StatusCreated, recorder.Code)
//		fmt.Println("step 5>>>>>>>>>>>>>-----------", response.Body)
//		// 	// assert.Equal(t, userJSON, recorder.Body.String())
//		// }
//	}
func TestGetAlltag(t *testing.T) {
	testCases := []struct {
		name          string
		jsond         []byte
		ResponseCode  int
		ResponceName  string
		ResponceTitle string
	}{
		{name: "ok",
			jsond:         []byte(`{"name":"name","title":"this title","description":"description"}`),
			ResponseCode:  http.StatusOK,
			ResponceName:  "name",
			ResponceTitle: "this title",
		},
		{name: "empty results",
			ResponseCode:  200,
			ResponceName:  "",
			ResponceTitle: "",
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			_ = CreateTestProduct(t, tc.jsond)
			e := echo.New()
			endpoint := "/api/tags"

			request, _ := http.NewRequest("GET", endpoint, nil)
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			addAuthorization(t, request, authorisationType, data, time.Minute)
			e.GET(endpoint, cont.GetAll, middle.PasetoAuthMiddleware)

			response := httptest.NewRecorder()
			e.ServeHTTP(response, request)
			// assert.EqualValues(t, 200, response.Code)
			tags := []Tag{}
			json.NewDecoder(response.Body).Decode(&tags)
			assert.EqualValues(t, tc.ResponseCode, response.Code)
			//cleanup
			service := NewtagService(NewtagRepo())
			if len(tags) > 0 {
				assert.EqualValues(t, tc.ResponceName, tags[0].Name)
				assert.EqualValues(t, tc.ResponceTitle, tags[0].Title)
				service.Delete(tags[0].Code)
			}
		})
	}
	e := echo.New()
	endpoint := "/api/tags"

	request, _ := http.NewRequest("GET", endpoint, nil)
	addAuthorization(t, request, authorisationType, data, time.Minute)
	e.GET(endpoint, cont.GetAll, middle.PasetoAuthMiddleware)

	response := httptest.NewRecorder()
	e.ServeHTTP(response, request)
	assert.EqualValues(t, 200, response.Code)
}

// func TestGetOnetag(t *testing.T) {
// 	testCases := []struct {
// 		name          string
// 		jsond         []byte
// 		ResponseCode  int
// 		ResponceName  string
// 		ResponceTitle string
// 	}{
// 		{name: "ok",
// 			jsond:        []byte(`{"name":"name","title":"this title","description":"description"}`),
// 			ResponseCode: http.StatusOK,
// 		},
// 		{name: "not found",
// 			ResponseCode: 404,
// 		},
// 	}
// 	for i := range testCases {
// 		tc := testCases[i]
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := CreateTestProduct(t, tc.jsond)
// 			tagi := &Tag{}
// 			json.NewDecoder(res.Body).Decode(&tagi)
// 			e := echo.New()
// 			fmt.Println("777777777777777", tagi)
// 			// tagi.Code = "tagCodeshopa1-16623"
// 			endpoint := "/api/tags/" + tagi.Code
// 			// values := map[string]string{"code": tagi.Code}
// 			// jsonData, _ := json.Marshal(values)
// 			// fmt.Println("777777777777777", tagi.Code)
// 			// request, _ := http.NewRequest("GET", endpoint, bytes.NewBuffer(jsonData))
// 			request, _ := http.NewRequest("GET", endpoint, nil)
// 			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 			addAuthorization(t, request, authorisationType, data, time.Minute)
// 			// q := request.URL.Query()
// 			// q.Add("code", tagi.Code)
// 			// request.URL.RawQuery = q.Encode()

// 			e.GET(endpoint, cont.GetOne, middle.PasetoAuthMiddleware)
// 			response := httptest.NewRecorder()
// 			e.ServeHTTP(response, request)

// 			// assert.EqualValues(t, 200, response.Code)
// 			tag := Tag{}
// 			json.NewDecoder(response.Body).Decode(&tag)
// 			fmt.Println(">>>>>>>>>>>>>>>ADSgdghhh", tag)
// 			assert.EqualValues(t, tc.ResponseCode, response.Code)
// 			//cleanup
// 			service := NewtagService(NewtagRepo())
// 			if tag.Code != "" {
// 				assert.EqualValues(t, tc.ResponceName, tag.Name)
// 				assert.EqualValues(t, tc.ResponceTitle, tag.Title)
// 				service.Delete(tag.Code)
// 			}
// 			service.Delete(tag.Code)
// 		})
// 	}
// }

func TestUpdatetag(t *testing.T) {
	e := echo.New()
	endpoint := "/api/tags"

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPut, endpoint, formdata())
	request.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
	addAuthorization(t, request, authorisationType, data, time.Minute)
	e.PUT(endpoint, cont.Create, middle.PasetoAuthMiddleware)
	client := &http.Client{}
	client.Do(request)
	require.Equal(t, http.StatusOK, recorder.Code)
}
func TestDeletetag(t *testing.T) {
	e := echo.New()
	endpoint := "/api/tags"

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, endpoint, nil)
	addAuthorization(t, request, authorisationType, data, time.Minute)
	e.DELETE(endpoint, cont.Create, middle.PasetoAuthMiddleware)
	require.Equal(t, http.StatusOK, recorder.Code)
	fmt.Println(">>>>>>>>>>>>>-----------", recorder)
}

// jsond := []byte(`{"name":"name","title":"this title","description":"description"}`)
// response := CreateTestProduct(t, jsond)
// assert.Equal(t, http.StatusCreated, response.Code)
// tag := &Tag{}
// json.NewDecoder(response.Body).Decode(&tag)
// assert.EqualValues(t, 201, response.Code)
// assert.EqualValues(t, tag.Name, "name")
// assert.EqualValues(t, tag.Title, "this title")
// //cleanup
// service := NewtagService(NewtagRepo())
// service.Delete(tag.Code)
