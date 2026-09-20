// Buggeon - SelfHosted service for bug and task tracking
// Copyright (C) 2026 DEVE corp.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package repositories

import (
	"buggeon/internal/db"
	"buggeon/internal/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectRepo struct {
	collection *mongo.Collection
}

func NewProjectRepo() *ProjectRepo {
	return &ProjectRepo{
		collection: db.GetCollection("buggeon", "projects"),
	}
}

func NewProjectRepoWithDbName(dbName string) *ProjectRepo {
	return &ProjectRepo{
		collection: db.GetCollection(dbName, "projects"),
	}
}

func (r *ProjectRepo) CreateProject(ctx context.Context, project *models.Project) error {

	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, project)
	return err

}

func (r *ProjectRepo) DeleteProject(ctx context.Context, projectID primitive.ObjectID) error {

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": projectID})

	if result.DeletedCount != 1 || err != nil {
		return errors.New("Failed to delete card")
	}

	return nil

}

func (r *ProjectRepo) UpdateProject(ctx context.Context, projectID primitive.ObjectID, newProjectData *models.Project) error {

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": projectID},
		bson.M{"$set": newProjectData},
	)

	return err

}

func (r *ProjectRepo) GetProject(ctx context.Context, projectID primitive.ObjectID) (models.Project, error) {

	var project models.Project

	err := r.collection.FindOne(ctx, bson.M{"_id": projectID}).Decode(&project)

	return project, err

}

func (r *ProjectRepo) GetProjectsByIDs(ctx context.Context, projectsIDs []primitive.ObjectID) ([]models.Project, error) {

	if len(projectsIDs) == 0 {
		return []models.Project{}, nil
	}

	filter := bson.M{"_id": bson.M{"$in": projectsIDs}}
	cursor, err := r.collection.Find(ctx, filter)

	if err != nil {
		return []models.Project{}, err
	}

	defer cursor.Close(ctx)

	var projects []models.Project

	err = cursor.All(ctx, &projects)

	return projects, err

}

func (r *ProjectRepo) AddBoard(ctx context.Context, projectID, boardID primitive.ObjectID) error {

	filter := bson.M{"_id": projectID}
	update := bson.M{"$push": bson.M{"boards": boardID}}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err

}

func (r *ProjectRepo) AddMember(ctx context.Context, projectID, memberID primitive.ObjectID) error {

	filter := bson.M{"_id": projectID}
	update := bson.M{"$push": bson.M{"members": memberID}}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err

}

func (r *ProjectRepo) AddLead(ctx context.Context, projectID, leadID primitive.ObjectID) error {

	filter := bson.M{"_id": projectID}
	update := bson.M{"$set": bson.M{"lead_id": leadID}}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err

}

func (r *ProjectRepo) SetProjectLogoUrl(ctx context.Context, projectID primitive.ObjectID, logoUrl string) error {

	filter := bson.M{"_id": projectID}
	update := bson.M{"$set": bson.M{"logo_url": logoUrl}}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err

}

func (r *ProjectRepo) DeleteBoard(ctx context.Context, projectID, boardID primitive.ObjectID) error {

	filter := bson.M{"_id": projectID}
	update := bson.M{"$pull": bson.M{"boards": boardID}}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err

}

func (r *ProjectRepo) AddSchema(ctx context.Context, projectID primitive.ObjectID, schemaUrl string) error {

	filter := bson.M{"_id": projectID}

	update := bson.M{"$push": bson.M{"schemas": schemaUrl}}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err

}

func (r *ProjectRepo) GetSchemas(ctx context.Context, schemaID primitive.ObjectID) (models.Schema, error) {

	filter := bson.M{"schemas": schemaID}

	var schema models.Schema

	err := r.collection.FindOne(ctx, filter).Decode(&schema)

	return schema, err

}
