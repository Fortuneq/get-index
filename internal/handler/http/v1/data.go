package v1

import (
	"getProject/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type IndexHandler struct {
	getIndexInteractor usecase.GetIndexInteractor
}

func (h *IndexHandler) parseIndex(ctx *fiber.Ctx) error {
	var p usecase.ParseIndexInputDTO
	if err := ctx.BodyParser(&p); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(newErrResp(err))
	}

	output, err := h.getIndexInteractor.Execute(ctx.Context(), p)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(newResp(output))
}

func (h *IndexHandler) GetItems(r fiber.Router) {
	r.Get("get-items", h.parseIndex)
}

func NewUserHandler(getIndexDataInteractor usecase.GetIndexInteractor) *IndexHandler {
	return &IndexHandler{getIndexDataInteractor}
}
