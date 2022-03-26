package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type IService interface {
	Retrieve(id string) ([]byte, *HTTPErr)
	GetMetaData(id string) (*MetaData, *HTTPErr)
	Delete(id string) error
	Create(ctx *fiber.Ctx, f *File) (float64, *HTTPErr)
}

type Controller struct {
	s IService
}

func NewController(s IService) *Controller {
	return &Controller{s}
}

func (c *Controller) Retrieve(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "favicon.ico" {
		return nil
	}

	md, hErr := c.s.GetMetaData(id)
	if hErr != nil {
		return ctx.Status(hErr.Status).Send([]byte(hErr.Msg + "\n"))
	}

	buff, hErr := c.s.Retrieve(id)
	if hErr != nil {
		return ctx.Status(hErr.Status).Send([]byte(hErr.Msg + "\n"))
	}

	if md.IsOneTime {
		err := c.s.Delete(id)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).Send(
				[]byte(err.Error() + "\n"),
			)
		}
	}

	return ctx.Status(fiber.StatusOK).Send(buff)
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	q := new(Query)
	err := ctx.QueryParser(q)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	f, hErr := parseForm(ctx)
	if hErr != nil {
		ctx.Status(hErr.Status)
		if q.JSON {
			return ctx.JSON(fiber.Map{"error": hErr.Msg})
		}

		return ctx.Send([]byte(hErr.Msg + "\n"))
	}

	id, err := genID(f.Ext)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		if q.JSON {
			return ctx.JSON(fiber.Map{"error": err.Error()})
		}
		return ctx.Send([]byte(err.Error() + "\n"))
	}

	f.ID = id

	exp, hErr := c.s.Create(ctx, f)
	if hErr != nil {
		ctx.Status(hErr.Status)
		if q.JSON {
			return ctx.JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Send([]byte(err.Error() + "\n"))
	}

	url := fileURL(f.ID)
	e := fmt.Sprintf("%0.fh", exp)

	ctx.Status(fiber.StatusCreated)
	if q.JSON {
		return ctx.JSON(fiber.Map{
			"url":    url,
			"expiry": e,
		})
	}

	return ctx.Send([]byte(
		fmt.Sprintf("%s\nFile will be deleted in %s\n", url, e)),
	)
}
