package db

import (
	"database/sql"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"

	"github.com/concourse/concourse/atc"
)

//go:generate counterfeiter . WorkerArtifact

type WorkerArtifact interface {
	ID() int
	Name() string
	BuildID() int
	CreatedAt() time.Time
	Volume(teamID int) (CreatedVolume, bool, error)
	AttachToBuild(build Build) error
	Initialized() bool
}

type artifact struct {
	conn Conn

	id          int
	name        string
	buildID     int
	createdAt   time.Time
	initialized bool
}

func (a *artifact) ID() int              { return a.id }
func (a *artifact) Name() string         { return a.name }
func (a *artifact) BuildID() int         { return a.buildID }
func (a *artifact) CreatedAt() time.Time { return a.createdAt }
func (a *artifact) Initialized() bool    { return a.initialized }

func (a *artifact) Volume(teamID int) (CreatedVolume, bool, error) {
	where := map[string]interface{}{
		"v.team_id":            teamID,
		"v.worker_artifact_id": a.id,
	}

	_, created, err := getVolume(a.conn, where)
	if err != nil {
		return nil, false, err
	}

	if created == nil {
		return nil, false, nil
	}

	return created, true, nil
}

func (a *artifact) AttachToBuild(build Build) error {
	result, err := psql.Update("worker_artifacts").
		Set("build_id", build.ID()).
		Set("initialized", true).
		Where(sq.Eq{"id": a.id}).
		RunWith(a.conn).Exec()

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("artifact not found")
	}
	return nil
}

func saveWorkerArtifact(tx Tx, conn Conn, atcArtifact atc.WorkerArtifact) (WorkerArtifact, error) {

	var artifactID int

	values := map[string]interface{}{
		"name": atcArtifact.Name,
	}

	if atcArtifact.BuildID != 0 {
		values["build_id"] = atcArtifact.BuildID
		values["initialized"] = true
	}

	err := psql.Insert("worker_artifacts").
		SetMap(values).
		Suffix("RETURNING id").
		RunWith(tx).
		QueryRow().
		Scan(&artifactID)

	if err != nil {
		return nil, err
	}

	artifact, found, err := getWorkerArtifact(tx, conn, artifactID)

	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("Not found")
	}

	return artifact, nil
}

func getWorkerArtifact(tx Tx, conn Conn, id int) (WorkerArtifact, bool, error) {
	var (
		createdAtTime pq.NullTime
		buildID       sql.NullInt64
	)

	artifact := &artifact{conn: conn}

	err := psql.Select("id", "created_at", "name", "build_id").
		From("worker_artifacts").
		Where(sq.Eq{
			"id": id,
		}).
		RunWith(tx).
		QueryRow().
		Scan(&artifact.id, &createdAtTime, &artifact.name, &buildID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}

		return nil, false, err
	}

	artifact.createdAt = createdAtTime.Time
	artifact.buildID = int(buildID.Int64)

	return artifact, true, nil
}
