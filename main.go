package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"task-web-dev-with-bootstrap/connection"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id         int
	Title      string
	Content    string
	StartDate  string
	EndDate    string
	Duration   string
	NodeJS     bool
	NextJS     bool
	ReactJS    bool
	TypeScript bool
	Image      string
}

// var dataProject = []Project{
// 	{
// 		Title:      "Dumbways Web Apps 2023",
// 		Content:    "Content",
// 		StartDate:  "2023-05-08",
// 		EndDate:    "2023-06-08",
// 		Duration:   "1 month",
// 		NodeJS:     true,
// 		NextJS:     true,
// 		ReactJS:    true,
// 		TypeScript: true,
// 	},
// 	{
// 		Title:      "Dumbways Web Apps 2023",
// 		Content:    "Content 2",
// 		StartDate:  "2023-05-08",
// 		EndDate:    "2023-06-08",
// 		Duration:   "1 month",
// 		NodeJS:     true,
// 		NextJS:     true,
// 		ReactJS:    true,
// 		TypeScript: true,
// 	},
// }

func main() {
	connection.DatabaseConnect()

	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/my-testimonials", testimonials)
	e.GET("/add-project", addProjectForm)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/update-project/:id", updateProjectForm)

	e.POST("/add-project", addProject)
	e.POST("/project-delete/:id", deleteProject)
	e.POST("/update-project/:id", updatedProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, start_date, end_date, duration, content, nodejs, nextjs, reactjs, typescript, image FROM tb_project")

	var result []Project
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.Id, &each.Title, &each.StartDate, &each.EndDate, &each.Duration, &each.Content, &each.NodeJS, &each.NextJS, &each.ReactJS, &each.TypeScript, &each.Image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		result = append(result, each)
	}

	var tmpl, errTemplate = template.ParseFiles("views/index.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemplate.Error()})
	}

	projects := map[string]interface{}{
		"Projects": result,
	}

	return tmpl.Execute(c.Response(), projects)

}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func testimonials(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/my-testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"mesasge": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addProjectForm(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func calculateDuration(startDateInput string, endDateInput string) string {
	startTime, _ := time.Parse("2006-01-02", startDateInput)
	endTime, _ := time.Parse("2006-01-02", endDateInput)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " years"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + " year"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " months"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " month"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " weeks"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " week"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " days"
				} else {
					duration = strconv.Itoa(durationDays) + " day"
				}
			}
		}
	}

	return duration
}

func addProject(c echo.Context) error {
	title := c.FormValue("inputProjectName")
	content := c.FormValue("inputDescription")
	startDateInput := c.FormValue("inputStartDate")
	endDateInput := c.FormValue("inputEndDate")
	duration := calculateDuration(startDateInput, endDateInput)
	var nodeJs bool
	if c.FormValue("NodeJs") == "yes" {
		nodeJs = true
	}

	var nextJs bool
	if c.FormValue("NextJs") == "yes" {
		nextJs = true
	}

	var reactJs bool
	if c.FormValue("ReactJs") == "yes" {
		reactJs = true
	}

	var typeScript bool
	if c.FormValue("TypeScript") == "yes" {
		typeScript = true
	}
	image := c.FormValue("imageUploud")

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (title, start_date, end_date, duration, content, nodejs, nextjs, reactjs, typescript, image) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", title, startDateInput, endDateInput, duration, content, nodeJs, nextJs, reactJs, typeScript, image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, start_date, end_date, duration, content, nodejs, nextjs, reactjs, typescript, image FROM tb_project WHERE id=$1", id).Scan(
		&ProjectDetail.Id, &ProjectDetail.Title, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Duration, &ProjectDetail.Content, &ProjectDetail.NodeJS, &ProjectDetail.NextJS, &ProjectDetail.ReactJS, &ProjectDetail.TypeScript, &ProjectDetail.Image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, errTemplate = template.ParseFiles("views/project-detail.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemplate.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Index : ", id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func updateProjectForm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, start_date, end_date, duration, content, nodejs, nextjs, reactjs, typescript, image FROM tb_project WHERE id=$1", id).Scan(
		&ProjectDetail.Id, &ProjectDetail.Title, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Duration, &ProjectDetail.Content, &ProjectDetail.NodeJS, &ProjectDetail.NextJS, &ProjectDetail.ReactJS, &ProjectDetail.TypeScript, &ProjectDetail.Image)

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, errTemplate = template.ParseFiles("views/update-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemplate.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func updatedProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	title := c.FormValue("inputProjectName")
	content := c.FormValue("inputDescription")
	startDateInput := c.FormValue("inputStartDate")
	endDateInput := c.FormValue("inputEndDate")
	duration := calculateDuration(startDateInput, endDateInput)
	var nodeJs bool
	if c.FormValue("NodeJs") == "yes" {
		nodeJs = true
	}

	var nextJs bool
	if c.FormValue("NextJs") == "yes" {
		nextJs = true
	}

	var reactJs bool
	if c.FormValue("ReactJs") == "yes" {
		reactJs = true
	}

	var typeScript bool
	if c.FormValue("TypeScript") == "yes" {
		typeScript = true
	}

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET title=$1, start_date=$2, end_date=$3, duration=$4, content=$5, nodejs=$6, nextjs=$7, reactjs=$8, typescript=$9 WHERE id=$10", title, startDateInput, endDateInput, duration, content, nodeJs, nextJs, reactJs, typeScript, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}
