package builds_test

import (
	"context"
	"testing"
	"time"

	"github.com/concourse/concourse/atc/builds"
	"github.com/concourse/concourse/atc/component"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/db/dbfakes"
	"github.com/concourse/concourse/atc/engine"
	"github.com/concourse/concourse/atc/engine/enginefakes"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TrackerSuite struct {
	suite.Suite
	*require.Assertions

	fakeBuildFactory *dbfakes.FakeBuildFactory
	fakeEngine       *enginefakes.FakeEngine

	tracker *builds.Tracker
}

func TestTracker(t *testing.T) {
	suite.Run(t, &TrackerSuite{
		Assertions: require.New(t),
	})
}

func (s *TrackerSuite) SetupTest() {
	s.fakeBuildFactory = new(dbfakes.FakeBuildFactory)
	s.fakeEngine = new(enginefakes.FakeEngine)

	s.tracker = builds.NewTracker(
		s.fakeBuildFactory,
		s.fakeEngine,
	)
}

func (s *TrackerSuite) TestTrackRunsStartedBuilds() {
	startedBuilds := []db.Build{}
	for i := 0; i < 3; i++ {
		fakeBuild := new(dbfakes.FakeBuild)
		fakeBuild.IDReturns(i + 1)
		startedBuilds = append(startedBuilds, fakeBuild)
	}

	s.fakeBuildFactory.GetAllStartedBuildsReturns(startedBuilds, nil)

	running := make(chan db.Build, 3)
	s.fakeEngine.NewBuildStub = func(build db.Build) engine.Runnable {
		engineBuild := new(enginefakes.FakeRunnable)
		engineBuild.RunStub = func(context.Context) {
			running <- build
		}

		return engineBuild
	}

	err := s.tracker.Run(context.TODO())
	s.NoError(err)

	ranBuilds := []db.Build{
		<-running,
		<-running,
		<-running,
	}
	s.ElementsMatch(startedBuilds, ranBuilds)
}

func (s *TrackerSuite) TestTrackerDoesntCrashWhenOneBuildPanic() {
	startedBuilds := []db.Build{}
	fakeBuild := new(dbfakes.FakeBuild)
	fakeBuild.IDReturns(1)
	startedBuilds = append(startedBuilds, fakeBuild)

	s.fakeBuildFactory.GetAllStartedBuildsReturns(startedBuilds, nil)

	fakeEngineBuild := new(enginefakes.FakeRunnable)
	fakeEngineBuild.RunStub = func(context.Context) {
		panic("something went wrong")
	}
	s.fakeEngine.NewBuildReturns(fakeEngineBuild)

	err := s.tracker.Run(context.TODO())
	s.NoError(err)

	s.Eventually(func() bool { return fakeEngineBuild.RunCallCount() == 1 }, time.Second, 10*time.Millisecond)
	s.Eventually(func() bool { return fakeBuild.FinishCallCount() == 1 }, time.Second, 10*time.Millisecond)
	s.Eventually(func() bool { return fakeBuild.FinishArgsForCall(0) == db.BuildStatusFailed }, time.Second, 10*time.Millisecond)
}

func (s *TrackerSuite) TestTrackDoesntTrackAlreadyRunningBuilds() {
	fakeBuild := new(dbfakes.FakeBuild)
	fakeBuild.IDReturns(1)
	s.fakeBuildFactory.GetAllStartedBuildsReturns([]db.Build{fakeBuild}, nil)

	wait := make(chan struct{})
	defer close(wait)

	running := make(chan db.Build, 3)
	s.fakeEngine.NewBuildStub = func(build db.Build) engine.Runnable {
		engineBuild := new(enginefakes.FakeRunnable)
		engineBuild.RunStub = func(context.Context) {
			running <- build
			<-wait
		}

		return engineBuild
	}

	err := s.tracker.Run(context.TODO())
	s.NoError(err)

	<-running

	err = s.tracker.Run(context.TODO())
	s.NoError(err)

	select {
	case <-running:
		s.Fail("another build was started!")
	case <-time.After(100 * time.Millisecond):
	}
}

func (s *TrackerSuite) TestTrackerDrainsEngine() {
	var _ component.Drainable = s.tracker

	ctx := context.TODO()
	s.tracker.Drain(ctx)
	s.Equal(1, s.fakeEngine.DrainCallCount())
	s.Equal(ctx, s.fakeEngine.DrainArgsForCall(0))
}
