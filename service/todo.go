package service

import (
	"context"
	"database/sql"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	result, err := s.db.ExecContext(ctx, insert, subject,description)

	if err != nil {

        return nil, err
    }

	id,err := result.LastInsertId();
	if err !=nil{
		return nil ,err
	}


	row := s.db.QueryRowContext(ctx,confirm,id)

	todo := &model.TODO{ID: id}
	
	if err := row.Scan(&todo.Subject, &todo.Description, &todo.CreatedAt,&todo.UpdatedAt); err != nil {
		return nil, err
	}

	return todo, nil
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)
	var todos []*model.TODO
	var rows *sql.Rows
	var error error

	if prevID > 0 {
		rows,error = s.db.QueryContext(ctx,readWithID,prevID,size)
	}else{
		rows,error = s.db.QueryContext(ctx,read,size)
	}
		

		if error != nil {
			return nil,error
		}
		defer rows.Close()
		
		for rows.Next(){
			todo := &model.TODO{}
			if err := rows.Scan(&todo.ID,&todo.Subject, &todo.Description, &todo.CreatedAt,&todo.UpdatedAt); err != nil {
				return nil, err
			}
			todos = append(todos, todo)
		}
		
		if err := rows.Err(); err != nil {
			return nil, err
		}
	
		  if len(todos) == 0 {
			return []*model.TODO{}, nil
		}
		
		return todos,nil
}

		


// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	 result,err := s.db.ExecContext(ctx,update,subject,description,id)
	 if err!=nil{
		return nil,err
	}
	num,err := result.RowsAffected();
	if err != nil{
		return nil,err
	}
	if(num == 0){
		err := &model.ErrNotFound{"Update件数が0でした"}
		return nil,err
	}

	row := s.db.QueryRowContext(ctx,confirm,id)
	todo := &model.TODO{ID: id}
	if err:= row.Scan(&todo.Subject,&todo.Description,&todo.CreatedAt,&todo.UpdatedAt); err !=nil{
		return nil,err
	}
		
	return todo, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
