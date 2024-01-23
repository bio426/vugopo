package task

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/bio426/vugopo/datasource"
	"github.com/bio426/vugopo/pkg/auth"
	"github.com/bio426/vugopo/pkg/core"
)

type TaskCtl core.Controller

func (ctl *TaskCtl) ListByUser(c echo.Context) error {
	userId := c.Get(auth.CtxUserKey).(string)

	// service
	rows, err := datasource.Postgres.QueryContext(
		c.Request().Context(),
		"select id,title,public,done,created_at from tasks where user_id = $1",
		userId,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	type row = struct {
		Id        int32     `json:"id"`
		Title     string    `json:"title"`
		Public    bool      `json:"public"`
		Done      bool      `json:"done"`
		CreatedAt time.Time `json:"createdAt"`
	}
	rowSlice := []row{}
	for rows.Next() {
		var row = row{}
		if err = rows.Scan(
			&row.Id,
			&row.Title,
			&row.Public,
			&row.Done,
			&row.CreatedAt,
		); err != nil {
			return err
		}
		rowSlice = append(rowSlice, row)
	}
	// ~
	res := struct {
		Data []row `json:"data"`
	}{
		Data: rowSlice,
	}

	return c.JSON(http.StatusOK, res)
}

func (ctl *TaskCtl) Create(c echo.Context) error {
	body := struct {
		Title  string `json:"title" validate:"required"`
		Public bool   `json:"public" validate:"required|boolean"`
	}{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	userId := c.Get(auth.CtxUserKey).(string)
	// service
	_, err := datasource.Postgres.ExecContext(
		c.Request().Context(),
		"insert into tasks(title,public,user_id) values ($1,$2,$3)",
		body.Title, body.Public, userId,
	)
	if err != nil {
		return err
	}
	// ~

	return c.NoContent(http.StatusCreated)
}

func (ctl *TaskCtl) ListPublic(c echo.Context) error {
	// service
	rows, err := datasource.Postgres.QueryContext(
		c.Request().Context(),
		`
        select 
          t.id, 
          t.title, 
          t.done, 
          count(ut.id) as interesteds, 
          u.username as owner, 
          t.created_at 
        from 
          tasks t 
          join jt_users_tasks ut on ut.task_id = t.id 
          join users u on u.id = t.user_id 
        group by 
          t.id, 
          owner;
        `,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	type row = struct {
		Id          int32     `json:"id"`
		Title       string    `json:"title"`
		Done        bool      `json:"done"`
		Interesteds int32     `json:"interesteds"`
		Owner       string    `json:"owner"`
		CreatedAt   time.Time `json:"createdAt"`
	}
	rowSlice := []row{}
	for rows.Next() {
		var row = row{}
		if err = rows.Scan(
			&row.Id,
			&row.Title,
			&row.Done,
			&row.Interesteds,
			&row.Owner,
			&row.CreatedAt,
		); err != nil {
			return err
		}
		rowSlice = append(rowSlice, row)
	}
	// ~
	res := struct {
		Data []row `json:"data"`
	}{
		Data: rowSlice,
	}

	return c.JSON(http.StatusOK, res)
}

func (ctl *TaskCtl) MyTest(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")

	myWriter := c.Response().Writer
	flusher, ok := c.Response().Writer.(http.Flusher)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// Simulate event generation (replace with your actual event source)

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second * 2)
		if i == 3 {
			msg := fmt.Sprintf("event: custom\ndata: other data %d\n\n", i)
			fmt.Fprintf(myWriter, msg)
		} else {
			fmt.Fprintf(myWriter, "data: Event number %d\n\n", i)
		}
		flusher.Flush()
	}
	c.Response().Writer.WriteHeader(200)

	return nil
}

var Controller = &TaskCtl{}
