package fiber

import (
	"errors"

	common_error "github.com/Bit-Bridge-Source/BitBridge-CommonService-Go/public/error"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/status"
)

func FiberErrorHandler(c *fiber.Ctx, err error) error {
	// After the next handler returns, check if there was an error
	if err != nil {
		// Check if the error is of type *ServiceError
		if serr, ok := err.(*common_error.ServiceError); ok {
			switch serr.Code {
			case common_error.InternalServerError:
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": serr.Message})
			case common_error.NotFound:
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": serr.Message})
			case common_error.BadRequest:
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": serr.Message})
			case common_error.Unauthorized:
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": serr.Message})
			case common_error.Forbidden:
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": serr.Message})
			case common_error.Conflict:
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": serr.Message})
			case common_error.TooManyRequests:
				return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": serr.Message})
			case common_error.ServiceUnavailable:
				return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": serr.Message})
			case common_error.NotImplemented:
				return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"error": serr.Message})
			case common_error.Timeout:
				return c.Status(fiber.StatusRequestTimeout).JSON(fiber.Map{"error": serr.Message})
			case common_error.UserCreationFailed:
				if mongo.IsDuplicateKeyError(serr.Cause) {
					return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": serr.Message})
				}
				if errors.Is(serr.Cause, mongo.ErrNoDocuments) {
					return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": serr.Message})
				}

				statusErr, ok := status.FromError(serr.Cause)
				if ok {
					return c.Status(int(statusErr.Code())).JSON(fiber.Map{"error": serr.Message})
				}
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": serr.Message})
			default:
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "An unexpected error occurred"})
			}
		}

		// Handle other types of errors (non-ServiceError)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return nil
}
