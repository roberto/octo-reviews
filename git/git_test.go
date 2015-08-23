package git

import (
	"gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type GitSuite struct {
	localInfo LocalInfo
}

var _ = check.Suite(&GitSuite{})

func (s *GitSuite) SetUpSuite(c *check.C) {
	s.localInfo = NewLocalInfo()
}

func (s *GitSuite) TestBranchName(c *check.C) {
	c.Assert(s.localInfo.Branch, check.Equals, "master")
}

func (s *GitSuite) TestRepositoryOwner(c *check.C) {
	c.Assert(s.localInfo.Owner, check.Equals, "roberto")
}

func (s *GitSuite) TestRepositoryProject(c *check.C) {
	c.Assert(s.localInfo.Project, check.Equals, "octo-reviews")
}
