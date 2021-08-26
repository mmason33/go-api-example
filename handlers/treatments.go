package treatments

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/weareenvoy/nuface-recommendation-server-golang/database"
	"github.com/weareenvoy/nuface-recommendation-server-golang/model"
)

// GetAllTreatments from db
func GetAllTreatments(c *fiber.Ctx) error {
	// query treatments table in the database
	rows, err := database.DB.Query("SELECT title, treatment_id, firebase_id, document_date, id FROM treatments")

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return err
	}

	defer rows.Close()

	result := model.Treatments{}

	for rows.Next() {
		treatment := model.Treatment{}
		err := rows.Scan(&treatment.Title, &treatment.Treatment_id, &treatment.Firebase_id, &treatment.Document_date, &treatment.Id)

		// Exit if we get an error
		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   "err",
			})
			return err
		}
		// Append Treatment to Treatments
		result.Data = append(result.Data, treatment)
	}

	// Return Treatments in JSON format
	if err := c.JSON(&fiber.Map{
		"success":    true,
		"treatments": result,
		"message":    "All treatments returned successfully",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}

	return nil
}

// GetSingleTreatment from db
func GetSingleTreatment(c *fiber.Ctx) error {
	id := c.Params("id")
	treatment := model.Treatment{}

	// query treatment database
	row, err := database.DB.Query("SELECT title, treatment_id, firebase_id, document_date, id FROM treatments WHERE id = ?", id)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}

	defer row.Close()

	// iterate through the values of the row
	for row.Next() {
		switch err := row.Scan(&treatment.Title, &treatment.Treatment_id, &treatment.Firebase_id, &treatment.Document_date, &treatment.Id); err {
		case sql.ErrNoRows:
			log.Println("No rows were returned!")
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		case nil:
			log.Println(treatment.Title, treatment.Treatment_id, treatment.Firebase_id, treatment.Document_date, treatment.Id)
		default:
			//   panic(err)
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}

	// return treatment in JSON format
	if err := c.JSON(&fiber.Map{
		"success":   true,
		"message":   "Successfully fetched treatment",
		"treatment": treatment,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}

	return nil
}

// CreateTreatment handler
func CreateTreatment(c *fiber.Ctx) error {
	// Instantiate new Treatmet struct
	t := new(model.Treatment)

	//  Parse body into Treatmet struct
	if err := c.BodyParser(t); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}

	// Insert Treatmet into database
	res, err := database.DB.Query("INSERT INTO treatments (title, treatment_id, firebase_id, document_date) VALUES (?, ?, ?, ?)", t.Title, t.Treatment_id, t.Firebase_id, t.Document_date)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	// Print result
	log.Println(res)

	// Return Treatmet in JSON format
	if err := c.JSON(&fiber.Map{
		"success":   true,
		"message":   "Treatment successfully created",
		"treatment": t,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating treatment",
		})
		return err
	}

	return nil
}

// DeleteTreatment from db
func DeleteTreatment(c *fiber.Ctx) error {
	id := c.Params("id")

	// query treatment table in database
	res, err := database.DB.Query("DELETE FROM treatments WHERE id = ?", id)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return err
	}
	// Print result
	log.Println(res)

	// return treatment in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "treatment deleted successfully",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return err
	}

	return nil
}
