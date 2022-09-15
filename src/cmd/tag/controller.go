package tag

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	httperrors "github.com/myrachanto/erroring"
	"github.com/myrachanto/imagery"
)

// tagController ...
var (
	TagController TagControllerInterface = tagController{}
	Bizname       string
)

type TagControllerInterface interface {
	Create(c echo.Context) error
	Create1(c echo.Context) error
	GetOne(c echo.Context) error
	GetAll(c echo.Context) error
	Featured(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type tagController struct {
	service TagServiceInterface
}

func NewtagController(ser TagServiceInterface) TagControllerInterface {
	return &tagController{
		ser,
	}
}

// ///////controllers/////////////////
func (controller tagController) Create(c echo.Context) error {
	Bizname = c.Get("bizname").(string)
	tag := &Tag{}
	// code := c.Param("majorcode")

	fmt.Println(">>>>>>>>>>>tag Bizname", Bizname, c.Request().Body)
	tag.Name = c.FormValue("name")
	tag.Description = c.FormValue("description")
	tag.Title = c.FormValue("title")
	tag.Shopalias = c.Get("bizname").(string)
	// fmt.Println(">>>>>>>>>>>tag create", tag)
	_, err1 := controller.service.Create(tag)
	if err1 != nil {
		return c.JSON(err1.Code(), err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}// ///////controllers/////////////////
func (controller tagController) Create1(c echo.Context) error {
	Bizname = c.Get("bizname").(string)
	tag := &Tag{}
	err := c.Bind(&tag); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	// fmt.Println(">>>>>>>>>>>tag Bizname create 1", tag)
	tag, err1 := controller.service.Create(tag)
	if err1 != nil {
		return c.JSON(err1.Code(), err1)
	}
	return c.JSON(http.StatusCreated, tag)
}

func (controller tagController) GetAll(c echo.Context) error {
	Bizname = c.Get("bizname").(string)
	// fmt.Println(">>>>>>>>>>>tag Bizname", Bizname)
	tags, err3 := controller.service.GetAll()
	if err3 != nil {
		return c.JSON(err3.Code(), err3)
	}
	return c.JSON(http.StatusOK, tags)
}
func (controller tagController) GetOne(c echo.Context) error {
	Bizname = c.Get("bizname").(string)
	id := c.Param("code")
	fmt.Println(">>>>>>>>>>>tag Bizname get one=====>", c.Request().Body)
	tag, problem := controller.service.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code(), problem)
	}
	return c.JSON(http.StatusOK, tag)
}
func (controller tagController) Featured(c echo.Context) error {
	Bizname = c.Get("bizname").(string)
	code := c.Param("code")
	status := c.FormValue("status")
	fmt.Println("ccccccccccccccccccccc", status)
	feat, err := strconv.ParseBool(status)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Unable to parse the status!")
	}
	fmt.Println("jjjjjjjjjjjjjjj", feat)
	problem := controller.service.Featured(code, feat)
	if problem != nil {
		return c.JSON(problem.Code(), problem)
	}
	return c.JSON(http.StatusOK, "updated succesifully")
}
func (controller tagController) Update(c echo.Context) error {
	Bizname = c.Get("bizname").(string)
	tag := &Tag{}
	tag.Name = c.FormValue("name")
	tag.Description = c.FormValue("description")
	tag.Title = c.FormValue("title")
	id := c.Param("code")
	fmt.Println("----------------", id)

	image := c.FormValue("image")
	pic, err2 := c.FormFile("picture")
	if image == "yes" {
		//    fmt.Println(pic.Filename)
		if err2 != nil {
			httperror := httperrors.NewBadRequestError("Invalid picture")
			return c.JSON(httperror.Code(), err2)
		}
		src, err := pic.Open()
		if err != nil {
			httperror := httperrors.NewBadRequestError("the picture is corrupted")
			return c.JSON(httperror.Code(), err)
		}
		defer src.Close()
		// filePath := "./public/imgs/blogs/"
		filePath := "./src/public/imgs/tags/" + Bizname[0:7] + pic.Filename
		filePath1 := "/imgs/tags/" + Bizname[0:7] + pic.Filename
		// Destination
		dst, err4 := os.Create(filePath)
		if err4 != nil {
			httperror := httperrors.NewBadRequestError("the Directory mess")
			return c.JSON(httperror.Code(), err4)
		}
		defer dst.Close()
		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			if err2 != nil {
				httperror := httperrors.NewBadRequestError("error filling")
				return c.JSON(httperror.Code(), httperror)
			}
		}
		imagery.Imageryrepository.Imagetype(filePath, filePath, 200, 200)

		tag.Picture = filePath1
		_, err1 := controller.service.Update(id, tag)
		if err1 != nil {
			return c.JSON(err1.Code(), err1)
		}
		if _, err = io.Copy(dst, src); err != nil {
			if err2 != nil {
				httperror := httperrors.NewBadRequestError("error filling")
				return c.JSON(httperror.Code(), httperror)
			}
		}
		return c.JSON(http.StatusCreated, "user created succesifully")
	}
	// fmt.Println("vvvvvvvv---------", id)
	_, problem := controller.service.Update(id, tag)
	if problem != nil {
		return c.JSON(problem.Code(), problem)
	}
	return c.JSON(http.StatusCreated, "updated successifuly")
}

func (controller tagController) Delete(c echo.Context) error {
	Bizname = c.Get("bizname").(string)
	id := c.Param("code")
	success, failure := controller.service.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code(), failure)
	}
	return c.JSON(http.StatusOK, success)

}
