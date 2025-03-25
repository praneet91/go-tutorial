package services

import (
	"fmt"
	internal "go-tutorial/internal/model"

	"gorm.io/gorm"
)

type NotesService struct {
	db *gorm.DB
}

func (ns *NotesService) Init(db *gorm.DB) {
	ns.db = db
	ns.db.AutoMigrate(&internal.Notes{})
}

type Note struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

func (ns *NotesService) GetNotesService(status *bool) ([]*internal.Notes, error) {
	var notes []*internal.Notes
	if ns.db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := ns.db
	if status != nil {
		query = query.Where("is_completed = ?", *status)
	}

	if err := query.Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

func (ns *NotesService) CreateNoteService(description string, isCompleted bool) (*internal.Notes, error) {
	if ns.db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	note := &internal.Notes{
		Description: description,
		IsCompleted: isCompleted,
	}

	if error := ns.db.Create(note).Error; error != nil {
		return nil, error
	}

	return note, nil
}

func (ns *NotesService) UpdateNoteService(description string, isCompleted bool, id int) (*internal.Notes, error) {
	if ns.db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	var note *internal.Notes

	if err := ns.db.Where("id = ?", id).First(&note).Error; err != nil {
		return nil, err
	}

	note.Description = description
	note.IsCompleted = isCompleted

	if error := ns.db.Save(&note).Error; error != nil {
		return nil, error
	}

	return note, nil
}

func (ns *NotesService) DeleteNoteService(id int) error {
	if ns.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	if err := ns.db.Where("id = ?", id).Delete(&internal.Notes{}).Error; err != nil {
		return err
	}

	return nil
}

func (ns *NotesService) GetNoteService(id int) (*internal.Notes, error) {
	if ns.db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	var note *internal.Notes
	if err := ns.db.Where("id = ?", id).First(&note).Error; err != nil {
		return nil, err
	}

	return note, nil
}
