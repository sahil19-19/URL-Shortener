package controllers

import (
	"database/sql"
	"url-shortener/models"
	"url-shortener/services"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Health OK",
	})
}

func CreateShortURL(c *fiber.Ctx) error {
	var url models.URL
	if err := c.BodyParser(&url); err != nil || url.OriginalURL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	if !govalidator.IsURL(url.OriginalURL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid URL",
		})
	}

	errCode := services.CheckSpecialCharacter(url.CustomURL)
	if errCode != 3 {
		if errCode == 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Custom URL cannot be longer than 10 digits",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "No special characters in custom URL are allowed",
		})
	}

	url.OriginalURL = services.EnforceHTTP(url.OriginalURL)

	if !services.CheckDomain(url.OriginalURL) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Cannot shorten this URL as DOMAIN matches URL",
		})
	}

	// need to check if entered customURL already redirects to some original URL
	if url.CustomURL != "" && services.IsURLTaken(url.CustomURL) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Specfied custom URL already in use",
		})
	}

	// check if we already have a shortURL for this originalURL
	if url.CustomURL == "" {
		if shortURL, _ := services.CheckURLExists(url.OriginalURL); shortURL != "" {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message":   "Short URL for given URL already exists",
				"short_url": shortURL,
			})
		}
	}

	// check if customURL doesnt contain special characters or if it is longer than 10 digits

	shortURL, err := services.GenerateAndStoreURL(url.OriginalURL, url.CustomURL)
	// shortURL, err = services.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not generate url",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "short URL generated",
		"short_url": shortURL,
	})
}

func RedirectURL(c *fiber.Ctx) error {

	shortURL := c.Params("shortURL")

	originalURL, err := services.GetOriginalURL(shortURL)

	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "URL not found",
			"error":   err,
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching original URL",
			"error":   err,
		})
	}

	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"message":      "URL fetched successfully",
	// 	"Original URL": originalURL,
	// })
	return c.Redirect(originalURL, 301) // permanent redirect
}
