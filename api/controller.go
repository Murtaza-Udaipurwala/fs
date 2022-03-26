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

	md, hErr := c.s.GetMetaData(id)
	if hErr != nil {
		return ctx.Status(hErr.Status).JSON(fiber.Map{
			"error": hErr.Msg,
		})
	}

	buff, hErr := c.s.Retrieve(id)
	if hErr != nil {
		return ctx.Status(hErr.Status).JSON(fiber.Map{
			"error": hErr.Msg,
		})
	}

	if md.IsOneTime {
		err := c.s.Delete(id)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return ctx.Status(fiber.StatusOK).Send(buff)
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	f, hErr := parseForm(ctx)
	if hErr != nil {
		return ctx.Status(hErr.Status).JSON(fiber.Map{"error": hErr.Msg})
	}

	id, err := genID(f.Ext)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": err.Error()},
		)
	}

	f.ID = id

	exp, hErr := c.s.Create(ctx, f)
	if hErr != nil {
		return ctx.Status(hErr.Status).JSON(
			fiber.Map{"error": err.Error()},
		)
	}

	url := fileURL(f.ID)
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"url":    url,
		"expiry": fmt.Sprintf("%0.fh", exp),
	})
}
